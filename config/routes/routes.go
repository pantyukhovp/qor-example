package routes

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/pantyukhovp/qor-example/app/controllers"
	"github.com/pantyukhovp/qor-example/config"
	"github.com/pantyukhovp/qor-example/config/auth"
	"github.com/pantyukhovp/qor-example/db"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"
)

var rootMux *http.ServeMux
var WildcardRouter *wildcard_router.WildcardRouter

func Router() *http.ServeMux {
	if rootMux == nil {
		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			tx := db.DB
			context := &qor.Context{Request: ctx.Request, Writer: ctx.Writer}
			if locale := utils.GetLocale(context); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			ctx.Set("DB", publish2.PreviewByDB(tx, context))
		})

		gin.SetMode(gin.DebugMode)

		router.GET("/", controllers.HomeIndex)
		router.GET("/products/:code", controllers.ProductShow)
		router.GET("/category/:code", controllers.CategoryShow)
		router.GET("/switch_locale", controllers.SwitchLocale)

		cartGroup := router.Group("/cart")
		{
			cartGroup.GET("/", controllers.ShowCartHandler)
			cartGroup.GET("/checkout", controllers.CheckoutCartHandler)
			cartGroup.POST("/", controllers.AddToCartHandler)
			cartGroup.POST("/checkout", controllers.OrderCartHandler)
			cartGroup.DELETE("/:id", controllers.RemoveFromCartHandler)
		}

		rootMux = http.NewServeMux()
		rootMux.Handle("/auth/", auth.Auth.NewRouter())
		publicDir := http.Dir(filepath.Join(config.Root, "public"))
		rootMux.Handle("/dist/", utils.FileServer(publicDir))
		rootMux.Handle("/vendors/", utils.FileServer(publicDir))
		rootMux.Handle("/images/", utils.FileServer(publicDir))
		rootMux.Handle("/fonts/", utils.FileServer(publicDir))

		WildcardRouter = wildcard_router.New()
		WildcardRouter.MountTo("/", rootMux)
		WildcardRouter.AddHandler(router)
	}
	return rootMux
}
