// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Creates a RegexMatchSet . You then use UpdateRegexMatchSet to
// identify the part of a web request that you want AWS WAF to inspect, such as the
// values of the User-Agent header or the query string. For example, you can
// create a RegexMatchSet that contains a RegexMatchTuple that looks for any
// requests with User-Agent headers that match a RegexPatternSet with pattern
// B[a@]dB[o0]t . You can then configure AWS WAF to reject those requests. To
// create and configure a RegexMatchSet , perform the following steps:
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of a CreateRegexMatchSet request.
//   - Submit a CreateRegexMatchSet request.
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateRegexMatchSet request.
//   - Submit an UpdateRegexMatchSet request to specify the part of the request
//     that you want AWS WAF to inspect (for example, the header or the URI) and the
//     value, using a RegexPatternSet , that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/)
// .
func (c *Client) CreateRegexMatchSet(ctx context.Context, params *CreateRegexMatchSetInput, optFns ...func(*Options)) (*CreateRegexMatchSetOutput, error) {
	if params == nil {
		params = &CreateRegexMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegexMatchSet", params, optFns, c.addOperationCreateRegexMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegexMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegexMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the RegexMatchSet . You can't change Name
	// after you create a RegexMatchSet .
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateRegexMatchSetOutput struct {

	// The ChangeToken that you used to submit the CreateRegexMatchSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	ChangeToken *string

	// A RegexMatchSet that contains no RegexMatchTuple objects.
	RegexMatchSet *types.RegexMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRegexMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRegexMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRegexMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRegexMatchSet"); err != nil {
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
	if err = addOpCreateRegexMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegexMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRegexMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRegexMatchSet",
	}
}
