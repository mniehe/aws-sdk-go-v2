// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the maintenance window targets or tasks that an
// instance is associated with.
func (c *Client) DescribeMaintenanceWindowsForTarget(ctx context.Context, params *DescribeMaintenanceWindowsForTargetInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowsForTargetOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowsForTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowsForTarget", params, optFns, addOperationDescribeMaintenanceWindowsForTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowsForTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowsForTargetInput struct {

	// The type of resource you want to retrieve information about. For example,
	// "INSTANCE".
	//
	// This member is required.
	ResourceType types.MaintenanceWindowResourceType

	// The instance ID or key/value pair to retrieve information about.
	//
	// This member is required.
	Targets []types.Target

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type DescribeMaintenanceWindowsForTargetOutput struct {

	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string

	// Information about the maintenance window targets and tasks an instance is
	// associated with.
	WindowIdentities []types.MaintenanceWindowIdentityForTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMaintenanceWindowsForTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowsForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowsForTarget{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowsForTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowsForTarget(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowsForTargetAPIClient is a client that implements the
// DescribeMaintenanceWindowsForTarget operation.
type DescribeMaintenanceWindowsForTargetAPIClient interface {
	DescribeMaintenanceWindowsForTarget(context.Context, *DescribeMaintenanceWindowsForTargetInput, ...func(*Options)) (*DescribeMaintenanceWindowsForTargetOutput, error)
}

var _ DescribeMaintenanceWindowsForTargetAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowsForTargetPaginatorOptions is the paginator options for
// DescribeMaintenanceWindowsForTarget
type DescribeMaintenanceWindowsForTargetPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowsForTargetPaginator is a paginator for
// DescribeMaintenanceWindowsForTarget
type DescribeMaintenanceWindowsForTargetPaginator struct {
	options   DescribeMaintenanceWindowsForTargetPaginatorOptions
	client    DescribeMaintenanceWindowsForTargetAPIClient
	params    *DescribeMaintenanceWindowsForTargetInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowsForTargetPaginator returns a new
// DescribeMaintenanceWindowsForTargetPaginator
func NewDescribeMaintenanceWindowsForTargetPaginator(client DescribeMaintenanceWindowsForTargetAPIClient, params *DescribeMaintenanceWindowsForTargetInput, optFns ...func(*DescribeMaintenanceWindowsForTargetPaginatorOptions)) *DescribeMaintenanceWindowsForTargetPaginator {
	options := DescribeMaintenanceWindowsForTargetPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeMaintenanceWindowsForTargetInput{}
	}

	return &DescribeMaintenanceWindowsForTargetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowsForTargetPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeMaintenanceWindowsForTarget page.
func (p *DescribeMaintenanceWindowsForTargetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowsForTargetOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeMaintenanceWindowsForTarget(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowsForTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowsForTarget",
	}
}
