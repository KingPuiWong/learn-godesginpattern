package strategy

import "fmt"

// 具体策略

type Fifo struct {
}

func (f Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}
