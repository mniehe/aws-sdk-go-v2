// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the readme file or descriptive text for a package version. The returned
// text might contain formatting. For example, it might contain formatting for
// Markdown or reStructuredText.
func (c *Client) GetPackageVersionReadme(ctx context.Context, params *GetPackageVersionReadmeInput, optFns ...func(*Options)) (*GetPackageVersionReadmeOutput, error) {
	if params == nil {
		params = &GetPackageVersionReadmeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPackageVersionReadme", params, optFns, c.addOperationGetPackageVersionReadmeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPackageVersionReadmeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPackageVersionReadmeInput struct {

	// The name of the domain that contains the repository that contains the package
	// version with the requested readme file.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the package version with the requested
	// readme file.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package version that contains the requested readme file.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2 ).
	//
	// This member is required.
	PackageVersion *string

	// The repository that contains the package with the requested readme file.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package version with the requested readme file. The
	// package version component that specifies its namespace depends on its type. For
	// example:
	//   - The namespace of an npm package version is its scope .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	Namespace *string

	noSmithyDocumentSerde
}

type GetPackageVersionReadmeOutput struct {

	// The format of the package with the requested readme file.
	Format types.PackageFormat

	// The namespace of the package version with the requested readme file. The
	// package version component that specifies its namespace depends on its type. For
	// example:
	//   - The namespace of a Maven package version is its groupId .
	//   - The namespace of an npm package version is its scope .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	Namespace *string

	// The name of the package that contains the returned readme file.
	Package *string

	// The text of the returned readme file.
	Readme *string

	// The version of the package with the requested readme file.
	Version *string

	// The current revision associated with the package version.
	VersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPackageVersionReadmeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPackageVersionReadme{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPackageVersionReadme{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPackageVersionReadme"); err != nil {
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
	if err = addOpGetPackageVersionReadmeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPackageVersionReadme(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPackageVersionReadme(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPackageVersionReadme",
	}
}
