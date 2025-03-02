// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deactivates an evaluation form in the specified Amazon Connect instance. After
// a form is deactivated, it is no longer available for users to start new
// evaluations based on the form.
func (c *Client) DeactivateEvaluationForm(ctx context.Context, params *DeactivateEvaluationFormInput, optFns ...func(*Options)) (*DeactivateEvaluationFormOutput, error) {
	if params == nil {
		params = &DeactivateEvaluationFormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeactivateEvaluationForm", params, optFns, c.addOperationDeactivateEvaluationFormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeactivateEvaluationFormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateEvaluationFormInput struct {

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// A version of the evaluation form. If the version property is not provided, the
	// latest version of the evaluation form is deactivated.
	//
	// This member is required.
	EvaluationFormVersion int32

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

type DeactivateEvaluationFormOutput struct {

	// The Amazon Resource Name (ARN) for the evaluation form resource.
	//
	// This member is required.
	EvaluationFormArn *string

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// The version of the deactivated evaluation form resource.
	//
	// This member is required.
	EvaluationFormVersion int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeactivateEvaluationFormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeactivateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeactivateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeactivateEvaluationForm"); err != nil {
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
	if err = addOpDeactivateEvaluationFormValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateEvaluationForm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeactivateEvaluationForm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeactivateEvaluationForm",
	}
}
