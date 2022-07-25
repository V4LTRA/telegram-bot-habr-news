package main

import (
	"TG/news-bot/actions"
	"fmt"
	"time"
)

func main() {
	for {
		habrTags := []string{"go", "python", "kubernetes"}
		for _, i := range habrTags {
			fmt.Println(i)
			actions.HabrGo(i)
			time.Sleep(time.Second * 5)
		}

	}
}
