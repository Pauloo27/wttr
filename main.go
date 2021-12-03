package main

import (
	"errors"
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

	bodyParts := strings.Split(bodyStr, ":")
	if len(bodyParts) != 2 {
		panic(errors.New("invalid body"))
	}

	wtr := strings.TrimSpace(strings.Split(bodyStr, ":")[1])
	fmt.Println(wtr)
}
