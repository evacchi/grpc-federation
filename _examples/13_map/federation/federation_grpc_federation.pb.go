// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	post "example/post"
	user "example/user"
)

// Org_Federation_GetPostsResponseArgument is argument for "org.federation.GetPostsResponse" message.
type Org_Federation_GetPostsResponseArgument[T any] struct {
	Ids    []string
	Posts  *Posts
	Client T
}

// Org_Federation_PostsArgument is argument for "org.federation.Posts" message.
type Org_Federation_PostsArgument[T any] struct {
	Ids     []string
	PostIds []string
	Posts   []*post.Post
	Res     *post.GetPostsResponse
	Users   []*User
	Client  T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
	Res    *user.GetUserResponse
	User   *user.User
	UserId string
	Client T
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
	// User_UserServiceClient create a gRPC Client to be used to call methods in user.UserService.
	User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
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
	Post_PostServiceClient post.PostServiceClient
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

type FederationServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

type FederationServiceCELPluginConfig struct {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_Post_PostService_GetPosts = "/post.PostService/GetPosts"
	FederationService_DependentMethod_User_UserService_GetUser  = "/user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
		Name:    "",
	})
	if err != nil {
		return nil, err
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
		Name:    "",
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
		"grpc.federation.private.GetPostsResponseArgument": {
			"ids": grpcfed.NewCELFieldType(celtypes.NewListType(celtypes.StringType), "Ids"),
		},
		"grpc.federation.private.PostsArgument": {
			"post_ids": grpcfed.NewCELFieldType(celtypes.NewListType(celtypes.StringType), "PostIds"),
		},
		"grpc.federation.private.UserArgument": {
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
	})
	envOpts := []cel.EnvOption{
		cel.StdLib(),
		cel.Lib(grpcfedcel.NewLibrary()),
		cel.CrossTypeNumericComparisons(true),
		cel.CustomTypeAdapter(celHelper.TypeAdapter()),
		cel.CustomTypeProvider(celHelper.TypeProvider()),
	}
	env, err := cel.NewCustomEnv(envOpts...)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// GetPosts implements "org.federation.FederationService/GetPosts" method.
