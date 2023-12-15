// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) aggregated statistical data about findings.
func (c *Client) GetFindingStatistics(ctx context.Context, params *GetFindingStatisticsInput, optFns ...func(*Options)) (*GetFindingStatisticsOutput, error) {
	if params == nil {
		params = &GetFindingStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingStatistics", params, optFns, c.addOperationGetFindingStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingStatisticsInput struct {

	// The finding property to use to group the query results. Valid values are:
	//   - classificationDetails.jobId - The unique identifier for the classification
	//   job that produced the finding.
	//   - resourcesAffected.s3Bucket.name - The name of the S3 bucket that the
	//   finding applies to.
	//   - severity.description - The severity level of the finding, such as High or
	//   Medium.
	//   - type - The type of finding, such as Policy:IAMUser/S3BucketPublic and
	//   SensitiveData:S3Object/Personal.
	//
	// This member is required.
	GroupBy types.GroupBy

	// The criteria to use to filter the query results.
	FindingCriteria *types.FindingCriteria

	// The maximum number of items to include in each page of the response.
	Size *int32

	// The criteria to use to sort the query results.
	SortCriteria *types.FindingStatisticsSortCriteria

	noSmithyDocumentSerde
}

type GetFindingStatisticsOutput struct {

	// An array of objects, one for each group of findings that matches the filter
	// criteria specified in the request.
	CountsByGroup []types.GroupCount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFindingStatistics"); err != nil {
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
	if err = addOpGetFindingStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFindingStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFindingStatistics",
	}
}
