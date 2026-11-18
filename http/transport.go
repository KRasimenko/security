package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	"example.com/employee-service/transport"
)

// MakeHTTPHandler создаёт HTTP-маршруты
func MakeHTTPHandler(endpoints transport.Endpoints) http.Handler {
	router := mux.NewRouter()

	router.Methods("POST").Path("/employees").Handler(
		decodeHandler(endpoints.Create, decodeCreateRequest),
	)

	router.Methods("GET").Path("/employees/{id}").Handler(
		decodeHandler(endpoints.GetByID, decodeGetByIDRequest),
	)

	router.Methods("PUT").Path("/employees/{id}/position").Handler(
		decodeHandler(endpoints.ChangePosition, decodeChangePositionRequest),
	)

	return router
}

// Универсальный обработчик Endpoint
func decodeHandler(
	endpoint endpoint.Endpoint,
	decode func(r *http.Request) (interface{}, error),
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request, err := decode(r)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		response, err := endpoint(r.Context(), request)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}

// Создание сотрудника
func decodeCreateRequest(r *http.Request) (interface{}, error) {
	var request transport.CreateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// Получение сотрудника
func decodeGetByIDRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	return transport.GetByIDRequest{
		ID: vars["id"],
	}, nil
}

// Изменение должности
func decodeChangePositionRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	var body struct {
		Position string `json:"position"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return transport.ChangePositionRequest{
		ID:       vars["id"],
		Position: body.Position,
	}, nil
}
