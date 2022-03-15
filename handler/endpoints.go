package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/store"
	"github.com/micro/services/pkg/tenant"

	pb "analytics/proto"
)

// Track inserts a new counter in the store
func (a *Analytics) Track(ctx context.Context, req *pb.AnalyticsRequest, rsp *pb.Empty) error {
	// Validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("analytics.track", "missing name")
	}

	defer func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		tnt, ok := tenant.FromContext(ctx)
		if !ok {
			tnt = "default"
		}

		key := fmt.Sprintf("%s:%s", tnt, req.Name)

		var counter *pb.Counter

		recs, err := store.Read(key)
		if err == store.ErrNotFound {
			t := time.Now().Format(time.RFC3339)
			counter = &pb.Counter{
				Name:    req.Name,
				Created: t,
				Value:   1,
			}
		} else if err == nil {
			if err := recs[0].Decode(&counter); err != nil {
				return
			}
			counter.Value = counter.Value + 1
		} else {
			return
		}

		rec := store.NewRecord(key, counter)

		if err = store.Write(rec); err != nil {
			return
		}
	}()

	return nil
}

// Get returns a single counter
func (a *Analytics) Get(ctx context.Context, req *pb.AnalyticsRequest, rsp *pb.AnalyticsResponse) error {
	// Validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("analytics.get", "missing name")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s", tnt, req.Name)

	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("analytics.get", "Counter not found")
	} else if err != nil {
		return errors.InternalServerError("analytics.get", "Error reading from store: %v", err.Error())
	}

	// Decode the counter
	var counter *pb.Counter
	if err := recs[0].Decode(&counter); err != nil {
		return errors.InternalServerError("analytics.get", "Error unmarshaling JSON: %v", err.Error())
	}

	rsp.Counter = counter

	return nil
}

// Delete removes the counter from the store
func (a *Analytics) Delete(ctx context.Context, req *pb.AnalyticsRequest, rsp *pb.AnalyticsResponse) error {
	// Validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("analytics.delete", "missing name")
	}

	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	key := fmt.Sprintf("%s:%s", tnt, req.Name)

	recs, err := store.Read(key)
	if err == store.ErrNotFound {
		return errors.NotFound("analytics.delete", "Counter not found")
	} else if err != nil {
		return errors.InternalServerError("analytics.delete", "Error reading from store: %v", err.Error())
	}

	// Decode the counter
	var counter *pb.Counter
	if err := recs[0].Decode(&counter); err != nil {
		return errors.InternalServerError("analytics.delete", "Error unmarshaling JSON: %v", err.Error())
	}

	// now delete it
	if err := store.Delete(key); err != nil && err != store.ErrNotFound {
		return errors.InternalServerError("analytics.delete", "Failed to delete counter")
	}

	rsp.Counter = counter

	return nil
}

// List returns all of the counters in the store
func (a *Analytics) List(ctx context.Context, req *pb.Empty, rsp *pb.Counters) error {
	tnt, ok := tenant.FromContext(ctx)
	if !ok {
		tnt = "default"
	}

	// Read all counters from the store
	recs, err := store.Read(tnt, store.ReadPrefix())
	if err != nil {
		return errors.InternalServerError("analytics.list", "Error reading from store: %v", err.Error())
	}

	// Initialize the response counters slice
	rsp.Counters = make([]*pb.Counter, len(recs))

	// Retrieve all of the records in the store
	for i, rec := range recs {

		// Unmarshal the counters into the response
		if err := rec.Decode(&rsp.Counters[i]); err != nil {
			return errors.InternalServerError("analytics.list", "Error decoding counter: %v", err.Error())
		}
	}

	return nil
}
