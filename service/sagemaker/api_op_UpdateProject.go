// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a machine learning (ML) project that is created from a template that
// sets up an ML pipeline from training to deploying an approved model. You must
// not update a project that is in use. If you update the
// ServiceCatalogProvisioningUpdateDetails of a project that is active or being
// created, or updated, you may lose resources already created by the project.
func (c *Client) UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) {
	if params == nil {
		params = &UpdateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProject", params, optFns, c.addOperationUpdateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProjectInput struct {

	// The name of the project.
	//
	// This member is required.
	ProjectName *string

	// The description for the project.
	ProjectDescription *string

	// The product ID and provisioning artifact ID to provision a service catalog. The
	// provisioning artifact ID will default to the latest provisioning artifact ID of
	// the product, if you don't provide the provisioning artifact ID. For more
	// information, see What is Amazon Web Services Service Catalog (https://docs.aws.amazon.com/servicecatalog/latest/adminguide/introduction.html)
	// .
	ServiceCatalogProvisioningUpdateDetails *types.ServiceCatalogProvisioningUpdateDetails

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// . In addition, the project must have tag update constraints set in order to
	// include this parameter in the request. For more information, see Amazon Web
	// Services Service Catalog Tag Update Constraints (https://docs.aws.amazon.com/servicecatalog/latest/adminguide/constraints-resourceupdate.html)
	// .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type UpdateProjectOutput struct {

	// The Amazon Resource Name (ARN) of the project.
	//
	// This member is required.
	ProjectArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProject{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProject"); err != nil {
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
	if err = addOpUpdateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProject",
	}
}
