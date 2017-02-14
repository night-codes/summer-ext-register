package main

import (
	"encoding/json"
	"gopkg.in/night-codes/summer.v1"
)

func toSummerUser(user *UserStruct) *summer.UsersStruct {
	ret := user.UsersStruct
	settings, _ := json.Marshal(user.Settings)
	json.Unmarshal(settings, &ret.Settings)
	return &ret
}
func fromSummerUser(summerUser *summer.UsersStruct) *UserStruct {
	user := &UserStruct{}
	user.UsersStruct = *summerUser
	settings, _ := json.Marshal(summerUser.Settings)
	json.Unmarshal(settings, &user.Settings)
	return user
}
func jsoner(object interface{}) string {
	j, _ := json.MarshalIndent(object, "", "\t")
	return string(j)
}
