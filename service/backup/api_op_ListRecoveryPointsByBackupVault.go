// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns detailed information about the recovery points stored in a backup vault.
func (c *Client) ListRecoveryPointsByBackupVault(ctx context.Context, params *ListRecoveryPointsByBackupVaultInput, optFns ...func(*Options)) (*ListRecoveryPointsByBackupVaultOutput, error) {
	if params == nil {
		params = &ListRecoveryPointsByBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecoveryPointsByBackupVault", params, optFns, c.addOperationListRecoveryPointsByBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecoveryPointsByBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecoveryPointsByBackupVaultInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens. Backup vault name might not be available when a
	// supported service creates the backup.
	//
	// This member is required.
	BackupVaultName *string

	// This parameter will sort the list of recovery points by account ID.
	BackupVaultAccountId *string

	// Returns only recovery points that match the specified backup plan ID.
	ByBackupPlanId *string

	// Returns only recovery points that were created after the specified timestamp.
	ByCreatedAfter *time.Time

	// Returns only recovery points that were created before the specified timestamp.
	ByCreatedBefore *time.Time

	// This returns only recovery points that match the specified parent (composite)
	// recovery point Amazon Resource Name (ARN).
	ByParentRecoveryPointArn *string

	// Returns only recovery points that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string

	// Returns only recovery points that match the specified resource type(s):
	//   - Aurora for Amazon Aurora
	//   - CloudFormation for CloudFormation
	//   - DocumentDB for Amazon DocumentDB (with MongoDB compatibility)
	//   - DynamoDB for Amazon DynamoDB
	//   - EBS for Amazon Elastic Block Store
	//   - EC2 for Amazon Elastic Compute Cloud
	//   - EFS for Amazon Elastic File System
	//   - FSx for Amazon FSx
	//   - Neptune for Amazon Neptune
	//   - Redshift for Amazon Redshift
	//   - RDS for Amazon Relational Database Service
	//   - SAP HANA on Amazon EC2 for SAP HANA databases
	//   - Storage Gateway for Storage Gateway
	//   - S3 for Amazon S3
	//   - Timestream for Amazon Timestream
	//   - VirtualMachine for virtual machines
	ByResourceType *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecoveryPointsByBackupVaultOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return MaxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// An array of objects that contain detailed information about recovery points
	// saved in a backup vault.
	RecoveryPoints []types.RecoveryPointByBackupVault

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecoveryPointsByBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecoveryPointsByBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecoveryPointsByBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecoveryPointsByBackupVault"); err != nil {
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
	if err = addOpListRecoveryPointsByBackupVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecoveryPointsByBackupVault(options.Region), middleware.Before); err != nil {
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

// ListRecoveryPointsByBackupVaultAPIClient is a client that implements the
// ListRecoveryPointsByBackupVault operation.
type ListRecoveryPointsByBackupVaultAPIClient interface {
	ListRecoveryPointsByBackupVault(context.Context, *ListRecoveryPointsByBackupVaultInput, ...func(*Options)) (*ListRecoveryPointsByBackupVaultOutput, error)
}

var _ ListRecoveryPointsByBackupVaultAPIClient = (*Client)(nil)

// ListRecoveryPointsByBackupVaultPaginatorOptions is the paginator options for
// ListRecoveryPointsByBackupVault
type ListRecoveryPointsByBackupVaultPaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecoveryPointsByBackupVaultPaginator is a paginator for
// ListRecoveryPointsByBackupVault
type ListRecoveryPointsByBackupVaultPaginator struct {
	options   ListRecoveryPointsByBackupVaultPaginatorOptions
	client    ListRecoveryPointsByBackupVaultAPIClient
	params    *ListRecoveryPointsByBackupVaultInput
	nextToken *string
	firstPage bool
}

// NewListRecoveryPointsByBackupVaultPaginator returns a new
// ListRecoveryPointsByBackupVaultPaginator
func NewListRecoveryPointsByBackupVaultPaginator(client ListRecoveryPointsByBackupVaultAPIClient, params *ListRecoveryPointsByBackupVaultInput, optFns ...func(*ListRecoveryPointsByBackupVaultPaginatorOptions)) *ListRecoveryPointsByBackupVaultPaginator {
	if params == nil {
		params = &ListRecoveryPointsByBackupVaultInput{}
	}

	options := ListRecoveryPointsByBackupVaultPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecoveryPointsByBackupVaultPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecoveryPointsByBackupVaultPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecoveryPointsByBackupVault page.
func (p *ListRecoveryPointsByBackupVaultPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecoveryPointsByBackupVaultOutput, error) {
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

	result, err := p.client.ListRecoveryPointsByBackupVault(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecoveryPointsByBackupVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecoveryPointsByBackupVault",
	}
}
