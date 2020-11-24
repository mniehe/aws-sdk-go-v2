// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get all rules for a detector (paginated) if ruleId and ruleVersion are not
// specified. Gets all rules for the detector and the ruleId if present
// (paginated). Gets a specific rule if both the ruleId and the ruleVersion are
// specified. This is a paginated API. Providing null maxResults results in
// retrieving maximum of 100 records per page. If you provide maxResults the value
// must be between 50 and 100. To get the next page result, a provide a pagination
// token from GetRulesResult as part of your request. Null pagination token fetches
// the records from the beginning.
func (c *Client) GetRules(ctx context.Context, params *GetRulesInput, optFns ...func(*Options)) (*GetRulesOutput, error) {
	if params == nil {
		params = &GetRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRules", params, optFns, addOperationGetRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRulesInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The maximum number of rules to return for the request.
	MaxResults *int32

	// The next page token.
	NextToken *string

	// The rule ID.
	RuleId *string

	// The rule version.
	RuleVersion *string
}

type GetRulesOutput struct {

	// The next page token to be used in subsequent requests.
	NextToken *string

	// The details of the requested rule.
	RuleDetails []types.RuleDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRules{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRules(options.Region), middleware.Before); err != nil {
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
	return nil
}

// GetRulesAPIClient is a client that implements the GetRules operation.
type GetRulesAPIClient interface {
	GetRules(context.Context, *GetRulesInput, ...func(*Options)) (*GetRulesOutput, error)
}

var _ GetRulesAPIClient = (*Client)(nil)

// GetRulesPaginatorOptions is the paginator options for GetRules
type GetRulesPaginatorOptions struct {
	// The maximum number of rules to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetRulesPaginator is a paginator for GetRules
type GetRulesPaginator struct {
	options   GetRulesPaginatorOptions
	client    GetRulesAPIClient
	params    *GetRulesInput
	nextToken *string
	firstPage bool
}

// NewGetRulesPaginator returns a new GetRulesPaginator
func NewGetRulesPaginator(client GetRulesAPIClient, params *GetRulesInput, optFns ...func(*GetRulesPaginatorOptions)) *GetRulesPaginator {
	options := GetRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetRulesInput{}
	}

	return &GetRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetRulesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetRules page.
func (p *GetRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetRules(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetRules",
	}
}
