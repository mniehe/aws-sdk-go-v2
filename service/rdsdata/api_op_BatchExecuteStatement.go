// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs a batch SQL statement over an array of data. You can run bulk update and
// insert operations for multiple records using a DML statement with different
// parameter sets. Bulk operations can provide a significant performance
// improvement over individual insert and update operations. If a call isn't part
// of a transaction because it doesn't include the transactionID parameter, changes
// that result from the call are committed automatically.
func (c *Client) BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error) {
	if params == nil {
		params = &BatchExecuteStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchExecuteStatement", params, optFns, addOperationBatchExecuteStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a SQL statement over an array of
// data.
type BatchExecuteStatementInput struct {

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// This member is required.
	ResourceArn *string

	// The name or ARN of the secret that enables access to the DB cluster.
	//
	// This member is required.
	SecretArn *string

	// The SQL statement to run.
	//
	// This member is required.
	Sql *string

	// The name of the database.
	Database *string

	// The parameter set for the batch operation. The SQL statement is executed as many
	// times as the number of parameter sets provided. To execute a SQL statement with
	// no parameters, use one of the following options:
	//
	// * Specify one or more empty
	// parameter sets.
	//
	// * Use the ExecuteStatement operation instead of the
	// BatchExecuteStatement operation.
	//
	// Array parameters are not supported.
	ParameterSets [][]types.SqlParameter

	// The name of the database schema.
	Schema *string

	// The identifier of a transaction that was started by using the BeginTransaction
	// operation. Specify the transaction ID of the transaction that you want to
	// include the SQL statement in. If the SQL statement is not part of a transaction,
	// don't set this parameter.
	TransactionId *string
}

// The response elements represent the output of a SQL statement over an array of
// data.
type BatchExecuteStatementOutput struct {

	// The execution results of each batch entry.
	UpdateResults []types.UpdateResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchExecuteStatement{}, middleware.After)
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
	if err = addOpBatchExecuteStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchExecuteStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds-data",
		OperationName: "BatchExecuteStatement",
	}
}
