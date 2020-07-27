package message

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	initOnceSQS sync.Once
	sqsSvc      *sqs.SQS
)

func NewSQS(profile string) *sqs.SQS {
	initOnceSQS.Do(func() {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           profile,
		}))
		sess.Config.Region = aws.String("ap-northeast-1")
		sqsSvc = sqs.New(sess)
	})

	return sqsSvc
}
