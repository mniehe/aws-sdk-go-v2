// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The CreateKeyspace operation adds a new keyspace to your account. In an Amazon
// Web Services account, keyspace names must be unique within each Region.
// CreateKeyspace is an asynchronous operation. You can monitor the creation status
// of the new keyspace by using the GetKeyspace operation. For more information,
// see Creating keyspaces (https://docs.aws.amazon.com/keyspaces/latest/devguide/working-with-keyspaces.html#keyspaces-create)
// in the Amazon Keyspaces Developer Guide.
func (c *Client) CreateKeyspace(ctx context.Context, params *CreateKeyspaceInput, optFns ...func(*Options)) (*CreateKeyspaceOutput, error) {
	if params == nil {
		params = &CreateKeyspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKeyspace", params, optFns, c.addOperationCreateKeyspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyspaceInput struct {

	// The name of the keyspace to be created.
	//
	// This member is required.
	KeyspaceName *string

	// The replication specification of the keyspace includes:
	//   - replicationStrategy - the required value is SINGLE_REGION or MULTI_REGION .
	//   - regionList - if the replicationStrategy is MULTI_REGION , the regionList
	//   requires the current Region and at least one additional Amazon Web Services
	//   Region where the keyspace is going to be replicated in. The maximum number of
	//   supported replication Regions including the current Region is six.
	ReplicationSpecification *types.ReplicationSpecification

	// A list of key-value pair tags to be attached to the keyspace. For more
	// information, see Adding tags and labels to Amazon Keyspaces resources (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
	// in the Amazon Keyspaces Developer Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateKeyspaceOutput struct {

	// The unique identifier of the keyspace in the format of an Amazon Resource Name
	// (ARN).
	//
	// This member is required.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateKeyspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateKeyspace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKeyspace"); err != nil {
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
	if err = addOpCreateKeyspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKeyspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateKeyspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKeyspace",
	}
}
