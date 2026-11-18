package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"example.com/employee-service/employee"
)

// Endpoints содержит все точки входа сервиса
type Endpoints struct {
	Create         endpoint.Endpoint
	GetByID        endpoint.Endpoint
	ChangePosition endpoint.Endpoint
}

// MakeServerEndpoints создаёт endpoints
func MakeServerEndpoints(s employee.Service) Endpoints {
	return Endpoints{
		Create:         makeCreateEndpoint(s),
		GetByID:        makeGetByIDEndpoint(s),
		ChangePosition: makeChangePositionEndpoint(s),
	}
}

// Запрос на создание сотрудника
type CreateRequest struct {
	Employee employee.Employee `json:"employee"`
}

// Ответ на создание сотрудника
type CreateResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

// Endpoint создания сотрудника
func makeCreateEndpoint(s employee.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)

		id, err := s.Create(ctx, req.Employee)
		if err != nil {
			return CreateResponse{Err: err.Error()}, nil
		}

		return CreateResponse{ID: id}, nil
	}
}

// Запрос на получение сотрудника
type GetByIDRequest struct {
	ID string `json:"id"`
}

// Ответ с данными сотрудника
type GetByIDResponse struct {
	Employee employee.Employee `json:"employee,omitempty"`
	Err      string            `json:"error,omitempty"`
}

// Endpoint получения сотрудника
func makeGetByIDEndpoint(s employee.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)

		e, err := s.GetByID(ctx, req.ID)
		if err != nil {
			return GetByIDResponse{Err: err.Error()}, nil
		}

		return GetByIDResponse{Employee: e}, nil
	}
}

// Запрос на изменение должности
type ChangePositionRequest struct {
	ID       string `json:"id"`
	Position string `json:"position"`
}

// Ответ на изменение должности
type ChangePositionResponse struct {
	Err string `json:"error,omitempty"`
}

// Endpoint изменения должности
func makeChangePositionEndpoint(s employee.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangePositionRequest)

		err := s.ChangePosition(ctx, req.ID, req.Position)
		if err != nil {
			return ChangePositionResponse{Err: err.Error()}, nil
		}

		return ChangePositionResponse{}, nil
	}
}
