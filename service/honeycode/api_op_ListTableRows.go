// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListTableRows API allows you to retrieve a list of all the rows in a table
// in a workbook.
func (c *Client) ListTableRows(ctx context.Context, params *ListTableRowsInput, optFns ...func(*Options)) (*ListTableRowsOutput, error) {
	if params == nil {
		params = &ListTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTableRows", params, optFns, c.addOperationListTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTableRowsInput struct {

	// The ID of the table whose rows are being retrieved. If a table with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook that contains the table whose rows are being retrieved.
	// If a workbook with the specified id could not be found, this API throws
	// ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The maximum number of rows to return in each page of the results.
	MaxResults *int32

	// This parameter is optional. If a nextToken is not specified, the API returns
	// the first page of data. Pagination tokens expire after 1 hour. If you use a
	// token that was returned more than an hour back, the API will throw
	// ValidationException.
	NextToken *string

	// This parameter is optional. If one or more row ids are specified in this list,
	// then only the specified row ids are returned in the result. If no row ids are
	// specified here, then all the rows in the table are returned.
	RowIds []string

	noSmithyDocumentSerde
}

type ListTableRowsOutput struct {

	// The list of columns in the table whose row data is returned in the result.
	//
	// This member is required.
	ColumnIds []string

	// The list of rows in the table. Note that this result is paginated, so this list
	// contains a maximum of 100 rows.
	//
	// This member is required.
	Rows []types.TableRow

	// Indicates the cursor of the workbook at which the data returned by this request
	// is read. Workbook cursor keeps increasing with every update and the increments
	// are not sequential.
	//
	// This member is required.
	WorkbookCursor int64

	// Provides the pagination token to load the next page if there are more results
	// matching the request. If a pagination token is not present in the response, it
	// means that all data matching the request has been loaded.
	NextToken *string

	// The list of row ids included in the request that were not found in the table.
	RowIdsNotFound []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTableRows"); err != nil {
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
	if err = addOpListTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTableRows(options.Region), middleware.Before); err != nil {
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

// ListTableRowsAPIClient is a client that implements the ListTableRows operation.
type ListTableRowsAPIClient interface {
	ListTableRows(context.Context, *ListTableRowsInput, ...func(*Options)) (*ListTableRowsOutput, error)
}

var _ ListTableRowsAPIClient = (*Client)(nil)

// ListTableRowsPaginatorOptions is the paginator options for ListTableRows
type ListTableRowsPaginatorOptions struct {
	// The maximum number of rows to return in each page of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTableRowsPaginator is a paginator for ListTableRows
type ListTableRowsPaginator struct {
	options   ListTableRowsPaginatorOptions
	client    ListTableRowsAPIClient
	params    *ListTableRowsInput
	nextToken *string
	firstPage bool
}

// NewListTableRowsPaginator returns a new ListTableRowsPaginator
func NewListTableRowsPaginator(client ListTableRowsAPIClient, params *ListTableRowsInput, optFns ...func(*ListTableRowsPaginatorOptions)) *ListTableRowsPaginator {
	if params == nil {
		params = &ListTableRowsInput{}
	}

	options := ListTableRowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTableRowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTableRowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTableRows page.
func (p *ListTableRowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTableRowsOutput, error) {
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

	result, err := p.client.ListTableRows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTableRows",
	}
}
