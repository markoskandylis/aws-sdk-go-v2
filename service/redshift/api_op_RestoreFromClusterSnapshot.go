// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cluster from a snapshot. By default, Amazon Redshift creates the
// resulting cluster with the same configuration as the original cluster from which
// the snapshot was created, except that the new cluster is created with the
// default cluster security and parameter groups. After Amazon Redshift creates the
// cluster, you can use the ModifyCluster API to associate a different security
// group and different parameter group with the restored cluster. If you are using
// a DS node type, you can also choose to change to another DS node type of the
// same size during restore. If you restore a cluster into a VPC, you must provide
// a cluster subnet group where you want the cluster restored. For more information
// about working with snapshots, go to Amazon Redshift Snapshots (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) RestoreFromClusterSnapshot(ctx context.Context, params *RestoreFromClusterSnapshotInput, optFns ...func(*Options)) (*RestoreFromClusterSnapshotOutput, error) {
	if params == nil {
		params = &RestoreFromClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreFromClusterSnapshot", params, optFns, c.addOperationRestoreFromClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreFromClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreFromClusterSnapshotInput struct {

	// The identifier of the cluster that will be created from restoring the snapshot.
	// Constraints:
	//   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	//   - Alphabetic characters must be lowercase.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//   - Must be unique for all clusters within an Amazon Web Services account.
	//
	// This member is required.
	ClusterIdentifier *string

	// Reserved.
	AdditionalInfo *string

	// If true , major version upgrades can be applied during the maintenance window to
	// the Amazon Redshift engine that is running on the cluster. Default: true
	AllowVersionUpgrade *bool

	// This parameter is retired. It does not set the AQUA configuration status.
	// Amazon Redshift automatically determines whether to use AQUA (Advanced Query
	// Accelerator).
	AquaConfigurationStatus types.AquaConfigurationStatus

	// The number of days that automated snapshots are retained. If the value is 0,
	// automated snapshots are disabled. Even if automated snapshots are disabled, you
	// can still create manual snapshots when you want with CreateClusterSnapshot . You
	// can't disable automated snapshots for RA3 node types. Set the automated
	// retention period from 1-35 days. Default: The value selected for the cluster
	// from which the snapshot was taken. Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *int32

	// The Amazon EC2 Availability Zone in which to restore the cluster. Default: A
	// random, system-chosen Availability Zone. Example: us-east-2a
	AvailabilityZone *string

	// The option to enable relocation for an Amazon Redshift cluster between
	// Availability Zones after the cluster is restored.
	AvailabilityZoneRelocation *bool

	// The name of the parameter group to be associated with this cluster. Default:
	// The default Amazon Redshift cluster parameter group. For information about the
	// default parameter group, go to Working with Amazon Redshift Parameter Groups (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	// . Constraints:
	//   - Must be 1 to 255 alphanumeric characters or hyphens.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	ClusterParameterGroupName *string

	// A list of security groups to be associated with this cluster. Default: The
	// default cluster security group for Amazon Redshift. Cluster security groups only
	// apply to clusters outside of VPCs.
	ClusterSecurityGroups []string

	// The name of the subnet group where you want to cluster restored. A snapshot of
	// cluster in VPC can be restored only in VPC. Therefore, you must provide subnet
	// group name where you want the cluster restored.
	ClusterSubnetGroupName *string

	// The Amazon Resource Name (ARN) for the IAM role that was set as default for the
	// cluster when the cluster was last modified while it was restored from a
	// snapshot.
	DefaultIamRoleArn *string

	// The Elastic IP (EIP) address for the cluster. Don't specify the Elastic IP
	// address for a publicly accessible cluster with availability zone relocation
	// turned on.
	ElasticIp *string

	// Enables support for restoring an unencrypted snapshot to a cluster encrypted
	// with Key Management Service (KMS) and a customer managed key.
	Encrypted *bool

	// An option that specifies whether to create the cluster with enhanced VPC
	// routing enabled. To create a cluster that uses enhanced VPC routing, the cluster
	// must be in a VPC. For more information, see Enhanced VPC Routing (https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html)
	// in the Amazon Redshift Cluster Management Guide. If this option is true ,
	// enhanced VPC routing is enabled. Default: false
	EnhancedVpcRouting *bool

	// Specifies the name of the HSM client certificate the Amazon Redshift cluster
	// uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string

	// Specifies the name of the HSM configuration that contains the information the
	// Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string

	// A list of Identity and Access Management (IAM) roles that can be used by the
	// cluster to access other Amazon Web Services services. You must supply the IAM
	// roles in their Amazon Resource Name (ARN) format. The maximum number of IAM
	// roles that you can associate is subject to a quota. For more information, go to
	// Quotas and limits (https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html)
	// in the Amazon Redshift Cluster Management Guide.
	IamRoles []string

	// The Key Management Service (KMS) key ID of the encryption key that encrypts
	// data in the cluster restored from a shared snapshot. You can also provide the
	// key ID when you restore from an unencrypted snapshot to an encrypted cluster in
	// the same account. Additionally, you can specify a new KMS key ID when you
	// restore from an encrypted snapshot in the same account in order to change it. In
	// that case, the restored cluster is encrypted with the new KMS key ID.
	KmsKeyId *string

	// The name of the maintenance track for the restored cluster. When you take a
	// snapshot, the snapshot inherits the MaintenanceTrack value from the cluster.
	// The snapshot might be on a different track than the cluster that was the source
	// for the snapshot. For example, suppose that you take a snapshot of a cluster
	// that is on the current track and then change the cluster to be on the trailing
	// track. In this case, the snapshot and the source cluster are on different
	// tracks.
	MaintenanceTrackName *string

	// If true , Amazon Redshift uses Secrets Manager to manage the restored cluster's
	// admin credentials. If ManageMasterPassword is false or not set, Amazon Redshift
	// uses the admin credentials the cluster had at the time the snapshot was taken.
	ManageMasterPassword *bool

	// The default number of days to retain a manual snapshot. If the value is -1, the
	// snapshot is retained indefinitely. This setting doesn't change the retention
	// period of existing snapshots. The value must be either -1 or an integer between
	// 1 and 3,653.
	ManualSnapshotRetentionPeriod *int32

	// The ID of the Key Management Service (KMS) key used to encrypt and store the
	// cluster's admin credentials secret. You can only use this parameter if
	// ManageMasterPassword is true.
	MasterPasswordSecretKmsKeyId *string

	// The node type that the restored cluster will be provisioned with. Default: The
	// node type of the cluster from which the snapshot was taken. You can modify this
	// if you are using any DS node type. In that case, you can choose to restore into
	// another DS node type of the same size. For example, you can restore ds1.8xlarge
	// into ds2.8xlarge, or ds1.xlarge into ds2.xlarge. If you have a DC instance type,
	// you must restore into that same instance type and size. In other words, you can
	// only restore a dc1.large instance type into another dc1.large instance type or
	// dc2.large instance type. You can't restore dc1.8xlarge to dc2.8xlarge. First
	// restore to a dc1.8xlarge cluster, then resize to a dc2.8large cluster. For more
	// information about node types, see About Clusters and Nodes (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-about-clusters-and-nodes)
	// in the Amazon Redshift Cluster Management Guide.
	NodeType *string

	// The number of nodes specified when provisioning the restored cluster.
	NumberOfNodes *int32

	// The Amazon Web Services account used to create or copy the snapshot. Required
	// if you are restoring a snapshot you do not own, optional if you own the
	// snapshot.
	OwnerAccount *string

	// The port number on which the cluster accepts connections. Default: The same
	// port as the original cluster. Constraints: Must be between 1115 and 65535 .
	Port *int32

	// The weekly time range (in UTC) during which automated cluster maintenance can
	// occur. Format: ddd:hh24:mi-ddd:hh24:mi Default: The value selected for the
	// cluster from which the snapshot was taken. For more information about the time
	// blocks for each region, see Maintenance Windows (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows)
	// in Amazon Redshift Cluster Management Guide. Valid Days: Mon | Tue | Wed | Thu |
	// Fri | Sat | Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// If true , the cluster can be accessed from a public network.
	PubliclyAccessible *bool

	// The identifier of the target reserved node offering.
	ReservedNodeId *string

	// The Amazon Resource Name (ARN) of the snapshot associated with the message to
	// restore from a cluster. You must specify this parameter or snapshotIdentifier ,
	// but not both.
	SnapshotArn *string

	// The name of the cluster the source snapshot was created from. This parameter is
	// required if your IAM user has a policy containing a snapshot resource element
	// that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string

	// The name of the snapshot from which to create the new cluster. This parameter
	// isn't case sensitive. You must specify this parameter or snapshotArn , but not
	// both. Example: my-snapshot-id
	SnapshotIdentifier *string

	// A unique identifier for the snapshot schedule.
	SnapshotScheduleIdentifier *string

	// The identifier of the target reserved node offering.
	TargetReservedNodeOfferingId *string

	// A list of Virtual Private Cloud (VPC) security groups to be associated with the
	// cluster. Default: The default VPC security group is associated with the cluster.
	// VPC security groups only apply to clusters in VPCs.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreFromClusterSnapshotOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreFromClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreFromClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreFromClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addRestoreFromClusterSnapshotResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRestoreFromClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreFromClusterSnapshot(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRestoreFromClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RestoreFromClusterSnapshot",
	}
}

type opRestoreFromClusterSnapshotResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opRestoreFromClusterSnapshotResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRestoreFromClusterSnapshotResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "redshift"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "redshift"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("redshift")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addRestoreFromClusterSnapshotResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opRestoreFromClusterSnapshotResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
