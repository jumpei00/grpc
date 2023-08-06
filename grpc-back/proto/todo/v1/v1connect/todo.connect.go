// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/jumpei00/grpc/grpcback/proto/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TodoServiceName is the fully-qualified name of the TodoService service.
	TodoServiceName = "todo.v1.TodoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TodoServiceIndexProcedure is the fully-qualified name of the TodoService's Index RPC.
	TodoServiceIndexProcedure = "/todo.v1.TodoService/Index"
	// TodoServiceShowProcedure is the fully-qualified name of the TodoService's Show RPC.
	TodoServiceShowProcedure = "/todo.v1.TodoService/Show"
	// TodoServiceCreateProcedure is the fully-qualified name of the TodoService's Create RPC.
	TodoServiceCreateProcedure = "/todo.v1.TodoService/Create"
	// TodoServiceUpdateProcedure is the fully-qualified name of the TodoService's Update RPC.
	TodoServiceUpdateProcedure = "/todo.v1.TodoService/Update"
	// TodoServiceDeleteProcedure is the fully-qualified name of the TodoService's Delete RPC.
	TodoServiceDeleteProcedure = "/todo.v1.TodoService/Delete"
)

// TodoServiceClient is a client for the todo.v1.TodoService service.
type TodoServiceClient interface {
	Index(context.Context, *connect_go.Request[v1.IndexRequest]) (*connect_go.Response[v1.IndexResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
}

// NewTodoServiceClient constructs a client for the todo.v1.TodoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTodoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TodoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &todoServiceClient{
		index: connect_go.NewClient[v1.IndexRequest, v1.IndexResponse](
			httpClient,
			baseURL+TodoServiceIndexProcedure,
			opts...,
		),
		show: connect_go.NewClient[v1.ShowRequest, v1.ShowResponse](
			httpClient,
			baseURL+TodoServiceShowProcedure,
			opts...,
		),
		create: connect_go.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+TodoServiceCreateProcedure,
			opts...,
		),
		update: connect_go.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+TodoServiceUpdateProcedure,
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+TodoServiceDeleteProcedure,
			opts...,
		),
	}
}

// todoServiceClient implements TodoServiceClient.
type todoServiceClient struct {
	index  *connect_go.Client[v1.IndexRequest, v1.IndexResponse]
	show   *connect_go.Client[v1.ShowRequest, v1.ShowResponse]
	create *connect_go.Client[v1.CreateRequest, v1.CreateResponse]
	update *connect_go.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Index calls todo.v1.TodoService.Index.
func (c *todoServiceClient) Index(ctx context.Context, req *connect_go.Request[v1.IndexRequest]) (*connect_go.Response[v1.IndexResponse], error) {
	return c.index.CallUnary(ctx, req)
}

// Show calls todo.v1.TodoService.Show.
func (c *todoServiceClient) Show(ctx context.Context, req *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return c.show.CallUnary(ctx, req)
}

// Create calls todo.v1.TodoService.Create.
func (c *todoServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls todo.v1.TodoService.Update.
func (c *todoServiceClient) Update(ctx context.Context, req *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls todo.v1.TodoService.Delete.
func (c *todoServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// TodoServiceHandler is an implementation of the todo.v1.TodoService service.
type TodoServiceHandler interface {
	Index(context.Context, *connect_go.Request[v1.IndexRequest]) (*connect_go.Response[v1.IndexResponse], error)
	Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error)
	Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error)
	Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
}

// NewTodoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTodoServiceHandler(svc TodoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	todoServiceIndexHandler := connect_go.NewUnaryHandler(
		TodoServiceIndexProcedure,
		svc.Index,
		opts...,
	)
	todoServiceShowHandler := connect_go.NewUnaryHandler(
		TodoServiceShowProcedure,
		svc.Show,
		opts...,
	)
	todoServiceCreateHandler := connect_go.NewUnaryHandler(
		TodoServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	todoServiceUpdateHandler := connect_go.NewUnaryHandler(
		TodoServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	todoServiceDeleteHandler := connect_go.NewUnaryHandler(
		TodoServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/todo.v1.TodoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TodoServiceIndexProcedure:
			todoServiceIndexHandler.ServeHTTP(w, r)
		case TodoServiceShowProcedure:
			todoServiceShowHandler.ServeHTTP(w, r)
		case TodoServiceCreateProcedure:
			todoServiceCreateHandler.ServeHTTP(w, r)
		case TodoServiceUpdateProcedure:
			todoServiceUpdateHandler.ServeHTTP(w, r)
		case TodoServiceDeleteProcedure:
			todoServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTodoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTodoServiceHandler struct{}

func (UnimplementedTodoServiceHandler) Index(context.Context, *connect_go.Request[v1.IndexRequest]) (*connect_go.Response[v1.IndexResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.Index is not implemented"))
}

func (UnimplementedTodoServiceHandler) Show(context.Context, *connect_go.Request[v1.ShowRequest]) (*connect_go.Response[v1.ShowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.Show is not implemented"))
}

func (UnimplementedTodoServiceHandler) Create(context.Context, *connect_go.Request[v1.CreateRequest]) (*connect_go.Response[v1.CreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.Create is not implemented"))
}

func (UnimplementedTodoServiceHandler) Update(context.Context, *connect_go.Request[v1.UpdateRequest]) (*connect_go.Response[v1.UpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.Update is not implemented"))
}

func (UnimplementedTodoServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("todo.v1.TodoService.Delete is not implemented"))
}
