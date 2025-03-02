// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the current state of block public access for snapshots setting for the
// account and Region. For more information, see Block public access for snapshots (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-public-access-snapshots.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) GetSnapshotBlockPublicAccessState(ctx context.Context, params *GetSnapshotBlockPublicAccessStateInput, optFns ...func(*Options)) (*GetSnapshotBlockPublicAccessStateOutput, error) {
	if params == nil {
		params = &GetSnapshotBlockPublicAccessStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshotBlockPublicAccessState", params, optFns, c.addOperationGetSnapshotBlockPublicAccessStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotBlockPublicAccessStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnapshotBlockPublicAccessStateInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type GetSnapshotBlockPublicAccessStateOutput struct {

	// The current state of block public access for snapshots. Possible values
	// include:
	//   - block-all-sharing - All public sharing of snapshots is blocked. Users in the
	//   account can't request new public sharing. Additionally, snapshots that were
	//   already publicly shared are treated as private and are not publicly available.
	//   - block-new-sharing - Only new public sharing of snapshots is blocked. Users
	//   in the account can't request new public sharing. However, snapshots that were
	//   already publicly shared, remain publicly available.
	//   - unblocked - Public sharing is not blocked. Users can publicly share
	//   snapshots.
	State types.SnapshotBlockPublicAccessState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSnapshotBlockPublicAccessStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetSnapshotBlockPublicAccessState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetSnapshotBlockPublicAccessState{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSnapshotBlockPublicAccessState"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshotBlockPublicAccessState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSnapshotBlockPublicAccessState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSnapshotBlockPublicAccessState",
	}
}
