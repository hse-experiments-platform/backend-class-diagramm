package models

import (
	"github.com/hse-experiments-platform/backend-class-diagramm/datasets"
)

type Model struct {
	ID             int64
	Name           string
	Description    string
	Problem        *Problem
	Hyperarameters []Hyperparameter
}

type Hyperparameter struct {
	ID           int64
	Name         string
	Type         string
	DefaultValue any
}

type Problem struct {
	ID          int64
	Name        string
	Description string
	Metrics     []Metric
}

type Metric struct {
	ID          int64
	Name        string
	Description string
}

type TrainedModel struct {
	ID   int64
	Name string

	Schema       *datasets.Schema
	TagretColumn string
	Model        *Model
	Weights      Weights // TODO: параметры кроме весов и гиперпараметров?
}

type Weights = []any // TODO: подумать
