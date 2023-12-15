// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation updates the contact information for a particular domain. You
// must specify information for at least one contact: registrant, administrator, or
// technical. If the update is successful, this method returns an operation ID that
// you can use to track the progress and completion of the operation. If the
// request is not completed successfully, the domain registrant will be notified by
// email.
func (c *Client) UpdateDomainContact(ctx context.Context, params *UpdateDomainContactInput, optFns ...func(*Options)) (*UpdateDomainContactOutput, error) {
	if params == nil {
		params = &UpdateDomainContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainContact", params, optFns, c.addOperationUpdateDomainContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The UpdateDomainContact request includes the following elements.
type UpdateDomainContactInput struct {

	// The name of the domain that you want to update contact information for.
	//
	// This member is required.
	DomainName *string

	// Provides detailed contact information.
	AdminContact *types.ContactDetail

	// Customer's consent for the owner change request. Required if the domain is not
	// free (consent price is more than $0.00).
	Consent *types.Consent

	// Provides detailed contact information.
	RegistrantContact *types.ContactDetail

	// Provides detailed contact information.
	TechContact *types.ContactDetail

	noSmithyDocumentSerde
}

// The UpdateDomainContact response includes the following element.
type UpdateDomainContactOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDomainContact"); err != nil {
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
	if err = addOpUpdateDomainContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainContact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDomainContact",
	}
}
