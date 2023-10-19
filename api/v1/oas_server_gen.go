// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddProject implements addProject operation.
	//
	// Add a single project.
	//
	// POST /add-project
	AddProject(ctx context.Context, req *AddProject) (AddProjectRes, error)
	// DeleteProject implements deleteProject operation.
	//
	// Delete a project.
	//
	// POST /delete-project
	DeleteProject(ctx context.Context, req *DeleteProjectRequest) (DeleteProjectRes, error)
	// GetProject implements getProject operation.
	//
	// Get a project.
	//
	// POST /get-project
	GetProject(ctx context.Context, req *GetProjectRequest) (GetProjectRes, error)
	// SuggestProjects implements suggestProjects operation.
	//
	// Search in projects.
	//
	// POST /suggest-projects
	SuggestProjects(ctx context.Context, req *SuggestProjectsRequest) (*SuggestProjectsResponse, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
