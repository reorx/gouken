package candidate

import (
	"testing"

	"github.com/reorx/gouken"
	"github.com/reorx/gouken/examples/poll"
	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/stretchr/testify/assert"
)

func init() {
	poll.Init()
}

func TestAddCandidate(t *testing.T) {
	c := poll.Client()
	name := "foo"

	// add ok
	resp, err := c.AddCandidate(
		gouken.Context(),
		&pb.AddCandidateRequest{
			Name: name,
		},
	)
	assert.Nil(t, err)
	t.Log("resp", resp)

	// add duplicate
	resp, err = c.AddCandidate(
		gouken.Context(),
		&pb.AddCandidateRequest{
			Name: name,
		},
	)
	assert.NotNil(t, err)
	t.Log("resp 2", resp)
}
