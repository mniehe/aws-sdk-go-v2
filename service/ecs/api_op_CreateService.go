// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs and maintains your desired number of tasks from a specified task
// definition. If the number of tasks running in a service drops below the
// desiredCount , Amazon ECS runs another copy of the task in the specified
// cluster. To update an existing service, see the UpdateService action. Starting
// April 15, 2023, Amazon Web Services will not onboard new customers to Amazon
// Elastic Inference (EI), and will help current customers migrate their workloads
// to options that offer better price and performance. After April 15, 2023, new
// customers will not be able to launch instances with Amazon EI accelerators in
// Amazon SageMaker, Amazon ECS, or Amazon EC2. However, customers who have used
// Amazon EI at least once during the past 30-day period are considered current
// customers and will be able to continue using the service. In addition to
// maintaining the desired count of tasks in your service, you can optionally run
// your service behind one or more load balancers. The load balancers distribute
// traffic across the tasks that are associated with the service. For more
// information, see Service load balancing (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html)
// in the Amazon Elastic Container Service Developer Guide. Tasks for services that
// don't use a load balancer are considered healthy if they're in the RUNNING
// state. Tasks for services that use a load balancer are considered healthy if
// they're in the RUNNING state and are reported as healthy by the load balancer.
// There are two service scheduler strategies available:
//   - REPLICA - The replica scheduling strategy places and maintains your desired
//     number of tasks across your cluster. By default, the service scheduler spreads
//     tasks across Availability Zones. You can use task placement strategies and
//     constraints to customize task placement decisions. For more information, see
//     Service scheduler concepts (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html)
//     in the Amazon Elastic Container Service Developer Guide.
//   - DAEMON - The daemon scheduling strategy deploys exactly one task on each
//     active container instance that meets all of the task placement constraints that
//     you specify in your cluster. The service scheduler also evaluates the task
//     placement constraints for running tasks. It also stops tasks that don't meet the
//     placement constraints. When using this strategy, you don't need to specify a
//     desired number of tasks, a task placement strategy, or use Service Auto Scaling
//     policies. For more information, see Service scheduler concepts (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html)
//     in the Amazon Elastic Container Service Developer Guide.
//
// You can optionally specify a deployment configuration for your service. The
// deployment is initiated by changing properties. For example, the deployment
// might be initiated by the task definition or by your desired count of a service.
// This is done with an UpdateService operation. The default value for a replica
// service for minimumHealthyPercent is 100%. The default value for a daemon
// service for minimumHealthyPercent is 0%. If a service uses the ECS deployment
// controller, the minimum healthy percent represents a lower limit on the number
// of tasks in a service that must remain in the RUNNING state during a
// deployment. Specifically, it represents it as a percentage of your desired
// number of tasks (rounded up to the nearest integer). This happens when any of
// your container instances are in the DRAINING state if the service contains
// tasks using the EC2 launch type. Using this parameter, you can deploy without
// using additional cluster capacity. For example, if you set your service to have
// desired number of four tasks and a minimum healthy percent of 50%, the scheduler
// might stop two existing tasks to free up cluster capacity before starting two
// new tasks. If they're in the RUNNING state, tasks for services that don't use a
// load balancer are considered healthy . If they're in the RUNNING state and
// reported as healthy by the load balancer, tasks for services that do use a load
// balancer are considered healthy . The default value for minimum healthy percent
// is 100%. If a service uses the ECS deployment controller, the maximum percent
// parameter represents an upper limit on the number of tasks in a service that are
// allowed in the RUNNING or PENDING state during a deployment. Specifically, it
// represents it as a percentage of the desired number of tasks (rounded down to
// the nearest integer). This happens when any of your container instances are in
// the DRAINING state if the service contains tasks using the EC2 launch type.
// Using this parameter, you can define the deployment batch size. For example, if
// your service has a desired number of four tasks and a maximum percent value of
// 200%, the scheduler may start four new tasks before stopping the four older
// tasks (provided that the cluster resources required to do this are available).
// The default value for maximum percent is 200%. If a service uses either the
// CODE_DEPLOY or EXTERNAL deployment controller types and tasks that use the EC2
// launch type, the minimum healthy percent and maximum percent values are used
// only to define the lower and upper limit on the number of the tasks in the
// service that remain in the RUNNING state. This is while the container instances
// are in the DRAINING state. If the tasks in the service use the Fargate launch
// type, the minimum healthy percent and maximum percent values aren't used. This
// is the case even if they're currently visible when describing your service. When
// creating a service that uses the EXTERNAL deployment controller, you can
// specify only parameters that aren't controlled at the task set level. The only
// required parameter is the service name. You control your services using the
// CreateTaskSet operation. For more information, see Amazon ECS deployment types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide. When the service
// scheduler launches new tasks, it determines task placement. For information
// about task placement and task placement strategies, see Amazon ECS task
// placement (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	if params == nil {
		params = &CreateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateService", params, optFns, c.addOperationCreateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {

	// The name of your service. Up to 255 letters (uppercase and lowercase), numbers,
	// underscores, and hyphens are allowed. Service names must be unique within a
	// cluster, but you can have similarly named services in multiple clusters within a
	// Region or across multiple Regions.
	//
	// This member is required.
	ServiceName *string

	// The capacity provider strategy to use for the service. If a
	// capacityProviderStrategy is specified, the launchType parameter must be
	// omitted. If no capacityProviderStrategy or launchType is specified, the
	// defaultCapacityProviderStrategy for the cluster is used. A capacity provider
	// strategy may contain a maximum of 6 capacity providers.
	CapacityProviderStrategy []types.CapacityProviderStrategyItem

	// An identifier that you provide to ensure the idempotency of the request. It
	// must be unique and is case sensitive. Up to 36 ASCII characters in the range of
	// 33-126 (inclusive) are allowed.
	ClientToken *string

	// The short name or full Amazon Resource Name (ARN) of the cluster that you run
	// your service on. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string

	// Optional deployment parameters that control how many tasks run during the
	// deployment and the ordering of stopping and starting tasks.
	DeploymentConfiguration *types.DeploymentConfiguration

	// The deployment controller to use for the service. If no deployment controller
	// is specified, the default value of ECS is used.
	DeploymentController *types.DeploymentController

	// The number of instantiations of the specified task definition to place and keep
	// running in your service. This is required if schedulingStrategy is REPLICA or
	// isn't specified. If schedulingStrategy is DAEMON then this isn't required.
	DesiredCount *int32

	// Specifies whether to turn on Amazon ECS managed tags for the tasks within the
	// service. For more information, see Tagging your Amazon ECS resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide. When you use Amazon ECS
	// managed tags, you need to set the propagateTags request parameter.
	EnableECSManagedTags bool

	// Determines whether the execute command functionality is turned on for the
	// service. If true , this enables execute command functionality on all containers
	// in the service tasks.
	EnableExecuteCommand bool

	// The period of time, in seconds, that the Amazon ECS service scheduler ignores
	// unhealthy Elastic Load Balancing target health checks after a task has first
	// started. This is only used when your service is configured to use a load
	// balancer. If your service has a load balancer defined and you don't specify a
	// health check grace period value, the default value of 0 is used. If you do not
	// use an Elastic Load Balancing, we recommend that you use the startPeriod in the
	// task definition health check parameters. For more information, see Health check (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html)
	// . If your service's tasks take a while to start and respond to Elastic Load
	// Balancing health checks, you can specify a health check grace period of up to
	// 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service
	// scheduler ignores health check status. This grace period can prevent the service
	// scheduler from marking tasks as unhealthy and stopping them before they have
	// time to come up.
	HealthCheckGracePeriodSeconds *int32

	// The infrastructure that you run your service on. For more information, see
	// Amazon ECS launch types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// in the Amazon Elastic Container Service Developer Guide. The FARGATE launch
	// type runs your tasks on Fargate On-Demand infrastructure. Fargate Spot
	// infrastructure is available for use but a capacity provider strategy must be
	// used. For more information, see Fargate capacity providers (https://docs.aws.amazon.com/AmazonECS/latest/userguide/fargate-capacity-providers.html)
	// in the Amazon ECS User Guide for Fargate. The EC2 launch type runs your tasks
	// on Amazon EC2 instances registered to your cluster. The EXTERNAL launch type
	// runs your tasks on your on-premises server or virtual machine (VM) capacity
	// registered to your cluster. A service can use either a launch type or a capacity
	// provider strategy. If a launchType is specified, the capacityProviderStrategy
	// parameter must be omitted.
	LaunchType types.LaunchType

	// A load balancer object representing the load balancers to use with your
	// service. For more information, see Service load balancing (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html)
	// in the Amazon Elastic Container Service Developer Guide. If the service uses the
	// rolling update ( ECS ) deployment controller and using either an Application
	// Load Balancer or Network Load Balancer, you must specify one or more target
	// group ARNs to attach to the service. The service-linked role is required for
	// services that use multiple target groups. For more information, see Using
	// service-linked roles for Amazon ECS (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html)
	// in the Amazon Elastic Container Service Developer Guide. If the service uses the
	// CODE_DEPLOY deployment controller, the service is required to use either an
	// Application Load Balancer or Network Load Balancer. When creating an CodeDeploy
	// deployment group, you specify two target groups (referred to as a
	// targetGroupPair ). During a deployment, CodeDeploy determines which task set in
	// your service has the status PRIMARY , and it associates one target group with
	// it. Then, it also associates the other target group with the replacement task
	// set. The load balancer can also have up to two listeners: a required listener
	// for production traffic and an optional listener that you can use to perform
	// validation tests with Lambda functions before routing production traffic to it.
	// If you use the CODE_DEPLOY deployment controller, these values can be changed
	// when updating the service. For Application Load Balancers and Network Load
	// Balancers, this object must contain the load balancer target group ARN, the
	// container name, and the container port to access from the load balancer. The
	// container name must be as it appears in a container definition. The load
	// balancer name parameter must be omitted. When a task from this service is placed
	// on a container instance, the container instance and port combination is
	// registered as a target in the target group that's specified here. For Classic
	// Load Balancers, this object must contain the load balancer name, the container
	// name , and the container port to access from the load balancer. The container
	// name must be as it appears in a container definition. The target group ARN
	// parameter must be omitted. When a task from this service is placed on a
	// container instance, the container instance is registered with the load balancer
	// that's specified here. Services with tasks that use the awsvpc network mode
	// (for example, those with the Fargate launch type) only support Application Load
	// Balancers and Network Load Balancers. Classic Load Balancers aren't supported.
	// Also, when you create any target groups for these services, you must choose ip
	// as the target type, not instance . This is because tasks that use the awsvpc
	// network mode are associated with an elastic network interface, not an Amazon EC2
	// instance.
	LoadBalancers []types.LoadBalancer

	// The network configuration for the service. This parameter is required for task
	// definitions that use the awsvpc network mode to receive their own elastic
	// network interface, and it isn't supported for other network modes. For more
	// information, see Task networking (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
	// in the Amazon Elastic Container Service Developer Guide.
	NetworkConfiguration *types.NetworkConfiguration

	// An array of placement constraint objects to use for tasks in your service. You
	// can specify a maximum of 10 constraints for each task. This limit includes
	// constraints in the task definition and those specified at runtime.
	PlacementConstraints []types.PlacementConstraint

	// The placement strategy objects to use for tasks in your service. You can
	// specify a maximum of 5 strategy rules for each service.
	PlacementStrategy []types.PlacementStrategy

	// The platform version that your tasks in the service are running on. A platform
	// version is specified only for tasks using the Fargate launch type. If one isn't
	// specified, the LATEST platform version is used. For more information, see
	// Fargate platform versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion *string

	// Specifies whether to propagate the tags from the task definition to the task.
	// If no value is specified, the tags aren't propagated. Tags can only be
	// propagated to the task during task creation. To add tags to a task after task
	// creation, use the TagResource (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html)
	// API action. The default is NONE .
	PropagateTags types.PropagateTags

	// The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon
	// ECS to make calls to your load balancer on your behalf. This parameter is only
	// permitted if you are using a load balancer with your service and your task
	// definition doesn't use the awsvpc network mode. If you specify the role
	// parameter, you must also specify a load balancer object with the loadBalancers
	// parameter. If your account has already created the Amazon ECS service-linked
	// role, that role is used for your service unless you specify a role here. The
	// service-linked role is required if your task definition uses the awsvpc network
	// mode or if the service is configured to use service discovery, an external
	// deployment controller, multiple target groups, or Elastic Inference accelerators
	// in which case you don't specify a role here. For more information, see Using
	// service-linked roles for Amazon ECS (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html)
	// in the Amazon Elastic Container Service Developer Guide. If your specified role
	// has a path other than / , then you must either specify the full role ARN (this
	// is recommended) or prefix the role name with the path. For example, if a role
	// with the name bar has a path of /foo/ then you would specify /foo/bar as the
	// role name. For more information, see Friendly names and paths (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names)
	// in the IAM User Guide.
	Role *string

	// The scheduling strategy to use for the service. For more information, see
	// Services (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html)
	// . There are two service scheduler strategies available:
	//   - REPLICA -The replica scheduling strategy places and maintains the desired
	//   number of tasks across your cluster. By default, the service scheduler spreads
	//   tasks across Availability Zones. You can use task placement strategies and
	//   constraints to customize task placement decisions. This scheduler strategy is
	//   required if the service uses the CODE_DEPLOY or EXTERNAL deployment controller
	//   types.
	//   - DAEMON -The daemon scheduling strategy deploys exactly one task on each
	//   active container instance that meets all of the task placement constraints that
	//   you specify in your cluster. The service scheduler also evaluates the task
	//   placement constraints for running tasks and will stop tasks that don't meet the
	//   placement constraints. When you're using this strategy, you don't need to
	//   specify a desired number of tasks, a task placement strategy, or use Service
	//   Auto Scaling policies. Tasks using the Fargate launch type or the CODE_DEPLOY
	//   or EXTERNAL deployment controller types don't support the DAEMON scheduling
	//   strategy.
	SchedulingStrategy types.SchedulingStrategy

	// The configuration for this service to discover and connect to services, and be
	// discovered by, and connected from, other services within a namespace. Tasks that
	// run in a namespace can use short names to connect to services in the namespace.
	// Tasks can connect to services across all of the clusters in the namespace. Tasks
	// connect through a managed proxy container that collects logs and metrics for
	// increased visibility. Only the tasks that Amazon ECS services create are
	// supported with Service Connect. For more information, see Service Connect (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html)
	// in the Amazon Elastic Container Service Developer Guide.
	ServiceConnectConfiguration *types.ServiceConnectConfiguration

	// The details of the service discovery registry to associate with this service.
	// For more information, see Service discovery (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html)
	// . Each service may be associated with one service registry. Multiple service
	// registries for each service isn't supported.
	ServiceRegistries []types.ServiceRegistry

	// The metadata that you apply to the service to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. When a service is deleted, the tags are deleted as well. The following
	// basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case-sensitive.
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	// The family and revision ( family:revision ) or full ARN of the task definition
	// to run in your service. If a revision isn't specified, the latest ACTIVE
	// revision is used. A task definition must be specified if the service uses either
	// the ECS or CODE_DEPLOY deployment controllers. For more information about
	// deployment types, see Amazon ECS deployment types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// .
	TaskDefinition *string

	noSmithyDocumentSerde
}

type CreateServiceOutput struct {

	// The full description of your service following the create call. A service will
	// return either a capacityProviderStrategy or launchType parameter, but not both,
	// depending where one was specified when it was created. If a service is using the
	// ECS deployment controller, the deploymentController and taskSets parameters
	// will not be returned. if the service uses the CODE_DEPLOY deployment
	// controller, the deploymentController , taskSets and deployments parameters will
	// be returned, however the deployments parameter will be an empty list.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateService"); err != nil {
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
	if err = addOpCreateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateService",
	}
}
