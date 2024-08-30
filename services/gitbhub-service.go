package services

import (
	"encoding/json"
	"fmt"
	"github-user-activity/models"
	"net/http"
)

type (
	GithubService interface {
		GetRecentActivity(userName string) ([]*models.Event, error)
	}

	githubService struct{}
)

func NewGithub() GithubService {
	return &githubService{}
}

func (g *githubService) GetRecentActivity(userName string) ([]*models.Event, error) {
	response, err := http.Get("https://api.github.com/users/" + userName + "/events")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var activities []*models.Event
	err = json.NewDecoder(response.Body).Decode(&activities)

	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (g *githubService) GetRecentActivityByType(userName string, eventType models.EventType) ([]*models.Event, error) {
	events, err := g.GetRecentActivity(userName)

	if err != nil {
		return nil, err
	}

	var filteredEvents []*models.Event
	for _, event := range events {
		if event.Type == eventType {
			filteredEvents = append(filteredEvents, event)
		}
	}

	return filteredEvents, nil
}
