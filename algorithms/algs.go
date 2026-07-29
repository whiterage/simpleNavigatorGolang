package algorithms

import (
	"errors"
)

var (
	ErrEmptyGraph        = errors.New("graph is empty")
	ErrVertex            = errors.New("invalid vertex")
	ErrDisconnectedGraph = errors.New("graph is disconnected")
)

type GraphAlgs struct{}

func NewGraphAlgs() *GraphAlgs {
	return &GraphAlgs{}
}
