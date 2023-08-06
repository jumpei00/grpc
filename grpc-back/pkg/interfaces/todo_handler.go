package interfaces

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/jumpei00/grpc/grpcback/pkg/application"
	todopb "github.com/jumpei00/grpc/grpcback/proto/todo/v1"
	todoconnect "github.com/jumpei00/grpc/grpcback/proto/todo/v1/v1connect"
)

type todoServer struct {
	todoApplication application.TodoApplication
}

func NewTodoServer(todoApplication application.TodoApplication) *todoServer {
	return &todoServer{todoApplication: todoApplication}
}

func (s *todoServer) HandleServer(mux *http.ServeMux) {
	path, handler := todoconnect.NewTodoServiceHandler(s)
	mux.Handle(path, handler)
}

func (s *todoServer) Index(ctx context.Context, req *connect.Request[todopb.IndexRequest]) (*connect.Response[todopb.IndexResponse], error) {
	todos, err := s.todoApplication.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	todoRes := make([]*todopb.Todo, len(todos))
	for i, t := range todos {
		todoRes[i] = &todopb.Todo{
			Key:     t.Key,
			Content: t.Content,
		}
	}

	res := connect.NewResponse(&todopb.IndexResponse{Todos: todoRes})

	return res, nil
}

func (s *todoServer) Show(ctx context.Context, req *connect.Request[todopb.ShowRequest]) (*connect.Response[todopb.ShowResponse], error) {
	todo, err := s.todoApplication.Get(ctx, req.Msg.Key)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&todopb.ShowResponse{Todo: &todopb.Todo{Key: todo.Key, Content: todo.Content}})

	return res, nil
}

func (s *todoServer) Create(ctx context.Context, req *connect.Request[todopb.CreateRequest]) (*connect.Response[todopb.CreateResponse], error) {
	todo, err := s.todoApplication.Create(ctx, req.Msg.Todo.Content)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&todopb.CreateResponse{Todo: &todopb.Todo{Key: todo.Key, Content: todo.Content}})

	return res, nil
}

func (s *todoServer) Update(ctx context.Context, req *connect.Request[todopb.UpdateRequest]) (*connect.Response[todopb.UpdateResponse], error) {
	todo, err := s.todoApplication.Update(ctx, req.Msg.Todo.Key, req.Msg.Todo.Content)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&todopb.UpdateResponse{Todo: &todopb.Todo{Key: todo.Key, Content: todo.Content}})

	return res, nil
}

func (s *todoServer) Delete(ctx context.Context, req *connect.Request[todopb.DeleteRequest]) (*connect.Response[todopb.DeleteResponse], error) {
	err := s.todoApplication.Delete(ctx, req.Msg.Key)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&todopb.DeleteResponse{})

	return res, nil
}
