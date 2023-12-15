// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a traffic policy, which you use to create multiple DNS resource record
// sets for one domain name (such as example.com) or one subdomain name (such as
// www.example.com).
func (c *Client) CreateTrafficPolicy(ctx context.Context, params *CreateTrafficPolicyInput, optFns ...func(*Options)) (*CreateTrafficPolicyOutput, error) {
	if params == nil {
		params = &CreateTrafficPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrafficPolicy", params, optFns, c.addOperationCreateTrafficPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the traffic policy that you want
// to create.
type CreateTrafficPolicyInput struct {

	// The definition of this traffic policy in JSON format. For more information, see
	// Traffic Policy Document Format (https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)
	// .
	//
	// This member is required.
	Document *string

	// The name of the traffic policy.
	//
	// This member is required.
	Name *string

	// (Optional) Any comments that you want to include about the traffic policy.
	Comment *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the
// CreateTrafficPolicy request.
type CreateTrafficPolicyOutput struct {

	// A unique URL that represents a new traffic policy.
	//
	// This member is required.
	Location *string

	// A complex type that contains settings for the new traffic policy.
	//
	// This member is required.
	TrafficPolicy *types.TrafficPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrafficPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTrafficPolicy"); err != nil {
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
	if err = addOpCreateTrafficPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrafficPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrafficPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTrafficPolicy",
	}
}
