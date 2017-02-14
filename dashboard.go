package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/night-codes/summer.v1"
)

type (
	DashboardModule struct {
		summer.Module
	}
)

var (
	dashboard = panel.AddModule(
		&summer.ModuleSettings{
			Name:   "dashboard",
			Title:  "Главная",
			Rights: summer.Rights{Groups: []string{"all"}},
		},
		&DashboardModule{},
	)
)

func (m *DashboardModule) Page(c *gin.Context) {
	summerUser := c.MustGet("user").(summer.UsersStruct)
	user := fromSummerUser(&summerUser)
	c.HTML(200, m.Settings.TemplateName+".html", obj{
		"summerUserJson": jsoner(summerUser),
		"userJson":       jsoner(user),
	})
}
