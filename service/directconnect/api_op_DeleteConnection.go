// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes the specified connection. Deleting a connection only stops the Direct
// Connect port hour and data transfer charges. If you are partnering with any
// third parties to connect with the Direct Connect location, you must cancel your
// service with them separately.
func (c *Client) DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) {
	if params == nil {
		params = &DeleteConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnection", params, optFns, c.addOperationDeleteConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectionInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	noSmithyDocumentSerde
}

// Information about an Direct Connect connection.
type DeleteConnectionOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//   - ordering : The initial state of a hosted connection provisioned on an
	//   interconnect. The connection stays in the ordering state until the owner of the
	//   hosted connection confirms or declines the connection order.
	//   - requested : The initial state of a standard connection. The connection stays
	//   in the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//   - pending : The connection has been approved and is being initialized.
	//   - available : The network link is up and the connection is ready for use.
	//   - down : The network link is down.
	//   - deleting : The connection is being deleted.
	//   - deleted : The connection has been deleted.
	//   - rejected : A hosted connection in the ordering state enters the rejected
	//   state if it is deleted by the customer.
	//   - unknown : The state of the connection is not available.
	ConnectionState types.ConnectionState

	// The MAC Security (MACsec) connection encryption mode. The valid values are
	// no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []types.MacSecKey

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string

	// The MAC Security (MACsec) port link status of the connection. The valid values
	// are Encryption Up , which means that there is an active Connection Key Name, or
	// Encryption Down .
	PortEncryptionStatus *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []types.Tag

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteConnection"); err != nil {
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
	if err = addOpDeleteConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteConnection",
	}
}
