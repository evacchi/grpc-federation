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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	user "example/user"
)

// Org_Federation_GetResponseArgument is argument for "org.federation.GetResponse" message.
type Org_Federation_GetResponseArgument[T any] struct {
	Uid    *UserID
	User   *User
	User2  *User
	Client T
}

// Org_Federation_SubArgument is argument for "org.federation.Sub" message.
type Org_Federation_SubArgument[T any] struct {
	Client T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
	Res    *user.GetUserResponse
	User   *user.User
	UserId string
	XDef2  *Sub
	Client T
}

// Org_Federation_UserIDArgument is argument for "org.federation.UserID" message.
type Org_Federation_UserIDArgument[T any] struct {
	Client T
}

// Org_Federation_User_NameArgument is custom resolver's argument for "name" field of "org.federation.User" message.
type Org_Federation_User_NameArgument[T any] struct {
	*Org_Federation_UserArgument[T]
	Client T
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
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
	// Resolve_Org_Federation_Sub implements resolver for "org.federation.Sub".
	Resolve_Org_Federation_Sub(context.Context, *Org_Federation_SubArgument[*FederationServiceDependentClientSet]) (*Sub, error)
	// Resolve_Org_Federation_User_Name implements resolver for "org.federation.User.name".
	Resolve_Org_Federation_User_Name(context.Context, *Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]) (string, error)
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// Resolve_Org_Federation_Sub resolve "org.federation.Sub".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_Sub(context.Context, *Org_Federation_SubArgument[*FederationServiceDependentClientSet]) (ret *Sub, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Org_Federation_Sub not implemented")
	return
}

// Resolve_Org_Federation_User_Name resolve "org.federation.User.name".
// This method always returns Unimplemented error.
func (FederationServiceUnimplementedResolver) Resolve_Org_Federation_User_Name(context.Context, *Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]) (ret string, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Org_Federation_User_Name not implemented")
	return
}

