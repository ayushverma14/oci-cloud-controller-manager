// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudformation_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudFormation_CancelUpdateStack() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.CancelUpdateStackInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.CancelUpdateStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ContinueUpdateRollback() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ContinueUpdateRollbackInput{
		StackName: aws.String("StackNameOrId"), // Required
		ResourcesToSkip: []*string{
			aws.String("ResourceToSkip"), // Required
			// More values...
		},
		RoleARN: aws.String("RoleARN"),
	}
	resp, err := svc.ContinueUpdateRollback(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_CreateChangeSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.CreateChangeSetInput{
		ChangeSetName: aws.String("ChangeSetName"), // Required
		StackName:     aws.String("StackNameOrId"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		ChangeSetType: aws.String("ChangeSetType"),
		ClientToken:   aws.String("ClientToken"),
		Description:   aws.String("Description"),
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		Parameters: []*cloudformation.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		ResourceTypes: []*string{
			aws.String("ResourceType"), // Required
			// More values...
		},
		RoleARN: aws.String("RoleARN"),
		Tags: []*cloudformation.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:        aws.String("TemplateBody"),
		TemplateURL:         aws.String("TemplateURL"),
		UsePreviousTemplate: aws.Bool(true),
	}
	resp, err := svc.CreateChangeSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_CreateStack() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.CreateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		DisableRollback: aws.Bool(true),
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		OnFailure: aws.String("OnFailure"),
		Parameters: []*cloudformation.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		ResourceTypes: []*string{
			aws.String("ResourceType"), // Required
			// More values...
		},
		RoleARN:         aws.String("RoleARN"),
		StackPolicyBody: aws.String("StackPolicyBody"),
		StackPolicyURL:  aws.String("StackPolicyURL"),
		Tags: []*cloudformation.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:     aws.String("TemplateBody"),
		TemplateURL:      aws.String("TemplateURL"),
		TimeoutInMinutes: aws.Int64(1),
	}
	resp, err := svc.CreateStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DeleteChangeSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DeleteChangeSetInput{
		ChangeSetName: aws.String("ChangeSetNameOrId"), // Required
		StackName:     aws.String("StackNameOrId"),
	}
	resp, err := svc.DeleteChangeSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DeleteStack() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DeleteStackInput{
		StackName: aws.String("StackName"), // Required
		RetainResources: []*string{
			aws.String("LogicalResourceId"), // Required
			// More values...
		},
		RoleARN: aws.String("RoleARN"),
	}
	resp, err := svc.DeleteStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeAccountLimits() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeAccountLimitsInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.DescribeAccountLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeChangeSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeChangeSetInput{
		ChangeSetName: aws.String("ChangeSetNameOrId"), // Required
		NextToken:     aws.String("NextToken"),
		StackName:     aws.String("StackNameOrId"),
	}
	resp, err := svc.DescribeChangeSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeStackEvents() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeStackEventsInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String("StackName"),
	}
	resp, err := svc.DescribeStackEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeStackResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeStackResourceInput{
		LogicalResourceId: aws.String("LogicalResourceId"), // Required
		StackName:         aws.String("StackName"),         // Required
	}
	resp, err := svc.DescribeStackResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeStackResources() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeStackResourcesInput{
		LogicalResourceId:  aws.String("LogicalResourceId"),
		PhysicalResourceId: aws.String("PhysicalResourceId"),
		StackName:          aws.String("StackName"),
	}
	resp, err := svc.DescribeStackResources(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_DescribeStacks() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.DescribeStacksInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String("StackName"),
	}
	resp, err := svc.DescribeStacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_EstimateTemplateCost() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.EstimateTemplateCostInput{
		Parameters: []*cloudformation.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.EstimateTemplateCost(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ExecuteChangeSet() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ExecuteChangeSetInput{
		ChangeSetName: aws.String("ChangeSetNameOrId"), // Required
		StackName:     aws.String("StackNameOrId"),
	}
	resp, err := svc.ExecuteChangeSet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_GetStackPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.GetStackPolicyInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.GetStackPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_GetTemplate() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.GetTemplateInput{
		ChangeSetName: aws.String("ChangeSetNameOrId"),
		StackName:     aws.String("StackName"),
		TemplateStage: aws.String("TemplateStage"),
	}
	resp, err := svc.GetTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_GetTemplateSummary() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.GetTemplateSummaryInput{
		StackName:    aws.String("StackNameOrId"),
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.GetTemplateSummary(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ListChangeSets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ListChangeSetsInput{
		StackName: aws.String("StackNameOrId"), // Required
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListChangeSets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ListExports() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ListExportsInput{
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListExports(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ListImports() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ListImportsInput{
		ExportName: aws.String("ExportName"), // Required
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListImports(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ListStackResources() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ListStackResourcesInput{
		StackName: aws.String("StackName"), // Required
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListStackResources(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ListStacks() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ListStacksInput{
		NextToken: aws.String("NextToken"),
		StackStatusFilter: []*string{
			aws.String("StackStatus"), // Required
			// More values...
		},
	}
	resp, err := svc.ListStacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_SetStackPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.SetStackPolicyInput{
		StackName:       aws.String("StackName"), // Required
		StackPolicyBody: aws.String("StackPolicyBody"),
		StackPolicyURL:  aws.String("StackPolicyURL"),
	}
	resp, err := svc.SetStackPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_SignalResource() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.SignalResourceInput{
		LogicalResourceId: aws.String("LogicalResourceId"),      // Required
		StackName:         aws.String("StackNameOrId"),          // Required
		Status:            aws.String("ResourceSignalStatus"),   // Required
		UniqueId:          aws.String("ResourceSignalUniqueId"), // Required
	}
	resp, err := svc.SignalResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_UpdateStack() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.UpdateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		Parameters: []*cloudformation.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		ResourceTypes: []*string{
			aws.String("ResourceType"), // Required
			// More values...
		},
		RoleARN:                     aws.String("RoleARN"),
		StackPolicyBody:             aws.String("StackPolicyBody"),
		StackPolicyDuringUpdateBody: aws.String("StackPolicyDuringUpdateBody"),
		StackPolicyDuringUpdateURL:  aws.String("StackPolicyDuringUpdateURL"),
		StackPolicyURL:              aws.String("StackPolicyURL"),
		Tags: []*cloudformation.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:        aws.String("TemplateBody"),
		TemplateURL:         aws.String("TemplateURL"),
		UsePreviousTemplate: aws.Bool(true),
	}
	resp, err := svc.UpdateStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFormation_ValidateTemplate() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := cloudformation.New(sess)

	params := &cloudformation.ValidateTemplateInput{
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.ValidateTemplate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
