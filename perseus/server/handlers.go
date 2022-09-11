package main

import (
	"context"

	perseus "github.com/ornlu-is/system-a/perseus"
)

// Add implements perseus.Persues.Add
func (s *server) Add(ctx context.Context, req *perseus.AddReq) (*perseus.AddResp, error) {
	resp, err := Add()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Delete implements perseus.Perseus.Delete
func (s *server) Delete(ctx context.Context, req *perseus.DeleteReq) (*perseus.DeleteResp, error) {
	resp, err := Delete()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get implements perseus.Perseus.Get
func (s *server) Get(ctx context.Context, req *perseus.GetReq) (*perseus.GetResp, error) {
	resp, err := Get()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
