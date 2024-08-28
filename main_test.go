package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("Should fetch the recent activity of a GitHub user", func(t *testing.T) {
		result, err := GetRecentActivity("kamranahmedse")

		asserts.NoError(err)
		asserts.NotEmpty(result)
		asserts.IsType([]*Event{}, result)
	})

	t.Run("Should fail to fetch the recent activity of a GitHub user", func(t *testing.T) {
		result, err := GetRecentActivity("inexistent-user")

		asserts.Error(err)
		asserts.Empty(result)
	})
}
