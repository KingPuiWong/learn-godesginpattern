package single

import (
	"fmt"
	"sync"
)

/**
另一个例子
1. init 函数，init函数仅在每个文件里调用一次，所以可以确定只会创建一个实例
2. sync.Once 仅会执行一次操作，如下面的例子
*/

var once sync.Once

type singleAnotherExample struct {
}

var singleInstanceAnotherExample *singleAnotherExample

func GetInstanceAnotherExample() *singleAnotherExample {
	if singleInstanceAnotherExample == nil {
		once.Do(
			func() {
				fmt.Println(`Creating single instance now.`)
				singleInstanceAnotherExample = &singleAnotherExample{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstanceAnotherExample
}
