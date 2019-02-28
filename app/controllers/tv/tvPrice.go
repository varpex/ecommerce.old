package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type TvPriceController struct {
	*revel.Controller
}

func (c TvPriceController) List() revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instances []TvPrice
	var total int

	findErr := db.Find(&instances).Count(&total).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	data := make(map[string]interface{})
	data["results"] = instances
	data["count"] = total
	return c.RenderJSON(data)
}

func (c TvPriceController) Retrieve(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvPrice
	findErr := db.Where("id = ?", id).Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	return c.RenderJSON(instance)
}

func (c TvPriceController) Post(TvId uint, Value uint) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	instance := &TvPrice{
		TvId:  TvId,
		Value: Value,
	}

	db.Create(instance).Scan(&instance)

	return c.RenderJSON(instance)
}

func (c TvPriceController) Patch(id int, TvId uint, Value uint) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvPrice
	findErr := db.Where("id = ?", id).Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	if _, ok := c.Params.Values["TvId"]; ok {
		instance.TvId = TvId
	}
	if _, ok := c.Params.Values["Value"]; ok {
		instance.Value = Value
	}

	saveErr := db.Save(&instance).Error
	if saveErr != nil {
		return c.RenderError(saveErr)
	}

	return c.RenderJSON(instance)
}

func (c TvPriceController) Delete(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvPrice
	db.Where("id = ?", id).Find(&instance)

	db.Delete(&instance)

	data := make(map[string]interface{})
	data["message"] = fmt.Sprintf("%d Deleted.", id)
	return c.RenderJSON(data)
}
