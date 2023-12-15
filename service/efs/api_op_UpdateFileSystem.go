// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the throughput mode or the amount of provisioned throughput of an
// existing file system.
func (c *Client) UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) {
	if params == nil {
		params = &UpdateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFileSystem", params, optFns, c.addOperationUpdateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFileSystemInput struct {

	// The ID of the file system that you want to update.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) The throughput, measured in mebibytes per second (MiBps), that you
	// want to provision for a file system that you're creating. Required if
	// ThroughputMode is set to provisioned . Valid values are 1-3414 MiBps, with the
	// upper limit depending on Region. To increase this limit, contact Amazon Web
	// Services Support. For more information, see Amazon EFS quotas that you can
	// increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in
	// the Amazon EFS User Guide.
	ProvisionedThroughputInMibps *float64

	// (Optional) Updates the file system's throughput mode. If you're not updating
	// your throughput mode, you don't need to provide this value in your request. If
	// you are changing the ThroughputMode to provisioned , you must also set a value
	// for ProvisionedThroughputInMibps .
	ThroughputMode types.ThroughputMode

	noSmithyDocumentSerde
}

// A description of the file system.
type UpdateFileSystemOutput struct {

	// The time that the file system was created, in seconds (since
	// 1970-01-01T00:00:00Z).
	//
	// This member is required.
	CreationTime *time.Time

	// The opaque string specified in the request.
	//
	// This member is required.
	CreationToken *string

	// The ID of the file system, assigned by Amazon EFS.
	//
	// This member is required.
	FileSystemId *string

	// The lifecycle phase of the file system.
	//
	// This member is required.
	LifeCycleState types.LifeCycleState

	// The current number of mount targets that the file system has. For more
	// information, see CreateMountTarget .
	//
	// This member is required.
	NumberOfMountTargets int32

	// The Amazon Web Services account that created the file system.
	//
	// This member is required.
	OwnerId *string

	// The Performance mode of the file system.
	//
	// This member is required.
	PerformanceMode types.PerformanceMode

	// The latest known metered size (in bytes) of data stored in the file system, in
	// its Value field, and the time at which that size was determined in its Timestamp
	// field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of a
	// consistent snapshot of the file system, but it is eventually consistent when
	// there are no writes to the file system. That is, SizeInBytes represents actual
	// size only if the file system is not modified for a period longer than a couple
	// of hours. Otherwise, the value is not the exact size that the file system was at
	// any point in time.
	//
	// This member is required.
	SizeInBytes *types.FileSystemSize

	// The tags associated with the file system, presented as an array of Tag objects.
	//
	// This member is required.
	Tags []types.Tag

	// The unique and consistent identifier of the Availability Zone in which the file
	// system is located, and is valid only for One Zone file systems. For example,
	// use1-az1 is an Availability Zone ID for the us-east-1 Amazon Web Services
	// Region, and it has the same location in every Amazon Web Services account.
	AvailabilityZoneId *string

	// Describes the Amazon Web Services Availability Zone in which the file system is
	// located, and is valid only for One Zone file systems. For more information, see
	// Using EFS storage classes (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html)
	// in the Amazon EFS User Guide.
	AvailabilityZoneName *string

	// A Boolean value that, if true, indicates that the file system is encrypted.
	Encrypted *bool

	// The Amazon Resource Name (ARN) for the EFS file system, in the format
	// arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id .
	// Example with sample data:
	// arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-01234567
	FileSystemArn *string

	// Describes the protection on the file system.
	FileSystemProtection *types.FileSystemProtectionDescription

	// The ID of an KMS key used to protect the encrypted file system.
	KmsKeyId *string

	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem . If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	Name *string

	// The amount of provisioned throughput, measured in MiBps, for the file system.
	// Valid for file systems using ThroughputMode set to provisioned .
	ProvisionedThroughputInMibps *float64

	// Displays the file system's throughput mode. For more information, see
	// Throughput modes (https://docs.aws.amazon.com/efs/latest/ug/performance.html#throughput-modes)
	// in the Amazon EFS User Guide.
	ThroughputMode types.ThroughputMode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFileSystem"); err != nil {
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
	if err = addOpUpdateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFileSystem",
	}
}
