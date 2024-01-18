package launcher

import "context"

type LauncherService interface {
	StartTraining(ctx context.Context, mt *ModelTraining) (int64, error)
	GetFullTrainingInfo(ctx context.Context, id int64) (*ModelTraining, error)

	StartPrediction(ctx context.Context, mp *ModelPrediction) (int64, error)
	GetFullPredictionInfo(ctx context.Context, id int64) (*ModelPrediction, error)
}
