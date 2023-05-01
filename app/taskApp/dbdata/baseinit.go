package dbdata

type InitDBData interface {
	TableName() string
	Initialize() error
	CheckDataExist() bool
}
