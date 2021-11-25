package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	res, err := http.Get("https://wttr.in?view=3")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	bodyStr := string(body)

	// TODO: handle arry with 0 size
	wtr := strings.TrimSpace(strings.Split(bodyStr, ":")[1])
	fmt.Println(wtr)
}
