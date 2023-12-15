// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a database user name and temporary password with temporary
// authorization to log in to Amazon Redshift Serverless. By default, the temporary
// credentials expire in 900 seconds. You can optionally specify a duration between
// 900 seconds (15 minutes) and 3600 seconds (60 minutes). The Identity and Access
// Management (IAM) user or role that runs GetCredentials must have an IAM policy
// attached that allows access to all necessary actions and resources. If the
// DbName parameter is specified, the IAM policy must allow access to the resource
// dbname for the specified database name.
func (c *Client) GetCredentials(ctx context.Context, params *GetCredentialsInput, optFns ...func(*Options)) (*GetCredentialsOutput, error) {
	if params == nil {
		params = &GetCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCredentials", params, optFns, c.addOperationGetCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCredentialsInput struct {

	// The custom domain name associated with the workgroup. The custom domain name or
	// the workgroup name must be included in the request.
	CustomDomainName *string

	// The name of the database to get temporary authorization to log on to.
	// Constraints:
	//   - Must be 1 to 64 alphanumeric characters or hyphens.
	//   - Must contain only uppercase or lowercase letters, numbers, underscore, plus
	//   sign, period (dot), at symbol (@), or hyphen.
	//   - The first character must be a letter.
	//   - Must not contain a colon ( : ) or slash ( / ).
	//   - Cannot be a reserved word. A list of reserved words can be found in
	//   Reserved Words  (https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html)
	//   in the Amazon Redshift Database Developer Guide
	DbName *string

	// The number of seconds until the returned temporary password expires. The
	// minimum is 900 seconds, and the maximum is 3600 seconds.
	DurationSeconds *int32

	// The name of the workgroup associated with the database.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type GetCredentialsOutput struct {

	// A temporary password that authorizes the user name returned by DbUser to log on
	// to the database DbName .
	DbPassword *string

	// A database user name that is authorized to log on to the database DbName using
	// the password DbPassword . If the specified DbUser exists in the database, the
	// new user name has the same database privileges as the the user named in DbUser .
	// By default, the user is added to PUBLIC.
	DbUser *string

	// The date and time the password in DbPassword expires.
	Expiration *time.Time

	// The date and time of when the DbUser and DbPassword authorization refreshes.
	NextRefreshTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCredentials"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCredentials",
	}
}
