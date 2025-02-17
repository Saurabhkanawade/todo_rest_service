package endpoints

import (
	"context"
	"github.com/Saurabhkanawade/todo_rest_service/internal/models"
	"github.com/Saurabhkanawade/todo_rest_service/internal/services"
	"github.com/go-kit/kit/endpoint"
	"github.com/gofrs/uuid"
)

type TaskEndpoints struct {
	CreateTask endpoint.Endpoint
	GetAllTask endpoint.Endpoint
	GetTask    endpoint.Endpoint
	UpdateTask endpoint.Endpoint
	DeleteTask endpoint.Endpoint
}

func MakeTaskEndpoints(taskService services.TaskService) TaskEndpoints {
	return TaskEndpoints{
		CreateTask: MakeCreateTaskEndpoint(taskService),
		GetAllTask: MakeGetAllTaskEndpoint(taskService),
		GetTask:    MakeGetTaskEndpoint(taskService),
		UpdateTask: MakeUpdateTaskEndpoint(taskService),
		DeleteTask: MakeDeleteTaskEndpoint(taskService),
	}
}

// CreateTaskRequest
// swagger:model CreateTaskRequest
type CreateTaskRequest struct {
	Task models.Task `json:"task"`
}

// CreateTaskResponse
// swagger:response CreateTaskResponse
type CreateTaskResponse struct {
	//in:body
	Task models.Task `json:"task"`
}

func MakeCreateTaskEndpoint(service services.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateTaskRequest)
		if !ok {
			return nil, err
		}

		task, err := service.Create(ctx, req.Task)
		if err != nil {
			return nil, err
		}

		return CreateTaskResponse{
			Task: *task,
		}, err
	}
}

type GetTaskByIDParamsRequest struct {
	TaskId uuid.UUID
}

// GetTaskByIDResponse
// swagger:response GetTaskByIDResponse
type GetTaskByIDResponse struct {
	//in:body
	Task models.Task `json:"task"`
}

func MakeGetTaskEndpoint(service services.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GetTaskByIDParamsRequest)
		if !ok {
			return nil, err
		}
		task, err := service.Get(ctx, req.TaskId)
		if err != nil {
			return nil, err
		}

		return GetTaskByIDResponse{
			Task: *task,
		}, nil
	}
}

type GetAllTaskRequest struct {
}

// GetAllTaskResponse
// swagger:response GetAllTaskResponse
type GetAllTaskResponse struct {
	//in:body
	Tasks []models.Task `json:"tasks"`
}

func MakeGetAllTaskEndpoint(service services.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		taskService, err := service.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		res := GetAllTaskResponse{
			Tasks: taskService,
		}
		return res, err
	}
}

// swagger:model UpdateTaskRequestBody
type UpdateTaskRequestBody struct {
	Task models.Task `json:"task"`
}

// UpdateTaskResponse
// swagger:response UpdateTaskResponse
type UpdateTaskResponse struct {
}

func MakeUpdateTaskEndpoint(service services.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(UpdateTaskRequestBody)
		if !ok {
			return nil, err
		}
		err = service.Update(ctx, req.Task.ID, req.Task)
		if err != nil {
			return nil, err
		}
		return UpdateTaskResponse{}, nil
	}
}

type DeleteTaskByIdRequest struct {
	TaskId uuid.UUID
}

// DeleteTaskByIdResponse
// swagger:response DeleteTaskByIdResponse
type DeleteTaskByIdResponse struct {
}

func MakeDeleteTaskEndpoint(service services.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(DeleteTaskByIdRequest)
		if !ok {
			return nil, err
		}
		err = service.Delete(ctx, req.TaskId)
		if err != nil {
			return nil, err
		}

		return nil, err
	}
}
