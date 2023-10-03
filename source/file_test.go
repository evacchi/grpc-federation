package source_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/mercari/grpc-federation/source"
)

func TestFile_FindLocationByPos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc     string
		pos      source.Position
		expected *source.Location
	}{
		{
			desc: "find nothing",
			pos: source.Position{
				Line: 21,
				Col:  1,
			},
			expected: nil,
		},
		{
			desc: "find a ServiceDependencyOption name",
			pos: source.Position{
				Line: 15,
				Col:  15,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Service: &source.Service{
					Name: "FederationService",
					Option: &source.ServiceOption{
						Dependencies: &source.ServiceDependencyOption{Name: true},
					},
				},
			},
		},
		{
			desc: "find a ServiceDependencyOption name",
			pos: source.Position{
				Line: 15,
				Col:  15,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Service: &source.Service{
					Name: "FederationService",
					Option: &source.ServiceOption{
						Dependencies: &source.ServiceDependencyOption{Name: true},
					},
				},
			},
		},
		{
			desc: "find a ServiceDependencyOption service",
			pos: source.Position{
				Line: 15,
				Col:  32,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Service: &source.Service{
					Name: "FederationService",
					Option: &source.ServiceOption{
						Dependencies: &source.ServiceDependencyOption{Service: true},
					},
				},
			},
		},
		{
			desc: "find a MethodOption string timeout",
			pos: source.Position{
				Line: 19,
				Col:  47,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Service: &source.Service{
					Name: "FederationService",
					Method: &source.Method{
						Name:   "GetPost",
						Option: &source.MethodOption{Timeout: true},
					},
				},
			},
		},
		{
			desc: "find a MethodOption message timeout",
			pos: source.Position{
				Line: 23,
				Col:  16,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Service: &source.Service{
					Name: "FederationService",
					Method: &source.Method{
						Name:   "GetPost2",
						Option: &source.MethodOption{Timeout: true},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption method",
			pos: source.Position{
				Line: 44,
				Col:  15,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{Method: true},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption request field",
			pos: source.Position{
				Line: 46,
				Col:  18,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{
							Request: &source.RequestOption{Field: true},
						},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption request by",
			pos: source.Position{
				Line: 46,
				Col:  28,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{
							Request: &source.RequestOption{By: true},
						},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption response name",
			pos: source.Position{
				Line: 48,
				Col:  27,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{
							Response: &source.ResponseOption{Name: true},
						},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption response field",
			pos: source.Position{
				Line: 48,
				Col:  42,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{
							Response: &source.ResponseOption{Field: true},
						},
					},
				},
			},
		},
		{
			desc: "find a ResolverOption response autobind",
			pos: source.Position{
				Line: 48,
				Col:  60,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Resolver: &source.ResolverOption{
							Response: &source.ResponseOption{AutoBind: true},
						},
					},
				},
			},
		},
		{
			desc: "find a FieldOption by",
			pos: source.Position{
				Line: 38,
				Col:  50,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "GetPostResponse",
					Field: &source.Field{
						Name:   "post",
						Option: &source.FieldOption{By: true},
					},
				},
			},
		},
		{
			desc: "find a MessageDependencyOption name",
			pos: source.Position{
				Line: 35,
				Col:  15,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "GetPostResponse",
					Option: &source.MessageOption{
						Messages: &source.MessageDependencyOption{
							Idx:  0,
							Name: true,
						},
					},
				},
			},
		},
		{
			desc: "find a MessageDependencyOption message",
			pos: source.Position{
				Line: 35,
				Col:  32,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "GetPostResponse",
					Option: &source.MessageOption{
						Messages: &source.MessageDependencyOption{
							Idx:     0,
							Message: true,
						},
					},
				},
			},
		},
		{
			desc: "find a MessageDependencyOption arg name",
			pos: source.Position{
				Line: 35,
				Col:  55,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "GetPostResponse",
					Option: &source.MessageOption{
						Messages: &source.MessageDependencyOption{
							Idx:  0,
							Args: &source.ArgumentOption{Name: true},
						},
					},
				},
			},
		},
		{
			desc: "find a MessageDependencyOption arg by",
			pos: source.Position{
				Line: 35,
				Col:  65,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "GetPostResponse",
					Option: &source.MessageOption{
						Messages: &source.MessageDependencyOption{
							Idx:  0,
							Args: &source.ArgumentOption{By: true},
						},
					},
				},
			},
		},
		{
			desc: "find a MessageDependencyOption arg inline",
			pos: source.Position{
				Line: 51,
				Col:  57,
			},
			expected: &source.Location{
				FileName: filepath.Base("testdata/service.proto"),
				Message: &source.Message{
					Name: "Post",
					Option: &source.MessageOption{
						Messages: &source.MessageDependencyOption{
							Idx:  0,
							Args: &source.ArgumentOption{Inline: true},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			path := filepath.Join("testdata", "service.proto")
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			sourceFile, err := source.NewFile(path, content)
			if err != nil {
				t.Fatal(err)
			}
			got := sourceFile.FindLocationByPos(tc.pos)
			if diff := cmp.Diff(got, tc.expected); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
