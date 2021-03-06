package migrations

import (
	"github.com/qor/activity"
	"github.com/qor/help"
	"github.com/qor/media/asset_manager"
	"github.com/pantyukhovp/qor-example/app/models"
	"github.com/pantyukhovp/qor-example/config/admin"
	"github.com/pantyukhovp/qor-example/config/seo"
	"github.com/pantyukhovp/qor-example/db"
	"github.com/qor/transition"
)

func init() {
	AutoMigrate(&asset_manager.AssetManager{})

	AutoMigrate(&models.Product{}, &models.ProductVariation{}, &models.ProductImage{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Material{}, &models.Category{}, &models.Collection{})

	AutoMigrate(&models.Address{})

	AutoMigrate(&models.Order{}, &models.OrderItem{})

	AutoMigrate(&models.Store{})

	AutoMigrate(&models.Setting{})

	AutoMigrate(&models.User{})

	AutoMigrate(&transition.StateChangeLog{})

	AutoMigrate(&activity.QorActivity{})

	AutoMigrate(&admin.QorWidgetSetting{})

	AutoMigrate(&models.Page{})

	AutoMigrate(&seo.MySEOSetting{})

	AutoMigrate(&models.MediaLibrary{})

	AutoMigrate(&models.Article{})

	AutoMigrate(&help.QorHelpEntry{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
