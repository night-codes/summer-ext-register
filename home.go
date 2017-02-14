package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/night-codes/summer.v1"
)

type (
	homeModule struct {
		summer.Module
	}
	Settings struct {
		Gender string `form:"gender" json:"gender" bson:"gender"`
		City   string `form:"city" json:"city" bson:"city"`
	}
	UserStruct struct {
		summer.UsersStruct
		Settings `bson:",inline"`
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
	user := &UserStruct{}
	c.Bind(user)
	user.Login = user.Email
	user.Rights = summer.Rights{Groups: []string{"all"}}

	// user.Disabled = true // если нужно, посылаем на модерацию !!!

	if err := panel.Users.Add(*toSummerUser(user)); err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, obj{"data": user})
}
