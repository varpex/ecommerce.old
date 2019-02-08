package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// This function will be called when the Controller is put back into the stack
func (c App) Destroy() {
	c.Controller.Destroy()
	// Clean up locally defined maps or items
}
