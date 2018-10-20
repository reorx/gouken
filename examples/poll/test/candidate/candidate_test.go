package candidate

import (
	"testing"

	"github.com/reorx/gouken/utils"

	"github.com/reorx/gouken/examples/poll"
	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/stretchr/testify/assert"
)

func TestAddCandidate(t *testing.T) {
	app := poll.NewApp()
	c := app.GRPCClient()
	name := "foo"

	// add ok
	resp, err := c.AddCandidate(
		utils.Context(),
		&pb.AddCandidateRequest{
			Name: name,
		},
	)
	assert.Nil(t, err)
	t.Log("resp", resp)

	// add duplicate
	resp, err = c.AddCandidate(
		utils.Context(),
		&pb.AddCandidateRequest{
			Name: name,
		},
	)
	assert.NotNil(t, err)
	t.Log("resp 2", resp)
}
