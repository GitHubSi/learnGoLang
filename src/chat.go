package main

//every one online person need to push a message
type Room struct {
	OnlineUser map[string]*OnlineUser
	Broadcast  chan Message
	CloseSign  chan bool
}

type Message struct {
	MType       string
	TestMessage string
	UserStatus  string
}

type OnlineUser struct {
	InRoom     *Room
	Connection *websocket.Conn
	UserInfo   *User
	Send       chan Message
}

func (this *Room) run() {
	for {
		select {
		case b := <-this.Broadcast:
			for _, online := range this.OnlineUser {
				online.Send <- b
			}
		case c := <-this.CloseSign:
			if c == true {
				close(this.Broadcast)
				close(this.CloseSign)
				return
			}
		}
	}
}

func (this *OnlineUser) PullFromClient() {
	for {
		var content string
		err := websocket.Message.Receive(this.Connection, &content)
		if err != nil {
			return
		}

		m := Message{
			MType:TEXT_MTYPE,
			TextMessage{
				Time:humanCreateAt(),
				Content:content,
			},

		}
	}
}

func (this *OnlineUser) PushToClient() {
	for b := range this.Send {
		err := websocket.Json.Send(this.Connection, b)
		if err != nil {
			break
		}
	}
}