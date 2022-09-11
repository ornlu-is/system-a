package main

import (
	perseus "github.com/ornlu-is/system-a/perseus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Add does something.
func Add() (*perseus.AddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}

// Delete does something.
func Delete() (*perseus.DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}

// Get does something.
func Get() (*perseus.GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
