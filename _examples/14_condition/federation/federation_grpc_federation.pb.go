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
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	post "example/post"
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
	Posts  []*post.Post
	Res    *post.GetPostResponse
	User   *User
	Users  []*User
	Client T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
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
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
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
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*celtypes.FieldType{
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.UserArgument": {
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
		cel.Lib(grpcfedcel.NewLibrary()),
		cel.CrossTypeNumericComparisons(true),
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
		tracer:       otel.Tracer("org.federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
		},
	}, nil
}

// GetPost implements "org.federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/GetPost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
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
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetPostResponse resolve "org.federation.GetPostResponse" message.
func (s *FederationService) resolve_Org_Federation_GetPostResponse(ctx context.Context, req *Org_Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetPostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.GetPostResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetPostResponseArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *Post
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetPostResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "post"
	     message {
	       name: "Post"
	       args { name: "id", by: "$.id" }
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("post", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_PostArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "id", by: "$.id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					valueMu.RUnlock()
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Id = value.(string)
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_Post(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*Post)
		valueMu.Lock()
		valuePost = value
		envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("org.federation.Post")))
		evalValues["post"] = valuePost
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	{
		value, err := grpcfed.EvalCEL(s.env, "post", envOpts, evalValues, reflect.TypeOf((*Post)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Post = value.(*Post)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetPostResponse", slog.Any("org.federation.GetPostResponse", s.logvalue_Org_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Post resolve "org.federation.Post" message.
func (s *FederationService) resolve_Org_Federation_Post(ctx context.Context, req *Org_Federation_PostArgument[*FederationServiceDependentClientSet]) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.Post")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.Post", slog.Any("message_args", s.logvalue_Org_Federation_PostArgument(req)))
	var (
		sg         singleflight.Group
		valueMu    sync.RWMutex
		valuePost  *post.Post
		valuePosts []*post.Post
		valueRes   *post.GetPostResponse
		valueUser  *User
		valueUsers []*User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.PostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	eg, ctx1 := errgroup.WithContext(ctx)
	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     if: "$.id != ''"
		     call {
		       method: "post.PostService/GetPost"
		       request { field: "id", by: "$.id" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("res", func() (any, error) {
				valueMu.RLock()
				ifValue, err := grpcfed.EvalCEL(s.env, "$.id != ''", envOpts, evalValues, reflect.TypeOf(false))
				valueMu.RUnlock()
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				if !ifValue.(bool) {
					return (*post.GetPostResponse)(nil), nil
				}
				valueMu.RLock()
				args := &post.GetPostRequest{}
				// { field: "id", by: "$.id" }
				{
					value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
						valueMu.RUnlock()
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.Id = value.(string)
				}
				valueMu.RUnlock()
				return s.client.Post_PostServiceClient.GetPost(ctx1, args)
			})
			if err != nil {
				if err := s.errorHandler(ctx1, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
			}
			value := valueIface.(*post.GetPostResponse)
			valueMu.Lock()
			valueRes = value
			envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("post.GetPostResponse")))
			evalValues["res"] = valueRes
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "post"
		     if: "res != null"
		     by: "res.post"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("post", func() (any, error) {
				valueMu.RLock()
				ifValue, err := grpcfed.EvalCEL(s.env, "res != null", envOpts, evalValues, reflect.TypeOf(false))
				valueMu.RUnlock()
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				if !ifValue.(bool) {
					return (*post.Post)(nil), nil
				}
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "res.post", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*post.Post)
			valueMu.Lock()
			valuePost = value
			envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("post.Post")))
			evalValues["post"] = valuePost
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "posts"
		     by: "[post]"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("posts", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "[post]", envOpts, evalValues, reflect.TypeOf([]*post.Post(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.([]*post.Post)
			valueMu.Lock()
			valuePosts = value
			envOpts = append(envOpts, cel.Variable("posts", cel.ListType(cel.ObjectType("post.Post"))))
			evalValues["posts"] = valuePosts
			valueMu.Unlock()
		}
		return nil, nil
	})
	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     if: "$.id != ''"
		     call {
		       method: "post.PostService/GetPost"
		       request { field: "id", by: "$.id" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("res", func() (any, error) {
				valueMu.RLock()
				ifValue, err := grpcfed.EvalCEL(s.env, "$.id != ''", envOpts, evalValues, reflect.TypeOf(false))
				valueMu.RUnlock()
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				if !ifValue.(bool) {
					return (*post.GetPostResponse)(nil), nil
				}
				valueMu.RLock()
				args := &post.GetPostRequest{}
				// { field: "id", by: "$.id" }
				{
					value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
						valueMu.RUnlock()
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.Id = value.(string)
				}
				valueMu.RUnlock()
				return s.client.Post_PostServiceClient.GetPost(ctx1, args)
			})
			if err != nil {
				if err := s.errorHandler(ctx1, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
			}
			value := valueIface.(*post.GetPostResponse)
			valueMu.Lock()
			valueRes = value
			envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("post.GetPostResponse")))
			evalValues["res"] = valueRes
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "post"
		     if: "res != null"
		     by: "res.post"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("post", func() (any, error) {
				valueMu.RLock()
				ifValue, err := grpcfed.EvalCEL(s.env, "res != null", envOpts, evalValues, reflect.TypeOf(false))
				valueMu.RUnlock()
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				if !ifValue.(bool) {
					return (*post.Post)(nil), nil
				}
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "res.post", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*post.Post)
			valueMu.Lock()
			valuePost = value
			envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("post.Post")))
			evalValues["post"] = valuePost
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user"
		     if: "post != null"
		     message {
		       name: "User"
		       args { name: "user_id", by: "post.user_id" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("user", func() (any, error) {
				valueMu.RLock()
				ifValue, err := grpcfed.EvalCEL(s.env, "post != null", envOpts, evalValues, reflect.TypeOf(false))
				valueMu.RUnlock()
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				if !ifValue.(bool) {
					return (*User)(nil), nil
				}
				valueMu.RLock()
				args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				// { name: "user_id", by: "post.user_id" }
				{
					value, err := grpcfed.EvalCEL(s.env, "post.user_id", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
						valueMu.RUnlock()
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.UserId = value.(string)
				}
				valueMu.RUnlock()
				return s.resolve_Org_Federation_User(ctx1, args)
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*User)
			valueMu.Lock()
			valueUser = value
			envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("org.federation.User")))
			evalValues["user"] = valueUser
			valueMu.Unlock()
		}
		return nil, nil
	})
	if err := eg.Wait(); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "users"
	     if: "user != null"
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
	{
		valueIface, err, _ := sg.Do("users", func() (any, error) {
			valueMu.RLock()
			ifValue, err := grpcfed.EvalCEL(s.env, "user != null", envOpts, evalValues, reflect.TypeOf(false))
			valueMu.RUnlock()
			if err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
			if !ifValue.(bool) {
				return []*User(nil), nil
			}
			valueMu.RLock()
			valueMu.RUnlock()
			env, err := s.env.Extend(cel.Variable("iter", cel.ObjectType("post.Post")))
			if err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
			valueMu.RLock()
			defer valueMu.RUnlock()
			var value []*User
			for _, iter := range valuePosts {
				iterValues := make(map[string]any)
				for k, v := range evalValues {
					iterValues[k] = v
				}
				iterValues["iter"] = iter
				args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				// { name: "user_id", by: "iter.user_id" }
				{
					value, err := grpcfed.EvalCEL(env, "iter.user_id", envOpts, iterValues, reflect.TypeOf(""))
					if err != nil {
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.UserId = value.(string)
				}
				iterValue, err := s.resolve_Org_Federation_User(ctx, args)
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				value = append(value, iterValue)
			}
			return value, nil
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.([]*User)
		valueMu.Lock()
		valueUsers = value
		envOpts = append(envOpts, cel.Variable("users", cel.ListType(cel.ObjectType("org.federation.User"))))
		evalValues["users"] = valueUsers
		valueMu.Unlock()
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "_def5"
	     if: "users.size() > 0"
	     validation {
	       error {
	         code: INVALID_ARGUMENT
	         if: "users[0].id == ''"
	       }
	     }
	   }
	*/
	{
		if _, err, _ := sg.Do("_def5", func() (any, error) {
			valueMu.RLock()
			ifValue, err := grpcfed.EvalCEL(s.env, "users.size() > 0", envOpts, evalValues, reflect.TypeOf(false))
			valueMu.RUnlock()
			if err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
			if !ifValue.(bool) {
				return false, nil
			}
			{
				err := func() error {
					valueMu.RLock()
					value, err := grpcfed.EvalCEL(s.env, "users[0].id == ''", envOpts, evalValues, reflect.TypeOf(false))
					valueMu.RUnlock()
					if err != nil {
						return err
					}
					if value.(bool) {
						return grpcstatus.Error(grpccodes.InvalidArgument, "")
					}
					return nil
				}()
				if err != nil {
					if _, ok := grpcstatus.FromError(err); ok {
						return nil, err
					}
					s.logger.ErrorContext(ctx, "failed running validations", slog.String("error", err.Error()))
					return nil, grpcstatus.Errorf(grpccodes.Internal, "failed running validations: %s", err)
				}
				return nil, nil
			}
		}); err != nil {
			return nil, err
		}
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost
	req.Posts = valuePosts
	req.Res = valueRes
	req.User = valueUser
	req.Users = valueUsers

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	// (grpc.federation.field).by = "post.id"
	{
		value, err := grpcfed.EvalCEL(s.env, "post.id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Id = value.(string)
	}
	// (grpc.federation.field).by = "post.title"
	{
		value, err := grpcfed.EvalCEL(s.env, "post.title", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Title = value.(string)
	}
	// (grpc.federation.field).by = "users[0]"
	{
		value, err := grpcfed.EvalCEL(s.env, "users[0]", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.User = value.(*User)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.Post", slog.Any("org.federation.Post", s.logvalue_Org_Federation_Post(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.user_id"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Id = value.(string)
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
