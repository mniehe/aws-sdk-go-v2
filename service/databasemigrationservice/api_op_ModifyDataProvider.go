// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified data provider using the provided settings. You must
// remove the data provider from all migration projects before you can modify it.
func (c *Client) ModifyDataProvider(ctx context.Context, params *ModifyDataProviderInput, optFns ...func(*Options)) (*ModifyDataProviderOutput, error) {
	if params == nil {
		params = &ModifyDataProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDataProvider", params, optFns, c.addOperationModifyDataProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDataProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDataProviderInput struct {

	// The identifier of the data provider. Identifiers must begin with a letter and
	// must contain only ASCII letters, digits, and hyphens. They can't end with a
	// hyphen, or contain two consecutive hyphens.
	//
	// This member is required.
	DataProviderIdentifier *string

	// The name of the data provider.
	DataProviderName *string

	// A user-friendly description of the data provider.
	Description *string

	// The type of database engine for the data provider. Valid values include "aurora"
	// , "aurora-postgresql" , "mysql" , "oracle" , "postgres" , "sqlserver" , redshift
	// , mariadb , mongodb , and docdb . A value of "aurora" represents Amazon Aurora
	// MySQL-Compatible Edition.
	Engine *string

	// If this attribute is Y, the current call to ModifyDataProvider replaces all
	// existing data provider settings with the exact settings that you specify in this
	// call. If this attribute is N, the current call to ModifyDataProvider does two
	// things:
	//   - It replaces any data provider settings that already exist with new values,
	//   for settings with the same names.
	//   - It creates new data provider settings that you specify in the call, for
	//   settings with different names.
	ExactSettings *bool

	// The settings in JSON format for a data provider.
	Settings types.DataProviderSettings

	noSmithyDocumentSerde
}

type ModifyDataProviderOutput struct {

	// The data provider that was modified.
	DataProvider *types.DataProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDataProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyDataProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyDataProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDataProvider"); err != nil {
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
	if err = addOpModifyDataProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDataProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDataProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDataProvider",
	}
}
