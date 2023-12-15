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
// and global use. Returns an array of RuleGroup objects that you are subscribed
// to.
func (c *Client) ListSubscribedRuleGroups(ctx context.Context, params *ListSubscribedRuleGroupsInput, optFns ...func(*Options)) (*ListSubscribedRuleGroupsOutput, error) {
	if params == nil {
		params = &ListSubscribedRuleGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscribedRuleGroups", params, optFns, c.addOperationListSubscribedRuleGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscribedRuleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscribedRuleGroupsInput struct {

	// Specifies the number of subscribed rule groups that you want AWS WAF to return
	// for this request. If you have more objects than the number you specify for Limit
	// , the response includes a NextMarker value that you can use to get another
	// batch of objects.
	Limit int32

	// If you specify a value for Limit and you have more ByteMatchSets subscribed rule
	// groups than the value of Limit , AWS WAF returns a NextMarker value in the
	// response that allows you to list another group of subscribed rule groups. For
	// the second and subsequent ListSubscribedRuleGroupsRequest requests, specify the
	// value of NextMarker from the previous response to get information about another
	// batch of subscribed rule groups.
	NextMarker *string

	noSmithyDocumentSerde
}

type ListSubscribedRuleGroupsOutput struct {

	// If you have more objects than the number that you specified for Limit in the
	// request, the response includes a NextMarker value. To list more objects, submit
	// another ListSubscribedRuleGroups request, and specify the NextMarker value from
	// the response in the NextMarker value in the next request.
	NextMarker *string

	// An array of RuleGroup objects.
	RuleGroups []types.SubscribedRuleGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscribedRuleGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSubscribedRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSubscribedRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubscribedRuleGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscribedRuleGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSubscribedRuleGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubscribedRuleGroups",
	}
}
