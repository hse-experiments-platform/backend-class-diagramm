package models

import "context"

type ModelsService interface {
	GetModel(ctx context.Context, id int64) (*Model, error)
	GetShortModelsInfo(ctx context.Context) ([]*Model, error)

	GetTrainedModel(ctx context.Context, id int64) (*TrainedModel, error)
	GetShortTrainedModelsInfo(ctx context.Context) ([]*TrainedModel, error)
	GetTrainedModelsByModel(ctx context.Context, modelID int64) ([]*TrainedModel, error)
}
