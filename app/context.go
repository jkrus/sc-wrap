package app

import "context"

// NewContext construct context of the application.
func NewContext() context.Context {
	return context.Background()
}
