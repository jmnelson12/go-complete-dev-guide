package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	checkErr(err)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// custom Write implementation
func (logWriter) Write(bs []byte) (n int, er error) {
	fmt.Println(string(bs))
	fmt.Println("Num of bytes written:", len(bs))
	return len(bs), nil
}
