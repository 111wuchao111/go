// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/beego/samples/WebIM/models"
	"github.com/gorilla/websocket"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

//客户端发来的消息体(注意：字段首字母必须大写，不然json解析不出来)
type ClientMsg struct {
	ClientId string
	Content  string
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	uname := this.GetString("uname")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.TplName = "websocket.html"
	this.Data["IsWebSocket"] = true
	this.Data["UserName"] = uname
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	clientId := ""
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	//为每个新用户生成clinetId
	uniqueId := models.UniqueId()
	// Join chat room.
	Join(uniqueId, uname, ws)

	defer Leave(uniqueId)
	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		msg := &ClientMsg{}
		json.Unmarshal(p, msg)
		clientId = msg.ClientId
		publish <- newEvent(clientId, models.EVENT_MESSAGE, uname, msg.Content)
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func sendMsgWithWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("errFail to marshal event:", err)
		return
	}
	for k, v := range userMap.relationMap {
		ws := v.Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- k
			}
		}
	}

}
