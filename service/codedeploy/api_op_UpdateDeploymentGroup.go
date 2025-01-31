// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes information about a deployment group.
func (c *Client) UpdateDeploymentGroup(ctx context.Context, params *UpdateDeploymentGroupInput, optFns ...func(*Options)) (*UpdateDeploymentGroupOutput, error) {
	if params == nil {
		params = &UpdateDeploymentGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDeploymentGroup", params, optFns, c.addOperationUpdateDeploymentGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDeploymentGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateDeploymentGroup operation.
type UpdateDeploymentGroupInput struct {

	// The application name that corresponds to the deployment group to update.
	//
	// This member is required.
	ApplicationName *string

	// The current name of the deployment group.
	//
	// This member is required.
	CurrentDeploymentGroupName *string

	// Information to add or change about Amazon CloudWatch alarms when the deployment
	// group is updated.
	AlarmConfiguration *types.AlarmConfiguration

	// Information for an automatic rollback configuration that is added or changed
	// when a deployment group is updated.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration

	// The replacement list of Auto Scaling groups to be included in the deployment
	// group, if you want to change them.
	//   - To keep the Auto Scaling groups, enter their names or do not specify this
	//   parameter.
	//   - To remove Auto Scaling groups, specify a non-null empty list of Auto
	//   Scaling group names to detach all CodeDeploy-managed Auto Scaling lifecycle
	//   hooks. For examples, see Amazon EC2 instances in an Amazon EC2 Auto Scaling
	//   group fail to launch and receive the error "Heartbeat Timeout" (https://docs.aws.amazon.com/codedeploy/latest/userguide/troubleshooting-auto-scaling.html#troubleshooting-auto-scaling-heartbeat)
	//   in the CodeDeploy User Guide.
	AutoScalingGroups []string

	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration *types.BlueGreenDeploymentConfiguration

	// The replacement deployment configuration name to use, if you want to change it.
	DeploymentConfigName *string

	// Information about the type of deployment, either in-place or blue/green, you
	// want to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle *types.DeploymentStyle

	// The replacement set of Amazon EC2 tags on which to filter, if you want to
	// change them. To keep the existing tags, enter their names. To remove tags, do
	// not enter any tag names.
	Ec2TagFilters []types.EC2TagFilter

	// Information about groups of tags applied to on-premises instances. The
	// deployment group includes only Amazon EC2 instances identified by all the tag
	// groups.
	Ec2TagSet *types.EC2TagSet

	// The target Amazon ECS services in the deployment group. This applies only to
	// deployment groups that use the Amazon ECS compute platform. A target Amazon ECS
	// service is specified as an Amazon ECS cluster and service name pair using the
	// format : .
	EcsServices []types.ECSService

	// Information about the load balancer used in a deployment.
	LoadBalancerInfo *types.LoadBalancerInfo

	// The new name of the deployment group, if you want to change it.
	NewDeploymentGroupName *string

	// The replacement set of on-premises instance tags on which to filter, if you
	// want to change them. To keep the existing tags, enter their names. To remove
	// tags, do not enter any tag names.
	OnPremisesInstanceTagFilters []types.TagFilter

	// Information about an on-premises instance tag set. The deployment group
	// includes only on-premises instances identified by all the tag groups.
	OnPremisesTagSet *types.OnPremisesTagSet

	// Indicates what happens when new Amazon EC2 instances are launched
	// mid-deployment and do not receive the deployed application revision. If this
	// option is set to UPDATE or is unspecified, CodeDeploy initiates one or more
	// 'auto-update outdated instances' deployments to apply the deployed application
	// revision to the new Amazon EC2 instances. If this option is set to IGNORE ,
	// CodeDeploy does not initiate a deployment to update the new Amazon EC2
	// instances. This may result in instances having different revisions.
	OutdatedInstancesStrategy types.OutdatedInstancesStrategy

	// A replacement ARN for the service role, if you want to change it.
	ServiceRoleArn *string

	// Information about triggers to change when the deployment group is updated. For
	// examples, see Edit a Trigger in a CodeDeploy Deployment Group (https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-edit.html)
	// in the CodeDeploy User Guide.
	TriggerConfigurations []types.TriggerConfig

	noSmithyDocumentSerde
}

// Represents the output of an UpdateDeploymentGroup operation.
type UpdateDeploymentGroupOutput struct {

	// If the output contains no data, and the corresponding deployment group
	// contained at least one Auto Scaling group, CodeDeploy successfully removed all
	// corresponding Auto Scaling lifecycle event hooks from the Amazon Web Services
	// account. If the output contains data, CodeDeploy could not remove some Auto
	// Scaling lifecycle event hooks from the Amazon Web Services account.
	HooksNotCleanedUp []types.AutoScalingGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDeploymentGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDeploymentGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDeploymentGroup{}, middleware.After)
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
	if err = addUpdateDeploymentGroupResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateDeploymentGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDeploymentGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDeploymentGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "UpdateDeploymentGroup",
	}
}

type opUpdateDeploymentGroupResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateDeploymentGroupResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateDeploymentGroupResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "codedeploy"
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
				signingName = "codedeploy"
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
				v4aScheme.SigningName = aws.String("codedeploy")
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

func addUpdateDeploymentGroupResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateDeploymentGroupResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
