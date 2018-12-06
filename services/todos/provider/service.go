package provider

import (
	"context"

	"github.com/koryaz/eskit-todo/generated/grpc/go/todos"
	"github.com/makkalot/eskit/generated/grpc/go/common"

	"github.com/makkalot/eskit/services/clients"

	"fmt"

	"github.com/makkalot/eskit/generated/grpc/go/crudstore"
	common2 "github.com/makkalot/eskit/services/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoServiceProvider struct {
	crud *clients.CrudStoreClient
}

func NewTodoServiceProvider(crudStoreEndpoint string) (*TodoServiceProvider, error) {
	ctx := context.Background()
	var crudConn *grpc.ClientConn

	if err := common2.RetryNormal(func() error {
		var err error
		crudConn, err = grpc.Dial(crudStoreEndpoint, grpc.WithInsecure())
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	crudGRPC := crudstore.NewCrudStoreServiceClient(crudConn)
	_, err := crudGRPC.RegisterType(ctx, &crudstore.RegisterTypeRequest{
		Spec: &crudstore.CrudEntitySpec{
			EntityType: clients.EntityTypeFromMsg(&todos.Todo{}),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("type registration for use entity type failed : %v", err)
	}

	crudClient, err := clients.NewCrudStoreWithActiveConn(ctx, crudGRPC)
	if err != nil {
		return nil, err
	}

	return &TodoServiceProvider{
		crud: crudClient,
	}, nil
}

func (t *TodoServiceProvider) Healtz(ctx context.Context, request *todos.HealthRequest) (*todos.HealthResponse, error) {
	return &todos.HealthResponse{}, nil
}

func (t *TodoServiceProvider) Create(ctx context.Context, request *todos.CreateRequest) (*todos.CreateResponse, error) {
	originator := &common.Originator{
		Id:      uuid.Must(uuid.NewV4()).String(),
		Version: "1",
	}

	todo := &todos.Todo{
		Originator: originator,
		Title:      request.Title,
		Url:        request.Url,
		Completed:  false,
		//TODO set oreder and text
		Order: "1",
		Text:  "",
	}

	createdOriginator, err := t.crud.Create(todo)
	if err != nil {
		return nil, status.Error(codes.Internal, "creation failed")
	}

	todo.Originator = createdOriginator

	return &todos.CreateResponse{
		Todo: todo,
	}, nil
}

func (u *TodoServiceProvider) Find(ctx context.Context, req *todos.FindRequest) (*todos.FindResponse, error) {
	if req.Originator == nil {
		return nil, status.Error(codes.InvalidArgument, "missing originator")
	}

	retrievedTodo := &todos.Todo{}
	//TODO fetchdeleted?
	var FetchDeleted bool
	if err := u.crud.Get(req.Originator, retrievedTodo, FetchDeleted); err != nil {
		return nil, err
	}

	return &todos.FindResponse{
		Todo: retrievedTodo,
	}, nil
}

func (u *TodoServiceProvider) Update(ctx context.Context, req *todos.UpdateRequest) (*todos.UpdateResponse, error) {
	if req.Originator == nil {
		return nil, status.Error(codes.InvalidArgument, "missing originator")
	}

	retrievedTodo := &todos.Todo{}
	if err := u.crud.Get(req.Originator, retrievedTodo, false); err != nil {
		return nil, err
	}

	if req.UpdatedTodo.Title != "" {
		retrievedTodo.Title = req.UpdatedTodo.Title
	}

	if req.UpdatedTodo.Url != "" {
		retrievedTodo.Url = req.UpdatedTodo.Url
	}

	if req.UpdatedTodo.Completed != retrievedTodo.Completed {
		retrievedTodo.Completed = req.UpdatedTodo.Completed
	}

	if req.UpdatedTodo.Text != "" {
		retrievedTodo.Text = req.UpdatedTodo.Text
	}

	updatedOriginator, err := u.crud.Update(retrievedTodo)
	if err != nil {
		return nil, err
	}

	retrievedTodo.Originator = updatedOriginator

	return &todos.UpdateResponse{
		Todo: retrievedTodo,
	}, nil

}

func (u *TodoServiceProvider) Delete(ctx context.Context, req *todos.DeleteRequest) (*todos.DeleteResponse, error) {
	if req.Originator == nil {
		return nil, status.Error(codes.InvalidArgument, "missing originator")
	}

	deletedOriginator, err := u.crud.Delete(req.Originator, &todos.Todo{})
	if err != nil {
		return nil, err
	}

	return &todos.DeleteResponse{
		Originator: deletedOriginator,
	}, nil
}

func (u *TodoServiceProvider) DeleteAll(ctx context.Context, req *todos.DeleteAllRequest) (*todos.DeleteAllResponse, error) {

	//TODO implement DeleteAll
	/*deletedOriginator, err := u.crud.(req.Originator, &todos.Todo{})
	if err != nil {
		return nil, err
	}*/

	return &todos.DeleteAllResponse{}, nil
}
