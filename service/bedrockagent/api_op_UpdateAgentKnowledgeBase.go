// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Knowledge Base associated to an Amazon Bedrock Agent
func (c *Client) UpdateAgentKnowledgeBase(ctx context.Context, params *UpdateAgentKnowledgeBaseInput, optFns ...func(*Options)) (*UpdateAgentKnowledgeBaseOutput, error) {
	if params == nil {
		params = &UpdateAgentKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAgentKnowledgeBase", params, optFns, c.addOperationUpdateAgentKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAgentKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Update Agent Knowledge Base Request
type UpdateAgentKnowledgeBaseInput struct {

	// Id generated at the server side when an Agent is created
	//
	// This member is required.
	AgentId *string

	// Draft Version of the Agent.
	//
	// This member is required.
	AgentVersion *string

	// Id generated at the server side when a Knowledge Base is associated to an Agent
	//
	// This member is required.
	KnowledgeBaseId *string

	// Description of the Resource.
	Description *string

	// State of the knowledge base; whether it is enabled or disabled
	KnowledgeBaseState types.KnowledgeBaseState

	noSmithyDocumentSerde
}

// Update Agent Knowledge Base Response
type UpdateAgentKnowledgeBaseOutput struct {

	// Contains the information of an Agent Knowledge Base.
	//
	// This member is required.
	AgentKnowledgeBase *types.AgentKnowledgeBase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAgentKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAgentKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAgentKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAgentKnowledgeBase"); err != nil {
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
	if err = addOpUpdateAgentKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAgentKnowledgeBase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAgentKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAgentKnowledgeBase",
	}
}
