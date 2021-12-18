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
		panic("error =(")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic("error =(")
	}
	bodyStr := string(body)

	bodyParts := strings.Split(bodyStr, ":")
	if len(bodyParts) != 2 {
		panic("error =(")
	}

	wtr := strings.TrimSpace(bodyParts[1])
	fmt.Println(wtr)
}
