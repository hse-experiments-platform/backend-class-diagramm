package datasets

import (
	"context"

	"github.com/hse-experiments-platform/backend-class-diagramm/enum"
)

type DatasetsService interface {
	GetDatasetBatch(ctx context.Context, ids []int64) ([]*Dataset, error)
	GetDatasetRows(ctx context.Context, id, from, to int64) (Rows, error)
	GetDatasetContent(ctx context.Context, id int64) ([]byte, error)
	GetDatasetUploadStatus(ctx context.Context, id int64) (enum.UploadStatus, error)

	SearchDatasets(ctx context.Context, query string, limit int64, offset int64) ([]int64, error)

	CreateDataset(ctx context.Context, dataset *Dataset) (int64, error)
	UploadDatasetContent(ctx context.Context, data []byte) (int64, error)
	UploadDatasetContentByLink(ctx context.Context, url string) (int64, error)
	DeleteDataset(ctx context.Context, id int64) error
}
