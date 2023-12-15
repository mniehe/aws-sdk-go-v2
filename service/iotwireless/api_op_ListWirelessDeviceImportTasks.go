// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List wireless devices that have been added to an import task.
func (c *Client) ListWirelessDeviceImportTasks(ctx context.Context, params *ListWirelessDeviceImportTasksInput, optFns ...func(*Options)) (*ListWirelessDeviceImportTasksOutput, error) {
	if params == nil {
		params = &ListWirelessDeviceImportTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWirelessDeviceImportTasks", params, optFns, c.addOperationListWirelessDeviceImportTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWirelessDeviceImportTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWirelessDeviceImportTasksInput struct {

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWirelessDeviceImportTasksOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// List of import tasks and summary information of onboarding status of devices in
	// each import task.
	WirelessDeviceImportTaskList []types.WirelessDeviceImportTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWirelessDeviceImportTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWirelessDeviceImportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWirelessDeviceImportTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWirelessDeviceImportTasks"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWirelessDeviceImportTasks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListWirelessDeviceImportTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWirelessDeviceImportTasks",
	}
}
