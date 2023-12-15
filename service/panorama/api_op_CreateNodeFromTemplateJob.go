// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a camera stream node.
func (c *Client) CreateNodeFromTemplateJob(ctx context.Context, params *CreateNodeFromTemplateJobInput, optFns ...func(*Options)) (*CreateNodeFromTemplateJobOutput, error) {
	if params == nil {
		params = &CreateNodeFromTemplateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNodeFromTemplateJob", params, optFns, c.addOperationCreateNodeFromTemplateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNodeFromTemplateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNodeFromTemplateJobInput struct {

	// A name for the node.
	//
	// This member is required.
	NodeName *string

	// An output package name for the node.
	//
	// This member is required.
	OutputPackageName *string

	// An output package version for the node.
	//
	// This member is required.
	OutputPackageVersion *string

	// Template parameters for the node.
	//
	// This member is required.
	TemplateParameters map[string]string

	// The type of node.
	//
	// This member is required.
	TemplateType types.TemplateType

	// Tags for the job.
	JobTags []types.JobResourceTags

	// A description for the node.
	NodeDescription *string

	noSmithyDocumentSerde
}

type CreateNodeFromTemplateJobOutput struct {

	// The job's ID.
	//
	// This member is required.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNodeFromTemplateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNodeFromTemplateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNodeFromTemplateJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNodeFromTemplateJob"); err != nil {
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
	if err = addOpCreateNodeFromTemplateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNodeFromTemplateJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNodeFromTemplateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNodeFromTemplateJob",
	}
}
