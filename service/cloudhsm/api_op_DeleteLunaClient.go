// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/) , the AWS
// CloudHSM Classic User Guide (https://docs.aws.amazon.com/cloudhsm/classic/userguide/)
// , and the AWS CloudHSM Classic API Reference (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/)
// . For information about the current version of AWS CloudHSM, see AWS CloudHSM (http://aws.amazon.com/cloudhsm/)
// , the AWS CloudHSM User Guide (https://docs.aws.amazon.com/cloudhsm/latest/userguide/)
// , and the AWS CloudHSM API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/)
// . Deletes a client.
//
// Deprecated: This API is deprecated.
func (c *Client) DeleteLunaClient(ctx context.Context, params *DeleteLunaClientInput, optFns ...func(*Options)) (*DeleteLunaClientOutput, error) {
	if params == nil {
		params = &DeleteLunaClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLunaClient", params, optFns, c.addOperationDeleteLunaClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLunaClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLunaClientInput struct {

	// The ARN of the client to delete.
	//
	// This member is required.
	ClientArn *string

	noSmithyDocumentSerde
}

type DeleteLunaClientOutput struct {

	// The status of the action.
	//
	// This member is required.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLunaClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLunaClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLunaClient{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteLunaClient"); err != nil {
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
	if err = addOpDeleteLunaClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLunaClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteLunaClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteLunaClient",
	}
}
