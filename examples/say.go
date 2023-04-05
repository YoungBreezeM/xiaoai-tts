package main

import (
	"bufio"
	"fmt"
	"os"

	xiaoaitts "github.com/qfyang-cn/xiaoai-tts"
)

func main() {
	m := &xiaoaitts.MiAccount{
		User: "username",
		Pwd:  "password",
	}
	x := xiaoaitts.NewXiaoAi(m)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please Input Text")
		text, _ := input.ReadString('\n')
		x.Say(text)
	}
}
