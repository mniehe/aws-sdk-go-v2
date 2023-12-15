// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/personalizeruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of recommended actions in sorted in descending order by
// prediction score. Use the GetActionRecommendations API if you have a custom
// campaign that deploys a solution version trained with a PERSONALIZED_ACTIONS
// recipe. For more information about PERSONALIZED_ACTIONS recipes, see
// PERSONALIZED_ACTIONS recipes (https://docs.aws.amazon.com/personalize/latest/dg/nexts-best-action-recipes.html)
// . For more information about getting action recommendations, see Getting action
// recommendations (https://docs.aws.amazon.com/personalize/latest/dg/get-action-recommendations.html)
// .
func (c *Client) GetActionRecommendations(ctx context.Context, params *GetActionRecommendationsInput, optFns ...func(*Options)) (*GetActionRecommendationsOutput, error) {
	if params == nil {
		params = &GetActionRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetActionRecommendations", params, optFns, c.addOperationGetActionRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetActionRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetActionRecommendationsInput struct {

	// The Amazon Resource Name (ARN) of the campaign to use for getting action
	// recommendations. This campaign must deploy a solution version trained with a
	// PERSONALIZED_ACTIONS recipe.
	CampaignArn *string

	// The ARN of the filter to apply to the returned recommendations. For more
	// information, see Filtering Recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter.html)
	// . When using this parameter, be sure the filter resource is ACTIVE .
	FilterArn *string

	// The values to use when filtering recommendations. For each placeholder
	// parameter in your filter expression, provide the parameter name (in matching
	// case) as a key and the filter value(s) as the corresponding value. Separate
	// multiple values for one parameter with a comma. For filter expressions that use
	// an INCLUDE element to include actions, you must provide values for all
	// parameters that are defined in the expression. For filters with expressions that
	// use an EXCLUDE element to exclude actions, you can omit the filter-values . In
	// this case, Amazon Personalize doesn't use that portion of the expression to
	// filter recommendations. For more information, see Filtering recommendations and
	// user segments (https://docs.aws.amazon.com/personalize/latest/dg/filter.html) .
	FilterValues map[string]string

	// The number of results to return. The default is 5. The maximum is 100.
	NumResults int32

	// The user ID of the user to provide action recommendations for.
	UserId *string

	noSmithyDocumentSerde
}

type GetActionRecommendationsOutput struct {

	// A list of action recommendations sorted in descending order by prediction
	// score. There can be a maximum of 100 actions in the list. For information about
	// action scores, see How action recommendation scoring works (https://docs.aws.amazon.com/personalize/latest/dg/how-action-recommendation-scoring-works.html)
	// .
	ActionList []types.PredictedAction

	// The ID of the recommendation.
	RecommendationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetActionRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetActionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetActionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetActionRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetActionRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetActionRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetActionRecommendations",
	}
}
