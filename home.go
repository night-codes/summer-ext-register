package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/night-codes/summer.v1"
)

type (
	homeModule struct {
		summer.Module
	}
)

var (
	_ = panel.AddModule(
		&summer.ModuleSettings{
			Name:           "home",
			DisableAuth:    true,
			OriginTemplate: true,
		},
		&homeModule{},
	)
)

func (m *homeModule) Register(c *gin.Context) {
	var user summer.UsersStruct
	c.Bind(&user)
	user.Login = user.Email
	// user.Disabled = true // if needs moderation
	if err := panel.Users.Add(user); err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, obj{"data": user})
}
