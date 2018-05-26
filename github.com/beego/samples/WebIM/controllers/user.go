package controllers

import (
	"encoding/json"
)

type UserController struct {
	baseController
}

type CidName struct {
	ClientId string
	UserName string
}

//get all online users

func (this *UserController) GetOnlineUsers() {
	var cidNameSlice []CidName
	for _, v := range userMap.relationMap {
		cidNameSlice = append(cidNameSlice, CidName{v.ClientId, v.Name})
	}
	JsonUser, _ := json.Marshal(cidNameSlice)
	this.Ctx.WriteString(string(JsonUser))
}
