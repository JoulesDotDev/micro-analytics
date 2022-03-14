package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/services/pkg/tenant"

	pb "analytics/proto"
)

// CreateProject inserts a new project in the store
func (e *Analytics) CreateProject(ctx context.Context, req *pb.CreateProjectRequest, rsp *pb.ProjectResponse) error {
	// Validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("project.create", "missing project name")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	// generate a key (uuid v4)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	t := time.Now().Format(time.RFC3339)

	project := &pb.Project{
		Id:          id.String(),
		Created:     t,
		Updated:     t,
		Name:        req.Name,
		Description: req.Description,
	}

	key := fmt.Sprintf("%s:%s:project", tnt, id)
	rec := store.NewRecord(key, project)

	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("project.created", "failed to create project")
	}

	rsp.Project = project

	return nil
}

// GetProject returns a project, looking up using ID
func (h *Analytics) GetProject(ctx context.Context, req *pb.RequestById, rsp *pb.ProjectResponse) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("project.update", "Missing project ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s:project", tnt, req.Id)

	// read the specific project
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("project.update", "Project not found")
	} else if err != nil {
		return errors.InternalServerError("project.update", "Error reading from store: %v", err.Error())
	}

	// Decode the project
	var project *pb.Project
	if err := recs[0].Decode(&project); err != nil {
		return errors.InternalServerError("project.update", "Error unmarshaling JSON: %v", err.Error())
	}

	rsp.Project = project

	return nil
}

// DeleteProject removes the project and related actions from the store, looking up using ID
func (h *Analytics) DeleteProject(ctx context.Context, req *pb.RequestById, rsp *pb.ProjectResponse) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("project.delete", "Missing project ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	// Check if project actions exist
	prefix := fmt.Sprintf("%s:%s", tnt, req.Id)
	suffix := ":action"

	keys, err := store.List(store.ListPrefix(prefix), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("project.delete", "Error reading from store: %v", err.Error())
	}

	if len(keys) > 0 {
		return errors.BadRequest("project.delete", "Project can't be deleted if it has actions")
	}

	key := fmt.Sprintf("%s:%s:project", tnt, req.Id)

	// read the specific project
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("project.delete", "Project not found")
	} else if err != nil {
		return errors.InternalServerError("project.delete", "Error reading from store: %v", err.Error())
	}

	// Decode the project
	var project *pb.Project
	if err := recs[0].Decode(&project); err != nil {
		return errors.InternalServerError("project.delete", "Error unmarshaling JSON: %v", err.Error())
	}

	// now delete it
	if err := store.Delete(key); err != nil && err != store.ErrNotFound {
		return errors.InternalServerError("project.delete", "Failed to delete project")
	}

	rsp.Project = project

	return nil
}

// UpdateProject is a unary API which updates a project in the store
func (h *Analytics) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest, rsp *pb.ProjectResponse) error {
	// Validate the request
	if req.Project == nil {
		return errors.BadRequest("project.update", "Missing project")
	}
	if len(req.Project.Id) == 0 {
		return errors.BadRequest("project.update", "Missing project ID")
	}
	if len(req.Project.Name) == 0 && len(req.Project.Description) == 0 {
		return errors.BadRequest("action.update", "Provide a property to update")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s:project", tnt, req.Project.Id)

	// read the specific project
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("project.update", "Project not found")
	} else if err != nil {
		return errors.InternalServerError("project.update", "Error reading from store: %v", err.Error())
	}

	// Decode the project
	var project *pb.Project
	if err := recs[0].Decode(&project); err != nil {
		return errors.InternalServerError("project.update", "Error unmarshaling JSON: %v", err.Error())
	}

	// Update the project's name and text
	if len(req.Project.Name) > 0 {
		project.Name = req.Project.Name
	}
	if len(req.Project.Description) > 0 {
		project.Description = req.Project.Description
	}
	project.Updated = time.Now().Format(time.RFC3339)

	rec := store.NewRecord(key, project)

	// Write the updated project to the store
	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("project.update", "Error writing to store: %v", err.Error())
	}

	rsp.Project = project

	return nil
}

// ListProjects returns all of the projects in the store
func (h *Analytics) ListProjects(ctx context.Context, req *pb.Empty, rsp *pb.ListProjectsResponse) error {
	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	prefix := fmt.Sprintf("%s:", tnt)
	suffix := ":project"

	// Find all project keys
	keys, err := store.List(store.ListPrefix(prefix), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	// Initialize the response projects slice
	rsp.Projects = make([]*pb.Project, len(keys))

	// Retrieve all of the records in the store
	for i, key := range keys {
		recs, err := store.Read(key)
		if err != nil {
			return errors.InternalServerError("projects.list", "Error reading from store: %v", err.Error())
		}
		// Unmarshal the projects into the response
		if err := recs[0].Decode(&rsp.Projects[i]); err != nil {
			return errors.InternalServerError("projects.list", "Error decoding project: %v", err.Error())
		}
	}

	return nil
}
