package main

import "net/http"

func main() {

}

func GetRecentActivity(userName string) (any, error) {
	response, err := http.Get("https://api.github.com/users/" + userName + "/events")

	if err != nil {
		return nil, err
	}

	return response, nil
}
