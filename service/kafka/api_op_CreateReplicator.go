// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the replicator.
func (c *Client) CreateReplicator(ctx context.Context, params *CreateReplicatorInput, optFns ...func(*Options)) (*CreateReplicatorOutput, error) {
	if params == nil {
		params = &CreateReplicatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicator", params, optFns, c.addOperationCreateReplicatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a replicator using the specified configuration.
type CreateReplicatorInput struct {

	// Kafka Clusters to use in setting up sources / targets for replication.
	//
	// This member is required.
	KafkaClusters []types.KafkaCluster

	// A list of replication configurations, where each configuration targets a given
	// source cluster to target cluster replication flow.
	//
	// This member is required.
	ReplicationInfoList []types.ReplicationInfo

	// The name of the replicator. Alpha-numeric characters with '-' are allowed.
	//
	// This member is required.
	ReplicatorName *string

	// The ARN of the IAM role used by the replicator to access resources in the
	// customer's account (e.g source and target clusters)
	//
	// This member is required.
	ServiceExecutionRoleArn *string

	// A summary description of the replicator.
	Description *string

	// List of tags to attach to created Replicator.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateReplicatorOutput struct {

	// The Amazon Resource Name (ARN) of the replicator.
	ReplicatorArn *string

	// Name of the replicator provided by the customer.
	ReplicatorName *string

	// State of the replicator.
	ReplicatorState types.ReplicatorState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReplicator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReplicator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplicator"); err != nil {
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
	if err = addOpCreateReplicatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplicator",
	}
}
