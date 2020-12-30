package main

import (
	"log"

	cw "github.com/baileywickham/gocuiw"
	c "github.com/jroimartin/gocui"
)

func layout(g *c.Gui) error {
	//if _, err := CreateLeftHalfView(g, "hi", "2"); err != nil {
	//	return err
	//}
	if _, err := cw.CreateTopRightQuarterView(g, "1"); err != nil {
		return err
	}
	if _, err := cw.CreateBottomRightQuarterView(g, "2"); err != nil {
		return err
	}
	if _, err := cw.CreateBottomLeftQuarterView(g, "3"); err != nil {
		return err
	}
	if _, err := cw.CreateTopLeftQuarterView(g, "4"); err != nil {
		return err
	}

	return nil
}

func main() {
	g, err := c.NewGui(c.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(layout)

	// cw.SetKeybindings must be called in addition to your local keybindings
	if err := cw.SetKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != c.ErrQuit {
		log.Panicln(err)
	}
}
