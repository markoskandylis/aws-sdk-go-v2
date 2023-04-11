// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You already have a DAX cluster with the given identifier.
type ClusterAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested cluster ID does not refer to an existing DAX cluster.
type ClusterNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have attempted to exceed the maximum number of DAX clusters for your AWS
// account.
type ClusterQuotaForCustomerExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterQuotaForCustomerExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There are not enough system resources to create the cluster you requested (or to
// resize an already-existing cluster).
type InsufficientClusterCapacityFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InsufficientClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientClusterCapacityFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InsufficientClusterCapacityFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InsufficientClusterCapacityFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Amazon Resource Name (ARN) supplied in the request is not valid.
type InvalidARNFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidARNFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidARNFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidARNFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidARNFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidARNFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested DAX cluster is not in the available state.
type InvalidClusterStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClusterStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidClusterStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Two or more incompatible parameters were specified.
type InvalidParameterCombinationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterCombinationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// One or more parameters in a parameter group are in an invalid state.
type InvalidParameterGroupStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterGroupStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterGroupStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value for a parameter is invalid.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid subnet identifier was specified.
type InvalidSubnet struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSubnet"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The VPC network is in an invalid state.
type InvalidVPCNetworkStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidVPCNetworkStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCNetworkStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCNetworkStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidVPCNetworkStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidVPCNetworkStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// None of the nodes in the cluster have the given node ID.
type NodeNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NodeNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NodeNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NodeNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have attempted to exceed the maximum number of nodes for a DAX cluster.
type NodeQuotaForClusterExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NodeQuotaForClusterExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForClusterExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForClusterExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NodeQuotaForClusterExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NodeQuotaForClusterExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have attempted to exceed the maximum number of nodes for your AWS account.
type NodeQuotaForCustomerExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NodeQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForCustomerExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NodeQuotaForCustomerExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NodeQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified parameter group already exists.
type ParameterGroupAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified parameter group does not exist.
type ParameterGroupNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have attempted to exceed the maximum number of parameter groups.
type ParameterGroupQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified service linked role (SLR) was not found.
type ServiceLinkedRoleNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceLinkedRoleNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceLinkedRoleNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the maximum number of x509 certificates that can be created for
// encrypted clusters in a 30 day period. Contact AWS customer support to discuss
// options for continuing to create encrypted clusters.
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified subnet group already exists.
type SubnetGroupAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified subnet group is currently in use.
type SubnetGroupInUseFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupInUseFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupInUseFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupInUseFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupInUseFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupInUseFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet group name does not refer to an existing subnet group.
type SubnetGroupNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// subnets in a subnet group.
type SubnetGroupQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested subnet is being used by another subnet group.
type SubnetInUse struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetInUse) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the allowed number of
// subnets in a subnet group.
type SubnetQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The tag does not exist.
type TagNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TagNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *TagNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of tags for this DAX cluster.
type TagQuotaPerResourceExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TagQuotaPerResourceExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagQuotaPerResourceExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagQuotaPerResourceExceeded) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagQuotaPerResourceExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *TagQuotaPerResourceExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
