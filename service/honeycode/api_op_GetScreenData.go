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

// The GetScreenData API allows retrieval of data from a screen in a Honeycode
// app. The API allows setting local variables in the screen to filter, sort or
// otherwise affect what will be displayed on the screen.
func (c *Client) GetScreenData(ctx context.Context, params *GetScreenDataInput, optFns ...func(*Options)) (*GetScreenDataOutput, error) {
	if params == nil {
		params = &GetScreenDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScreenData", params, optFns, c.addOperationGetScreenDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetScreenDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetScreenDataInput struct {

	// The ID of the app that contains the screen.
	//
	// This member is required.
	AppId *string

	// The ID of the screen.
	//
	// This member is required.
	ScreenId *string

	// The ID of the workbook that contains the screen.
	//
	// This member is required.
	WorkbookId *string

	// The number of results to be returned on a single page. Specify a number between
	// 1 and 100. The maximum value is 100. This parameter is optional. If you don't
	// specify this parameter, the default page size is 100.
	MaxResults *int32

	// This parameter is optional. If a nextToken is not specified, the API returns
	// the first page of data. Pagination tokens expire after 1 hour. If you use a
	// token that was returned more than an hour back, the API will throw
	// ValidationException.
	NextToken *string

	// Variables are optional and are needed only if the screen requires them to
	// render correctly. Variables are specified as a map where the key is the name of
	// the variable as defined on the screen. The value is an object which currently
	// has only one property, rawValue, which holds the value of the variable to be
	// passed to the screen.
	Variables map[string]types.VariableValue

	noSmithyDocumentSerde
}

type GetScreenDataOutput struct {

	// A map of all the rows on the screen keyed by block name.
	//
	// This member is required.
	Results map[string]types.ResultSet

	// Indicates the cursor of the workbook at which the data returned by this
	// workbook is read. Workbook cursor keeps increasing with every update and the
	// increments are not sequential.
	//
	// This member is required.
	WorkbookCursor int64

	// Provides the pagination token to load the next page if there are more results
	// matching the request. If a pagination token is not present in the response, it
	// means that all data matching the query has been loaded.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetScreenDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetScreenData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetScreenData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetScreenData"); err != nil {
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
	if err = addOpGetScreenDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScreenData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetScreenData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetScreenData",
	}
}
