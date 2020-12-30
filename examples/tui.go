package main

import (
	"log"

	cw "github.com/baileywickham/gocuiw"
	c "github.com/jroimartin/gocui"
)

func keybindings(g *c.Gui) error {
	if err := g.SetKeybinding("main", c.KeyArrowDown, c.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", c.KeyArrowUp, c.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("", c.KeyCtrlC, c.ModNone,
		func(g *c.Gui, v *c.View) error { return c.ErrQuit }); err != nil {
		return err
	}

	if err := g.SetKeybinding("msg", c.KeyEnter, c.ModNone, delMsg); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", c.KeyEnter, c.ModNone, getLine); err != nil {
		return err
	}

	return nil
}

var PROMPT_VIEW = " Prompt View "

func layout(g *c.Gui) error {
	//if _, err := CreateLeftHalfView(g, "hi", "2"); err != nil {
	//	return err
	//}
	if _, err := cw.CreateTopRightQuarterView(g, "tr", "3"); err != nil {
		return err
	}
	if _, err := cw.CreateBottomRightQuarterView(g, "br", "4"); err != nil {
		return err
	}
	if _, err := cw.CreateBottomLeftQuarterView(g, "bl", "3"); err != nil {
		return err
	}
	if _, err := cw.CreateTopLeftQuarterView(g, "tl", "4"); err != nil {
		return err
	}

	return nil
	//maxX, maxY := g.Size()
	//if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
	//	if err != c.ErrUnknownView {
	//		return err
	//	}
	//	// highlight items
	//	v.Frame = true
	//	v.Title = " main "
	//	v.Highlight = true
	//	v.SelBgColor = c.ColorGreen
	//	v.SelFgColor = c.ColorBlack

	//	fmt.Fprintln(v, "Item 1")
	//	fmt.Fprintln(v, "Item 2")
	//	fmt.Fprintln(v, "Item 3")
	//	fmt.Fprint(v, "\rWill be")
	//	fmt.Fprint(v, "deleted\rItem 4\nItem 5")

	//	if _, err := g.SetCurrentView("main"); err != nil {
	//		return err
	//	}

	//}

	//return nil
}

func main() {
	g, err := c.NewGui(c.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != c.ErrQuit {
		log.Panicln(err)
	}
}
