// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the auth policy for the specified service or
// service network.
func (c *Client) GetAuthPolicy(ctx context.Context, params *GetAuthPolicyInput, optFns ...func(*Options)) (*GetAuthPolicyOutput, error) {
	if params == nil {
		params = &GetAuthPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAuthPolicy", params, optFns, c.addOperationGetAuthPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAuthPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAuthPolicyInput struct {

	// The ID or Amazon Resource Name (ARN) of the service network or service.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type GetAuthPolicyOutput struct {

	// The date and time that the auth policy was created, specified in ISO-8601
	// format.
	CreatedAt *time.Time

	// The date and time that the auth policy was last updated, specified in ISO-8601
	// format.
	LastUpdatedAt *time.Time

	// The auth policy.
	Policy *string

	// The state of the auth policy. The auth policy is only active when the auth type
	// is set to Amazon Web Services_IAM . If you provide a policy, then authentication
	// and authorization decisions are made based on this policy and the client's IAM
	// policy. If the auth type is NONE , then any auth policy you provide will remain
	// inactive. For more information, see Create a service network (https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-networks.html#create-service-network)
	// in the Amazon VPC Lattice User Guide.
	State types.AuthPolicyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAuthPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAuthPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAuthPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAuthPolicy"); err != nil {
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
	if err = addOpGetAuthPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAuthPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAuthPolicy",
	}
}
