// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: project.proto

package pb

import (
	fmt "fmt"
	math "math"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetProjectStatisticsRequest) Validate() error {
	return nil
}
func (this *GetProjectStatisticsResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ProjectStatistics) Validate() error {
	return nil
}
func (this *GetProjectOverviewRequest) Validate() error {
	return nil
}
func (this *GetProjectOverviewResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ProjectOverviewList) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ProjectOverview) Validate() error {
	return nil
}
func (this *GetProjectRequest) Validate() error {
	if this.ProjectID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProjectID))
	}
	return nil
}
func (this *GetProjectResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetProjectsRequest) Validate() error {
	return nil
}
func (this *GetProjectsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetProjectsTenantsIDsRequest) Validate() error {
	return nil
}
func (this *GetProjectsTenantsIDsResponse) Validate() error {
	return nil
}
func (this *DeleteProjectRequest) Validate() error {
	return nil
}
func (this *DeleteProjectResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *CreateProjectRequest) Validate() error {
	return nil
}
func (this *CreateProjectResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UpdateProjectRequest) Validate() error {
	return nil
}
func (this *UpdateProjectResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Project) Validate() error {
	for _, item := range this.Relationship {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Relationship", err)
			}
		}
	}
	return nil
}
func (this *TenantRelationship) Validate() error {
	return nil
}
