// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a stack to start streaming applications to users. A stack consists of
// an associated fleet, user access policies, and storage configurations.
func (c *Client) CreateStack(ctx context.Context, params *CreateStackInput, optFns ...func(*Options)) (*CreateStackOutput, error) {
	if params == nil {
		params = &CreateStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStack", params, optFns, c.addOperationCreateStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStackInput struct {

	// The name of the stack.
	//
	// This member is required.
	Name *string

	// The list of interface VPC endpoint (interface endpoint) objects. Users of the
	// stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []types.AccessEndpoint

	// The persistent application settings for users of a stack. When these settings
	// are enabled, changes that users make to applications and Windows settings are
	// automatically saved after each session and applied to the next session.
	ApplicationSettings *types.ApplicationSettings

	// The description to display.
	Description *string

	// The stack name to display.
	DisplayName *string

	// The domains where AppStream 2.0 streaming sessions can be embedded in an
	// iframe. You must approve the domains that you want to host embedded AppStream
	// 2.0 streaming sessions.
	EmbedHostDomains []string

	// The URL that users are redirected to after they click the Send Feedback link.
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string

	// The storage connectors to enable.
	StorageConnectors []types.StorageConnector

	// The streaming protocol you want your stack to prefer. This can be UDP or TCP.
	// Currently, UDP is only supported in the Windows native client.
	StreamingExperienceSettings *types.StreamingExperienceSettings

	// The tags to associate with the stack. A tag is a key-value pair, and the value
	// is optional. For example, Environment=Test. If you do not specify a value,
	// Environment=. If you do not specify a value, the value is set to an empty
	// string. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: _ . : / = + \ - @
	// For more information about tags, see Tagging Your Resources (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default, these actions are enabled.
	UserSettings []types.UserSetting

	noSmithyDocumentSerde
}

type CreateStackOutput struct {

	// Information about the stack.
	Stack *types.Stack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStack{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateStack"); err != nil {
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
	if err = addOpCreateStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateStack",
	}
}
