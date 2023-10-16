// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	post "example/post"
)

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler FederationServiceErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// Org_Post_PostServiceClient create a gRPC Client to be used to call methods in org.post.PostService.
	Org_Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
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

// FederationServiceDependencyServiceClient has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependencyServiceClient struct {
	Org_Post_PostServiceClient post.PostServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// FederationServiceErrorHandler Federation Service often needs to convert errors received from downstream services.
// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
type FederationServiceErrorHandler func(ctx context.Context, methodName string, err error) error

const (
	FederationService_DependentMethod_Org_Post_PostService_CreatePost = "/org.post.PostService/CreatePost"
)

// FederationServiceRecoveredError represents recovered error.
type FederationServiceRecoveredError struct {
	Message string
	Stack   []string
}

func (e *FederationServiceRecoveredError) Error() string {
	return fmt.Sprintf("recovered error: %s", e.Message)
}

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler FederationServiceErrorHandler
	env          *cel.Env
	client       *FederationServiceDependencyServiceClient
}

// Org_Federation_CreatePostArgument is argument for "org.federation.CreatePost" message.
type Org_Federation_CreatePostArgument struct {
	Content string
	Title   string
	UserId  string
	Client  *FederationServiceDependencyServiceClient
}

// Org_Federation_CreatePostResponseArgument is argument for "org.federation.CreatePostResponse" message.
type Org_Federation_CreatePostResponseArgument struct {
	Content string
	Cp      *CreatePost
	P       *post.Post
	Title   string
	UserId  string
	Client  *FederationServiceDependencyServiceClient
}

// FederationServiceCELTypeHelper
type FederationServiceCELTypeHelper struct {
	celRegistry    *celtypes.Registry
	structFieldMap map[string]map[string]*celtypes.FieldType
	mapMu          sync.RWMutex
}

func (h *FederationServiceCELTypeHelper) TypeProvider() celtypes.Provider {
	return h
}

func (h *FederationServiceCELTypeHelper) TypeAdapter() celtypes.Adapter {
	return h.celRegistry
}

func (h *FederationServiceCELTypeHelper) EnumValue(enumName string) ref.Val {
	return h.celRegistry.EnumValue(enumName)
}

func (h *FederationServiceCELTypeHelper) FindIdent(identName string) (ref.Val, bool) {
	return h.celRegistry.FindIdent(identName)
}

func (h *FederationServiceCELTypeHelper) FindStructType(structType string) (*celtypes.Type, bool) {
	if st, found := h.celRegistry.FindStructType(structType); found {
		return st, found
	}
	h.mapMu.RLock()
	defer h.mapMu.RUnlock()
	if _, exists := h.structFieldMap[structType]; exists {
		return celtypes.NewObjectType(structType), true
	}
	return nil, false
}

func (h *FederationServiceCELTypeHelper) FindStructFieldNames(structType string) ([]string, bool) {
	if names, found := h.celRegistry.FindStructFieldNames(structType); found {
		return names, found
	}

	h.mapMu.RLock()
	defer h.mapMu.RUnlock()
	fieldMap, exists := h.structFieldMap[structType]
	if !exists {
		return nil, false
	}
	fieldNames := make([]string, 0, len(fieldMap))
	for fieldName := range fieldMap {
		fieldNames = append(fieldNames, fieldName)
	}
	sort.Strings(fieldNames)
	return fieldNames, true
}

func (h *FederationServiceCELTypeHelper) FindStructFieldType(structType, fieldName string) (*celtypes.FieldType, bool) {
	if field, found := h.celRegistry.FindStructFieldType(structType, fieldName); found {
		return field, found
	}

	h.mapMu.RLock()
	defer h.mapMu.RUnlock()
	fieldMap, exists := h.structFieldMap[structType]
	if !exists {
		return nil, false
	}
	field, found := fieldMap[fieldName]
	return field, found
}

func (h *FederationServiceCELTypeHelper) NewValue(structType string, fields map[string]ref.Val) ref.Val {
	return h.celRegistry.NewValue(structType, fields)
}

