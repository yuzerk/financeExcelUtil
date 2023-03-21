package syn

import (
	"fmt"
	"my/syn/base"
	"sync"
)

const (
	A_BUSINESS base.BusinessType = 1
	B_BUSINESS base.BusinessType = 2
)

var strategy map[base.BusinessType]base.PriceHandle

func init() {
	pa := NewAHandleA(make([]string, 0), "aaa")
	pb := NewAHandleB("name", "bbb")
	strategy = map[base.BusinessType]base.PriceHandle{
		A_BUSINESS: pa,
		B_BUSINESS: pb,
	}
}

func Exe() {
	group := new(sync.WaitGroup)
	group.Add(2)
	for _, v := range strategy {
		go v.Handle(group)
	}
	group.Wait()
	price := float64(0)
	for _, v := range strategy {
		price += v.GetResult()
	}
	fmt.Println(price)
}
