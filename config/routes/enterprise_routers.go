// +build enterprise

package routes

import "github.com/pantyukhovp/qor-example/config/admin"

func init() {
	Router()
	WildcardRouter.AddHandler(admin.MicroSite)
}
