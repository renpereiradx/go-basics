package mypackage

import (
	"sync"
	"time"
)

func Say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	println(message)
}

func SaySomething(text string, c chan<- string) {
	time.Sleep(2 * time.Second)
	c <- text
}
