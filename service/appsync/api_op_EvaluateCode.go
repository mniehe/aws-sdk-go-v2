// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates the given code and returns the response. The code definition
// requirements depend on the specified runtime. For APPSYNC_JS runtimes, the code
// defines the request and response functions. The request function takes the
// incoming request after a GraphQL operation is parsed and converts it into a
// request configuration for the selected data source operation. The response
// function interprets responses from the data source and maps it to the shape of
// the GraphQL field output type.
func (c *Client) EvaluateCode(ctx context.Context, params *EvaluateCodeInput, optFns ...func(*Options)) (*EvaluateCodeOutput, error) {
	if params == nil {
		params = &EvaluateCodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateCode", params, optFns, c.addOperationEvaluateCodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluateCodeInput struct {

	// The code definition to be evaluated. Note that code and runtime are both
	// required for this action. The runtime value must be APPSYNC_JS .
	//
	// This member is required.
	Code *string

	// The map that holds all of the contextual information for your resolver
	// invocation. A context is required for this action.
	//
	// This member is required.
	Context *string

	// The runtime to be used when evaluating the code. Currently, only the APPSYNC_JS
	// runtime is supported.
	//
	// This member is required.
	Runtime *types.AppSyncRuntime

	// The function within the code to be evaluated. If provided, the valid values are
	// request and response .
	Function *string

	noSmithyDocumentSerde
}

type EvaluateCodeOutput struct {

	// Contains the payload of the response error.
	Error *types.EvaluateCodeErrorDetail

	// The result of the evaluation operation.
	EvaluationResult *string

	// A list of logs that were generated by calls to util.log.info and util.log.error
	// in the evaluated code.
	Logs []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateCodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEvaluateCode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEvaluateCode{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EvaluateCode"); err != nil {
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
	if err = addOpEvaluateCodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateCode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEvaluateCode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EvaluateCode",
	}
}
