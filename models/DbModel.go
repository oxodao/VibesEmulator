package models

type DbModel interface {
	GetTableName() string
	GetTableCreationScript() string
}
