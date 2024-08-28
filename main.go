package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Event struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

type Actor struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Payload struct {
	Action string `json:"action"`
}

type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {

}

func GetRecentActivity(userName string) ([]*Event, error) {
	response, err := http.Get("https://api.github.com/users/" + userName + "/events")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var activities []*Event

	err = json.NewDecoder(response.Body).Decode(&activities)

	if err != nil {
		return nil, err
	}

	return activities, nil

}
