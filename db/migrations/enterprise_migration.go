// +build enterprise

package migrations

import (
	"github.com/pantyukhovp/qor-example/config/admin"
)

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
