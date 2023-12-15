// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtain information about the email-sending status and capabilities of your
// Amazon SES account in the current Amazon Web Services Region.
func (c *Client) GetAccount(ctx context.Context, params *GetAccountInput, optFns ...func(*Options)) (*GetAccountOutput, error) {
	if params == nil {
		params = &GetAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccount", params, optFns, c.addOperationGetAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain information about the email-sending capabilities of your
// Amazon SES account.
type GetAccountInput struct {
	noSmithyDocumentSerde
}

// A list of details about the email-sending capabilities of your Amazon SES
// account in the current Amazon Web Services Region.
type GetAccountOutput struct {

	// Indicates whether or not the automatic warm-up feature is enabled for dedicated
	// IP addresses that are associated with your account.
	DedicatedIpAutoWarmupEnabled bool

	// An object that defines your account details.
	Details *types.AccountDetails

	// The reputation status of your Amazon SES account. The status can be one of the
	// following:
	//   - HEALTHY – There are no reputation-related issues that currently impact your
	//   account.
	//   - PROBATION – We've identified potential issues with your Amazon SES account.
	//   We're placing your account under review while you work on correcting these
	//   issues.
	//   - SHUTDOWN – Your account's ability to send email is currently paused because
	//   of an issue with the email sent from your account. When you correct the issue,
	//   you can contact us and request that your account's ability to send email is
	//   resumed.
	EnforcementStatus *string

	// Indicates whether or not your account has production access in the current
	// Amazon Web Services Region. If the value is false , then your account is in the
	// sandbox. When your account is in the sandbox, you can only send email to
	// verified identities. Additionally, the maximum number of emails you can send in
	// a 24-hour period (your sending quota) is 200, and the maximum number of emails
	// you can send per second (your maximum sending rate) is 1. If the value is true ,
	// then your account has production access. When your account has production
	// access, you can send email to any address. The sending quota and maximum sending
	// rate for your account vary based on your specific use case.
	ProductionAccessEnabled bool

	// An object that contains information about the per-day and per-second sending
	// limits for your Amazon SES account in the current Amazon Web Services Region.
	SendQuota *types.SendQuota

	// Indicates whether or not email sending is enabled for your Amazon SES account
	// in the current Amazon Web Services Region.
	SendingEnabled bool

	// An object that contains information about the email address suppression
	// preferences for your account in the current Amazon Web Services Region.
	SuppressionAttributes *types.SuppressionAttributes

	// The VDM attributes that apply to your Amazon SES account.
	VdmAttributes *types.VdmAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccount{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccount"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccount",
	}
}
