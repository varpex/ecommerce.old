package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type TvController struct {
	*revel.Controller
}

type TV struct {
	gorm.Model
	Title       string      `gorm:"index:idx_tvs_title"`
	Biography   string      `gorm:"index:idx_tvs_biography"`
	Description string      `gorm:"index:idx_tvs_description"`
	Prices      []TvPrice   `gorm:"ForeignKey:TvId"`
	Features    []TvFeature `gorm:"ForeignKey:TvId"`
}

type TvPrice struct {
	gorm.Model
	Value uint `gorm:"index:idx_tv_prices_value"`
	TvId  uint `gorm:"index:idx_tv_prices_tv_id"`
}

type TvBaseFeature struct {
	gorm.Model
	Title string `gorm:"index:idx_tv_base_features_title"`
}

type TvFeature struct {
	gorm.Model
	Feature     TvBaseFeature `gorm:"foreignkey:ParentRefer"`
	ParentRefer uint          `gorm:"index:idx_tv_features_parent_refer"`
	TvId        uint          `gorm:"index:idx_tv_features_tv_id"`
	Value       string        `gorm:"index:idx_tv_features_value"`
}

var (
	tvConnectionString = "host=server.saeid.us user=postgres password=1369s1r3d691369 dbname=tv sslmode=disable"
)

func initialMigration(c *revel.Controller) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	db.AutoMigrate(&TV{}, &TvPrice{}, &TvBaseFeature{}, &TvFeature{})

	return nil
}

func init() {
	revel.InterceptFunc(initialMigration, revel.BEFORE, &TvController{})
}

func (c TvController) List(offset int, limit int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instances []TV
	var total int

	findErr := db.Preload("Prices").Find(&instances).Count(&total).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	data := make(map[string]interface{})
	data["results"] = instances
	data["count"] = total
	// return c.RenderJSON(data)
	return c.Render()
}

func (c TvController) Retrieve(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	if err != nil {
		return c.RenderError(err)
	}
	defer db.Close()

	var instance TV
	findErr := db.Where("id = ?", id).Preload("Prices").Preload("Features").Find(&instance).Error
	if findErr != nil {
		return c.RenderError(findErr)
	}

	return c.RenderJSON(instance)
}

func (c TvController) Post(Title string, Biography string, Description string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	errorProcess(c, err)
	defer db.Close()

	instance := &TV{
		Title:       Title,
		Biography:   Biography,
		Description: Description,
	}

	db.Create(instance).Scan(&instance)

	return c.RenderJSON(instance)
}

func (c TvController) Patch(id int, Title string, Biography string, Description string) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	errorProcess(c, err)
	defer db.Close()

	var instance TV
	findErr := db.Where("id = ?", id).Find(&instance).Error
	errorProcess(c, findErr)

	if _, ok := c.Params.Values["Title"]; ok {
		instance.Title = Title
	}
	if _, ok := c.Params.Values["Biography"]; ok {
		instance.Biography = Biography
	}
	if _, ok := c.Params.Values["Description"]; ok {
		instance.Description = Description
	}

	saveErr := db.Save(&instance).Error
	errorProcess(c, saveErr)

	return c.RenderJSON(instance)
}

func (c TvController) Delete(id int) revel.Result {
	db, err := gorm.Open("postgres", tvConnectionString)
	errorProcess(c, err)
	defer db.Close()

	var instance TV
	db.Where("id = ?", id).Find(&instance)

	db.Delete(&instance)

	data := make(map[string]interface{})
	data["message"] = fmt.Sprintf("%d Deleted.", id)
	return c.RenderJSON(data)
}

func errorProcess(c TvController, err error) revel.Result {
	if err != nil {
		return c.RenderError(err)
	}
	return nil
}
