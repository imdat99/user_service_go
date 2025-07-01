package services

import (
	"app/internal/utils"
	"app/pkg/database/ent"
	"context"
	"errors"
	"runtime"

	"github.com/sirupsen/logrus"
)

type HealthCheckService interface {
	DBcheck() error
	MemoryHeapCheck() error
}

type healthCheckService struct {
	Log *logrus.Logger
	DB  *ent.Client
	ctx *context.Context
}

func NewHealthCheckService(db *ent.Client, ctx *context.Context) HealthCheckService {
	return &healthCheckService{
		Log: utils.Log,
		DB:  db,
		ctx: ctx,
	}
}

func (s *healthCheckService) DBcheck() error {
	sqlDB, errDB := s.DB.Tx(*s.ctx) // Get the underlying SQL database connection from the ent.Client
	if errDB != nil {
		s.Log.Errorf("failed to access the database connection pool: %v", errDB)
		return errDB
	}

	// Try to query a single user to check database connectivity
	_, err := sqlDB.Client().User.Query().First(*s.ctx)
	if err != nil && !ent.IsNotFound(err) {
		s.Log.Errorf("failed to query the user table: %v", err)
		return err
	}

	return nil
}

// MemoryHeapCheck checks if heap memory usage exceeds a threshold
func (s *healthCheckService) MemoryHeapCheck() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats) // Collect memory statistics

	heapAlloc := memStats.HeapAlloc            // Heap memory currently allocated
	heapThreshold := uint64(300 * 1024 * 1024) // Example threshold: 300 MB

	s.Log.Infof("Heap Memory Allocation: %v bytes", heapAlloc)

	// If the heap allocation exceeds the threshold, return an error
	if heapAlloc > heapThreshold {
		s.Log.Errorf("Heap memory usage exceeds threshold: %v bytes", heapAlloc)
		return errors.New("heap memory usage too high")
	}

	return nil
}
