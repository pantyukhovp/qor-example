package admin

import (
	"github.com/qor/filebox"
	"github.com/pantyukhovp/qor-example/config"
	"github.com/pantyukhovp/qor-example/config/auth"
	"github.com/qor/roles"
)

var Filebox *filebox.Filebox

func init() {
	Filebox = filebox.New(config.Root + "/public/downloads")
	Filebox.SetAuth(auth.AdminAuth{})
	dir := Filebox.AccessDir("/")
	dir.SetPermission(roles.Allow(roles.Read, "admin"))
}
