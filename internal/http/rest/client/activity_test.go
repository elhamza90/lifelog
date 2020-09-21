package client_test

import (
	"testing"
	"time"

	"github.com/elhamza90/lifelog/internal/domain"
	"github.com/elhamza90/lifelog/internal/http/rest/client"
)

func TestPostActivity(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o"
	testAct := domain.Activity{
		Label:    "Do smth",
		Place:    "Somewhere",
		Desc:     "With Details",
		Time:     time.Now().Add(time.Duration(-1 * time.Hour)),
		Duration: time.Duration(time.Minute * 20),
		Tags:     []domain.Tag{},
	}
	_, err := client.PostActivity(testAct, token)
	if err != nil {
		t.Fatal(err)
	}
}