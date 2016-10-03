package candidate

import (
	"testing"

	"github.com/reorx/gouken"
	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/examples/poll/test"
	"github.com/stretchr/testify/assert"
)

func TestAddCandidate(t *testing.T) {
	c := test.GetClient()
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
