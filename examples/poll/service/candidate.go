package service

import (
	"errors"

	"golang.org/x/net/context"

	pb "github.com/reorx/gouken/examples/poll/proto"
)

type candidatesStore map[string]pb.Candidate

var candidates candidatesStore

func add(n string) (pb.Candidate, error) {
	if candidates == nil {
		candidates = make(candidatesStore)
	}
	c := pb.Candidate{}
	_, ok := candidates[n]
	if ok {
		return c, errors.New("already exists")
	}
	c.Name = n
	candidates[n] = c
	return c, nil
}

// AddCandidate ...
func (p *Poll) AddCandidate(ctx context.Context,
	req *pb.AddCandidateRequest) (resp *pb.CandidateResponse, ferr error) {
	resp = new(pb.CandidateResponse)
	ferr = nil
	c, err := add(req.Name)
	if err != nil {
		ferr = err
		return
	}
	resp.Data = &c
	return
}

// DeleteCandidate ...
func (p *Poll) DeleteCandidate(ctx context.Context,
	req *pb.DeleteCandidateRequest) (resp *pb.DeleteCandidateResponse, ferr error) {
	return nil, nil
}

// ListCandidate ...
func (p *Poll) ListCandidate(ctx context.Context,
	req *pb.ListCandidateRequest) (resp *pb.ListCandidateResponse, ferr error) {
	return nil, nil
}
