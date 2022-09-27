package main

import (
	"cadena/blocke/blocke"
	"fmt"
)

func main() {
	var blockes = blocke.NewBlockes()
	var b = blocke.NewBlocke()

	nblock := b.GenerateNewBlock(b, 53)
	n2block := b.GenerateNewBlock(nblock, 35)

	n3block := b.GenerateNewBlock(n2block, 40)
	n4block := b.GenerateNewBlock(n3block, 1053)

	blockes.AddBlock(*b, *nblock, *n2block, *n3block, *n4block)

	fmt.Println("=========")
	for _, v := range blockes.B {
		fmt.Println("index ", v.Index)
		fmt.Println("time  ", v.Time)
		fmt.Println("Bpm ", v.Bpm)
		fmt.Println("hash ", v.Hash)
		fmt.Println("prev ", v.PrevHash)
		fmt.Println("========")
	}
}
