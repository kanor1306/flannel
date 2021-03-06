// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sns_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSNS_AddPermission() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.AddPermissionInput{
		AWSAccountId: []*string{ // Required
			aws.String("delegate"), // Required
			// More values...
		},
		ActionName: []*string{ // Required
			aws.String("action"), // Required
			// More values...
		},
		Label:    aws.String("label"),    // Required
		TopicArn: aws.String("topicARN"), // Required
	}
	resp, err := svc.AddPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_CheckIfPhoneNumberIsOptedOut() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.CheckIfPhoneNumberIsOptedOutInput{
		PhoneNumber: aws.String("PhoneNumber"), // Required
	}
	resp, err := svc.CheckIfPhoneNumberIsOptedOut(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ConfirmSubscription() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ConfirmSubscriptionInput{
		Token:                     aws.String("token"),    // Required
		TopicArn:                  aws.String("topicARN"), // Required
		AuthenticateOnUnsubscribe: aws.String("authenticateOnUnsubscribe"),
	}
	resp, err := svc.ConfirmSubscription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_CreatePlatformApplication() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.CreatePlatformApplicationInput{
		Attributes: map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		Name:     aws.String("String"), // Required
		Platform: aws.String("String"), // Required
	}
	resp, err := svc.CreatePlatformApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_CreatePlatformEndpoint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String("String"), // Required
		Token: aws.String("String"), // Required
		Attributes: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		CustomUserData: aws.String("String"),
	}
	resp, err := svc.CreatePlatformEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_CreateTopic() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.CreateTopicInput{
		Name: aws.String("topicName"), // Required
	}
	resp, err := svc.CreateTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_DeleteEndpoint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.DeleteEndpointInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_DeletePlatformApplication() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.DeletePlatformApplicationInput{
		PlatformApplicationArn: aws.String("String"), // Required
	}
	resp, err := svc.DeletePlatformApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_DeleteTopic() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.DeleteTopicInput{
		TopicArn: aws.String("topicARN"), // Required
	}
	resp, err := svc.DeleteTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_GetEndpointAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.GetEndpointAttributesInput{
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.GetEndpointAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_GetPlatformApplicationAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.GetPlatformApplicationAttributesInput{
		PlatformApplicationArn: aws.String("String"), // Required
	}
	resp, err := svc.GetPlatformApplicationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_GetSMSAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.GetSMSAttributesInput{
		Attributes: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.GetSMSAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_GetSubscriptionAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.GetSubscriptionAttributesInput{
		SubscriptionArn: aws.String("subscriptionARN"), // Required
	}
	resp, err := svc.GetSubscriptionAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_GetTopicAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.GetTopicAttributesInput{
		TopicArn: aws.String("topicARN"), // Required
	}
	resp, err := svc.GetTopicAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListEndpointsByPlatformApplication() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListEndpointsByPlatformApplicationInput{
		PlatformApplicationArn: aws.String("String"), // Required
		NextToken:              aws.String("String"),
	}
	resp, err := svc.ListEndpointsByPlatformApplication(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListPhoneNumbersOptedOut() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListPhoneNumbersOptedOutInput{
		NextToken: aws.String("string"),
	}
	resp, err := svc.ListPhoneNumbersOptedOut(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListPlatformApplications() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListPlatformApplicationsInput{
		NextToken: aws.String("String"),
	}
	resp, err := svc.ListPlatformApplications(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListSubscriptions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListSubscriptionsInput{
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListSubscriptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListSubscriptionsByTopic() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListSubscriptionsByTopicInput{
		TopicArn:  aws.String("topicARN"), // Required
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListSubscriptionsByTopic(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_ListTopics() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.ListTopicsInput{
		NextToken: aws.String("nextToken"),
	}
	resp, err := svc.ListTopics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_OptInPhoneNumber() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.OptInPhoneNumberInput{
		PhoneNumber: aws.String("PhoneNumber"), // Required
	}
	resp, err := svc.OptInPhoneNumber(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_Publish() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.PublishInput{
		Message: aws.String("message"), // Required
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"Key": { // Required
				DataType:    aws.String("String"), // Required
				BinaryValue: []byte("PAYLOAD"),
				StringValue: aws.String("String"),
			},
			// More values...
		},
		MessageStructure: aws.String("messageStructure"),
		PhoneNumber:      aws.String("String"),
		Subject:          aws.String("subject"),
		TargetArn:        aws.String("String"),
		TopicArn:         aws.String("topicARN"),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_RemovePermission() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.RemovePermissionInput{
		Label:    aws.String("label"),    // Required
		TopicArn: aws.String("topicARN"), // Required
	}
	resp, err := svc.RemovePermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_SetEndpointAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SetEndpointAttributesInput{
		Attributes: map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		EndpointArn: aws.String("String"), // Required
	}
	resp, err := svc.SetEndpointAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_SetPlatformApplicationAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SetPlatformApplicationAttributesInput{
		Attributes: map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
		PlatformApplicationArn: aws.String("String"), // Required
	}
	resp, err := svc.SetPlatformApplicationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_SetSMSAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SetSMSAttributesInput{
		Attributes: map[string]*string{ // Required
			"Key": aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.SetSMSAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_SetSubscriptionAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SetSubscriptionAttributesInput{
		AttributeName:   aws.String("attributeName"),   // Required
		SubscriptionArn: aws.String("subscriptionARN"), // Required
		AttributeValue:  aws.String("attributeValue"),
	}
	resp, err := svc.SetSubscriptionAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_SetTopicAttributes() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SetTopicAttributesInput{
		AttributeName:  aws.String("attributeName"), // Required
		TopicArn:       aws.String("topicARN"),      // Required
		AttributeValue: aws.String("attributeValue"),
	}
	resp, err := svc.SetTopicAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_Subscribe() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.SubscribeInput{
		Protocol: aws.String("protocol"), // Required
		TopicArn: aws.String("topicARN"), // Required
		Endpoint: aws.String("endpoint"),
	}
	resp, err := svc.Subscribe(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleSNS_Unsubscribe() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := sns.New(sess)

	params := &sns.UnsubscribeInput{
		SubscriptionArn: aws.String("subscriptionARN"), // Required
	}
	resp, err := svc.Unsubscribe(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
