package components

import (
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/gobugs/app/model"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type BugsComponent struct {
	component.BaseComponent
}

func NewBugsComponent(x, y, width, height int32) *BugsComponent {
	result := &BugsComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)

	return result
}

func (c *BugsComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(0, 128, 0, 255)
	for x := 0; x < model.Bugs.BoardWidth; x++ {
		for y := 0; y < model.Bugs.BoardHeight; y++ {
			if model.Bugs.Board[x][y] {
				r.DrawRect(&sdl.Rect{
					X: int32(x*3) + c.X + 1,
					Y: int32(y*3) + c.Y + 1,
					W: 2,
					H: 2,
				})
			}
		}
	}

	for _, bug := range model.Bugs.Bugs {
		x, y := int32(bug.X-1)*3+c.X+1, int32(bug.Y-1)*3+c.Y+1
		switch bug.Class {
		case 0:
			gfx.RoundedBoxRGBA(r, x, y, x+8, y+8, 3, 255, 0, 0, 255)
		case 1:
			gfx.RoundedBoxRGBA(r, x, y, x+8, y+8, 3, 0, 0, 255, 255)
		case 2:
			gfx.RoundedBoxRGBA(r, x, y, x+8, y+8, 3, 100, 100, 255, 255)
		case 3:
			gfx.RoundedBoxRGBA(r, x, y, x+8, y+8, 3, 255, 255, 255, 255)
		}
	}

	return nil
}

func (c *BugsComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
