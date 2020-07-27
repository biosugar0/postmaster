package message

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/biosugar0/postmaster/domain/entity"
	"github.com/biosugar0/postmaster/domain/repository"
)

type publishRepository struct {
	svc *sns.SNS
}

func NewPublishRepository(svc *sns.SNS) repository.PublishRepository {
	return &publishRepository{svc: svc}
}

func (r *publishRepository) Create(m entity.Publish) (string, error) {
	result, err := r.svc.CreateTopic(&sns.CreateTopicInput{
		Name: &m.Topic,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return "", errors.New(awsErr.Code())
		}
	}
	if result.TopicArn != nil {
		return *result.TopicArn, nil
	}
	return "", nil
}
