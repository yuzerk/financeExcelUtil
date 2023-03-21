package syn

import (
	"my/syn/base"
	"sync"
	"time"
)

type BPriceHandler struct {
	base.HandleBaseData
	BName string
	BDesc string
}

func NewAHandleB(name string, bDesc string) *BPriceHandler {
	handle := &BPriceHandler{}
	handle.Business = B_BUSINESS
	handle.BusinessId = "id=2"
	handle.BName = name
	handle.BDesc = bDesc
	return handle
}

func (handle *BPriceHandler) Handle(group *sync.WaitGroup) {
	//do handle
	//get the result
	handle.HandleBaseData.PriceResult = 22.22
	time.Sleep(10 * time.Second)
	group.Done()
}

func (handle *BPriceHandler) GetResult() float64 {
	return handle.HandleBaseData.PriceResult
}
