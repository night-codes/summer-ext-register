package main

import (
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/night-codes/summer.v1"
)

// get user POST-data with extra fields
func getUserPost(c *gin.Context) (ret summer.UsersStruct) {
	user := userStruct{}
	c.Bind(&user)
	user.Login = user.Email
	user.Rights = summer.Rights{Groups: []string{"all"}}
	ret = user.UsersStruct
	settings, _ := json.Marshal(user.Settings)
	json.Unmarshal(settings, &ret.Settings)
	return
}

// get user info by ID from summer with extra fields
func getUser(login string) *userStruct {
	user := &userStruct{}
	panelUser, _ := panel.Users.GetByLogin(login)
	user.UsersStruct = *panelUser
	settings, _ := json.Marshal(panelUser.Settings)
	json.Unmarshal(settings, &user.Settings)
	return user
}
