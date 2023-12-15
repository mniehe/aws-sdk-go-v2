// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the specified recommendation run that was used to generate rules.
func (c *Client) GetDataQualityRuleRecommendationRun(ctx context.Context, params *GetDataQualityRuleRecommendationRunInput, optFns ...func(*Options)) (*GetDataQualityRuleRecommendationRunOutput, error) {
	if params == nil {
		params = &GetDataQualityRuleRecommendationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataQualityRuleRecommendationRun", params, optFns, c.addOperationGetDataQualityRuleRecommendationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataQualityRuleRecommendationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataQualityRuleRecommendationRunInput struct {

	// The unique run identifier associated with this run.
	//
	// This member is required.
	RunId *string

	noSmithyDocumentSerde
}

type GetDataQualityRuleRecommendationRunOutput struct {

	// The date and time when this run was completed.
	CompletedOn *time.Time

	// The name of the ruleset that was created by the run.
	CreatedRulesetName *string

	// The data source (an Glue table) associated with this run.
	DataSource *types.DataSource

	// The error strings that are associated with the run.
	ErrorString *string

	// The amount of time (in seconds) that the run consumed resources.
	ExecutionTime int32

	// A timestamp. The last point in time when this data quality rule recommendation
	// run was modified.
	LastModifiedOn *time.Time

	// The number of G.1X workers to be used in the run. The default is 5.
	NumberOfWorkers *int32

	// When a start rule recommendation run completes, it creates a recommended
	// ruleset (a set of rules). This member has those rules in Data Quality Definition
	// Language (DQDL) format.
	RecommendedRuleset *string

	// An IAM role supplied to encrypt the results of the run.
	Role *string

	// The unique run identifier associated with this run.
	RunId *string

	// The date and time when this run started.
	StartedOn *time.Time

	// The status for this run.
	Status types.TaskStatusType

	// The timeout for a run in minutes. This is the maximum time that a run can
	// consume resources before it is terminated and enters TIMEOUT status. The
	// default is 2,880 minutes (48 hours).
	Timeout *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataQualityRuleRecommendationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDataQualityRuleRecommendationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDataQualityRuleRecommendationRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataQualityRuleRecommendationRun"); err != nil {
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
	if err = addOpGetDataQualityRuleRecommendationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataQualityRuleRecommendationRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataQualityRuleRecommendationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataQualityRuleRecommendationRun",
	}
}
