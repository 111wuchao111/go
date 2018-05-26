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
	"container/list"
	"time"

	"sync"

	"github.com/astaxie/beego"
	"github.com/beego/samples/WebIM/models"
	"github.com/gorilla/websocket"
)

type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

func newEvent(clientId string, ep models.EventType, user string, msg string) models.Event {
	return models.Event{ep, clientId, user, int(time.Now().Unix()), msg}
}

func Join(uniqueId, user string, ws *websocket.Conn) {
	subscribe <- Subscriber{ClientId: uniqueId, Name: user, Conn: ws}
}

func Leave(clientId string) {
	unsubscribe <- clientId
}

//clinetId与用户信息结构映射
type UserMap struct {
	UserMapLock sync.Mutex
	relationMap map[string]Subscriber
}

type Subscriber struct {
	ClientId string
	Name     string
	Conn     *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan string, 10)
	// Send events here to publish them.
	publish = make(chan models.Event, 10)
	// Long polling waiting list.
	waitingList = list.New()

	userMap = UserMap{relationMap: make(map[string]Subscriber)}
)

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			userMap.UserMapLock.Lock()
			userMap.relationMap[sub.ClientId] = sub
			userMap.UserMapLock.Unlock()
			// Publish a JOIN event.
			publish <- newEvent(sub.ClientId, models.EVENT_JOIN, sub.Name, "")
			beego.Info("New user:", sub.Name, "|clentId:", sub.ClientId, ";WebSocket:", sub.Conn != nil)
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			sendMsgWithWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case clientId := <-unsubscribe:
			beego.Info("clientId:", clientId, "exit")
			userMap.UserMapLock.Lock()
			delete(userMap.relationMap, clientId)
			userMap.UserMapLock.Unlock()
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