func newFederationServiceCELTypeHelper() *FederationServiceCELTypeHelper {
	celRegistry := celtypes.NewEmptyRegistry()
	protoregistry.GlobalFiles.RangeFiles(func(f protoreflect.FileDescriptor) bool {
		if err := celRegistry.RegisterDescriptor(f); err != nil {
			return false
		}
		return true
	})
	newFieldType := func(typ *celtypes.Type, fieldName string) *celtypes.FieldType {
		isSet := func(v any, fieldName string) bool {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
			}
			if rv.Kind() != reflect.Struct {
				return false
			}
			return rv.FieldByName(fieldName).IsValid()
		}
		getFrom := func(v any, fieldName string) (any, error) {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
			}
			if rv.Kind() != reflect.Struct {
				return nil, fmt.Errorf("%T is not struct type", v)
			}
			value := rv.FieldByName(fieldName)
			return value.Interface(), nil
		}
		return &celtypes.FieldType{
			Type: typ,
			IsSet: func(v any) bool {
				return isSet(v, fieldName)
			},
			GetFrom: func(v any) (any, error) {
				return getFrom(v, fieldName)
			},
		}
	}
	return &FederationServiceCELTypeHelper{
		celRegistry: celRegistry,
		structFieldMap: map[string]map[string]*celtypes.FieldType{
			"grpc.federation.private.CreatePostArgument": map[string]*celtypes.FieldType{
				"title":   newFieldType(celtypes.StringType, "Title"),
				"content": newFieldType(celtypes.StringType, "Content"),
				"user_id": newFieldType(celtypes.StringType, "UserId"),
			},
			"grpc.federation.private.CreatePostResponseArgument": map[string]*celtypes.FieldType{
				"title":   newFieldType(celtypes.StringType, "Title"),
				"content": newFieldType(celtypes.StringType, "Content"),
				"user_id": newFieldType(celtypes.StringType, "UserId"),
			},
		},
	}
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if err := validateFederationServiceConfig(cfg); err != nil {
		return nil, err
	}
	Org_Post_PostServiceClient, err := cfg.Client.Org_Post_PostServiceClient(FederationServiceClientConfig{
		Service: "org.post.PostService",
		Name:    "post_service",
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
	celHelper := newFederationServiceCELTypeHelper()
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
		client: &FederationServiceDependencyServiceClient{
			Org_Post_PostServiceClient: Org_Post_PostServiceClient,
		},
	}, nil
}

func validateFederationServiceConfig(cfg FederationServiceConfig) error {
	if cfg.Client == nil {
		return fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	return nil
}

func withTimeoutFederationService[T any](ctx context.Context, method string, timeout time.Duration, fn func(context.Context) (*T, error)) (*T, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var (
		ret   *T
		errch = make(chan error)
	)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errch <- recoverErrorFederationService(r, debug.Stack())
			}
		}()

		res, err := fn(ctx)
		ret = res
		errch <- err
	}()
	select {
	case <-ctx.Done():
		status := grpcstatus.New(grpccodes.DeadlineExceeded, ctx.Err().Error())
		withDetails, err := status.WithDetails(&errdetails.ErrorInfo{
			Metadata: map[string]string{
				"method":  method,
				"timeout": timeout.String(),
			},
		})
		if err != nil {
			return nil, status.Err()
		}
		return nil, withDetails.Err()
	case err := <-errch:
		return ret, err
	}
}

func withRetryFederationService[T any](b backoff.BackOff, fn func() (*T, error)) (*T, error) {
	var res *T
	if err := backoff.Retry(func() (err error) {
		res, err = fn()
		return
	}, b); err != nil {
		return nil, err
	}
	return res, nil
}

func recoverErrorFederationService(v interface{}, rawStack []byte) *FederationServiceRecoveredError {
	msg := fmt.Sprint(v)
	lines := strings.Split(msg, "\n")
	if len(lines) <= 1 {
		lines := strings.Split(string(rawStack), "\n")
		stack := make([]string, 0, len(lines))
		for _, line := range lines {
			if line == "" {
				continue
			}
			stack = append(stack, strings.TrimPrefix(line, "\t"))
		}
		return &FederationServiceRecoveredError{
			Message: msg,
			Stack:   stack,
		}
	}
	// If panic occurs under singleflight, singleflight's recover catches the error and gives a stack trace.
	// Therefore, once the stack trace is removed.
	stack := make([]string, 0, len(lines))
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		stack = append(stack, strings.TrimPrefix(line, "\t"))
	}
	return &FederationServiceRecoveredError{
		Message: lines[0],
		Stack:   stack,
	}
}

