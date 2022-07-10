package main

import (
	"fmt"
	"gitlab.com/wshaman/hw-concurrency/lib/batch"
)

func main() {
	ul := batch.GetBatch(40, 4)

	for _, u := range ul {
		fmt.Println(u.ID)
	}
}
