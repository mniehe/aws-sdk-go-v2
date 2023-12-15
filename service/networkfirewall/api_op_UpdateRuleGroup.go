// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the rule settings for the specified rule group. You use a rule group by
// reference in one or more firewall policies. When you modify a rule group, you
// modify all firewall policies that use the rule group. To update a rule group,
// first call DescribeRuleGroup to retrieve the current RuleGroup object, update
// the object as needed, and then provide the updated object to this call.
func (c *Client) UpdateRuleGroup(ctx context.Context, params *UpdateRuleGroupInput, optFns ...func(*Options)) (*UpdateRuleGroupOutput, error) {
	if params == nil {
		params = &UpdateRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRuleGroup", params, optFns, c.addOperationUpdateRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleGroupInput struct {

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the rule group. The token marks the state of the rule group
	// resource at the time of the request. To make changes to the rule group, you
	// provide the token in your request. Network Firewall uses the token to ensure
	// that the rule group hasn't changed since you last retrieved it. If it has
	// changed, the operation fails with an InvalidTokenException . If this happens,
	// retrieve the rule group again to get a current copy of it with a current token.
	// Reapply your changes as needed, then try the operation again using the new
	// token.
	//
	// This member is required.
	UpdateToken *string

	// Indicates whether you want Network Firewall to analyze the stateless rules in
	// the rule group for rule behavior such as asymmetric routing. If set to TRUE ,
	// Network Firewall runs the analysis and then updates the rule group for you. To
	// run the stateless rule group analyzer without updating the rule group, set
	// DryRun to TRUE .
	AnalyzeRuleGroup bool

	// A description of the rule group.
	Description *string

	// Indicates whether you want Network Firewall to just check the validity of the
	// request, rather than run the request. If set to TRUE , Network Firewall checks
	// whether the request can run successfully, but doesn't actually make the
	// requested changes. The call returns the value that the request would return if
	// you ran it with dry run set to FALSE , but doesn't make additions or changes to
	// your resources. This option allows you to make sure that you have the required
	// permissions to run the request and that your request parameters are valid. If
	// set to FALSE , Network Firewall makes the requested changes to your resources.
	DryRun bool

	// A complex type that contains settings for encryption of your rule group
	// resources.
	EncryptionConfiguration *types.EncryptionConfiguration

	// An object that defines the rule group rules. You must provide either this rule
	// group setting or a Rules setting, but not both.
	RuleGroup *types.RuleGroup

	// The Amazon Resource Name (ARN) of the rule group. You must specify the ARN or
	// the name, and you can specify both.
	RuleGroupArn *string

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it. You must specify the ARN or the name, and you can
	// specify both.
	RuleGroupName *string

	// A string containing stateful rule group rules specifications in Suricata flat
	// format, with one rule per line. Use this to import your existing Suricata
	// compatible rule groups. You must provide either this rules setting or a
	// populated RuleGroup setting, but not both. You can provide your rule group
	// specification in Suricata flat format through this setting when you create or
	// update your rule group. The call response returns a RuleGroup object that
	// Network Firewall has populated from your string.
	Rules *string

	// A complex type that contains metadata about the rule group that your own rule
	// group is copied from. You can use the metadata to keep track of updates made to
	// the originating rule group.
	SourceMetadata *types.SourceMetadata

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules. This setting is required for requests that do not include the
	// RuleGroupARN .
	Type types.RuleGroupType

	noSmithyDocumentSerde
}

type UpdateRuleGroupOutput struct {

	// The high-level properties of a rule group. This, along with the RuleGroup ,
	// define the rule group. You can retrieve all objects for a rule group by calling
	// DescribeRuleGroup .
	//
	// This member is required.
	RuleGroupResponse *types.RuleGroupResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the rule group. The token marks the state of the rule group
	// resource at the time of the request. To make changes to the rule group, you
	// provide the token in your request. Network Firewall uses the token to ensure
	// that the rule group hasn't changed since you last retrieved it. If it has
	// changed, the operation fails with an InvalidTokenException . If this happens,
	// retrieve the rule group again to get a current copy of it with a current token.
	// Reapply your changes as needed, then try the operation again using the new
	// token.
	//
	// This member is required.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateRuleGroup"); err != nil {
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
	if err = addOpUpdateRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateRuleGroup",
	}
}
