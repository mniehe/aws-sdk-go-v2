// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a layer. For more information, see How to Create a Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-create.html)
// . You should use CreateLayer for noncustom layer types such as PHP App Server
// only if the stack does not have an existing layer of that type. A stack can have
// at most one instance of each noncustom layer; if you attempt to create a second
// instance, CreateLayer fails. A stack can have an arbitrary number of custom
// layers, so you can call CreateLayer as many times as you like for that layer
// type. Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html)
// .
func (c *Client) CreateLayer(ctx context.Context, params *CreateLayerInput, optFns ...func(*Options)) (*CreateLayerOutput, error) {
	if params == nil {
		params = &CreateLayerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLayer", params, optFns, c.addOperationCreateLayerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLayerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLayerInput struct {

	// The layer name, which is used by the console.
	//
	// This member is required.
	Name *string

	// For custom layers only, use this parameter to specify the layer's short name,
	// which is used internally by AWS OpsWorks Stacks and by Chef recipes. The short
	// name is also used as the name for the directory where your app files are
	// installed. It can have a maximum of 200 characters, which are limited to the
	// alphanumeric characters, '-', '_', and '.'. The built-in layers' short names are
	// defined by AWS OpsWorks Stacks. For more information, see the Layer Reference (https://docs.aws.amazon.com/opsworks/latest/userguide/layers.html)
	// .
	//
	// This member is required.
	Shortname *string

	// The layer stack ID.
	//
	// This member is required.
	StackId *string

	// The layer type. A stack cannot have more than one built-in layer of the same
	// type. It can have any number of custom layers. Built-in layers are not available
	// in Chef 12 stacks.
	//
	// This member is required.
	Type types.LayerType

	// One or more user-defined key-value pairs to be added to the stack attributes.
	// To create a cluster layer, set the EcsClusterArn attribute to the cluster's ARN.
	Attributes map[string]string

	// Whether to automatically assign an Elastic IP address (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// to the layer's instances. For more information, see How to Edit a Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html)
	// .
	AutoAssignElasticIps *bool

	// For stacks that are running in a VPC, whether to automatically assign a public
	// IP address to the layer's instances. For more information, see How to Edit a
	// Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-edit.html)
	// .
	AutoAssignPublicIps *bool

	// Specifies CloudWatch Logs configuration options for the layer. For more
	// information, see CloudWatchLogsLogStream .
	CloudWatchLogsConfiguration *types.CloudWatchLogsConfiguration

	// The ARN of an IAM profile to be used for the layer's EC2 instances. For more
	// information about IAM ARNs, see Using Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// .
	CustomInstanceProfileArn *string

	// A JSON-formatted string containing custom stack configuration and deployment
	// attributes to be installed on the layer's instances. For more information, see
	// Using Custom JSON (https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook-json-override.html)
	// . This feature is supported as of version 1.7.42 of the AWS CLI.
	CustomJson *string

	// A LayerCustomRecipes object that specifies the layer custom recipes.
	CustomRecipes *types.Recipes

	// An array containing the layer custom security group IDs.
	CustomSecurityGroupIds []string

	// Whether to disable auto healing for the layer.
	EnableAutoHealing *bool

	// Whether to install operating system and package updates when the instance
	// boots. The default value is true . To control when updates are installed, set
	// this value to false . You must then update your instances manually by using
	// CreateDeployment to run the update_dependencies stack command or by manually
	// running yum (Amazon Linux) or apt-get (Ubuntu) on the instances. To ensure that
	// your instances have the latest security updates, we strongly recommend using the
	// default value of true .
	InstallUpdatesOnBoot *bool

	// A LifeCycleEventConfiguration object that you can use to configure the Shutdown
	// event to specify an execution timeout and enable or disable Elastic Load
	// Balancer connection draining.
	LifecycleEventConfiguration *types.LifecycleEventConfiguration

	// An array of Package objects that describes the layer packages.
	Packages []string

	// Whether to use Amazon EBS-optimized instances.
	UseEbsOptimizedInstances *bool

	// A VolumeConfigurations object that describes the layer's Amazon EBS volumes.
	VolumeConfigurations []types.VolumeConfiguration

	noSmithyDocumentSerde
}

// Contains the response to a CreateLayer request.
type CreateLayerOutput struct {

	// The layer ID.
	LayerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLayerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLayer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLayer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLayer"); err != nil {
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
	if err = addOpCreateLayerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLayer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLayer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLayer",
	}
}
