package main

import (
	"embed"
	"fmt"

	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/hculpan/gobugs/app/controllers"
)

func main() {
	component.SetupSDL()

	if err := resources.FontsInit(embed.FS{}); err != nil {
		fmt.Println(err)
		return
	}

	// Since our cells are all 3 pixels with a 1 pixel barrier
	// around them, we want to make sure our widht/height is
	// a divisor of 4
	var gameWidth int32 = 906
	var gameHeight int32 = 946

	gamecontroller := controllers.NewBugsController(gameWidth, gameHeight)
	if err := gamecontroller.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
