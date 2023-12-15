// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a connection string for a user to connect to a kdb cluster. You must
// call this API using the same role that you have defined while creating a user.
func (c *Client) GetKxConnectionString(ctx context.Context, params *GetKxConnectionStringInput, optFns ...func(*Options)) (*GetKxConnectionStringOutput, error) {
	if params == nil {
		params = &GetKxConnectionStringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKxConnectionString", params, optFns, c.addOperationGetKxConnectionStringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKxConnectionStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKxConnectionStringInput struct {

	// A name of the kdb cluster.
	//
	// This member is required.
	ClusterName *string

	// A unique identifier for the kdb environment.
	//
	// This member is required.
	EnvironmentId *string

	// The Amazon Resource Name (ARN) that identifies the user. For more information
	// about ARNs and how to use ARNs in policies, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in the IAM User Guide.
	//
	// This member is required.
	UserArn *string

	noSmithyDocumentSerde
}

type GetKxConnectionStringOutput struct {

	// The signed connection string that you can use to connect to clusters.
	SignedConnectionString *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKxConnectionStringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetKxConnectionString{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetKxConnectionString{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetKxConnectionString"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpGetKxConnectionStringValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKxConnectionString(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKxConnectionString(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetKxConnectionString",
	}
}
