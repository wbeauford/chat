package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)


type msg struct {
        Message string
        RoomId int
        User string
}


type ChatRoom struct {
        Name string
        Id int
        Users []string
}

type room struct {
        Name string
        Username string
}

var rooms []ChatRoom



func renderIndex(c *gin.Context) {
	render(c, gin.H{}, "index.html")
}


func getRooms(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        json.NewEncoder(c.Writer).Encode(rooms)
}

func createRoom(c *gin.Context) {
	r := c.Request
	roomName := r.FormValue("roomName")
	userName := r.FormValue("userName")

	var cR ChatRoom

	cR.Id = rand.Int()
	cR.Name = roomName
	cR.Users = append(cR.Users, userName)
	log.Println(cR)
	rooms = append(rooms, cR)
	jumpRoom(c, cR.Id, cR.Name, userName)

}

func jumpRoom(c *gin.Context, roomNumber int, roomName string, userName string) {
	render(c, gin.H{"roomNumber": roomNumber, "roomName": roomName, "username": userName}, "room.html")

}

func joinRoom(c *gin.Context) {
	r := c.Request
	roomNumber, err := strconv.Atoi(r.FormValue("roomID"))
	if err != nil {
		fmt.Println(err)
	}
	userName := r.FormValue("userName")
	roomName := r.FormValue("roomName")
	jumpRoom(c, roomNumber, roomName, userName)

}


func wsHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer
        if r.Header.Get("Origin") != "http://"+r.Host {
                http.Error(w, "Origin not allowed", 403)
                return
        }
        conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
        if err != nil {
                http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
                return
        }

        go echo(conn, w)
}

func echo(conn *websocket.Conn, w http.ResponseWriter) {
        for {
                m := msg{}
                err := conn.ReadJSON(&m)
                if err != nil {
                        fmt.Println("Error reading json.", err)
                }

                if err = conn.WriteJSON(m); err != nil {
                        fmt.Println(err)
                }


        }
}




