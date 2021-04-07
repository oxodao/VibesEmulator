package models

// Picture represent an uploaded picture
type Picture struct {
	Name       string
	UploadedBy string
}

func (p Picture) GetTableName() string {
	return "APP_PICTURES"
}

func (p Picture) GetTableCreationScript() string {
	return  `
		CREATE TABLE APP_PICTURES (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			UPLOADED_AT INTEGER DEFAULT (cast(strftime('%s','now') as int)),
			NAME VARCHAR(255),
			UPLOADED_BY VARCHAR(255)
		);
`
}