// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a compute resource to an Amazon GameLift Anywhere fleet. With
// Anywhere fleets you can incorporate your own computing hardware into an Amazon
// GameLift game hosting solution. To register a compute to a fleet, give the
// compute a name (must be unique within the fleet) and specify the compute
// resource's DNS name or IP address. Provide the Anywhere fleet ID and a fleet
// location to associate with the compute being registered. You can optionally
// include the path to a TLS certificate on the compute resource. If successful,
// this operation returns the compute details, including an Amazon GameLift SDK
// endpoint. Game server processes that run on the compute use this endpoint to
// communicate with the Amazon GameLift service. Each server process includes the
// SDK endpoint in its call to the Amazon GameLift server SDK action InitSDK() .
// Learn more
//   - Create an Anywhere fleet (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-anywhere.html)
//   - Test your integration (https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing.html)
//   - Server SDK reference guides (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html)
//     (for version 5.x)
func (c *Client) RegisterCompute(ctx context.Context, params *RegisterComputeInput, optFns ...func(*Options)) (*RegisterComputeOutput, error) {
	if params == nil {
		params = &RegisterComputeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCompute", params, optFns, c.addOperationRegisterComputeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterComputeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterComputeInput struct {

	// A descriptive label for the compute resource.
	//
	// This member is required.
	ComputeName *string

	// A unique identifier for the fleet to register the compute to. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The path to a TLS certificate on your compute resource. Amazon GameLift doesn't
	// validate the path and certificate.
	CertificatePath *string

	// The DNS name of the compute resource. Amazon GameLift requires either a DNS
	// name or IP address.
	DnsName *string

	// The IP address of the compute resource. Amazon GameLift requires either a DNS
	// name or IP address.
	IpAddress *string

	// The name of a custom location to associate with the compute resource being
	// registered.
	Location *string

	noSmithyDocumentSerde
}

type RegisterComputeOutput struct {

	// The details of the compute resource you registered.
	Compute *types.Compute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterComputeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterCompute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterCompute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterCompute"); err != nil {
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
	if err = addOpRegisterComputeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCompute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterCompute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterCompute",
	}
}
