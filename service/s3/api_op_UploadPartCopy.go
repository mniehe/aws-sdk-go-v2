// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/mniehe/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/mniehe/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Uploads a part by copying data from an existing object as data source. To
// specify the data source, you add the request header x-amz-copy-source in your
// request. To specify a byte range, you add the request header
// x-amz-copy-source-range in your request. For information about maximum and
// minimum part sizes and other multipart upload specifications, see Multipart
// upload limits (https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html)
// in the Amazon S3 User Guide. Instead of copying data from an existing object as
// part data, you might use the UploadPart (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
// action to upload new data as a part of an object in your request. You must
// initiate a multipart upload before you can upload any part. In response to your
// initiate request, Amazon S3 returns the upload ID, a unique identifier that you
// must include in your upload part request. For conceptual information about
// multipart uploads, see Uploading Objects Using Multipart Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html)
// in the Amazon S3 User Guide. For information about copying objects using a
// single atomic action vs. a multipart upload, see Operations on Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html)
// in the Amazon S3 User Guide. Directory buckets - For directory buckets, you must
// make requests for this API operation to the Zonal endpoint. These endpoints
// support virtual-hosted-style requests in the format
// https://bucket_name.s3express-az_id.region.amazonaws.com/key-name . Path-style
// requests are not supported. For more information, see Regional and Zonal
// endpoints (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-Regions-and-Zones.html)
// in the Amazon S3 User Guide. Authentication and authorization All UploadPartCopy
// requests must be authenticated and signed by using IAM credentials (access key
// ID and secret access key for the IAM identities). All headers with the x-amz-
// prefix, including x-amz-copy-source , must be signed. For more information, see
// REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html)
// . Directory buckets - You must use IAM credentials to authenticate and authorize
// your access to the UploadPartCopy API operation, instead of using the temporary
// security credentials through the CreateSession API operation. Amazon Web
// Services CLI or SDKs handles authentication and authorization on your behalf.
// Permissions You must have READ access to the source object and WRITE access to
// the destination bucket.
//   - General purpose bucket permissions - You must have the permissions in a
//     policy based on the bucket types of your source bucket and destination bucket in
//     an UploadPartCopy operation.
//   - If the source object is in a general purpose bucket, you must have the
//     s3:GetObject permission to read the source object that is being copied.
//   - If the destination bucket is a general purpose bucket, you must have the
//     s3:PubObject permission to write the object copy to the destination bucket.
//     For information about permissions required to use the multipart upload API, see
//     Multipart Upload and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html)
//     in the Amazon S3 User Guide.
//   - Directory bucket permissions - You must have permissions in a bucket policy
//     or an IAM identity-based policy based on the source and destination bucket types
//     in an UploadPartCopy operation.
//   - If the source object that you want to copy is in a directory bucket, you
//     must have the s3express:CreateSession permission in the Action element of a
//     policy to read the object . By default, the session is in the ReadWrite mode.
//     If you want to restrict the access, you can explicitly set the
//     s3express:SessionMode condition key to ReadOnly on the copy source bucket.
//   - If the copy destination is a directory bucket, you must have the
//     s3express:CreateSession permission in the Action element of a policy to write
//     the object to the destination. The s3express:SessionMode condition key cannot
//     be set to ReadOnly on the copy destination. For example policies, see Example
//     bucket policies for S3 Express One Zone (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html)
//     and Amazon Web Services Identity and Access Management (IAM) identity-based
//     policies for S3 Express One Zone (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html)
//     in the Amazon S3 User Guide.
//
// Encryption
//   - General purpose buckets - For information about using server-side
//     encryption with customer-provided encryption keys with the UploadPartCopy
//     operation, see CopyObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html)
//     and UploadPart (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//     .
//   - Directory buckets - For directory buckets, only server-side encryption with
//     Amazon S3 managed keys (SSE-S3) ( AES256 ) is supported.
//
// Special errors
//   - Error Code: NoSuchUpload
//   - Description: The specified multipart upload does not exist. The upload ID
//     might be invalid, or the multipart upload might have been aborted or completed.
//   - HTTP Status Code: 404 Not Found
//   - Error Code: InvalidRequest
//   - Description: The specified copy source is not supported as a byte-range
//     copy source.
//   - HTTP Status Code: 400 Bad Request
//
// HTTP Host header syntax Directory buckets - The HTTP Host header syntax is
// Bucket_name.s3express-az_id.region.amazonaws.com . The following operations are
// related to UploadPartCopy :
//   - CreateMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//   - UploadPart (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//   - CompleteMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//   - AbortMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
//   - ListParts (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//   - ListMultipartUploads (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html)
func (c *Client) UploadPartCopy(ctx context.Context, params *UploadPartCopyInput, optFns ...func(*Options)) (*UploadPartCopyOutput, error) {
	if params == nil {
		params = &UploadPartCopyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadPartCopy", params, optFns, c.addOperationUploadPartCopyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadPartCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UploadPartCopyInput struct {

	// The bucket name. Directory buckets - When you use this operation with a
	// directory bucket, you must use virtual-hosted-style requests in the format
	// Bucket_name.s3express-az_id.region.amazonaws.com . Path-style requests are not
	// supported. Directory bucket names must be unique in the chosen Availability
	// Zone. Bucket names must follow the format bucket_base_name--az-id--x-s3 (for
	// example, DOC-EXAMPLE-BUCKET--usw2-az2--x-s3 ). For information about bucket
	// naming restrictions, see Directory bucket naming rules (https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-bucket-naming-rules.html)
	// in the Amazon S3 User Guide. Access points - When you use this action with an
	// access point, you must provide the alias of the access point in place of the
	// bucket name or specify the access point ARN. When using the access point ARN,
	// you must direct requests to the access point hostname. The access point hostname
	// takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this action with an access point through the Amazon Web Services
	// SDKs, you provide the access point ARN in place of the bucket name. For more
	// information about access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. Access points and Object Lambda access points are
	// not supported by directory buckets. S3 on Outposts - When you use this action
	// with Amazon S3 on Outposts, you must direct requests to the S3 on Outposts
	// hostname. The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts? (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// Specifies the source object for the copy operation. You specify the value in
	// one of two formats, depending on whether you want to access the source object
	// through an access point (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html)
	// :
	//   - For objects not accessed through an access point, specify the name of the
	//   source bucket and key of the source object, separated by a slash (/). For
	//   example, to copy the object reports/january.pdf from the bucket
	//   awsexamplebucket , use awsexamplebucket/reports/january.pdf . The value must
	//   be URL-encoded.
	//   - For objects accessed through access points, specify the Amazon Resource
	//   Name (ARN) of the object as accessed through the access point, in the format
	//   arn:aws:s3:::accesspoint//object/ . For example, to copy the object
	//   reports/january.pdf through access point my-access-point owned by account
	//   123456789012 in Region us-west-2 , use the URL encoding of
	//   arn:aws:s3:us-west-2:123456789012:accesspoint/my-access-point/object/reports/january.pdf
	//   . The value must be URL encoded.
	//   - Amazon S3 supports copy operations using Access points only when the source
	//   and destination buckets are in the same Amazon Web Services Region.
	//   - Access points are not supported by directory buckets. Alternatively, for
	//   objects accessed through Amazon S3 on Outposts, specify the ARN of the object as
	//   accessed in the format arn:aws:s3-outposts:::outpost//object/ . For example,
	//   to copy the object reports/january.pdf through outpost my-outpost owned by
	//   account 123456789012 in Region us-west-2 , use the URL encoding of
	//   arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/object/reports/january.pdf
	//   . The value must be URL-encoded.
	// If your bucket has versioning enabled, you could have multiple versions of the
	// same object. By default, x-amz-copy-source identifies the current version of
	// the source object to copy. To copy a specific version of the source object to
	// copy, append ?versionId= to the x-amz-copy-source request header (for example,
	// x-amz-copy-source:
	// /awsexamplebucket/reports/january.pdf?versionId=QUpfdndhfd8438MNFDN93jdnJFkdmqnh893
	// ). If the current version is a delete marker and you don't specify a versionId
	// in the x-amz-copy-source request header, Amazon S3 returns a 404 Not Found
	// error, because the object does not exist. If you specify versionId in the
	// x-amz-copy-source and the versionId is a delete marker, Amazon S3 returns an
	// HTTP 400 Bad Request error, because you are not allowed to specify a delete
	// marker as a version for the x-amz-copy-source . Directory buckets - S3
	// Versioning isn't enabled and supported for directory buckets.
	//
	// This member is required.
	CopySource *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Part number of part being copied. This is a positive integer between 1 and
	// 10,000.
	//
	// This member is required.
	PartNumber *int32

	// Upload ID identifying the multipart upload whose part is being copied.
	//
	// This member is required.
	UploadId *string

	// Copies the object if its entity tag (ETag) matches the specified tag. If both
	// of the x-amz-copy-source-if-match and x-amz-copy-source-if-unmodified-since
	// headers are present in the request as follows: x-amz-copy-source-if-match
	// condition evaluates to true , and; x-amz-copy-source-if-unmodified-since
	// condition evaluates to false ; Amazon S3 returns 200 OK and copies the data.
	CopySourceIfMatch *string

	// Copies the object if it has been modified since the specified time. If both of
	// the x-amz-copy-source-if-none-match and x-amz-copy-source-if-modified-since
	// headers are present in the request as follows: x-amz-copy-source-if-none-match
	// condition evaluates to false , and; x-amz-copy-source-if-modified-since
	// condition evaluates to true ; Amazon S3 returns 412 Precondition Failed
	// response code.
	CopySourceIfModifiedSince *time.Time

	// Copies the object if its entity tag (ETag) is different than the specified
	// ETag. If both of the x-amz-copy-source-if-none-match and
	// x-amz-copy-source-if-modified-since headers are present in the request as
	// follows: x-amz-copy-source-if-none-match condition evaluates to false , and;
	// x-amz-copy-source-if-modified-since condition evaluates to true ; Amazon S3
	// returns 412 Precondition Failed response code.
	CopySourceIfNoneMatch *string

	// Copies the object if it hasn't been modified since the specified time. If both
	// of the x-amz-copy-source-if-match and x-amz-copy-source-if-unmodified-since
	// headers are present in the request as follows: x-amz-copy-source-if-match
	// condition evaluates to true , and; x-amz-copy-source-if-unmodified-since
	// condition evaluates to false ; Amazon S3 returns 200 OK and copies the data.
	CopySourceIfUnmodifiedSince *time.Time

	// The range of bytes to copy from the source object. The range value must use the
	// form bytes=first-last, where the first and last are the zero-based byte offsets
	// to copy. For example, bytes=0-9 indicates that you want to copy the first 10
	// bytes of the source. You can copy a range only if the source object is greater
	// than 5 MB.
	CopySourceRange *string

	// Specifies the algorithm to use when decrypting the source object (for example,
	// AES256 ). This functionality is not supported when the source object is in a
	// directory bucket.
	CopySourceSSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one that
	// was used when the source object was created. This functionality is not supported
	// when the source object is in a directory bucket.
	CopySourceSSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error. This functionality is not
	// supported when the source object is in a directory bucket.
	CopySourceSSECustomerKeyMD5 *string

	// The account ID of the expected destination bucket owner. If the account ID that
	// you provide does not match the actual owner of the destination bucket, the
	// request fails with the HTTP status code 403 Forbidden (access denied).
	ExpectedBucketOwner *string

	// The account ID of the expected source bucket owner. If the account ID that you
	// provide does not match the actual owner of the source bucket, the request fails
	// with the HTTP status code 403 Forbidden (access denied).
	ExpectedSourceBucketOwner *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. If either the
	// source or destination S3 bucket has Requester Pays enabled, the requester will
	// pay for corresponding charges to copy the object. For information about
	// downloading objects from Requester Pays buckets, see Downloading Objects in
	// Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide. This functionality is not supported for directory
	// buckets.
	RequestPayer types.RequestPayer

	// Specifies the algorithm to use when encrypting the object (for example,
	// AES256). This functionality is not supported when the destination bucket is a
	// directory bucket.
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header. This must be the same
	// encryption key specified in the initiate multipart upload request. This
	// functionality is not supported when the destination bucket is a directory
	// bucket.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error. This functionality is not
	// supported when the destination bucket is a directory bucket.
	SSECustomerKeyMD5 *string

	noSmithyDocumentSerde
}

func (in *UploadPartCopyInput) bindEndpointParams(p *EndpointParameters) {
	p.Bucket = in.Bucket
	p.DisableS3ExpressSessionAuth = ptr.Bool(true)
}

type UploadPartCopyOutput struct {

	// Indicates whether the multipart upload uses an S3 Bucket Key for server-side
	// encryption with Key Management Service (KMS) keys (SSE-KMS). This functionality
	// is not supported for directory buckets.
	BucketKeyEnabled *bool

	// Container for all response elements.
	CopyPartResult *types.CopyPartResult

	// The version of the source object that was copied, if you have enabled
	// versioning on the source bucket. This functionality is not supported when the
	// source object is in a directory bucket.
	CopySourceVersionId *string

	// If present, indicates that the requester was successfully charged for the
	// request. This functionality is not supported for directory buckets.
	RequestCharged types.RequestCharged

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header to confirm the encryption
	// algorithm that's used. This functionality is not supported for directory
	// buckets.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was
	// requested, the response will include this header to provide the round-trip
	// message integrity verification of the customer-provided encryption key. This
	// functionality is not supported for directory buckets.
	SSECustomerKeyMD5 *string

	// If present, indicates the ID of the Key Management Service (KMS) symmetric
	// encryption customer managed key that was used for the object. This functionality
	// is not supported for directory buckets.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when you store this object in Amazon
	// S3 (for example, AES256 , aws:kms ). For directory buckets, only server-side
	// encryption with Amazon S3 managed keys (SSE-S3) ( AES256 ) is supported.
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUploadPartCopyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpUploadPartCopy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUploadPartCopy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UploadPartCopy"); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUploadPartCopyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadPartCopy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addUploadPartCopyUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = s3cust.HandleResponseErrorWith200Status(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *UploadPartCopyInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opUploadPartCopy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UploadPartCopy",
	}
}

// getUploadPartCopyBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getUploadPartCopyBucketMember(input interface{}) (*string, bool) {
	in := input.(*UploadPartCopyInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addUploadPartCopyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getUploadPartCopyBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
