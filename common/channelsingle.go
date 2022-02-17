package common

import "sync"

type singleton chan [2]string

var channel *singleton
var once sync.Once

//var Size int32 = 12

func GetChannel(size int) *singleton {
	once.Do(func() {
		// 容量视库存大小而定，这里设为size
		ret := make(singleton, 100000)
		channel = &ret
	})
	return channel
}
