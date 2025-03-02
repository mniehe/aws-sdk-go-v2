// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all Investigations.
func (c *Client) ListInvestigations(ctx context.Context, params *ListInvestigationsInput, optFns ...func(*Options)) (*ListInvestigationsOutput, error) {
	if params == nil {
		params = &ListInvestigationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvestigations", params, optFns, c.addOperationListInvestigationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvestigationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvestigationsInput struct {

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	// Filter the investigation results based on a criteria.
	FilterCriteria *types.FilterCriteria

	// List the maximum number of investigations in a page.
	MaxResults *int32

	// List if there are more results available. The value of nextToken is a unique
	// pagination token for each page. Repeat the call using the returned token to
	// retrieve the next page. Keep all other arguments unchanged. Each pagination
	// token expires after 24 hours. Using an expired pagination token will return a
	// Validation Exception error.
	NextToken *string

	// Sorts the investigation results based on a criteria.
	SortCriteria *types.SortCriteria

	noSmithyDocumentSerde
}

type ListInvestigationsOutput struct {

	// Investigations details lists the summary of uncommon behavior or malicious
	// activity which indicates a compromise.
	InvestigationDetails []types.InvestigationDetail

	// List if there are more results available. The value of nextToken is a unique
	// pagination token for each page. Repeat the call using the returned token to
	// retrieve the next page. Keep all other arguments unchanged. Each pagination
	// token expires after 24 hours. Using an expired pagination token will return an
	// HTTP 400 InvalidToken error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInvestigationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInvestigations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvestigations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInvestigations"); err != nil {
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
	if err = addOpListInvestigationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvestigations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListInvestigations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInvestigations",
	}
}
