package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Block struct {
	Id int
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

func BlocksPage(c echo.Context) error {

	blocks := []Block{}

	startStr := c.QueryParam("start")
	start, err := strconv.Atoi(startStr)

	if err != nil {
		start = 0
	}

	for i := start; i < start+10; i++ {
		blocks = append(blocks, Block{Id: i})
	}

	template := "blocks"

	if start == 0 {
		template = "blocks-index"
	}

	fmt.Println(start, blocks, "test")

	return c.Render(
		http.StatusOK,
		template,
		Blocks{
			Start:  start,
			Next:   start + 10,
			More:   start+10 < 100,
			Blocks: blocks,
		})
}
