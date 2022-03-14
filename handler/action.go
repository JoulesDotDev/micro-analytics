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

// Create inserts a new action in the store
func (e *Analytics) CreateAction(ctx context.Context, req *pb.CreateActionRequest, rsp *pb.ActionResponse) error {
	// Validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("action.create", "missing action name")
	}
	if len(req.Project) == 0 {
		return errors.BadRequest("action.create", "missing project ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s:project", tnt, req.Project)

	// find the parent project
	_, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.create", "Project not found")
	} else if err != nil {
		return errors.InternalServerError("action.create", "Error reading from store: %v", err.Error())
	}

	// generate a key (uuid v4)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	t := time.Now().Format(time.RFC3339)

	action := &pb.Action{
		Id:          id.String(),
		Created:     t,
		Updated:     t,
		Name:        req.Name,
		Description: req.Description,
		Value:       0,
	}

	key = fmt.Sprintf("%s:%s:%s:action", tnt, req.Project, id)
	rec := store.NewRecord(key, action)

	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("action.created", "failed to create action")
	}

	rsp.Action = action

	return nil
}

// Update is a unary API which updates a action in the store
func (h *Analytics) GetAction(ctx context.Context, req *pb.RequestById, rsp *pb.ActionResponse) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("action.update", "Missing action ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	suffix := fmt.Sprintf("%s:action", req.Id)

	keys, err := store.List(store.ListPrefix(tnt), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	key := keys[0]

	// read the specific action
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.update", "Action not found")
	} else if err != nil {
		return errors.InternalServerError("action.update", "Error reading from store: %v", err.Error())
	}

	// Decode the action
	var action *pb.Action
	if err := recs[0].Decode(&action); err != nil {
		return errors.InternalServerError("action.update", "Error unmarshaling JSON: %v", err.Error())
	}

	rsp.Action = action

	return nil
}

// Delete removes the action from the store, looking up using ID
func (h *Analytics) DeleteAction(ctx context.Context, req *pb.RequestById, rsp *pb.ActionResponse) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("action.delete", "Missing action ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	suffix := fmt.Sprintf("%s:action", req.Id)

	keys, err := store.List(store.ListPrefix(tnt), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	key := keys[0]

	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.delete", "Action not found")
	} else if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	// Decode the action
	var action *pb.Action
	if err := recs[0].Decode(&action); err != nil {
		return errors.InternalServerError("action.delete", "Error unmarshaling JSON: %v", err.Error())
	}

	// now delete it
	if err := store.Delete(key); err != nil && err != store.ErrNotFound {
		return errors.InternalServerError("action.delete", "Failed to delete action")
	}

	rsp.Action = action

	return nil
}

// TriggerAction increments the value of the action by 1
func (h *Analytics) TriggerAction(ctx context.Context, req *pb.RequestById, rsp *pb.Empty) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("action.reset", "Missing action ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	suffix := fmt.Sprintf("%s:action", req.Id)

	keys, err := store.List(store.ListPrefix(tnt), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.reset", "Error reading from store: %v", err.Error())
	}

	key := keys[0]

	// read the specific action
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.reset", "Action not found")
	} else if err != nil {
		return errors.InternalServerError("action.reset", "Error reading from store: %v", err.Error())
	}

	// Decode the action
	var action *pb.Action
	if err := recs[0].Decode(&action); err != nil {
		return errors.InternalServerError("action.reset", "Error unmarshaling JSON: %v", err.Error())
	}

	// Reset the action's value
	action.Value = action.Value + 1

	rec := store.NewRecord(key, action)

	// Write the updated action to the store
	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("action.reset", "Error writing to store: %v", err.Error())
	}

	return nil
}

