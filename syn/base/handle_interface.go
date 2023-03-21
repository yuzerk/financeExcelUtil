package base

import "sync"

type PriceHandle interface {
	Handle(group *sync.WaitGroup)
	GetResult() float64
}
