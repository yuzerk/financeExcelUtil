package syn

import (
	"my/syn/base"
	"sync"
	"time"
)

type APriceHandler struct {
	base.HandleBaseData
	Flow  []string
	ADesc string
}

func NewAHandleA(flow []string, aDesc string) *APriceHandler {
	handle := &APriceHandler{}
	handle.Business = A_BUSINESS
	handle.BusinessId = "id=1"
	handle.Flow = flow
	handle.ADesc = aDesc
	return handle
}

func (handle *APriceHandler) Handle(group *sync.WaitGroup) {
	//do handle
	//get the result
	handle.HandleBaseData.PriceResult = 11.11
	time.Sleep(5 * time.Second)
	group.Done()
}

func (handle *APriceHandler) GetResult() float64 {
	return handle.HandleBaseData.PriceResult
}
