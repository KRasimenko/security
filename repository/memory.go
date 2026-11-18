package repository

import (
	"context"
	"sync"

	"example.com/employee-service/employee"
)

// MemoryRepository — хранилище сотрудников в оперативной памяти
type MemoryRepository struct {
	mu        sync.RWMutex
	employees map[string]employee.Employee
}

// NewMemoryRepository создаёт новое хранилище
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		employees: make(map[string]employee.Employee),
	}
}

// CreateEmployee сохраняет сотрудника
func (r *MemoryRepository) CreateEmployee(ctx context.Context, e employee.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.employees[e.ID] = e
	return nil
}

// GetEmployeeByID получает сотрудника по ID
func (r *MemoryRepository) GetEmployeeByID(ctx context.Context, id string) (employee.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	e, exists := r.employees[id]
	if !exists {
		return employee.Employee{}, employee.ErrEmployeeNotFound
	}

	return e, nil
}

// ChangeEmployeePosition изменяет должность сотрудника
func (r *MemoryRepository) ChangeEmployeePosition(ctx context.Context, id string, position string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	e, exists := r.employees[id]
	if !exists {
		return employee.ErrEmployeeNotFound
	}

	e.Position = position
	r.employees[id] = e

	return nil
}
