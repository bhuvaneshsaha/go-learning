// app/services/emailService/emailService.go
package emailService

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/bhuvaneshsaha/final-demo-website/app/config"
)

// const (
// 	Sender    = "bhuvaneshmib@gmail.com"
// 	Recipient = "bhuvaneshs@nallas.com"
// 	Subject   = "Amazon SES Test (AWS SDK for Go)"
// 	CharSet   = "UTF-8"
// )

func SendEmail(message string) {

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.AWS.Region),
		Credentials: credentials.NewStaticCredentials(conf.AWS.AccessKey, conf.AWS.SecretKey, ""),
	},
	)

	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(conf.Email.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(conf.Email.CharSet),
					Data:    aws.String(message),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(conf.Email.CharSet),
				Data:    aws.String(conf.Email.Subject),
			},
		},
		Source: aws.String(conf.Email.Sender),
	}

	_, err = svc.SendEmail(input)

	if err != nil {
		log.Println("Error sending message", err)
		return
	}

	log.Println("Email sent!")
}