func (s *FederationService) evalCEL(expr string, vars []cel.EnvOption, args map[string]any, outType reflect.Type) (any, error) {
	env, err := s.env.Extend(vars...)
	if err != nil {
		return nil, err
	}
	expr = strings.Replace(expr, "$", grpcfed.MessageArgumentVariableName, -1)
	ast, iss := env.Compile(expr)
	if iss.Err() != nil {
		return nil, iss.Err()
	}
	program, err := env.Program(ast)
	if err != nil {
		return nil, err
	}
	out, _, err := program.Eval(args)
	if err != nil {
		return nil, err
	}
	if outType != nil {
		return out.ConvertToNative(outType)
	}
	return out.Value(), nil
}

func (s *FederationService) goWithRecover(eg *errgroup.Group, fn func() (interface{}, error)) {
	eg.Go(func() (e error) {
		defer func() {
			if r := recover(); r != nil {
				e = recoverErrorFederationService(r, debug.Stack())
			}
		}()
		_, err := fn()
		return err
	})
}

func (s *FederationService) outputErrorLog(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if status, ok := grpcstatus.FromError(err); ok {
		s.logger.ErrorContext(ctx, status.Message(),
			slog.Group("grpc_status",
				slog.String("code", status.Code().String()),
				slog.Any("details", status.Details()),
			),
		)
		return
	}
	var recoveredErr *FederationServiceRecoveredError
	if errors.As(err, &recoveredErr) {
		trace := make([]interface{}, 0, len(recoveredErr.Stack))
		for idx, stack := range recoveredErr.Stack {
			trace = append(trace, slog.String(fmt.Sprint(idx+1), stack))
		}
		s.logger.ErrorContext(ctx, recoveredErr.Message, slog.Group("stack_trace", trace...))
		return
	}
	s.logger.ErrorContext(ctx, err.Error())
}

