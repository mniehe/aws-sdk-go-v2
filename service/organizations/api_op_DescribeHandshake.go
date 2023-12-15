// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a previously requested handshake. The handshake ID
// comes from the response to the original InviteAccountToOrganization operation
// that generated the handshake. You can access handshakes that are ACCEPTED ,
// DECLINED , or CANCELED for only 30 days after they change to that state.
// They're then deleted and no longer accessible. This operation can be called from
// any account in the organization.
func (c *Client) DescribeHandshake(ctx context.Context, params *DescribeHandshakeInput, optFns ...func(*Options)) (*DescribeHandshakeOutput, error) {
	if params == nil {
		params = &DescribeHandshakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHandshake", params, optFns, c.addOperationDescribeHandshakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHandshakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHandshakeInput struct {

	// The unique identifier (ID) of the handshake that you want information about.
	// You can get the ID from the original call to InviteAccountToOrganization , or
	// from a call to ListHandshakesForAccount or ListHandshakesForOrganization . The
	// regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string requires
	// "h-" followed by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	HandshakeId *string

	noSmithyDocumentSerde
}

type DescribeHandshakeOutput struct {

	// A structure that contains information about the specified handshake.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHandshakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHandshake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHandshake{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHandshake"); err != nil {
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
	if err = addOpDescribeHandshakeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHandshake(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHandshake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHandshake",
	}
}
