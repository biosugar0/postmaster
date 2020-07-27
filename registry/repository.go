// Package registry define repository for resolve dependency injection
package registry

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/biosugar0/postmaster/domain/repository"
	"github.com/biosugar0/postmaster/infrastructure/message"
)

type (
	Repository interface {
		PublishRepository() repository.PublishRepository
	}
	repositoryImpl struct {
		snsSvc *sns.SNS
		sqsSvc *sqs.SQS
	}
)

func NewRepository() Repository {
	return &repositoryImpl{
		snsSvc: message.NewSNS(),
		sqsSvc: message.NewSQS(),
	}
}

func (r *repositoryImpl) PublishRepository() repository.PublishRepository {
	return message.NewPublishRepository(r.snsSvc)
}
