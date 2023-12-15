// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an RDS event notification subscription. This operation requires a topic
// Amazon Resource Name (ARN) created by either the RDS console, the SNS console,
// or the SNS API. To obtain an ARN with SNS, you must create a topic in Amazon SNS
// and subscribe to the topic. The ARN is displayed in the SNS console. You can
// specify the type of source ( SourceType ) that you want to be notified of and
// provide a list of RDS sources ( SourceIds ) that triggers the events. You can
// also provide a list of event categories ( EventCategories ) for events that you
// want to be notified of. For example, you can specify SourceType = db-instance ,
// SourceIds = mydbinstance1 , mydbinstance2 and EventCategories = Availability ,
// Backup . If you specify both the SourceType and SourceIds , such as SourceType
// = db-instance and SourceIds = myDBInstance1 , you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify SourceIds , you receive notice of the events for that source type
// for all your RDS sources. If you don't specify either the SourceType or the
// SourceIds , you are notified of events generated from all RDS sources belonging
// to your customer account. For more information about subscribing to an event for
// RDS DB engines, see Subscribing to Amazon RDS event notification (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Subscribing.html)
// in the Amazon RDS User Guide. For more information about subscribing to an event
// for Aurora DB engines, see Subscribing to Amazon RDS event notification (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Subscribing.html)
// in the Amazon Aurora User Guide.
func (c *Client) CreateEventSubscription(ctx context.Context, params *CreateEventSubscriptionInput, optFns ...func(*Options)) (*CreateEventSubscriptionOutput, error) {
	if params == nil {
		params = &CreateEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventSubscription", params, optFns, c.addOperationCreateEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventSubscriptionInput struct {

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	//
	// This member is required.
	SnsTopicArn *string

	// The name of the subscription. Constraints: The name must be less than 255
	// characters.
	//
	// This member is required.
	SubscriptionName *string

	// Specifies whether to activate the subscription. If the event notification
	// subscription isn't activated, the subscription is created but not active.
	Enabled *bool

	// A list of event categories for a particular source type ( SourceType ) that you
	// want to subscribe to. You can see a list of the categories for a given source
	// type in the "Amazon RDS event categories and event messages" section of the
	// Amazon RDS User Guide  (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html)
	// or the Amazon Aurora User Guide  (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html)
	// . You can also see this list by using the DescribeEventCategories operation.
	EventCategories []string

	// The list of identifiers of the event sources for which events are returned. If
	// not specified, then all sources are included in the response. An identifier must
	// begin with a letter and must contain only ASCII letters, digits, and hyphens. It
	// can't end with a hyphen or contain two consecutive hyphens. Constraints:
	//   - If SourceIds are supplied, SourceType must also be provided.
	//   - If the source type is a DB instance, a DBInstanceIdentifier value must be
	//   supplied.
	//   - If the source type is a DB cluster, a DBClusterIdentifier value must be
	//   supplied.
	//   - If the source type is a DB parameter group, a DBParameterGroupName value
	//   must be supplied.
	//   - If the source type is a DB security group, a DBSecurityGroupName value must
	//   be supplied.
	//   - If the source type is a DB snapshot, a DBSnapshotIdentifier value must be
	//   supplied.
	//   - If the source type is a DB cluster snapshot, a DBClusterSnapshotIdentifier
	//   value must be supplied.
	//   - If the source type is an RDS Proxy, a DBProxyName value must be supplied.
	SourceIds []string

	// The type of source that is generating the events. For example, if you want to
	// be notified of events generated by a DB instance, you set this parameter to
	// db-instance . For RDS Proxy events, specify db-proxy . If this value isn't
	// specified, all events are returned. Valid Values: db-instance | db-cluster |
	// db-parameter-group | db-security-group | db-snapshot | db-cluster-snapshot |
	// db-proxy
	SourceType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html)
	// in the Amazon RDS User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEventSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEventSubscription"); err != nil {
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
	if err = addOpCreateEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEventSubscription",
	}
}
