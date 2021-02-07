package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type message string

func (m message) isValid() bool {
	isNumber := false
	for _, r := range m {
		if '0' <= r && r <= '9' {
			isNumber = false
			continue
		}
		return true
	}
	return isNumber
}
func (m message) toUpper() string {
	return strings.ToUpper(string(m))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputChan := make(chan message)
	outputChan := make(chan string)
	errChan := make(chan string)
	for scanner.Scan() {
		// `Text` は、入力から現在のトークン、
		// ここでは次の行、を返します。
		go convertText(inputChan, outputChan, errChan)
		inputChan <- message(scanner.Text())
		select {
		case msg1 := <-outputChan:
			fmt.Println("received", msg1)
		case msg2 := <-errChan:
			fmt.Println("received", msg2)
		}
	}

	// `Scan` 中にエラーがなかったかを確認します。
	// EOF (ファイルの末尾) が期待され、その場合は
	// `Scan` にエラーとして報告されません。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func convertText(inputChan chan message, outputChan, errorChan chan string) {
	text := <-inputChan
	if !text.isValid() {
		errorChan <- "error: >>> invalid format"
		return
	}
	outputChan <- "output: >>> " + text.toUpper()
}
