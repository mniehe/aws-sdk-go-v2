// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes the specified graph. Graphs cannot be deleted if delete-protection is
// enabled.
func (c *Client) DeleteGraph(ctx context.Context, params *DeleteGraphInput, optFns ...func(*Options)) (*DeleteGraphOutput, error) {
	if params == nil {
		params = &DeleteGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGraph", params, optFns, c.addOperationDeleteGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGraphInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// Determines whether a final graph snapshot is created before the graph is
	// deleted. If true is specified, no graph snapshot is created. If false is
	// specified, a graph snapshot is created before the graph is deleted.
	//
	// This member is required.
	SkipSnapshot *bool

	noSmithyDocumentSerde
}

func (in *DeleteGraphInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type DeleteGraphOutput struct {

	// The ARN associated with the graph.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph.
	//
	// This member is required.
	Id *string

	// The name of the graph.
	//
	// This member is required.
	Name *string

	// The build number associated with the graph.
	BuildNumber *string

	// The time at which the graph was created.
	CreateTime *time.Time

	// If true , deletion protection was enabled for the graph.
	DeletionProtection *bool

	// The graph endpoint.
	Endpoint *string

	// The ID of the KMS key used to encrypt and decrypt graph data.
	KmsKeyIdentifier *string

	// The number of memory-optimized Neptune Capacity Units (m-NCUs) allocated to the
	// graph.
	ProvisionedMemory *int32

	// If true , the graph has a public endpoint, otherwise not.
	PublicConnectivity *bool

	// The number of replicas for the graph.
	ReplicaCount *int32

	// The ID of the snapshot from which the graph was created, if the graph was
	// recovered from a snapshot.
	SourceSnapshotId *string

	// The status of the graph.
	Status types.GraphStatus

	// The reason for the status of the graph.
	StatusReason *string

	// Specifies the number of dimensions for vector embeddings loaded into the graph.
	// Max = 65535
	VectorSearchConfiguration *types.VectorSearchConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGraph{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteGraph"); err != nil {
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
	if err = addOpDeleteGraphValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGraph(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteGraph",
	}
}
