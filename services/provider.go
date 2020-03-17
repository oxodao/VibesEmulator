package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	gonanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/argon2"
)

// Provider is a provider
type Provider struct {
	DB          *gorm.DB
	ArgonParams *ArgonParams
}

// NewProvider returns an instance of Provider
func NewProvider() *Provider {
	db, err := gorm.Open("sqlite3", "./the_test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	argonParams := &ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return &Provider{
		DB:          db,
		ArgonParams: argonParams,
	}
}

// GenerateUID function generates an ID for the pictures
func (prv *Provider) GenerateUID(size int) string {
	str, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ", size)
	return str
}

// HashPassword will generate a hash of a password ready to be stored in the database
func (prv *Provider) HashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(prv.ArgonParams.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, prv.ArgonParams.Iterations, prv.ArgonParams.Memory, prv.ArgonParams.Parallelism, prv.ArgonParams.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, prv.ArgonParams.Memory, prv.ArgonParams.Iterations, prv.ArgonParams.Parallelism, b64Salt, b64Hash), nil
}

// VerifyPassword will check for the validity of the password in the database
func (prv *Provider) VerifyPassword(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		// @TODO Log error in DB
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

// ArgonParams represents all the parameters needed to hash the passwords
type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

/** Thanks to https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go **/
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func decodeHash(encodedHash string) (p *ArgonParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &ArgonParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
