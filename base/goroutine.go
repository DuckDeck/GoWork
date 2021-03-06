package base

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	for val := range ch {
		if val == "EOF" {
			break
		}
		fmt.Printf("Input is %s\n", val)
	}
}

func TestChannel() {
	ch := make(chan string)
	go printInput(ch)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Printf("End the game!")
			break
		}
	}
	defer close(ch)
}
