package main

import (
	"fmt"
	"godesignpattern/behavior_patterns/iterator"
)

// 迭代器是一种行为设计模式，让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素

func main() {
	user1 := &iterator.User{
		Name: "a",
		Age:  20,
	}

	user2 := &iterator.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{user1, user2},
	}

	createIterator := userCollection.CreateIterator()

	for createIterator.HasNext() {
		user := createIterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
