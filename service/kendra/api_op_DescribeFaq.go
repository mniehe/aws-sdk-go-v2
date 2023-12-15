// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an FAQ list.
func (c *Client) DescribeFaq(ctx context.Context, params *DescribeFaqInput, optFns ...func(*Options)) (*DescribeFaqOutput, error) {
	if params == nil {
		params = &DescribeFaqInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFaq", params, optFns, c.addOperationDescribeFaqMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFaqOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFaqInput struct {

	// The identifier of the FAQ you want to get information on.
	//
	// This member is required.
	Id *string

	// The identifier of the index for the FAQ.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeFaqOutput struct {

	// The Unix timestamp when the FAQ was created.
	CreatedAt *time.Time

	// The description of the FAQ that you provided when it was created.
	Description *string

	// If the Status field is FAILED , the ErrorMessage field contains the reason why
	// the FAQ failed.
	ErrorMessage *string

	// The file format used by the input files for the FAQ.
	FileFormat types.FaqFileFormat

	// The identifier of the FAQ.
	Id *string

	// The identifier of the index for the FAQ.
	IndexId *string

	// The code for a language. This shows a supported language for the FAQ document.
	// English is supported by default. For more information on supported languages,
	// including their codes, see Adding documents in languages other than English (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html)
	// .
	LanguageCode *string

	// The name that you gave the FAQ when it was created.
	Name *string

	// The Amazon Resource Name (ARN) of the role that provides access to the S3
	// bucket containing the input files for the FAQ.
	RoleArn *string

	// Information required to find a specific file in an Amazon S3 bucket.
	S3Path *types.S3Path

	// The status of the FAQ. It is ready to use when the status is ACTIVE .
	Status types.FaqStatus

	// The Unix timestamp when the FAQ was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFaqMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFaq{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFaq{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFaq"); err != nil {
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
	if err = addOpDescribeFaqValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFaq(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFaq(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFaq",
	}
}
