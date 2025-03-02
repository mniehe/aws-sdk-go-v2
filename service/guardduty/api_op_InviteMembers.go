// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invites Amazon Web Services accounts to become members of an organization
// administered by the Amazon Web Services account that invokes this API. If you
// are using Amazon Web Services Organizations to manage your GuardDuty
// environment, this step is not needed. For more information, see Managing
// accounts with organizations (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html)
// . To invite Amazon Web Services accounts, the first step is to ensure that
// GuardDuty has been enabled in the potential member accounts. You can now invoke
// this API to add accounts by invitation. The invited accounts can either accept
// or decline the invitation from their GuardDuty accounts. Each invited Amazon Web
// Services account can choose to accept the invitation from only one Amazon Web
// Services account. For more information, see Managing GuardDuty accounts by
// invitation (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_invitations.html)
// . After the invite has been accepted and you choose to disassociate a member
// account (by using DisassociateMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DisassociateMembers.html)
// ) from your account, the details of the member account obtained by invoking
// CreateMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html)
// , including the associated email addresses, will be retained. This is done so
// that you can invoke InviteMembers without the need to invoke CreateMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html)
// again. To remove the details associated with a member account, you must also
// invoke DeleteMembers (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html)
// .
func (c *Client) InviteMembers(ctx context.Context, params *InviteMembersInput, optFns ...func(*Options)) (*InviteMembersOutput, error) {
	if params == nil {
		params = &InviteMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteMembers", params, optFns, c.addOperationInviteMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteMembersInput struct {

	// A list of account IDs of the accounts that you want to invite to GuardDuty as
	// members.
	//
	// This member is required.
	AccountIds []string

	// The unique ID of the detector of the GuardDuty account that you want to invite
	// members with.
	//
	// This member is required.
	DetectorId *string

	// A Boolean value that specifies whether you want to disable email notification
	// to the accounts that you are inviting to GuardDuty as members.
	DisableEmailNotification *bool

	// The invitation message that you want to send to the accounts that you're
	// inviting to GuardDuty as members.
	Message *string

	noSmithyDocumentSerde
}

type InviteMembersOutput struct {

	// A list of objects that contain the unprocessed account and a result string that
	// explains why it was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInviteMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInviteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InviteMembers"); err != nil {
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
	if err = addOpInviteMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInviteMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInviteMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InviteMembers",
	}
}
