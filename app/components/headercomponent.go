package components

import (
	"fmt"

	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/gobugs/app/model"
	"github.com/veandco/go-sdl2/sdl"
)

type HeaderComponent struct {
	component.BaseComponent
}

func NewHeaderComponent(x, y, width, height int32) *HeaderComponent {
	result := &HeaderComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)

	result.AddChild(component.NewLabelComponent(x+5, y+5, 170, 30, 24, func() string {
		return fmt.Sprintf("Time : %d", model.Bugs.Cycle)
	}))
	result.AddChild(component.NewLabelComponent(x+200, y+5, 170, 30, 24, func() string {
		return fmt.Sprintf("Bugs : %d", len(model.Bugs.Bugs))
	}))
	result.AddChild(component.NewLabelComponent(x+375, y+5, 170, 30, 24, func() string {
		return fmt.Sprintf("Bacteria : %d", model.Bugs.BacteriaCount)
	}))
	result.AddChild(component.NewButtonComponent(width-320, y+5, 100, 30,
		"Pause",
		sdl.Color{R: 25, G: 25, B: 25, A: 255},
		sdl.Color{R: 50, G: 255, B: 50, A: 255},
		sdl.Color{R: 0, G: 0, B: 0, A: 255},
		func() {
			model.Bugs.Stop()
		}))
	result.AddChild(component.NewButtonComponent(width-215, y+5, 100, 30,
		"Reset",
		sdl.Color{R: 25, G: 25, B: 25, A: 255},
		sdl.Color{R: 50, G: 255, B: 50, A: 255},
		sdl.Color{R: 0, G: 0, B: 0, A: 255},
		func() {
			model.Bugs.Reset()
		}))
	result.AddChild(component.NewButtonComponent(width-110, y+5, 100, 30,
		"Start",
		sdl.Color{R: 25, G: 25, B: 25, A: 255},
		sdl.Color{R: 50, G: 255, B: 50, A: 255},
		sdl.Color{R: 0, G: 0, B: 0, A: 255},
		func() {
			model.Bugs.Start()
		}))

	return result
}

func (c *HeaderComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(25, 25, 25, 255)
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}
	r.FillRect(&rect)

	//	r.SetDrawColor(60, 60, 60, 0)
	//	r.DrawLine(c.X, c.Height, c.Width, c.Height)

	return nil
}

func (c *HeaderComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
