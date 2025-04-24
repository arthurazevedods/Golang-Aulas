package main

import (
	"fmt"
	"time"
)

func main() {
	data := time.Now()
	fmt.Println(data.Format(("02/Jan/2006 15:04:05")))
}
