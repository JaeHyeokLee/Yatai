package services

import (
	"context"

	"github.com/bentoml/yatai/schemas/modelschemas"

	"github.com/pkg/errors"

	"github.com/bentoml/yatai/api-server/models"
)

type resourceService struct{}

var ResourceService = resourceService{}

func (m *resourceService) Get(ctx context.Context, resourceType modelschemas.ResourceType, resourceId uint) (models.IResource, error) {
	switch resourceType {
	case modelschemas.ResourceTypeUser:
		user, err := UserService.Get(ctx, resourceId)
		return user, err
	case modelschemas.ResourceTypeOrganization:
		org, err := OrganizationService.Get(ctx, resourceId)
		return org, err
	case modelschemas.ResourceTypeCluster:
		cluster, err := ClusterService.Get(ctx, resourceId)
		return cluster, err
	case modelschemas.ResourceTypeBento:
		bento, err := BentoService.Get(ctx, resourceId)
		return bento, err
	case modelschemas.ResourceTypeBentoVersion:
		bentoVersion, err := BentoVersionService.Get(ctx, resourceId)
		return bentoVersion, err
	case modelschemas.ResourceTypeDeployment:
		deployment, err := DeploymentService.Get(ctx, resourceId)
		return deployment, err
	case modelschemas.ResourceTypeDeploymentRevision:
		deploymentRevision, err := DeploymentRevisionService.Get(ctx, resourceId)
		return deploymentRevision, err
	case modelschemas.ResourceTypeTerminalRecord:
		terminalRecord, err := TerminalRecordService.Get(ctx, resourceId)
		return terminalRecord, err
	case modelschemas.ResourceTypeModel:
		model, err := ModelService.Get(ctx, resourceId)
		return model, err
	case modelschemas.ResourceTypeModelVersion:
		modelVersion, err := ModelVersionService.Get(ctx, resourceId)
		return modelVersion, err
	default:
		return nil, errors.Errorf("cannot recognize this resource type: %s", resourceType)
	}
}
