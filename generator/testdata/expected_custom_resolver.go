// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"
	"sync"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"golang.org/x/sync/singleflight"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	post "example/post"
	user "example/user"
)

// Org_Federation_GetPostResponseArgument is argument for "org.federation.GetPostResponse" message.
type Org_Federation_GetPostResponseArgument[T any] struct {
	Id     string
	Post   *Post
	Client T
}

// Org_Federation_PostArgument is argument for "org.federation.Post" message.
type Org_Federation_PostArgument[T any] struct {
	Id     string
	Post   *post.Post
	User   *User
	Client T
}

// Org_Federation_Post_UserArgument is custom resolver's argument for "user" field of "org.federation.Post" message.
type Org_Federation_Post_UserArgument[T any] struct {
	*Org_Federation_PostArgument[T]
	Client T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
	Content string
	Id      string
	Title   string
	U       *user.User
	UserId  string
	Client  T
}

// Org_Federation_User_NameArgument is custom resolver's argument for "name" field of "org.federation.User" message.
type Org_Federation_User_NameArgument[T any] struct {
	*Org_Federation_UserArgument[T]
	Org_Federation_User *User
	Client              T
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver FederationServiceResolver // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// Org_Post_PostServiceClient create a gRPC Client to be used to call methods in org.post.PostService.
	Org_Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
	// Org_User_UserServiceClient create a gRPC Client to be used to call methods in org.user.UserService.
	Org_User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
}

// FederationServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
	Org_Post_PostServiceClient post.PostServiceClient
	Org_User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
	// Resolve_Org_Federation_Post_User implements resolver for "org.federation.Post.user".
	Resolve_Org_Federation_Post_User(context.Context, *Org_Federation_Post_UserArgument[*FederationServiceDependentClientSet]) (*User, error)
	// Resolve_Org_Federation_User implements resolver for "org.federation.User".
	Resolve_Org_Federation_User(context.Context, *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error)
	// Resolve_Org_Federation_User_Name implements resolver for "org.federation.User.name".
	Resolve_Org_Federation_User_Name(context.Context, *Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]) (string, error)
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// Resolve_Org_Federation_Post_User resolve "org.federation.Post.user".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_Post_User(context.Context, *Org_Federation_Post_UserArgument[*FederationServiceDependentClientSet]) (ret *User, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Org_Federation_Post_User not implemented")
	return
}

// Resolve_Org_Federation_User resolve "org.federation.User".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_User(context.Context, *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (ret *User, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Org_Federation_User not implemented")
	return
}

// Resolve_Org_Federation_User_Name resolve "org.federation.User.name".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_User_Name(context.Context, *Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]) (ret string, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Org_Federation_User_Name not implemented")
	return
}

