package main

import (
	"log"
	"net/http"
	"time"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/server-sdk-go"
)

func getJoinToken(room, identity string) string {
	at := auth.NewAccessToken("devkey", "secret")
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room:     room,
	}
	at.AddGrant(grant).
		SetIdentity(identity).
		SetValidFor(time.Hour)

	token, _ := at.ToJWT()
	return token
}

func main() {
	http.HandleFunc("/getToken", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getJoinToken("my-room", "identity")))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
