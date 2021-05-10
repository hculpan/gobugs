package controllers

import (
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/game"
	"github.com/hculpan/gobugs/app/model"
	"github.com/hculpan/gobugs/app/pages"
	"github.com/veandco/go-sdl2/sdl"
)

type BugsController struct {
	game.GameController
}

var Controller *BugsController

func NewBugsController(windowWidth, windowHeight int32) BugsController {
	result := BugsController{}

	windowBackground := sdl.Color{R: 0, G: 0, B: 0, A: 0}

	result.Game = model.NewBugsGame(int((windowWidth-6)/3), int((windowHeight-46)/3), 40, 0.05, .5)
	result.Window = component.NewWindow(windowWidth, windowHeight, "GoBugs - an implementation of Palmiter's Protozoa", windowBackground)

	result.RegisterPages()

	Controller = &result

	return result
}

func (s *BugsController) RegisterPages() {
	component.RegisterPage(pages.NewGamePage("GamePage", 0, 0, s.Window.Width, s.Window.Height))
}
