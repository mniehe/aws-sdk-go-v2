// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is used by SaaS partners to write events to a customer's partner event
// bus. Amazon Web Services customers do not use this operation. For information on
// calculating event batch size, see Calculating EventBridge PutEvents event entry
// size (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-putevent-size.html)
// in the EventBridge User Guide.
func (c *Client) PutPartnerEvents(ctx context.Context, params *PutPartnerEventsInput, optFns ...func(*Options)) (*PutPartnerEventsOutput, error) {
	if params == nil {
		params = &PutPartnerEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPartnerEvents", params, optFns, c.addOperationPutPartnerEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPartnerEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPartnerEventsInput struct {

	// The list of events to write to the event bus.
	//
	// This member is required.
	Entries []types.PutPartnerEventsRequestEntry

	noSmithyDocumentSerde
}

type PutPartnerEventsOutput struct {

	// The results for each event entry the partner submitted in this request. If the
	// event was successfully submitted, the entry has the event ID in it. Otherwise,
	// you can use the error code and error message to identify the problem with the
	// entry. For each record, the index of the response element is the same as the
	// index in the request array.
	Entries []types.PutPartnerEventsResultEntry

	// The number of events from this operation that could not be written to the
	// partner event bus.
	FailedEntryCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPartnerEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPartnerEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPartnerEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutPartnerEvents"); err != nil {
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
	if err = addOpPutPartnerEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPartnerEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPartnerEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutPartnerEvents",
	}
}
