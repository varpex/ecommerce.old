package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type TvBaseFeatureController struct {
	*revel.Controller
}

func (c TvBaseFeatureController) List() revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instances []TvBaseFeature
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

func (c TvBaseFeatureController) Retrieve(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvBaseFeature
	findErr := db.Where("id = ?", id).Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	return c.RenderJSON(instance)
}

func (c TvBaseFeatureController) Post(Title string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	instance := &TvBaseFeature{
		Title: Title,
	}

	db.Create(instance).Scan(&instance)

	return c.RenderJSON(instance)
}

func (c TvBaseFeatureController) Patch(id int, Title string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvBaseFeature
	findErr := db.Where("id = ?", id).Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	if _, ok := c.Params.Values["Title"]; ok {
		instance.Title = Title
	}

	saveErr := db.Save(&instance).Error
	if saveErr != nil {
		return c.RenderError(saveErr)
	}

	return c.RenderJSON(instance)
}

func (c TvBaseFeatureController) Delete(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TvBaseFeature
	db.Where("id = ?", id).Find(&instance)

	db.Delete(&instance)

	data := make(map[string]interface{})
	data["message"] = fmt.Sprintf("%d Deleted.", id)
	return c.RenderJSON(data)
}

func tvBaseFeaturetvBaseFeatureErrorProcess(c TvBaseFeatureController, err error) revel.Result {
	if err != nil {
		return c.RenderError(err)
	}
	return nil
}
