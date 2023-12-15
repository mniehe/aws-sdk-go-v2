// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Amazon ECS service recommendations. Compute Optimizer generates
// recommendations for Amazon ECS services on Fargate that meet a specific set of
// requirements. For more information, see the Supported resources and requirements (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html)
// in the Compute Optimizer User Guide.
func (c *Client) GetECSServiceRecommendations(ctx context.Context, params *GetECSServiceRecommendationsInput, optFns ...func(*Options)) (*GetECSServiceRecommendationsOutput, error) {
	if params == nil {
		params = &GetECSServiceRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetECSServiceRecommendations", params, optFns, c.addOperationGetECSServiceRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetECSServiceRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetECSServiceRecommendationsInput struct {

	// Return the Amazon ECS service recommendations to the specified Amazon Web
	// Services account IDs. If your account is the management account or the delegated
	// administrator of an organization, use this parameter to return the Amazon ECS
	// service recommendations to specific member accounts. You can only specify one
	// account ID per request.
	AccountIds []string

	// An array of objects to specify a filter that returns a more specific list of
	// Amazon ECS service recommendations.
	Filters []types.ECSServiceRecommendationFilter

	// The maximum number of Amazon ECS service recommendations to return with a
	// single request. To retrieve the remaining results, make another request with the
	// returned nextToken value.
	MaxResults *int32

	// The token to advance to the next page of Amazon ECS service recommendations.
	NextToken *string

	// The ARN that identifies the Amazon ECS service. The following is the format of
	// the ARN: arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name
	ServiceArns []string

	noSmithyDocumentSerde
}

type GetECSServiceRecommendationsOutput struct {

	// An array of objects that describe the Amazon ECS service recommendations.
	EcsServiceRecommendations []types.ECSServiceRecommendation

	// An array of objects that describe errors of the request.
	Errors []types.GetRecommendationError

	// The token to advance to the next page of Amazon ECS service recommendations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetECSServiceRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetECSServiceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetECSServiceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetECSServiceRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetECSServiceRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetECSServiceRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetECSServiceRecommendations",
	}
}
