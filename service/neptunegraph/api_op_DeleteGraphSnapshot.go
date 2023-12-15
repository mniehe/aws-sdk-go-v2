// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes the specifed graph snapshot.
func (c *Client) DeleteGraphSnapshot(ctx context.Context, params *DeleteGraphSnapshotInput, optFns ...func(*Options)) (*DeleteGraphSnapshotOutput, error) {
	if params == nil {
		params = &DeleteGraphSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGraphSnapshot", params, optFns, c.addOperationDeleteGraphSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGraphSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGraphSnapshotInput struct {

	// ID of the graph snapshot to be deleted.
	//
	// This member is required.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

func (in *DeleteGraphSnapshotInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type DeleteGraphSnapshotOutput struct {

	// The ARN of the graph snapshot.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph snapshot.
	//
	// This member is required.
	Id *string

	// The snapshot name. For example: my-snapshot-1 . The name must contain from 1 to
	// 63 letters, numbers, or hyphens, and its first character must be a letter. It
	// cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	Name *string

	// The ID of the KMS key used to encrypt and decrypt the snapshot.
	KmsKeyIdentifier *string

	// The time when the snapshot was created.
	SnapshotCreateTime *time.Time

	// The graph identifier for the graph from which the snapshot was created.
	SourceGraphId *string

	// The status of the graph snapshot.
	Status types.SnapshotStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGraphSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteGraphSnapshot"); err != nil {
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
	if err = addOpDeleteGraphSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGraphSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGraphSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteGraphSnapshot",
	}
}
