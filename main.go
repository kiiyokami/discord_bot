package main

import (
	"discord_bot/services"
	"fmt"
)

func main() {
	testUrl := services.AutoEmbed("https://www.instagram.com/reel/DMXv3zVhgoG/")
	fmt.Println(testUrl)
}
