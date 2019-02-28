package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Product struct {
	gorm.Model
	Test string
}

func (c App) Index() revel.Result {
	return c.Render()
}

// This function will be called when the Controller is put back into the stack
func (c App) Destroy() {
	c.Controller.Destroy()
	// Clean up locally defined maps or items
}
