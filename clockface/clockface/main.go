package main

import (
	"os"
	"time"

	"github.com/jaymorelli96/learn-go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
