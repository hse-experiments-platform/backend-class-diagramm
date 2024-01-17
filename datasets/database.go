package datasets

type UploadStatus uint8

type Dataset struct {
	ID int64
	Name string
	Version uint64
	UploadStatus UploadStatus
	RowsCount int64
}

type Rows []Row

type Row struct {
	ID int64
	Meta map[string]any
	Values []Value
}

type Value = any