package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.OpenFile("D-test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	f.WriteString(strconv.Itoa(n) + "\n")
	for i := 0; i < n; i++ {
		cmd := make([]string, 0)
		cmd = append(cmd, strconv.Itoa(rand.Intn(2)))
		if cmd[0] == "0" {
			cmd = append(cmd, strconv.Itoa(rand.Intn(100)))
		}
		// write cmd in file
		for _, v := range cmd {
			f.WriteString(v + " ")
		}
		f.WriteString("\n")
	}
}
