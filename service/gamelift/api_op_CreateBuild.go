// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon GameLift build resource for your game server binary files.
// Combine game server binaries into a zip file for use with Amazon GameLift. When
// setting up a new game build for Amazon GameLift, we recommend using the CLI
// command upload-build (https://docs.aws.amazon.com/cli/latest/reference/gamelift/upload-build.html)
// . This helper command combines two tasks: (1) it uploads your build files from a
// file directory to an Amazon GameLift Amazon S3 location, and (2) it creates a
// new build resource. You can use the CreateBuild operation in the following
// scenarios:
//   - Create a new game build with build files that are in an Amazon S3 location
//     under an Amazon Web Services account that you control. To use this option, you
//     give Amazon GameLift access to the Amazon S3 bucket. With permissions in place,
//     specify a build name, operating system, and the Amazon S3 storage location of
//     your game build.
//   - Upload your build files to a Amazon GameLift Amazon S3 location. To use
//     this option, specify a build name and operating system. This operation creates a
//     new build resource and also returns an Amazon S3 location with temporary access
//     credentials. Use the credentials to manually upload your build files to the
//     specified Amazon S3 location. For more information, see Uploading Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/UploadingObjects.html)
//     in the Amazon S3 Developer Guide. After you upload build files to the Amazon
//     GameLift Amazon S3 location, you can't update them.
//
// If successful, this operation creates a new build resource with a unique build
// ID and places it in INITIALIZED status. A build must be in READY status before
// you can create fleets with it. Learn more Uploading Your Game (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Create a Build with Files in Amazon S3 (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build)
// All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateBuild(ctx context.Context, params *CreateBuildInput, optFns ...func(*Options)) (*CreateBuildOutput, error) {
	if params == nil {
		params = &CreateBuildInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBuild", params, optFns, c.addOperationCreateBuildMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBuildInput struct {

	// A descriptive label associated with a build. Build names don't need to be
	// unique. You can change this value later.
	Name *string

	// The operating system that your game server binaries run on. This value
	// determines the type of fleet resources that you use for this build. If your game
	// build contains multiple executables, they all must run on the same operating
	// system. You must specify a valid operating system in this request. There is no
	// default value. You can't change a build's operating system later. If you have
	// active fleets using the Windows Server 2012 operating system, you can continue
	// to create new builds using this OS until October 10, 2023, when Microsoft ends
	// its support. All others must use Windows Server 2016 when creating new
	// Windows-based builds.
	OperatingSystem types.OperatingSystem

	// A server SDK version you used when integrating your game server build with
	// Amazon GameLift. For more information see Integrate games with custom game
	// servers (https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-custom-intro.html)
	// . By default Amazon GameLift sets this value to 4.0.2 .
	ServerSdkVersion *string

	// Information indicating where your game build files are stored. Use this
	// parameter only when creating a build with files stored in an Amazon S3 bucket
	// that you own. The storage location must specify an Amazon S3 bucket name and
	// key. The location must also specify a role ARN that you set up to allow Amazon
	// GameLift to access your Amazon S3 bucket. The S3 bucket and your new build must
	// be in the same Region. If a StorageLocation is specified, the size of your file
	// can be found in your Amazon S3 bucket. Amazon GameLift will report a SizeOnDisk
	// of 0.
	StorageLocation *types.S3Location

	// A list of labels to assign to the new build resource. Tags are developer
	// defined key-value pairs. Tagging Amazon Web Services resources are useful for
	// resource management, access management and cost allocation. For more
	// information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference. Once the resource is created, you
	// can use TagResource (https://docs.aws.amazon.com/gamelift/latest/apireference/API_TagResource.html)
	// , UntagResource (https://docs.aws.amazon.com/gamelift/latest/apireference/API_UntagResource.html)
	// , and ListTagsForResource (https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListTagsForResource.html)
	// to add, remove, and view tags. The maximum tag limit may be lower than stated.
	// See the Amazon Web Services General Reference for actual tagging limits.
	Tags []types.Tag

	// Version information associated with a build or script. Version strings don't
	// need to be unique. You can change this value later.
	Version *string

	noSmithyDocumentSerde
}

type CreateBuildOutput struct {

	// The newly created build resource, including a unique build IDs and status.
	Build *types.Build

	// Amazon S3 location for your game build file, including bucket name and key.
	StorageLocation *types.S3Location

	// This element is returned only when the operation is called without a storage
	// location. It contains credentials to use when you are uploading a build file to
	// an Amazon S3 bucket that is owned by Amazon GameLift. Credentials have a limited
	// life span. To refresh these credentials, call RequestUploadCredentials (https://docs.aws.amazon.com/gamelift/latest/apireference/API_RequestUploadCredentials.html)
	// .
	UploadCredentials *types.AwsCredentials

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBuildMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBuild{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBuild{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBuild"); err != nil {
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
	if err = addOpCreateBuildValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBuild(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBuild(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBuild",
	}
}
