// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: openapi_rule.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/openapi_rule/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.openapi_rule",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType                  = reflect.TypeOf((*Client)(nil)).Elem()
	openapiRuleServiceClientType = reflect.TypeOf((*pb.OpenapiRuleServiceClient)(nil)).Elem()
	openapiRuleServiceServerType = reflect.TypeOf((*pb.OpenapiRuleServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.openapi_rule-client":
		return p.client
	case "erda.core.hepa.openapi_rule.OpenapiRuleService":
		return &openapiRuleServiceWrapper{client: p.client.OpenapiRuleService(), opts: opts}
	case "erda.core.hepa.openapi_rule.OpenapiRuleService.client":
		return p.client.OpenapiRuleService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case openapiRuleServiceClientType:
		return p.client.OpenapiRuleService()
	case openapiRuleServiceServerType:
		return &openapiRuleServiceWrapper{client: p.client.OpenapiRuleService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.openapi_rule-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.openapi_rule.OpenapiRuleService",
			"erda.core.hepa.openapi_rule-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			openapiRuleServiceClientType,
			// server types
			openapiRuleServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}