// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the deployment groups for an application registered with the Amazon Web
// Services user or Amazon Web Services account.
func (c *Client) ListDeploymentGroups(ctx context.Context, params *ListDeploymentGroupsInput, optFns ...func(*Options)) (*ListDeploymentGroupsOutput, error) {
	if params == nil {
		params = &ListDeploymentGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentGroups", params, optFns, c.addOperationListDeploymentGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeploymentGroups operation.
type ListDeploymentGroupsInput struct {

	// The name of an CodeDeploy application associated with the user or Amazon Web
	// Services account.
	//
	// This member is required.
	ApplicationName *string

	// An identifier returned from the previous list deployment groups call. It can be
	// used to return the next set of deployment groups in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListDeploymentGroups operation.
type ListDeploymentGroupsOutput struct {

	// The application name.
	ApplicationName *string

	// A list of deployment group names.
	DeploymentGroups []string

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment groups call to return the next
	// set of deployment groups in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDeploymentGroups"); err != nil {
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
	if err = addOpListDeploymentGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentGroups(options.Region), middleware.Before); err != nil {
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

// ListDeploymentGroupsAPIClient is a client that implements the
// ListDeploymentGroups operation.
type ListDeploymentGroupsAPIClient interface {
	ListDeploymentGroups(context.Context, *ListDeploymentGroupsInput, ...func(*Options)) (*ListDeploymentGroupsOutput, error)
}

var _ ListDeploymentGroupsAPIClient = (*Client)(nil)

// ListDeploymentGroupsPaginatorOptions is the paginator options for
// ListDeploymentGroups
type ListDeploymentGroupsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentGroupsPaginator is a paginator for ListDeploymentGroups
type ListDeploymentGroupsPaginator struct {
	options   ListDeploymentGroupsPaginatorOptions
	client    ListDeploymentGroupsAPIClient
	params    *ListDeploymentGroupsInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentGroupsPaginator returns a new ListDeploymentGroupsPaginator
func NewListDeploymentGroupsPaginator(client ListDeploymentGroupsAPIClient, params *ListDeploymentGroupsInput, optFns ...func(*ListDeploymentGroupsPaginatorOptions)) *ListDeploymentGroupsPaginator {
	if params == nil {
		params = &ListDeploymentGroupsInput{}
	}

	options := ListDeploymentGroupsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeploymentGroups page.
func (p *ListDeploymentGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListDeploymentGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDeploymentGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDeploymentGroups",
	}
}
