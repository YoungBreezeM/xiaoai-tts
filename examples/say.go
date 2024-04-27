package main

import (
	"bufio"
	"fmt"
	"os"

	xiaoaitts "github.com/YoungBreezeM/xiaoai-tts"
)

func main() {
	m := &xiaoaitts.MiAccount{
		User: "xxxx",
		Pwd:  "xxxxx",
	}
	x := xiaoaitts.NewXiaoAi(m)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please Input Text")
		text, _ := input.ReadString('\n')
		x.Say(text)
	}
}
