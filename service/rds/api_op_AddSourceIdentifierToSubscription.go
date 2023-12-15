// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a source identifier to an existing RDS event notification subscription.
func (c *Client) AddSourceIdentifierToSubscription(ctx context.Context, params *AddSourceIdentifierToSubscriptionInput, optFns ...func(*Options)) (*AddSourceIdentifierToSubscriptionOutput, error) {
	if params == nil {
		params = &AddSourceIdentifierToSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddSourceIdentifierToSubscription", params, optFns, c.addOperationAddSourceIdentifierToSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddSourceIdentifierToSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddSourceIdentifierToSubscriptionInput struct {

	// The identifier of the event source to be added. Constraints:
	//   - If the source type is a DB instance, a DBInstanceIdentifier value must be
	//   supplied.
	//   - If the source type is a DB cluster, a DBClusterIdentifier value must be
	//   supplied.
	//   - If the source type is a DB parameter group, a DBParameterGroupName value
	//   must be supplied.
	//   - If the source type is a DB security group, a DBSecurityGroupName value must
	//   be supplied.
	//   - If the source type is a DB snapshot, a DBSnapshotIdentifier value must be
	//   supplied.
	//   - If the source type is a DB cluster snapshot, a DBClusterSnapshotIdentifier
	//   value must be supplied.
	//   - If the source type is an RDS Proxy, a DBProxyName value must be supplied.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the RDS event notification subscription you want to add a source
	// identifier to.
	//
	// This member is required.
	SubscriptionName *string

	noSmithyDocumentSerde
}

type AddSourceIdentifierToSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddSourceIdentifierToSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddSourceIdentifierToSubscription"); err != nil {
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
	if err = addOpAddSourceIdentifierToSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddSourceIdentifierToSubscription",
	}
}
