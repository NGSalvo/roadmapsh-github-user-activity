package models

import "time"

type (
	Event struct {
		ID        string    `json:"id"`
		Type      EventType `json:"type"`
		Actor     Actor     `json:"actor"`
		Repo      Repo      `json:"repo"`
		Payload   Payload   `json:"payload"`
		Public    bool      `json:"public"`
		CreatedAt time.Time `json:"created_at"`
	}

	Actor struct {
		ID           int64  `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	}

	Payload struct {
		Action string `json:"action"`
	}

	Repo struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	EventType string
)

const (
	CommitCommentEvent            = EventType("CommitCommentEvent")
	CreateEvent                   = EventType("CreateEvent")
	DeleteEvent                   = EventType("DeleteEvent")
	ForkEvent                     = EventType("ForkEvent")
	GollumEvent                   = EventType("GollumEvent")
	IssueCommentEvent             = EventType("IssueCommentEvent")
	IssuesEvent                   = EventType("IssuesEvent")
	MemberEvent                   = EventType("MemberEvent")
	PublicEvent                   = EventType("PublicEvent")
	PullRequestEvent              = EventType("PullRequestEvent")
	PullRequestReviewEvent        = EventType("PullRequestReviewEvent")
	PullRequestReviewCommentEvent = EventType("PullRequestReviewCommentEvent")
	PullRequestThreadEvent        = EventType("PullRequestThreadEvent")
	PushEvent                     = EventType("PushEvent")
	ReleaseEvent                  = EventType("ReleaseEvent")
	SponsorshipEvent              = EventType("SponsorshipEvent")
	WatchEvent                    = EventType("WatchEvent")
)
