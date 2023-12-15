// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfigdata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a configuration session used to retrieve a deployed configuration. For
// more information about this API action and to view example CLI commands that
// show how to use it with the GetLatestConfiguration API action, see Retrieving
// the configuration (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration)
// in the AppConfig User Guide.
func (c *Client) StartConfigurationSession(ctx context.Context, params *StartConfigurationSessionInput, optFns ...func(*Options)) (*StartConfigurationSessionOutput, error) {
	if params == nil {
		params = &StartConfigurationSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartConfigurationSession", params, optFns, c.addOperationStartConfigurationSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartConfigurationSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartConfigurationSessionInput struct {

	// The application ID or the application name.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The configuration profile ID or the configuration profile name.
	//
	// This member is required.
	ConfigurationProfileIdentifier *string

	// The environment ID or the environment name.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// Sets a constraint on a session. If you specify a value of, for example, 60
	// seconds, then the client that established the session can't call
	// GetLatestConfiguration more frequently than every 60 seconds.
	RequiredMinimumPollIntervalInSeconds *int32

	noSmithyDocumentSerde
}

type StartConfigurationSessionOutput struct {

	// Token encapsulating state about the configuration session. Provide this token
	// to the GetLatestConfiguration API to retrieve configuration data. This token
	// should only be used once in your first call to GetLatestConfiguration . You must
	// use the new token in the GetLatestConfiguration response (
	// NextPollConfigurationToken ) in each subsequent call to GetLatestConfiguration .
	// The InitialConfigurationToken and NextPollConfigurationToken should only be
	// used once. To support long poll use cases, the tokens are valid for up to 24
	// hours. If a GetLatestConfiguration call uses an expired token, the system
	// returns BadRequestException .
	InitialConfigurationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartConfigurationSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartConfigurationSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartConfigurationSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartConfigurationSession"); err != nil {
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
	if err = addOpStartConfigurationSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartConfigurationSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartConfigurationSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartConfigurationSession",
	}
}
