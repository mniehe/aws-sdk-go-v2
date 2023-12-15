// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/mniehe/aws-sdk-go-v2/service/entityresolution/document"
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An object containing an error message, if there was an error.
type ErrorDetails struct {

	// The error message from the job, if there is one.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// An object containing InputRecords , TotalRecordsProcessed , MatchIDs , and
// RecordsNotProcessed .
type IdMappingJobMetrics struct {

	// The total number of input records.
	InputRecords *int32

	// The total number of records that did not get processed.
	RecordsNotProcessed *int32

	// The total number of records processed.
	TotalRecordsProcessed *int32

	noSmithyDocumentSerde
}

// An object which defines the ID mapping techniques and provider configurations.
type IdMappingTechniques struct {

	// The type of ID mapping.
	//
	// This member is required.
	IdMappingType IdMappingType

	// An object which defines any additional configurations required by the provider
	// service.
	//
	// This member is required.
	ProviderProperties *ProviderProperties

	noSmithyDocumentSerde
}

// An object containing InputSourceARN and SchemaName .
type IdMappingWorkflowInputSource struct {

	// An Gluetable ARN for the input source table.
	//
	// This member is required.
	InputSourceARN *string

	// The name of the schema to be retrieved.
	//
	// This member is required.
	SchemaName *string

	noSmithyDocumentSerde
}

// The output source for the ID mapping workflow.
type IdMappingWorkflowOutputSource struct {

	// The S3 path to which Entity Resolution will write the output table.
	//
	// This member is required.
	OutputS3Path *string

	// Customer KMS ARN for encryption at rest. If not provided, system will use an
	// Entity Resolution managed KMS key.
	KMSArn *string

	noSmithyDocumentSerde
}

// A list of IdMappingWorkflowSummary objects, each of which contain the fields
// WorkflowName , WorkflowArn , CreatedAt , and UpdatedAt .
type IdMappingWorkflowSummary struct {

	// The timestamp of when the workflow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The timestamp of when the workflow was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// IdMappingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

// An object which defines an incremental run type and has only incrementalRunType
// as a field.
type IncrementalRunConfig struct {

	// The type of incremental run. It takes only one value: IMMEDIATE .
	IncrementalRunType IncrementalRunType

	noSmithyDocumentSerde
}

// An object containing InputSourceARN , SchemaName , and ApplyNormalization .
type InputSource struct {

	// An Glue table ARN for the input source table.
	//
	// This member is required.
	InputSourceARN *string

	// The name of the schema to be retrieved.
	//
	// This member is required.
	SchemaName *string

	// Normalizes the attributes defined in the schema in the input data. For example,
	// if an attribute has an AttributeType of PHONE_NUMBER , and the data in the input
	// table is in a format of 1234567890, Entity Resolution will normalize this field
	// in the output to (123)-456-7890.
	ApplyNormalization *bool

	noSmithyDocumentSerde
}

// The Amazon S3 location that temporarily stores your data while it processes.
// Your information won't be saved permanently.
type IntermediateSourceConfiguration struct {

	// The Amazon S3 location (bucket and prefix). For example:
	// s3://provider_bucket/DOC-EXAMPLE-BUCKET
	//
	// This member is required.
	IntermediateS3Path *string

	noSmithyDocumentSerde
}

// An object containing InputRecords , TotalRecordsProcessed , MatchIDs , and
// RecordsNotProcessed .
type JobMetrics struct {

	// The total number of input records.
	InputRecords *int32

	// The total number of matchID s generated.
	MatchIDs *int32

	// The total number of records that did not get processed.
	RecordsNotProcessed *int32

	// The total number of records processed.
	TotalRecordsProcessed *int32

	noSmithyDocumentSerde
}

// An object containing the JobId , Status , StartTime , and EndTime of a job.
type JobSummary struct {

	// The ID of the job.
	//
	// This member is required.
	JobId *string

	// The time at which the job was started.
	//
	// This member is required.
	StartTime *time.Time

	// The current status of the job.
	//
	// This member is required.
	Status JobStatus

	// The time at which the job has finished.
	EndTime *time.Time

	noSmithyDocumentSerde
}

// A list of MatchingWorkflowSummary objects, each of which contain the fields
// WorkflowName , WorkflowArn , CreatedAt , UpdatedAt .
type MatchingWorkflowSummary struct {

	// The timestamp of when the workflow was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The method that has been specified for data matching, either using matching
	// provided by Entity Resolution or through a provider service.
	//
	// This member is required.
	ResolutionType ResolutionType

	// The timestamp of when the workflow was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// MatchingWorkflow .
	//
	// This member is required.
	WorkflowArn *string

	// The name of the workflow.
	//
	// This member is required.
	WorkflowName *string

	noSmithyDocumentSerde
}

// A list of OutputAttribute objects, each of which have the fields Name and Hashed
// . Each of these objects selects a column to be included in the output table, and
// whether the values of the column should be hashed.
type OutputAttribute struct {

	// A name of a column to be written to the output. This must be an InputField name
	// in the schema mapping.
	//
	// This member is required.
	Name *string

	// Enables the ability to hash the column values in the output.
	Hashed *bool

	noSmithyDocumentSerde
}

// A list of OutputAttribute objects, each of which have the fields Name and Hashed
// . Each of these objects selects a column to be included in the output table, and
// whether the values of the column should be hashed.
type OutputSource struct {

	// A list of OutputAttribute objects, each of which have the fields Name and Hashed
	// . Each of these objects selects a column to be included in the output table, and
	// whether the values of the column should be hashed.
	//
	// This member is required.
	Output []OutputAttribute

	// The S3 path to which Entity Resolution will write the output table.
	//
	// This member is required.
	OutputS3Path *string

	// Normalizes the attributes defined in the schema in the input data. For example,
	// if an attribute has an AttributeType of PHONE_NUMBER , and the data in the input
	// table is in a format of 1234567890, Entity Resolution will normalize this field
	// in the output to (123)-456-7890.
	ApplyNormalization *bool

	// Customer KMS ARN for encryption at rest. If not provided, system will use an
	// Entity Resolution managed KMS key.
	KMSArn *string

	noSmithyDocumentSerde
}

// The required configuration fields to use with the provider service.
//
// The following types satisfy this interface:
//
//	ProviderEndpointConfigurationMemberMarketplaceConfiguration
type ProviderEndpointConfiguration interface {
	isProviderEndpointConfiguration()
}

// The identifiers of the provider service, from Data Exchange.
type ProviderEndpointConfigurationMemberMarketplaceConfiguration struct {
	Value ProviderMarketplaceConfiguration

	noSmithyDocumentSerde
}

func (*ProviderEndpointConfigurationMemberMarketplaceConfiguration) isProviderEndpointConfiguration() {
}

// The required configuration fields to give intermediate access to a provider
// service.
type ProviderIntermediateDataAccessConfiguration struct {

	// The Amazon Web Services account that provider can use to read or write data
	// into the customer's intermediate S3 bucket.
	AwsAccountIds []string

	// The S3 bucket actions that the provider requires permission for.
	RequiredBucketActions []string

	noSmithyDocumentSerde
}

// The identifiers of the provider service, from Data Exchange.
type ProviderMarketplaceConfiguration struct {

	// The asset ID on Data Exchange.
	//
	// This member is required.
	AssetId *string

	// The dataset ID on Data Exchange.
	//
	// This member is required.
	DataSetId *string

	// The listing ID on Data Exchange.
	//
	// This member is required.
	ListingId *string

	// The revision ID on Data Exchange.
	//
	// This member is required.
	RevisionId *string

	noSmithyDocumentSerde
}

// An object containing the providerServiceARN , intermediateSourceConfiguration ,
// and providerConfiguration .
type ProviderProperties struct {

	// The ARN of the provider service.
	//
	// This member is required.
	ProviderServiceArn *string

	// The Amazon S3 location that temporarily stores your data while it processes.
	// Your information won't be saved permanently.
	IntermediateSourceConfiguration *IntermediateSourceConfiguration

	// The required configuration fields to use with the provider service.
	ProviderConfiguration document.Interface

	noSmithyDocumentSerde
}

// A list of ProviderService objects, each of which contain the fields providerName
// , providerServiceArn , providerServiceName , and providerServiceType .
type ProviderServiceSummary struct {

	// The name of the provider. This name is typically the company name.
	//
	// This member is required.
	ProviderName *string

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// providerService .
	//
	// This member is required.
	ProviderServiceArn *string

	// The display name of the provider service.
	//
	// This member is required.
	ProviderServiceDisplayName *string

	// The name of the product that the provider service provides.
	//
	// This member is required.
	ProviderServiceName *string

	// The type of provider service.
	//
	// This member is required.
	ProviderServiceType ServiceType

	noSmithyDocumentSerde
}

// An object which defines the resolutionType and the ruleBasedProperties .
type ResolutionTechniques struct {

	// The type of matching. There are two types of matching: RULE_MATCHING and
	// ML_MATCHING .
	//
	// This member is required.
	ResolutionType ResolutionType

	// The properties of the provider service.
	ProviderProperties *ProviderProperties

	// An object which defines the list of matching rules to run and has a field Rules
	// , which is a list of rule objects.
	RuleBasedProperties *RuleBasedProperties

	noSmithyDocumentSerde
}

// An object containing RuleName , and MatchingKeys .
type Rule struct {

	// A list of MatchingKeys . The MatchingKeys must have been defined in the
	// SchemaMapping . Two records are considered to match according to this rule if
	// all of the MatchingKeys match.
	//
	// This member is required.
	MatchingKeys []string

	// A name for the matching rule.
	//
	// This member is required.
	RuleName *string

	noSmithyDocumentSerde
}

// An object which defines the list of matching rules to run and has a field Rules
// , which is a list of rule objects.
type RuleBasedProperties struct {

	// The comparison type. You can either choose ONE_TO_ONE or MANY_TO_MANY as the
	// AttributeMatchingModel. When choosing MANY_TO_MANY , the system can match
	// attributes across the sub-types of an attribute type. For example, if the value
	// of the Email field of Profile A and the value of BusinessEmail field of Profile
	// B matches, the two profiles are matched on the Email type. When choosing
	// ONE_TO_ONE ,the system can only match if the sub-types are exact matches. For
	// example, only when the value of the Email field of Profile A and the value of
	// the Email field of Profile B matches, the two profiles are matched on the Email
	// type.
	//
	// This member is required.
	AttributeMatchingModel AttributeMatchingModel

	// A list of Rule objects, each of which have fields RuleName and MatchingKeys .
	//
	// This member is required.
	Rules []Rule

	noSmithyDocumentSerde
}

// An object containing FieldName , Type , GroupName , and MatchKey .
type SchemaInputAttribute struct {

	// A string containing the field name.
	//
	// This member is required.
	FieldName *string

	// The type of the attribute, selected from a list of values.
	//
	// This member is required.
	Type SchemaAttributeType

	// Instruct Entity Resolution to combine several columns into a unified column
	// with the identical attribute type. For example, when working with columns such
	// as first_name, middle_name, and last_name, assigning them a common GroupName
	// will prompt Entity Resolution to concatenate them into a single value.
	GroupName *string

	// A key that allows grouping of multiple input attributes into a unified matching
	// group. For example, let's consider a scenario where the source table contains
	// various addresses, such as business_address and shipping_address . By assigning
	// the MatchKey Address to both attributes, Entity Resolution will match records
	// across these fields to create a consolidated matching group. If no MatchKey is
	// specified for a column, it won't be utilized for matching purposes but will
	// still be included in the output table.
	MatchKey *string

	// The subtype of the attribute, selected from a list of values.
	SubType *string

	noSmithyDocumentSerde
}

// An object containing SchemaName , SchemaArn , CreatedAt , and UpdatedAt .
type SchemaMappingSummary struct {

	// The timestamp of when the SchemaMapping was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Specifies whether the schema mapping has been applied to a workflow.
	//
	// This member is required.
	HasWorkflows *bool

	// The ARN (Amazon Resource Name) that Entity Resolution generated for the
	// SchemaMapping .
	//
	// This member is required.
	SchemaArn *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	// The timestamp of when the SchemaMapping was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isProviderEndpointConfiguration() {}
