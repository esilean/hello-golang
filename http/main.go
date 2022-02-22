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
	if err != nil {
		fmt.Println("Error reaching google:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	io.Copy(logWriter{}, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
