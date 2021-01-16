package patcher

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Patcher(apk, url string) {
	if url[:4] != "http" {
		url = "http://" + url
		fmt.Println("The url did not have the protocol so the new URL is now ", url, ". If you want https please specify it directly.")
	}

	if _, err := os.Stat(apk); os.IsNotExist(err) {
		fmt.Println("The given APK file does not exists")
		return
	}

	if _, err := os.Stat("patching"); !os.IsNotExist(err) {
		fmt.Println("The patching folder already exists. Please remove it since it's old junk from the APK patcher")
		return
	}

	f, err := os.Open(apk)
	if err != nil {
		fmt.Println("Something went wrong checking the APK: ", err)
		return
	}

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println("Something went wrong checking the APK: ", err)
		return
	}

	f.Close()

	hash := fmt.Sprintf("%x", h.Sum(nil))
	if hash != "954379f65df8b666f0ab3ded1f8bc3b00249eef8" {
		fmt.Println("The APK does not match the correct sha1 for version 1.2.1 (954379f65df8b666f0ab3ded1f8bc3b00249eef8)")
		return
	}

	_ = os.Mkdir("patching", os.ModePerm)

	_, err = copy(apk, "patching/the_test.apk")
	if err != nil {
		fmt.Println("Could not copy the APK!", err)
		return
	}

	fmt.Println("- Disassembling the APK")
	cmd := exec.Command("apktool", "d", "the_test.apk")
	cmd.Dir = "./patching"
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not disassemble the APK. Do you have APKTOOL installed ?")
		fmt.Println(string(out), err)
		return
	}

	fmt.Println("- Changing the urls in the files")
	couldUpdate := updateURL("patching/the_test/smali/de/lotumapps/vibes/Vibes.smali", url)
	couldUpdate = couldUpdate && updateURL("patching/the_test/smali/de/lotumapps/vibes/api/ApiSessionCookieStore.smali", url)
	couldUpdate = couldUpdate && updateURL("patching/the_test/smali/de/lotumapps/vibes/Build.smali", url)

	if !couldUpdate {
		fmt.Println("Could not update all files...")
		return
	}

	fmt.Println("- Reassembling the APK")
	cmd = exec.Command("apktool", "b", "the_test", "-o", "the_test_patched.apk")
	cmd.Dir = "./patching"
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Could not re-assemble the APK...", err)
		return
	}

	fmt.Println("- Generating a key to sign the APK")
	cmd = exec.Command("keytool", "-genkey", "-v", "-keystore", "key.keystore", "-alias", "thetest_key", "-keyalg", "RSA", "-keysize", "2048", "-validity", "10000", "-keypass", "qwerty", "-storepass", "qwerty", "-dname", "CN=VibesEmulator, OU=VibesEmulator, O=VibesEmulator, L=VibesEmulator, ST=VibesEmulator, C=VibesEmulator")
	cmd.Dir = "./patching"
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Could not generate the key to sign the APK. Do you have keytool installed ?", err)
		return
	}

	fmt.Println("- Signing the APK")
	cmd = exec.Command("jarsigner", "-sigalg", "SHA1withRSA", "-digestalg", "SHA1", "-keystore", "key.keystore", "-storepass", "qwerty", "the_test_patched.apk", "thetest_key")
	cmd.Dir = "./patching"
	out, err = cmd.Output()
	if err != nil {
		fmt.Println("Could not sign the APK. Do you have jarsigner installed ?", string(out), err)
		return
	}

	fmt.Println("- Moving it back & cleaning up temporary files")
	dest := path.Dir(apk) + "/the_test_patched.apk"
	_, err = copy("patching/the_test_patched.apk", dest)
	if err != nil {
		fmt.Println("Could not copy the file back!", err)
		return
	}
	err = os.RemoveAll("patching")
	if err != nil {
		fmt.Println("Could not remove the temporary files stored (the folder \"patching\")", err)
	}

	fmt.Println("Congratulation, you can now use the the_test_patched.apk file on your phone")
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func updateURL(file, newURL string) bool {
	read, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Could not update urls in ", file)
		return false
	}

	content := strings.ReplaceAll(string(read), "https://vibes.lotum.com", newURL)
	err = ioutil.WriteFile(file, []byte(content), 0)

	if err != nil {
		fmt.Println("Could not update urls in ", file)
		return false
	}

	return true
}