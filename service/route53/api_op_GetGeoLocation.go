// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/mniehe/aws-sdk-go-v2/aws/middleware"
	"github.com/mniehe/aws-sdk-go-v2/aws/signer/v4"
	"github.com/mniehe/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about whether a specified geographic location is supported for
// Amazon Route 53 geolocation resource record sets. Route 53 does not perform
// authorization for this API because it retrieves information that is already
// available to the public. Use the following syntax to determine whether a
// continent is supported for geolocation: GET
// /2013-04-01/geolocation?continentcode=two-letter abbreviation for a continent
// Use the following syntax to determine whether a country is supported for
// geolocation: GET /2013-04-01/geolocation?countrycode=two-character country code
// Use the following syntax to determine whether a subdivision of a country is
// supported for geolocation: GET
// /2013-04-01/geolocation?countrycode=two-character country
// code&subdivisioncode=subdivision code
func (c *Client) GetGeoLocation(ctx context.Context, params *GetGeoLocationInput, optFns ...func(*Options)) (*GetGeoLocationOutput, error) {
	if params == nil {
		params = &GetGeoLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGeoLocation", params, optFns, c.addOperationGetGeoLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGeoLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for information about whether a specified geographic location is
// supported for Amazon Route 53 geolocation resource record sets.
type GetGeoLocationInput struct {

	// For geolocation resource record sets, a two-letter abbreviation that identifies
	// a continent. Amazon Route 53 supports the following continent codes:
	//   - AF: Africa
	//   - AN: Antarctica
	//   - AS: Asia
	//   - EU: Europe
	//   - OC: Oceania
	//   - NA: North America
	//   - SA: South America
	ContinentCode *string

	// Amazon Route 53 uses the two-letter country codes that are specified in ISO
	// standard 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) .
	// Route 53 also supports the contry code UA forr Ukraine.
	CountryCode *string

	// The code for the subdivision, such as a particular state within the United
	// States. For a list of US state abbreviations, see Appendix B: Two–Letter State
	// and Possession Abbreviations (https://pe.usps.com/text/pub28/28apb.htm) on the
	// United States Postal Service website. For a list of all supported subdivision
	// codes, use the ListGeoLocations (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListGeoLocations.html)
	// API.
	SubdivisionCode *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the specified
// geolocation code.
type GetGeoLocationOutput struct {

	// A complex type that contains the codes and full continent, country, and
	// subdivision names for the specified geolocation code.
	//
	// This member is required.
	GeoLocationDetails *types.GeoLocationDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGeoLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetGeoLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetGeoLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGeoLocation"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeoLocation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGeoLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGeoLocation",
	}
}