const (
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
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
		"grpc.federation.private.GetResponseArgument": {},
		"grpc.federation.private.SubArgument":         {},
		"grpc.federation.private.UserArgument": {
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
		"grpc.federation.private.UserIDArgument": {},
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
		tracer:       otel.Tracer("org.federation.FederationService"),
		resolver:     cfg.Resolver,
		client: &FederationServiceDependentClientSet{
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// Get implements "org.federation.FederationService/Get" method.
func (s *FederationService) Get(ctx context.Context, req *GetRequest) (res *GetResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.FederationService/Get")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetResponse(ctx, &Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]{
		Client: s.client,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetResponse resolve "org.federation.GetResponse" message.
func (s *FederationService) resolve_Org_Federation_GetResponse(ctx context.Context, req *Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]) (*GetResponse, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.GetResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.GetResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetResponseArgument(req)))
	var (
		sg         singleflight.Group
		valueMu    sync.RWMutex
		valueUid   *UserID
		valueUser  *User
		valueUser2 *User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}
	// A tree view of message dependencies is shown below.
	/*
	   uid ─┐
	         user ─┐
	   uid ─┐      │
	        user2 ─┤
	*/
	eg, ctx1 := errgroup.WithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "uid"
		     message {
		       name: "UserID"
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("uid", func() (any, error) {
				valueMu.RLock()
				args := &Org_Federation_UserIDArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				valueMu.RUnlock()
				return s.resolve_Org_Federation_UserID(ctx1, args)
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*UserID)
			valueMu.Lock()
			valueUid = value
			envOpts = append(envOpts, cel.Variable("uid", cel.ObjectType("org.federation.UserID")))
			evalValues["uid"] = valueUid
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user"
		     message {
		       name: "User"
		       args { name: "user_id", by: "uid.value" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("user", func() (any, error) {
				valueMu.RLock()
				args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				// { name: "user_id", by: "uid.value" }
				{
					value, err := grpcfed.EvalCEL(s.env, "uid.value", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
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

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "uid"
		     message {
		       name: "UserID"
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("uid", func() (any, error) {
				valueMu.RLock()
				args := &Org_Federation_UserIDArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				valueMu.RUnlock()
				return s.resolve_Org_Federation_UserID(ctx1, args)
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*UserID)
			valueMu.Lock()
			valueUid = value
			envOpts = append(envOpts, cel.Variable("uid", cel.ObjectType("org.federation.UserID")))
			evalValues["uid"] = valueUid
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user2"
		     message {
		       name: "User"
		       args { name: "user_id", by: "uid.value" }
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("user2", func() (any, error) {
				valueMu.RLock()
				args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				// { name: "user_id", by: "uid.value" }
				{
					value, err := grpcfed.EvalCEL(s.env, "uid.value", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
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
			valueUser2 = value
			envOpts = append(envOpts, cel.Variable("user2", cel.ObjectType("org.federation.User")))
			evalValues["user2"] = valueUser2
			valueMu.Unlock()
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Uid = valueUid
	req.User = valueUser
	req.User2 = valueUser2

	// create a message value to be returned.
	ret := &GetResponse{}

	// field binding section.
	// (grpc.federation.field).by = "user"
	{
		value, err := grpcfed.EvalCEL(s.env, "user", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.User = value.(*User)
	}
	// (grpc.federation.field).by = "user2"
	{
		value, err := grpcfed.EvalCEL(s.env, "user2", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.User2 = value.(*User)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetResponse", slog.Any("org.federation.GetResponse", s.logvalue_Org_Federation_GetResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_Sub resolve "org.federation.Sub" message.
func (s *FederationService) resolve_Org_Federation_Sub(ctx context.Context, req *Org_Federation_SubArgument[*FederationServiceDependentClientSet]) (*Sub, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.Sub")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.Sub", slog.Any("message_args", s.logvalue_Org_Federation_SubArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ret, err := s.resolver.Resolve_Org_Federation_Sub(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved org.federation.Sub", slog.Any("org.federation.Sub", s.logvalue_Org_Federation_Sub(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	var (
		sg         singleflight.Group
		valueMu    sync.RWMutex
		valueRes   *user.GetUserResponse
		valueUser  *user.User
		value_Def2 *Sub
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}
	// A tree view of message dependencies is shown below.
	/*
	        _def2 ─┐
	   res ─┐      │
	         user ─┤
	*/
	eg, ctx1 := errgroup.WithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def2"
		     message {
		       name: "Sub"
		     }
		   }
		*/
		{
			valueIface, err, _ := sg.Do("_def2", func() (any, error) {
				valueMu.RLock()
				args := &Org_Federation_SubArgument[*FederationServiceDependentClientSet]{
					Client: s.client,
				}
				valueMu.RUnlock()
				return s.resolve_Org_Federation_Sub(ctx1, args)
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*Sub)
			valueMu.Lock()
			value_Def2 = value
			envOpts = append(envOpts, cel.Variable("_def2", cel.ObjectType("org.federation.Sub")))
			evalValues["_def2"] = value_Def2
			valueMu.Unlock()
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

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
		{
			valueIface, err, _ := sg.Do("res", func() (any, error) {
				valueMu.RLock()
				args := &user.GetUserRequest{}
				// { field: "id", by: "$.user_id" }
				{
					value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
					if err != nil {
						grpcfed.RecordErrorToSpan(ctx, err)
						return nil, err
					}
					args.Id = value.(string)
				}
				valueMu.RUnlock()
				return s.client.User_UserServiceClient.GetUser(ctx1, args)
			})
			if err != nil {
				if err := s.errorHandler(ctx1, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
			}
			value := valueIface.(*user.GetUserResponse)
			valueMu.Lock()
			valueRes = value
			envOpts = append(envOpts, cel.Variable("res", cel.ObjectType("user.GetUserResponse")))
			evalValues["res"] = valueRes
			valueMu.Unlock()
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user"
		     autobind: true
		     by: "res.user"
		   }
		*/
		{
			valueIface, err, _ := sg.Do("user", func() (any, error) {
				valueMu.RLock()
				valueMu.RUnlock()
				valueMu.RLock()
				value, err := grpcfed.EvalCEL(s.env, "res.user", envOpts, evalValues, reflect.TypeOf((*user.User)(nil)))
				valueMu.RUnlock()
				return value, err
			})
			if err != nil {
				return nil, err
			}
			value := valueIface.(*user.User)
			valueMu.Lock()
			valueUser = value
			envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("user.User")))
			evalValues["user"] = valueUser
			valueMu.Unlock()
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Res = valueRes
	req.User = valueUser
	req.XDef2 = value_Def2

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	ret.Id = valueUser.GetId() // { name: "user", autobind: true }
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.Name, err = s.resolver.Resolve_Org_Federation_User_Name(ctx, &Org_Federation_User_NameArgument[*FederationServiceDependentClientSet]{
			Client:                      s.client,
			Org_Federation_UserArgument: req,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

// resolve_Org_Federation_UserID resolve "org.federation.UserID" message.
func (s *FederationService) resolve_Org_Federation_UserID(ctx context.Context, req *Org_Federation_UserIDArgument[*FederationServiceDependentClientSet]) (*UserID, error) {
	ctx, span := s.tracer.Start(ctx, "org.federation.UserID")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve org.federation.UserID", slog.Any("message_args", s.logvalue_Org_Federation_UserIDArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
	)

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "_def0"
	     message {
	       name: "Sub"
	     }
	   }
	*/
	{
		if _, err, _ := sg.Do("_def0", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_SubArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_Sub(ctx, args)
		}); err != nil {
			return nil, err
		}
	}

	// create a message value to be returned.
	ret := &UserID{}

	// field binding section.
	ret.Value = "xxx" // (grpc.federation.field).string = "xxx"

	s.logger.DebugContext(ctx, "resolved org.federation.UserID", slog.Any("org.federation.UserID", s.logvalue_Org_Federation_UserID(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetResponse(v *GetResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user", s.logvalue_Org_Federation_User(v.GetUser())),
		slog.Any("user2", s.logvalue_Org_Federation_User(v.GetUser2())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetResponseArgument(v *Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_Sub(v *Sub) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_SubArgument(v *Org_Federation_SubArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
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

func (s *FederationService) logvalue_Org_Federation_UserID(v *UserID) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("value", v.GetValue()),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserIDArgument(v *Org_Federation_UserIDArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}
