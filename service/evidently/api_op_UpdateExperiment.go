// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Evidently experiment. Don't use this operation to update an
// experiment's tag. Instead, use TagResource (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_TagResource.html)
// .
func (c *Client) UpdateExperiment(ctx context.Context, params *UpdateExperimentInput, optFns ...func(*Options)) (*UpdateExperimentOutput, error) {
	if params == nil {
		params = &UpdateExperimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateExperiment", params, optFns, c.addOperationUpdateExperimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateExperimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateExperimentInput struct {

	// The name of the experiment to update.
	//
	// This member is required.
	Experiment *string

	// The name or ARN of the project that contains the experiment that you want to
	// update.
	//
	// This member is required.
	Project *string

	// An optional description of the experiment.
	Description *string

	// An array of structures that defines the metrics used for the experiment, and
	// whether a higher or lower value for each metric is the goal.
	MetricGoals []types.MetricGoalConfig

	// A structure that contains the configuration of which variation o use as the
	// "control" version. The "control" version is used for comparison with other
	// variations. This structure also specifies how much experiment traffic is
	// allocated to each variation.
	OnlineAbConfig *types.OnlineAbConfig

	// When Evidently assigns a particular user session to an experiment, it must use
	// a randomization ID to determine which variation the user session is served. This
	// randomization ID is a combination of the entity ID and randomizationSalt . If
	// you omit randomizationSalt , Evidently uses the experiment name as the
	// randomizationSalt .
	RandomizationSalt *string

	// Removes a segment from being used in an experiment. You can't use this
	// parameter if the experiment is currently running.
	RemoveSegment bool

	// The portion of the available audience that you want to allocate to this
	// experiment, in thousandths of a percent. The available audience is the total
	// audience minus the audience that you have allocated to overrides or current
	// launches of this feature. This is represented in thousandths of a percent. For
	// example, specify 20,000 to allocate 20% of the available audience.
	SamplingRate *int64

	// Adds an audience segment to an experiment. When a segment is used in an
	// experiment, only user sessions that match the segment pattern are used in the
	// experiment. You can't use this parameter if the experiment is currently running.
	Segment *string

	// An array of structures that define the variations being tested in the
	// experiment.
	Treatments []types.TreatmentConfig

	noSmithyDocumentSerde
}

type UpdateExperimentOutput struct {

	// A structure containing the configuration details of the experiment that was
	// updated.
	//
	// This member is required.
	Experiment *types.Experiment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateExperimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateExperiment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateExperiment"); err != nil {
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
	if err = addOpUpdateExperimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateExperiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateExperiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateExperiment",
	}
}
