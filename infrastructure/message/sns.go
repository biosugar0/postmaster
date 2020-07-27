package message

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var (
	initOnceSNS sync.Once
	snsSvc      *sns.SNS
)

func NewSNS(profile string) *sns.SNS {
	initOnceSQS.Do(func() {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           profile,
		}))
		sess.Config.Region = aws.String("ap-northeast-1")
		snsSvc = sns.New(sess)
	})

	return snsSvc
}
