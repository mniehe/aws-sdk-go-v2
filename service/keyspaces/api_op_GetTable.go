// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the table, including the table's name and current
// status, the keyspace name, configuration settings, and metadata. To read table
// metadata using GetTable , Select action permissions for the table and system
// tables are required to complete the operation.
func (c *Client) GetTable(ctx context.Context, params *GetTableInput, optFns ...func(*Options)) (*GetTableOutput, error) {
	if params == nil {
		params = &GetTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTable", params, optFns, c.addOperationGetTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTableInput struct {

	// The name of the keyspace that the table is stored in.
	//
	// This member is required.
	KeyspaceName *string

	// The name of the table.
	//
	// This member is required.
	TableName *string

	noSmithyDocumentSerde
}

type GetTableOutput struct {

	// The name of the keyspace that the specified table is stored in.
	//
	// This member is required.
	KeyspaceName *string

	// The Amazon Resource Name (ARN) of the specified table.
	//
	// This member is required.
	ResourceArn *string

	// The name of the specified table.
	//
	// This member is required.
	TableName *string

	// The read/write throughput capacity mode for a table. The options are:
	//   - throughputMode:PAY_PER_REQUEST
	//   - throughputMode:PROVISIONED
	CapacitySpecification *types.CapacitySpecificationSummary

	// The client-side timestamps setting of the table.
	ClientSideTimestamps *types.ClientSideTimestamps

	// The the description of the specified table.
	Comment *types.Comment

	// The creation timestamp of the specified table.
	CreationTimestamp *time.Time

	// The default Time to Live settings in seconds of the specified table.
	DefaultTimeToLive *int32

	// The encryption settings of the specified table.
	EncryptionSpecification *types.EncryptionSpecification

	// The point-in-time recovery status of the specified table.
	PointInTimeRecovery *types.PointInTimeRecoverySummary

	// The schema definition of the specified table.
	SchemaDefinition *types.SchemaDefinition

	// The current status of the specified table.
	Status types.TableStatus

	// The custom Time to Live settings of the specified table.
	Ttl *types.TimeToLive

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTable"); err != nil {
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
	if err = addOpGetTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTable",
	}
}
