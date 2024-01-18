package datasets

import "github.com/hse-experiments-platform/backend-class-diagramm/enum"

type Dataset struct {
	ID           int64
	Name         string
	Version      uint64
	UploadStatus enum.UploadStatus
	RowsCount    int64
	Schema       Schema
}

type Rows []Row

type Row struct {
	ID     int64
	Meta   map[string]any
	Values []any
}

type Schema = []Column

type Column struct {
	ID           int64
	Name         string
	DefaultValue any
	IsIndex      bool
}