// CreatePost implements "org.federation.FederationService/CreatePost" method.
func (s *FederationService) CreatePost(ctx context.Context, req *CreatePostRequest) (res *CreatePostResponse, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = recoverErrorFederationService(r, debug.Stack())
			s.outputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Org_Federation_CreatePostResponse(ctx, &Org_Federation_CreatePostResponseArgument{
		Client:  s.client,
		Title:   req.Title,
		Content: req.Content,
		UserId:  req.UserId,
	})
	if err != nil {
		s.outputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_CreatePost resolve "org.federation.CreatePost" message.
func (s *FederationService) resolve_Org_Federation_CreatePost(ctx context.Context, req *Org_Federation_CreatePostArgument) (*CreatePost, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.CreatePost", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostArgument(req)))
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.CreatePostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// create a message value to be returned.
	ret := &CreatePost{}

	// field binding section.
	// (grpc.federation.field).by = "$.title"
	{
		_value, err := s.evalCEL("$.title", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			return nil, err
		}
		ret.Title = _value.(string)
	}
	// (grpc.federation.field).by = "$.content"
	{
		_value, err := s.evalCEL("$.content", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			return nil, err
		}
		ret.Content = _value.(string)
	}
	// (grpc.federation.field).by = "$.user_id"
	{
		_value, err := s.evalCEL("$.user_id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			return nil, err
		}
		ret.UserId = _value.(string)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePost", slog.Any("org.federation.CreatePost", s.logvalue_Org_Federation_CreatePost(ret)))
	return ret, nil
}

// resolve_Org_Federation_CreatePostResponse resolve "org.federation.CreatePostResponse" message.
func (s *FederationService) resolve_Org_Federation_CreatePostResponse(ctx context.Context, req *Org_Federation_CreatePostResponseArgument) (*CreatePostResponse, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.CreatePostResponse", slog.Any("message_args", s.logvalue_Org_Federation_CreatePostResponseArgument(req)))
	var (
		sg      singleflight.Group
		valueCp *CreatePost
		valueMu sync.RWMutex
		valueP  *post.Post
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.CreatePostResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "cp"
	     message: "CreatePost"
	     args: [
	       { name: "title", by: "$.title" },
	       { name: "content", by: "$.content" },
	       { name: "user_id", by: "$.user_id" }
	     ]
	   }
	*/
	resCreatePostIface, err, _ := sg.Do("cp_org.federation.CreatePost", func() (interface{}, error) {
		valueMu.RLock()
		args := &Org_Federation_CreatePostArgument{
			Client: s.client,
		}
		// { name: "title", by: "$.title" }
		{
			_value, err := s.evalCEL("$.title", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Title = _value.(string)
		}
		// { name: "content", by: "$.content" }
		{
			_value, err := s.evalCEL("$.content", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Content = _value.(string)
		}
		// { name: "user_id", by: "$.user_id" }
		{
			_value, err := s.evalCEL("$.user_id", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.UserId = _value.(string)
		}
		valueMu.RUnlock()
		return s.resolve_Org_Federation_CreatePost(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resCreatePost := resCreatePostIface.(*CreatePost)
	valueMu.Lock()
	valueCp = resCreatePost // { name: "cp", message: "CreatePost" ... }
	envOpts = append(envOpts, cel.Variable("cp", cel.ObjectType("org.federation.CreatePost")))
	evalValues["cp"] = valueCp
	valueMu.Unlock()

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "org.post.PostService/CreatePost"
	     request { field: "post", by: "cp" }
	     response { name: "p", field: "post" }
	   }
	*/
	resCreatePostResponseIface, err, _ := sg.Do("org.post.PostService/CreatePost", func() (interface{}, error) {
		valueMu.RLock()
		args := &post.CreatePostRequest{}
		// { field: "post", by: "cp" }
		{
			_value, err := s.evalCEL("cp", envOpts, evalValues, reflect.TypeOf((*CreatePost)(nil)))
			if err != nil {
				return nil, err
			}
			args.Post = s.cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(_value.(*CreatePost))
		}
		valueMu.RUnlock()
		return s.client.Org_Post_PostServiceClient.CreatePost(ctx, args)
	})
	if err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_Org_Post_PostService_CreatePost, err); err != nil {
			return nil, err
		}
	}
	resCreatePostResponse := resCreatePostResponseIface.(*post.CreatePostResponse)
	valueMu.Lock()
	valueP = resCreatePostResponse.GetPost() // { name: "p", field: "post" }
	envOpts = append(envOpts, cel.Variable("p", cel.ObjectType("org.post.Post")))
	evalValues["p"] = valueP
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Cp = valueCp
	req.P = valueP

	// create a message value to be returned.
	ret := &CreatePostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "p"
	{
		_value, err := s.evalCEL("p", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
		if err != nil {
			return nil, err
		}
		ret.Post = s.cast_Org_Post_Post__to__Org_Federation_Post(_value.(*post.Post))
	}

	s.logger.DebugContext(ctx, "resolved org.federation.CreatePostResponse", slog.Any("org.federation.CreatePostResponse", s.logvalue_Org_Federation_CreatePostResponse(ret)))
	return ret, nil
}

// cast_Org_Federation_CreatePost__to__Org_Post_CreatePost cast from "org.federation.CreatePost" to "org.post.CreatePost".
func (s *FederationService) cast_Org_Federation_CreatePost__to__Org_Post_CreatePost(from *CreatePost) *post.CreatePost {
	if from == nil {
		return nil
	}

	return &post.CreatePost{
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

// cast_Org_Post_Post__to__Org_Federation_Post cast from "org.post.Post" to "org.federation.Post".
func (s *FederationService) cast_Org_Post_Post__to__Org_Federation_Post(from *post.Post) *Post {
	if from == nil {
		return nil
	}

	return &Post{
		Id:      from.GetId(),
		Title:   from.GetTitle(),
		Content: from.GetContent(),
		UserId:  from.GetUserId(),
	}
}

func (s *FederationService) logvalue_Org_Federation_CreatePost(v *CreatePost) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.String("user_id", v.GetUserId()),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostArgument(v *Org_Federation_CreatePostArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponse(v *CreatePostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Org_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Org_Federation_CreatePostResponseArgument(v *Org_Federation_CreatePostResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
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
		slog.String("user_id", v.GetUserId()),
	)
}