const (
	FederationService_DependentMethod_Org_Post_PostService_GetPost = "/org.post.PostService/GetPost"
	FederationService_DependentMethod_Org_User_UserService_GetUser = "/org.user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	resolver     FederationServiceResolver
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	if cfg.Resolver == nil {
		return nil, fmt.Errorf("Resolver field in FederationServiceConfig is not set. this field must be set")
	}
	Org_Post_PostServiceClient, err := cfg.Client.Org_Post_PostServiceClient(FederationServiceClientConfig{
		Service: "org.post.PostService",
		Name:    "post_service",
	})
	if err != nil {
		return nil, err
	}
	Org_User_UserServiceClient, err := cfg.Client.Org_User_UserServiceClient(FederationServiceClientConfig{
		Service: "org.user.UserService",
		Name:    "user_service",
	})
	if err != nil {
		return nil, err
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*celtypes.FieldType{
		"grpc.federation.private.GetPostResponseArgument": map[string]*celtypes.FieldType{
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.PostArgument": map[string]*celtypes.FieldType{
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.UserArgument": map[string]*celtypes.FieldType{
			"id":      grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
			"title":   grpcfed.NewCELFieldType(celtypes.StringType, "Title"),
			"content": grpcfed.NewCELFieldType(celtypes.StringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
		cel.CustomTypeAdapter(celHelper.TypeAdapter()),
		cel.CustomTypeProvider(celHelper.TypeProvider()),
	)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		resolver:     cfg.Resolver,
		client: &FederationServiceDependentClientSet{
			Org_Post_PostServiceClient: Org_Post_PostServiceClient,
			Org_User_UserServiceClient: Org_User_UserServiceClient,
		},
	}, nil
}

// GetPost implements "org.federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetPostResponse(ctx, &Org_Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]{
		Client: s.client,
		Id:     req.Id,
	})
	if err != nil {
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetPostResponse resolve "org.federation.GetPostResponse" message.
func (s *FederationService) resolve_Org_Federation_GetPostResponse(ctx context.Context, req *Org_Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) (*GetPostResponse, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.GetPostResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetPostResponseArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *Post
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetPostResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "post"
	     message: "Post"
	     args { name: "id", by: "$.id" }
	   }
	*/
	resPostIface, err, _ := sg.Do("post_org.federation.Post", func() (interface{}, error) {
		valueMu.RLock()
		args := &Org_Federation_PostArgument[*FederationServiceDependentClientSet]{
			Client: s.client,
		}
		// { name: "id", by: "$.id" }
		{
			_value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Id = _value.(string)
		}
		valueMu.RUnlock()
		return s.resolve_Org_Federation_Post(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resPost := resPostIface.(*Post)
	valueMu.Lock()
	valuePost = resPost // { name: "post", message: "Post" ... }
	envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("org.federation.Post")))
	evalValues["post"] = valuePost
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	{
		_value, err := grpcfed.EvalCEL(s.env, "post", envOpts, evalValues, reflect.TypeOf((*Post)(nil)))
		if err != nil {
			return nil, err
		}
		ret.Post = _value.(*Post)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetPostResponse", slog.Any("org.federation.GetPostResponse", s.logvalue_Org_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Post resolve "org.federation.Post" message.
func (s *FederationService) resolve_Org_Federation_Post(ctx context.Context, req *Org_Federation_PostArgument[*FederationServiceDependentClientSet]) (*Post, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.Post", slog.Any("message_args", s.logvalue_Org_Federation_PostArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *post.Post
		valueUser *User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.PostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "org.post.PostService/GetPost"
	     request { field: "id", by: "$.id" }
	     response { name: "post", field: "post", autobind: true }
	   }
	*/
	resGetPostResponseIface, err, _ := sg.Do("org.post.PostService/GetPost", func() (interface{}, error) {
		valueMu.RLock()
		args := &post.GetPostRequest{}
		// { field: "id", by: "$.id" }
		{
			_value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Id = _value.(string)
		}
		valueMu.RUnlock()
		return s.client.Org_Post_PostServiceClient.GetPost(ctx, args)
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_Post_PostService_GetPost, err); err != nil {
			return nil, err
		}
	}
	resGetPostResponse := resGetPostResponseIface.(*post.GetPostResponse)
	valueMu.Lock()
	valuePost = resGetPostResponse.GetPost() // { name: "post", field: "post", autobind: true }
	envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("org.post.Post")))
	evalValues["post"] = valuePost
	valueMu.Unlock()

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "user"
	     message: "User"
	     args { inline: "post" }
	   }
	*/
	resUserIface, err, _ := sg.Do("user_org.federation.User", func() (interface{}, error) {
		valueMu.RLock()
		args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
			Client: s.client,
		}
		// { inline: "post" }
		{
			_value, err := grpcfed.EvalCEL(s.env, "post", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
			if err != nil {
				return nil, err
			}
			_inlineValue := _value.(*post.Post)
			args.Id = _inlineValue.GetId()
			args.Title = _inlineValue.GetTitle()
			args.Content = _inlineValue.GetContent()
			args.UserId = _inlineValue.GetUserId()
		}
		valueMu.RUnlock()
		return s.resolve_Org_Federation_User(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resUser := resUserIface.(*User)
	valueMu.Lock()
	valueUser = resUser // { name: "user", message: "User" ... }
	envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("org.federation.User")))
	evalValues["user"] = valueUser
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost
	req.User = valueUser

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = valuePost.GetId()           // { name: "post", autobind: true }
	ret.Title = valuePost.GetTitle()     // { name: "post", autobind: true }
	ret.Content = valuePost.GetContent() // { name: "post", autobind: true }
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.User, err = s.resolver.Resolve_Org_Federation_Post_User(ctx, &Org_Federation_Post_UserArgument[*FederationServiceDependentClientSet]{
			Client:                      s.client,
			Org_Federation_PostArgument: req,
		})
		if err != nil {
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved org.federation.Post", slog.Any("org.federation.Post", s.logvalue_Org_Federation_Post(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
		valueU  *user.User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "org.user.UserService/GetUser"
	     request { field: "id", by: "$.user_id" }
	     response { name: "u", field: "user" }
	   }
	*/
	resGetUserResponseIface, err, _ := sg.Do("org.user.UserService/GetUser", func() (interface{}, error) {
		valueMu.RLock()
		args := &user.GetUserRequest{}
		// { field: "id", by: "$.user_id" }
		{
			_value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Id = _value.(string)
		}
		valueMu.RUnlock()
		return s.client.Org_User_UserServiceClient.GetUser(ctx, args)
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_User_UserService_GetUser, err); err != nil {
			return nil, err
		}
	}
	resGetUserResponse := resGetUserResponseIface.(*user.GetUserResponse)
	valueMu.Lock()
	valueU = resGetUserResponse.GetUser() // { name: "u", field: "user" }
	envOpts = append(envOpts, cel.Variable("u", cel.ObjectType("org.user.User")))
	evalValues["u"] = valueU
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.U = valueU

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ret, err := s.resolver.Resolve_Org_Federation_User(ctx, req)
	if err != nil {
		return nil, err
	}

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.Name, err = s.resolver.Resolve_Org_Federation_User_Name(ctx, &Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]{
			Client:                      s.client,
			Org_Federation_UserArgument: req,
			Org_Federation_User:         ret,
		})
		if err != nil {
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Org_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetPostResponseArgument(v *Org_Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Org_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Org_Federation_User(v.GetUser())),
	)
}

func (s *FederationService) logvalue_Org_Federation_PostArgument(v *Org_Federation_PostArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Org_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserArgument(v *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}
