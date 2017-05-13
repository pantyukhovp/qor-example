package i18n

import (
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
	"github.com/pantyukhovp/qor-example/config"
	"github.com/pantyukhovp/qor-example/db"
	"path/filepath"
)

var I18n *i18n.I18n

func init() {
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
}
