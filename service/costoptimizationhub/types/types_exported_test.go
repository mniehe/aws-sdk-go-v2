// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/mniehe/aws-sdk-go-v2/service/costoptimizationhub/types"
)

func ExampleResourceDetails_outputUsage() {
	var union types.ResourceDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ResourceDetailsMemberComputeSavingsPlans:
		_ = v.Value // Value is types.ComputeSavingsPlans

	case *types.ResourceDetailsMemberEbsVolume:
		_ = v.Value // Value is types.EbsVolume

	case *types.ResourceDetailsMemberEc2AutoScalingGroup:
		_ = v.Value // Value is types.Ec2AutoScalingGroup

	case *types.ResourceDetailsMemberEc2Instance:
		_ = v.Value // Value is types.Ec2Instance

	case *types.ResourceDetailsMemberEc2InstanceSavingsPlans:
		_ = v.Value // Value is types.Ec2InstanceSavingsPlans

	case *types.ResourceDetailsMemberEc2ReservedInstances:
		_ = v.Value // Value is types.Ec2ReservedInstances

	case *types.ResourceDetailsMemberEcsService:
		_ = v.Value // Value is types.EcsService

	case *types.ResourceDetailsMemberElastiCacheReservedInstances:
		_ = v.Value // Value is types.ElastiCacheReservedInstances

	case *types.ResourceDetailsMemberLambdaFunction:
		_ = v.Value // Value is types.LambdaFunction

	case *types.ResourceDetailsMemberOpenSearchReservedInstances:
		_ = v.Value // Value is types.OpenSearchReservedInstances

	case *types.ResourceDetailsMemberRdsReservedInstances:
		_ = v.Value // Value is types.RdsReservedInstances

	case *types.ResourceDetailsMemberRedshiftReservedInstances:
		_ = v.Value // Value is types.RedshiftReservedInstances

	case *types.ResourceDetailsMemberSageMakerSavingsPlans:
		_ = v.Value // Value is types.SageMakerSavingsPlans

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.ComputeSavingsPlans
var _ *types.Ec2AutoScalingGroup
var _ *types.RdsReservedInstances
var _ *types.OpenSearchReservedInstances
var _ *types.Ec2ReservedInstances
var _ *types.SageMakerSavingsPlans
var _ *types.Ec2Instance
var _ *types.Ec2InstanceSavingsPlans
var _ *types.EcsService
var _ *types.RedshiftReservedInstances
var _ *types.ElastiCacheReservedInstances
var _ *types.LambdaFunction
var _ *types.EbsVolume