func (s *FederationService) GetPosts(ctx context.Context, req *GetPostsRequest) (res *GetPostsResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/GetPosts")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetPostsResponse(ctx, &Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]{
		Client: s.client,
		Ids:    req.Ids,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetPostsResponse resolve "org.federation.GetPostsResponse" message.
func (s *FederationService) resolve_Org_Federation_GetPostsResponse(ctx context.Context, req *Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]) (*GetPostsResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetPostsResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.GetPostsResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetPostsResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			posts *Posts
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.GetPostsResponseArgument", req)}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "posts"
	     message {
	       name: "Posts"
	       args { name: "post_ids", by: "$.ids" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*Posts, *localValueType]{
		Name:   "posts",
		Type:   cel.ObjectType("org.federation.Posts"),
		Setter: func(value *localValueType, v *Posts) { value.vars.posts = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Org_Federation_PostsArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "post_ids", by: "$.ids" }
			if err := grpcfed.SetCELValue(ctx, value, "$.ids", func(v []string) {
				args.PostIds = v
			}); err != nil {
				return nil, err
			}
			return s.resolve_Org_Federation_Posts(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Posts = value.vars.posts

	// create a message value to be returned.
	ret := &GetPostsResponse{}

	// field binding section.
	// (grpc.federation.field).by = "posts"
	if err := grpcfed.SetCELValue(ctx, value, "posts", func(v *Posts) { ret.Posts = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetPostsResponse", slog.Any("org.federation.GetPostsResponse", s.logvalue_Org_Federation_GetPostsResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Posts resolve "org.federation.Posts" message.
func (s *FederationService) resolve_Org_Federation_Posts(ctx context.Context, req *Org_Federation_PostsArgument[*FederationServiceDependentClientSet]) (*Posts, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.Posts")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.Posts", slog.Any("message_args", s.logvalue_Org_Federation_PostsArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			ids   []string
			posts []*post.Post
			res   *post.GetPostsResponse
			users []*User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.PostsArgument", req)}
	// A tree view of message dependencies is shown below.
	/*
	   res ─┐
	        posts ─┐
	                 ids ─┐
	   res ─┐             │
	        posts ─┐      │
	               users ─┤
	*/
	eg, ctx1 := errgroup.WithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     call {
		       method: "post.PostService/GetPosts"
		       request { field: "ids", by: "$.post_ids" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.GetPostsResponse, *localValueType]{
			Name:   "res",
			Type:   cel.ObjectType("post.GetPostsResponse"),
			Setter: func(value *localValueType, v *post.GetPostsResponse) { value.vars.res = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &post.GetPostsRequest{}
				// { field: "ids", by: "$.post_ids" }
				if err := grpcfed.SetCELValue(ctx, value, "$.post_ids", func(v []string) {
					args.Ids = v
				}); err != nil {
					return nil, err
				}
				return s.client.Post_PostServiceClient.GetPosts(ctx, args)
			},
		}); err != nil {
			if err := s.errorHandler(ctx1, FederationService_DependentMethod_Post_PostService_GetPosts, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx1, err)
				return nil, err
			}
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "posts"
		     by: "res.posts"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[[]*post.Post, *localValueType]{
			Name:   "posts",
			Type:   cel.ListType(cel.ObjectType("post.Post")),
			Setter: func(value *localValueType, v []*post.Post) { value.vars.posts = v },
			By:     "res.posts",
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "ids"
		     map {
		       iterator {
		         name: "post"
		         src: "posts"
		       }
		       by: "post.id"
		     }
		   }
		*/
		if err := grpcfed.EvalDefMap(ctx1, value, grpcfed.DefMap[[]string, *post.Post, *localValueType]{
			Name:           "ids",
			Type:           cel.ListType(celtypes.StringType),
			Setter:         func(value *localValueType, v []string) { value.vars.ids = v },
			IteratorName:   "post",
			IteratorType:   cel.ObjectType("post.Post"),
			IteratorSource: func(value *localValueType) []*post.Post { return value.vars.posts },
			Iterator: func(ctx context.Context, value *grpcfed.MapIteratorValue) (any, error) {
				return grpcfed.EvalCEL(ctx, value, "post.id", reflect.TypeOf(""))
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     call {
		       method: "post.PostService/GetPosts"
		       request { field: "ids", by: "$.post_ids" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.GetPostsResponse, *localValueType]{
			Name:   "res",
			Type:   cel.ObjectType("post.GetPostsResponse"),
			Setter: func(value *localValueType, v *post.GetPostsResponse) { value.vars.res = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &post.GetPostsRequest{}
				// { field: "ids", by: "$.post_ids" }
				if err := grpcfed.SetCELValue(ctx, value, "$.post_ids", func(v []string) {
					args.Ids = v
				}); err != nil {
					return nil, err
				}
				return s.client.Post_PostServiceClient.GetPosts(ctx, args)
			},
		}); err != nil {
			if err := s.errorHandler(ctx1, FederationService_DependentMethod_Post_PostService_GetPosts, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx1, err)
				return nil, err
			}
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "posts"
		     by: "res.posts"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[[]*post.Post, *localValueType]{
			Name:   "posts",
			Type:   cel.ListType(cel.ObjectType("post.Post")),
			Setter: func(value *localValueType, v []*post.Post) { value.vars.posts = v },
			By:     "res.posts",
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "users"
		     map {
		       iterator {
		         name: "iter"
		         src: "posts"
		       }
		       message {
		         name: "User"
		         args { name: "user_id", by: "iter.user_id" }
		       }
		     }
		   }
		*/
		if err := grpcfed.EvalDefMap(ctx1, value, grpcfed.DefMap[[]*User, *post.Post, *localValueType]{
			Name:           "users",
			Type:           cel.ListType(cel.ObjectType("org.federation.User")),
			Setter:         func(value *localValueType, v []*User) { value.vars.users = v },
			IteratorName:   "iter",
			IteratorType:   cel.ObjectType("post.Post"),
			IteratorSource: func(value *localValueType) []*post.Post { return value.vars.posts },
			Iterator: func(ctx context.Context, value *grpcfed.MapIteratorValue) (any, error) {
				args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				// { name: "user_id", by: "iter.user_id" }
				if err := grpcfed.SetCELValue(ctx, value, "iter.user_id", func(v string) {
					args.UserId = v
				}); err != nil {
					return nil, err
				}
				return s.resolve_Org_Federation_User(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Ids = value.vars.ids
	req.Posts = value.vars.posts
	req.Res = value.vars.res
	req.Users = value.vars.users

	// create a message value to be returned.
	ret := &Posts{}

	// field binding section.
	// (grpc.federation.field).by = "ids"
	if err := grpcfed.SetCELValue(ctx, value, "ids", func(v []string) { ret.Ids = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "posts.map(post, post.title)"
	if err := grpcfed.SetCELValue(ctx, value, "posts.map(post, post.title)", func(v []string) { ret.Titles = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "posts.map(post, post.content)"
	if err := grpcfed.SetCELValue(ctx, value, "posts.map(post, post.content)", func(v []string) { ret.Contents = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "users"
	if err := grpcfed.SetCELValue(ctx, value, "users", func(v []*User) { ret.Users = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved org.federation.Posts", slog.Any("org.federation.Posts", s.logvalue_Org_Federation_Posts(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			res  *user.GetUserResponse
			user *user.User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.UserArgument", req)}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "user.UserService/GetUser"
	       request { field: "id", by: "$.user_id" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.GetUserResponse, *localValueType]{
		Name:   "res",
		Type:   cel.ObjectType("user.GetUserResponse"),
		Setter: func(value *localValueType, v *user.GetUserResponse) { value.vars.res = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			if err := grpcfed.SetCELValue(ctx, value, "$.user_id", func(v string) {
				args.Id = v
			}); err != nil {
				return nil, err
			}
			return s.client.User_UserServiceClient.GetUser(ctx, args)
		},
	}); err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "user"
	     autobind: true
	     by: "res.user"
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.User, *localValueType]{
		Name:   "user",
		Type:   cel.ObjectType("user.User"),
		Setter: func(value *localValueType, v *user.User) { value.vars.user = v },
		By:     "res.user",
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Res = value.vars.res
	req.User = value.vars.user

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	ret.Id = value.vars.user.GetId()     // { name: "user", autobind: true }
	ret.Name = value.vars.user.GetName() // { name: "user", autobind: true }

	s.logger.DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetPostsResponse(v *GetPostsResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("posts", s.logvalue_Org_Federation_Posts(v.GetPosts())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetPostsResponseArgument(v *Org_Federation_GetPostsResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.Ids),
	)
}

func (s *FederationService) logvalue_Org_Federation_Posts(v *Posts) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
		slog.Any("titles", v.GetTitles()),
		slog.Any("contents", v.GetContents()),
		slog.Any("users", s.logvalue_repeated_Org_Federation_User(v.GetUsers())),
	)
}

func (s *FederationService) logvalue_Org_Federation_PostsArgument(v *Org_Federation_PostsArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post_ids", v.PostIds),
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
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_repeated_Org_Federation_User(v []*User) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for idx, vv := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(idx),
			Value: s.logvalue_Org_Federation_User(vv),
		})
	}
	return slog.GroupValue(attrs...)
}
