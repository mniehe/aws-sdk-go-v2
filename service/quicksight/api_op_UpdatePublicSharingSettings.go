// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use the UpdatePublicSharingSettings operation to turn on or turn off the public
// sharing settings of an Amazon QuickSight dashboard. To use this operation, turn
// on session capacity pricing for your Amazon QuickSight account. Before you can
// turn on public sharing on your account, make sure to give public sharing
// permissions to an administrative user in the Identity and Access Management
// (IAM) console. For more information on using IAM with Amazon QuickSight, see
// Using Amazon QuickSight with IAM (https://docs.aws.amazon.com/quicksight/latest/user/security_iam_service-with-iam.html)
// in the Amazon QuickSight User Guide.
func (c *Client) UpdatePublicSharingSettings(ctx context.Context, params *UpdatePublicSharingSettingsInput, optFns ...func(*Options)) (*UpdatePublicSharingSettingsOutput, error) {
	if params == nil {
		params = &UpdatePublicSharingSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePublicSharingSettings", params, optFns, c.addOperationUpdatePublicSharingSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePublicSharingSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePublicSharingSettingsInput struct {

	// The Amazon Web Services account ID associated with your Amazon QuickSight
	// subscription.
	//
	// This member is required.
	AwsAccountId *string

	// A Boolean value that indicates whether public sharing is turned on for an
	// Amazon QuickSight account.
	PublicSharingEnabled bool

	noSmithyDocumentSerde
}

type UpdatePublicSharingSettingsOutput struct {

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePublicSharingSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePublicSharingSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePublicSharingSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePublicSharingSettings"); err != nil {
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
	if err = addOpUpdatePublicSharingSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePublicSharingSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePublicSharingSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePublicSharingSettings",
	}
}
