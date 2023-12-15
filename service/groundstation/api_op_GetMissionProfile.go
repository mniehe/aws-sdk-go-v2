// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a mission profile.
func (c *Client) GetMissionProfile(ctx context.Context, params *GetMissionProfileInput, optFns ...func(*Options)) (*GetMissionProfileOutput, error) {
	if params == nil {
		params = &GetMissionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMissionProfile", params, optFns, c.addOperationGetMissionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMissionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMissionProfileInput struct {

	// UUID of a mission profile.
	//
	// This member is required.
	MissionProfileId *string

	noSmithyDocumentSerde
}

type GetMissionProfileOutput struct {

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPostPassDurationSeconds *int32

	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	ContactPrePassDurationSeconds *int32

	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config and a
	// to Config .
	DataflowEdges [][]string

	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than this
	// duration.
	MinimumViableContactDurationSeconds *int32

	// ARN of a mission profile.
	MissionProfileArn *string

	// UUID of a mission profile.
	MissionProfileId *string

	// Name of a mission profile.
	Name *string

	// Region of a mission profile.
	Region *string

	// KMS key to use for encrypting streams.
	StreamsKmsKey types.KmsKey

	// Role to use for encrypting streams with KMS key.
	StreamsKmsRole *string

	// Tags assigned to a mission profile.
	Tags map[string]string

	// ARN of a tracking Config .
	TrackingConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMissionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMissionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMissionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMissionProfile"); err != nil {
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
	if err = addOpGetMissionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMissionProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMissionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMissionProfile",
	}
}
