package launcher

import (
	"time"

	"github.com/hse-experiments-platform/backend-class-diagramm/datasets"
	"github.com/hse-experiments-platform/backend-class-diagramm/enum"
	"github.com/hse-experiments-platform/backend-class-diagramm/models"
)

type ModelTraining struct {
	ID int64

	Model                *models.Model
	HyperparameterValues HyperparameterValues
	Dataset              *datasets.Dataset
	TrainTestSplit       float64
	Loss                 *Loss // TODO: подумать

	TrainingStatus enum.TrainingStatus
	StartedAt      time.Time
	FinishedAt     *time.Time
	TrainedModelID int64
	MetricsValues  MetricsValues
	Error          string
}

type Loss struct {
	ID          int64
	Name        string
	Description string
}

type HyperparameterValues = map[int64]any

type MetricsValues = map[int64]any

type ModelPrediction struct {
	ID int64

	TrainedModel *models.TrainedModel
	InputDataset *datasets.Dataset

	StartedAt          time.Time
	FinishedAt         *time.Time
	ResultDataset      *datasets.Dataset
	ResultUploadStatus enum.UploadStatus
	Error              string
}
