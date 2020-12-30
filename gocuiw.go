package gocuiw

import (
	"fmt"
	"log"

	c "github.com/jroimartin/gocui"
)

func cursorDown(g *c.Gui, v *c.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *c.Gui, v *c.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func delMsg(g *c.Gui, v *c.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("main"); err != nil {
		return err
	}
	return nil
}
func getLine(g *c.Gui, v *c.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != c.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		v.Editable = true
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func CreatePromptView(g *c.Gui, title string) error {
	var PROMPT_VIEW = " Prompt View "
	tw, th := g.Size()
	v, err := g.SetView(PROMPT_VIEW, tw/6, (th/2)-1, (tw*5)/6, (th/2)+1)
	if err != nil && err != c.ErrUnknownView {
		return err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true
	_, err = g.SetCurrentView(PROMPT_VIEW)

	return err
}

func CreateLeftHalfView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	v, err := g.SetView("lh", 0, 0, (tw/2)-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("lh")

	return v, err

}

func CreateRightHalfView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	v, err := g.SetView("rh", (tw/2)+1, 0, tw-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("rh")

	return v, err
}

func CreateTopRightQuarterView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	// x0, y0, x1, y1
	v, err := g.SetView("tr", (tw/2)+1, 0, tw-1, (th/2)-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("tr")

	return v, err
}

func CreateBottomRightQuarterView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	// x0, y0, x1, y1
	v, err := g.SetView("br", (tw/2)+1, (th/2)+1, tw-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("br")

	return v, err
}

func CreateTopLeftQuarterView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	// x0, y0, x1, y1
	v, err := g.SetView("tl", 0, 0, (tw/2)-1, (th/2)-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("tl")

	return v, err
}

func CreateBottomLeftQuarterView(g *c.Gui, title string) (*c.View, error) {
	tw, th := g.Size()
	// x0, y0, x1, y1
	v, err := g.SetView("bl", 0, (th/2)+1, (tw/2)-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		return nil, err
	}
	v.Editable = true
	v.Title = title
	g.Cursor = true

	v, err = g.SetCurrentView("bl")

	return v, err
}

func setTopWindowTitle(g *c.Gui, view_name, title string) {
	v, err := g.View(view_name)
	if err != nil {
		log.Println("Error on setTopWindowTitle", err)
		return
	}
	v.Title = fmt.Sprintf("%v (Ctrl-q to close)", title)
}

func keybindings(g *c.Gui) error {
	// Add some sane defaults, add window manager

	return nil
}
