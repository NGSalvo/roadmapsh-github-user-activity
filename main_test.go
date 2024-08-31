package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// asserts := assert.New(t)

}

/*type githubServiceSucessMock struct{}

func (g *githubServiceSucessMock) GetRecentActivity(userName string) ([]*Event, error) {
	return []*Event{
		{
			ID:   "1",
			Type: "PushEvent",
			Actor: Actor{
				ID:           1,
				Login:        "kamranahmedse",
				DisplayLogin: "kamranahmedse",
				GravatarID:   "",
				URL:          "",
				AvatarURL:    "",
			},
			Repo: Repo{
				ID:   1,
				Name: "developer-roadmap",
				URL:  "",
			},
			Payload: Payload{
				Action: "pushed",
			},
			Public:    true,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:   "2",
			Type: "PullRequestEvent",
			Actor: Actor{
				ID:           1,
				Login:        "kamranahmedse",
				DisplayLogin: "kamranahmedse",
				GravatarID:   "",
				URL:          "",
				AvatarURL:    "",
			},
			Repo: Repo{
				ID:   1,
				Name: "developer-roadmap",
				URL:  "",
			},
			Payload: Payload{
				Action: "pushed",
			},
			Public:    true,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}
*/
