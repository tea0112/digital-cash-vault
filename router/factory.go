package router

import "errors"

func New(kind string) (Router, error) {
	switch kind {
	case "chi":
		return NewChi(), nil
	// sample for flexible usage of router library
	case "mux":
		return NewMux(), nil
	default:
		return nil, errors.New("unknown router")
	}
}
