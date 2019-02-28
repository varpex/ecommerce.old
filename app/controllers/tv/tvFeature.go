package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type TvFeatureController struct {
	*revel.Controller
}

func (c TvFeatureController) List() revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instances []TvFeature
	var total int

	findErr := db.Preload("Feature").Find(&instances).Count(&total).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	data := make(map[string]interface{})
	data["results"] = instances
	data["count"] = total
	return c.RenderJSON(data)
}

func (c TvFeatureController) Retrieve(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvFeature
	findErr := db.Where("id = ?", id).Preload("Feature").Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	return c.RenderJSON(instance)
}

func (c TvFeatureController) Post(ParentRefer uint, TvId uint, Value string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	instance := &TvFeature{
		ParentRefer: ParentRefer,
		TvId:        TvId,
		Value:       Value,
	}

	db.Create(instance).Scan(&instance)

	return c.RenderJSON(instance)
}

func (c TvFeatureController) Patch(id int, ParentRefer uint, TvId uint, Value string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvFeature
	findErr := db.Where("id = ?", id).Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	if _, ok := c.Params.Values["ParentRefer"]; ok {
		instance.ParentRefer = ParentRefer
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

func (c TvFeatureController) Delete(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvFeature
	db.Where("id = ?", id).Find(&instance)

	db.Delete(&instance)

	data := make(map[string]interface{})
	data["message"] = fmt.Sprintf("%d Deleted.", id)
	return c.RenderJSON(data)
}
