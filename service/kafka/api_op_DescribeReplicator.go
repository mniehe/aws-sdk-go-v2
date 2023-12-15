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
	"time"
)

// Describes a replicator.
func (c *Client) DescribeReplicator(ctx context.Context, params *DescribeReplicatorInput, optFns ...func(*Options)) (*DescribeReplicatorOutput, error) {
	if params == nil {
		params = &DescribeReplicatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicator", params, optFns, c.addOperationDescribeReplicatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicatorInput struct {

	// The Amazon Resource Name (ARN) of the replicator to be described.
	//
	// This member is required.
	ReplicatorArn *string

	noSmithyDocumentSerde
}

type DescribeReplicatorOutput struct {

	// The time when the replicator was created.
	CreationTime *time.Time

	// The current version number of the replicator.
	CurrentVersion *string

	// Whether this resource is a replicator reference.
	IsReplicatorReference *bool

	// Kafka Clusters used in setting up sources / targets for replication.
	KafkaClusters []types.KafkaClusterDescription

	// A list of replication configurations, where each configuration targets a given
	// source cluster to target cluster replication flow.
	ReplicationInfoList []types.ReplicationInfoDescription

	// The Amazon Resource Name (ARN) of the replicator.
	ReplicatorArn *string

	// The description of the replicator.
	ReplicatorDescription *string

	// The name of the replicator.
	ReplicatorName *string

	// The Amazon Resource Name (ARN) of the replicator resource in the region where
	// the replicator was created.
	ReplicatorResourceArn *string

	// State of the replicator.
	ReplicatorState types.ReplicatorState

	// The Amazon Resource Name (ARN) of the IAM role used by the replicator to access
	// resources in the customer's account (e.g source and target clusters)
	ServiceExecutionRoleArn *string

	// Details about the state of the replicator.
	StateInfo *types.ReplicationStateInfo

	// List of tags attached to the Replicator.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReplicator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReplicator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReplicator"); err != nil {
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
	if err = addOpDescribeReplicatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReplicator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReplicator",
	}
}