// ResetAction resets value of the action to 0
func (h *Analytics) ResetAction(ctx context.Context, req *pb.RequestById, rsp *pb.ActionResponse) error {
	// Validate the request
	if len(req.Id) == 0 {
		return errors.BadRequest("action.reset", "Missing action ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	suffix := fmt.Sprintf("%s:action", req.Id)

	keys, err := store.List(store.ListPrefix(tnt), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.reset", "Error reading from store: %v", err.Error())
	}

	key := keys[0]

	// read the specific action
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.reset", "Action not found")
	} else if err != nil {
		return errors.InternalServerError("action.reset", "Error reading from store: %v", err.Error())
	}

	// Decode the action
	var action *pb.Action
	if err := recs[0].Decode(&action); err != nil {
		return errors.InternalServerError("action.reset", "Error unmarshaling JSON: %v", err.Error())
	}

	// Reset the action's value
	action.Value = 0

	rec := store.NewRecord(key, action)

	// Write the updated action to the store
	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("action.reset", "Error writing to store: %v", err.Error())
	}

	rsp.Action = action

	return nil
}

// Update is a unary API which updates a action in the store
func (h *Analytics) UpdateAction(ctx context.Context, req *pb.UpdateActionRequest, rsp *pb.ActionResponse) error {
	// Validate the request
	if req.Action == nil {
		return errors.BadRequest("action.update", "Missing action")
	}
	if len(req.Action.Id) == 0 {
		return errors.BadRequest("action.update", "Missing action ID")
	}
	if len(req.Action.Name) == 0 && len(req.Action.Description) == 0 {
		return errors.BadRequest("action.update", "Provide a property to update")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	suffix := fmt.Sprintf("%s:action", req.Action.Id)

	keys, err := store.List(store.ListPrefix(tnt), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	key := keys[0]

	// read the specific action
	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.update", "Action not found")
	} else if err != nil {
		return errors.InternalServerError("action.update", "Error reading from store: %v", err.Error())
	}

	// Decode the action
	var action *pb.Action
	if err := recs[0].Decode(&action); err != nil {
		return errors.InternalServerError("action.update", "Error unmarshaling JSON: %v", err.Error())
	}

	// Update the action's name and description
	if len(req.Action.Name) > 0 {
		action.Name = req.Action.Name
	}
	if len(req.Action.Description) > 0 {
		action.Description = req.Action.Description
	}
	action.Updated = time.Now().Format(time.RFC3339)

	rec := store.NewRecord(key, action)

	// Write the updated action to the store
	if err = store.Write(rec); err != nil {
		return errors.InternalServerError("action.update", "Error writing to store: %v", err.Error())
	}

	rsp.Action = action

	return nil
}

// List returns all of the actions in the store
func (h *Analytics) ListActions(ctx context.Context, req *pb.ListActionsRequest, rsp *pb.ListActionsResponse) error {
	//validate the request
	if len(req.Project) == 0 {
		return errors.BadRequest("action.list", "Missing project ID")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s:project", tnt, req.Project)

	// find the parent project
	_, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("action.list", "Project not found")
	} else if err != nil {
		return errors.InternalServerError("action.list", "Error reading from store: %v", err.Error())
	}

	// Find all action keys
	prefix := fmt.Sprintf("%s:%s", tnt, req.Project)
	suffix := ":action"

	keys, err := store.List(store.ListPrefix(prefix), store.ListSuffix(suffix))
	if err != nil {
		return errors.InternalServerError("action.delete", "Error reading from store: %v", err.Error())
	}

	// Initialize the response actions slice
	rsp.Actions = make([]*pb.Action, len(keys))

	// Retrieve all of the records in the store
	for i, key := range keys {
		recs, err := store.Read(key)
		if err != nil {
			return errors.InternalServerError("actions.list", "Error reading from store: %v", err.Error())
		}
		// Unmarshal the actions into the response
		if err := recs[0].Decode(&rsp.Actions[i]); err != nil {
			return errors.InternalServerError("actions.list", "Error decoding action: %v", err.Error())
		}
	}
	return nil
}
