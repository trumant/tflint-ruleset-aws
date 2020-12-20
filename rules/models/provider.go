// This file generated by `rules/models/generator`. DO NOT EDIT

package models

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Rules is a list of rules generated from aws-sdk-go
var Rules = []tflint.Rule{
	NewAwsAcmCertificateInvalidCertificateBodyRule(),
	NewAwsAcmCertificateInvalidCertificateChainRule(),
	NewAwsAcmCertificateInvalidPrivateKeyRule(),
	NewAwsAcmpcaCertificateAuthorityInvalidTypeRule(),
	NewAwsALBInvalidIPAddressTypeRule(),
	NewAwsALBInvalidLoadBalancerTypeRule(),
	NewAwsALBListenerInvalidProtocolRule(),
	NewAwsALBTargetGroupInvalidProtocolRule(),
	NewAwsALBTargetGroupInvalidTargetTypeRule(),
	NewAwsAMIInvalidArchitectureRule(),
	NewAwsAPIGatewayAuthorizerInvalidTypeRule(),
	NewAwsAPIGatewayGatewayResponseInvalidResponseTypeRule(),
	NewAwsAPIGatewayGatewayResponseInvalidStatusCodeRule(),
	NewAwsAPIGatewayIntegrationInvalidConnectionTypeRule(),
	NewAwsAPIGatewayIntegrationInvalidContentHandlingRule(),
	NewAwsAPIGatewayIntegrationInvalidTypeRule(),
	NewAwsAPIGatewayIntegrationResponseInvalidContentHandlingRule(),
	NewAwsAPIGatewayIntegrationResponseInvalidStatusCodeRule(),
	NewAwsAPIGatewayMethodResponseInvalidStatusCodeRule(),
	NewAwsAPIGatewayRestAPIInvalidAPIKeySourceRule(),
	NewAwsAPIGatewayStageInvalidCacheClusterSizeRule(),
	NewAwsAppautoscalingPolicyInvalidPolicyTypeRule(),
	NewAwsAppautoscalingPolicyInvalidScalableDimensionRule(),
	NewAwsAppautoscalingPolicyInvalidServiceNamespaceRule(),
	NewAwsAppautoscalingScheduledActionInvalidScalableDimensionRule(),
	NewAwsAppautoscalingScheduledActionInvalidServiceNamespaceRule(),
	NewAwsAppautoscalingTargetInvalidScalableDimensionRule(),
	NewAwsAppautoscalingTargetInvalidServiceNamespaceRule(),
	NewAwsAppmeshMeshInvalidNameRule(),
	NewAwsAppmeshRouteInvalidMeshNameRule(),
	NewAwsAppmeshRouteInvalidNameRule(),
	NewAwsAppmeshRouteInvalidVirtualRouterNameRule(),
	NewAwsAppmeshVirtualNodeInvalidMeshNameRule(),
	NewAwsAppmeshVirtualNodeInvalidNameRule(),
	NewAwsAppmeshVirtualRouterInvalidMeshNameRule(),
	NewAwsAppmeshVirtualRouterInvalidNameRule(),
	NewAwsAppmeshVirtualServiceInvalidMeshNameRule(),
	NewAwsAppmeshVirtualServiceInvalidNameRule(),
	NewAwsAppsyncDatasourceInvalidNameRule(),
	NewAwsAppsyncDatasourceInvalidTypeRule(),
	NewAwsAppsyncFunctionInvalidDataSourceRule(),
	NewAwsAppsyncFunctionInvalidNameRule(),
	NewAwsAppsyncFunctionInvalidRequestMappingTemplateRule(),
	NewAwsAppsyncFunctionInvalidResponseMappingTemplateRule(),
	NewAwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule(),
	NewAwsAppsyncResolverInvalidDataSourceRule(),
	NewAwsAppsyncResolverInvalidFieldRule(),
	NewAwsAppsyncResolverInvalidRequestTemplateRule(),
	NewAwsAppsyncResolverInvalidResponseTemplateRule(),
	NewAwsAppsyncResolverInvalidTypeRule(),
	NewAwsAthenaDatabaseInvalidNameRule(),
	NewAwsAthenaNamedQueryInvalidDatabaseRule(),
	NewAwsAthenaNamedQueryInvalidDescriptionRule(),
	NewAwsAthenaNamedQueryInvalidNameRule(),
	NewAwsAthenaNamedQueryInvalidQueryRule(),
	NewAwsAthenaWorkgroupInvalidDescriptionRule(),
	NewAwsAthenaWorkgroupInvalidNameRule(),
	NewAwsAthenaWorkgroupInvalidStateRule(),
	NewAwsBackupSelectionInvalidNameRule(),
	NewAwsBackupVaultInvalidNameRule(),
	NewAwsBatchComputeEnvironmentInvalidStateRule(),
	NewAwsBatchComputeEnvironmentInvalidTypeRule(),
	NewAwsBatchJobDefinitionInvalidTypeRule(),
	NewAwsBatchJobQueueInvalidStateRule(),
	NewAwsBudgetsBudgetInvalidAccountIDRule(),
	NewAwsBudgetsBudgetInvalidBudgetTypeRule(),
	NewAwsBudgetsBudgetInvalidNameRule(),
	NewAwsBudgetsBudgetInvalidTimeUnitRule(),
	NewAwsCloud9EnvironmentEc2InvalidDescriptionRule(),
	NewAwsCloud9EnvironmentEc2InvalidInstanceTypeRule(),
	NewAwsCloud9EnvironmentEc2InvalidNameRule(),
	NewAwsCloud9EnvironmentEc2InvalidOwnerArnRule(),
	NewAwsCloud9EnvironmentEc2InvalidSubnetIDRule(),
	NewAwsCloudformationStackInvalidIAMRoleArnRule(),
	NewAwsCloudformationStackInvalidOnFailureRule(),
	NewAwsCloudformationStackInvalidPolicyBodyRule(),
	NewAwsCloudformationStackInvalidPolicyURLRule(),
	NewAwsCloudformationStackInvalidTemplateURLRule(),
	NewAwsCloudformationStackSetInstanceInvalidAccountIDRule(),
	NewAwsCloudformationStackSetInvalidAdministrationRoleArnRule(),
	NewAwsCloudformationStackSetInvalidDescriptionRule(),
	NewAwsCloudformationStackSetInvalidExecutionRoleNameRule(),
	NewAwsCloudformationStackSetInvalidTemplateURLRule(),
	NewAwsCloudfrontDistributionInvalidHTTPVersionRule(),
	NewAwsCloudfrontDistributionInvalidPriceClassRule(),
	NewAwsCloudhsmV2ClusterInvalidHsmTypeRule(),
	NewAwsCloudhsmV2ClusterInvalidSourceBackupIdentifierRule(),
	NewAwsCloudhsmV2HsmInvalidAvailabilityZoneRule(),
	NewAwsCloudhsmV2HsmInvalidClusterIDRule(),
	NewAwsCloudhsmV2HsmInvalidIPAddressRule(),
	NewAwsCloudhsmV2HsmInvalidSubnetIDRule(),
	NewAwsCloudwatchEventPermissionInvalidActionRule(),
	NewAwsCloudwatchEventPermissionInvalidPrincipalRule(),
	NewAwsCloudwatchEventPermissionInvalidStatementIDRule(),
	NewAwsCloudwatchEventRuleInvalidDescriptionRule(),
	NewAwsCloudwatchEventRuleInvalidNameRule(),
	NewAwsCloudwatchEventRuleInvalidRoleArnRule(),
	NewAwsCloudwatchEventRuleInvalidScheduleExpressionRule(),
	NewAwsCloudwatchEventTargetInvalidArnRule(),
	NewAwsCloudwatchEventTargetInvalidInputRule(),
	NewAwsCloudwatchEventTargetInvalidInputPathRule(),
	NewAwsCloudwatchEventTargetInvalidRoleArnRule(),
	NewAwsCloudwatchEventTargetInvalidRuleRule(),
	NewAwsCloudwatchEventTargetInvalidTargetIDRule(),
	NewAwsCloudwatchLogDestinationInvalidNameRule(),
	NewAwsCloudwatchLogDestinationPolicyInvalidDestinationNameRule(),
	NewAwsCloudwatchLogGroupInvalidKmsKeyIDRule(),
	NewAwsCloudwatchLogGroupInvalidNameRule(),
	NewAwsCloudwatchLogMetricFilterInvalidLogGroupNameRule(),
	NewAwsCloudwatchLogMetricFilterInvalidNameRule(),
	NewAwsCloudwatchLogMetricFilterInvalidPatternRule(),
	NewAwsCloudwatchLogResourcePolicyInvalidPolicyDocumentRule(),
	NewAwsCloudwatchLogStreamInvalidLogGroupNameRule(),
	NewAwsCloudwatchLogStreamInvalidNameRule(),
	NewAwsCloudwatchLogSubscriptionFilterInvalidDistributionRule(),
	NewAwsCloudwatchLogSubscriptionFilterInvalidFilterPatternRule(),
	NewAwsCloudwatchLogSubscriptionFilterInvalidLogGroupNameRule(),
	NewAwsCloudwatchLogSubscriptionFilterInvalidNameRule(),
	NewAwsCloudwatchMetricAlarmInvalidAlarmDescriptionRule(),
	NewAwsCloudwatchMetricAlarmInvalidAlarmNameRule(),
	NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule(),
	NewAwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule(),
	NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule(),
	NewAwsCloudwatchMetricAlarmInvalidMetricNameRule(),
	NewAwsCloudwatchMetricAlarmInvalidNamespaceRule(),
	NewAwsCloudwatchMetricAlarmInvalidStatisticRule(),
	NewAwsCloudwatchMetricAlarmInvalidTreatMissingDataRule(),
	NewAwsCloudwatchMetricAlarmInvalidUnitRule(),
	NewAwsCodebuildProjectInvalidDescriptionRule(),
	NewAwsCodebuildSourceCredentialInvalidAuthTypeRule(),
	NewAwsCodebuildSourceCredentialInvalidServerTypeRule(),
	NewAwsCodecommitRepositoryInvalidDefaultBranchRule(),
	NewAwsCodecommitRepositoryInvalidDescriptionRule(),
	NewAwsCodecommitRepositoryInvalidRepositoryNameRule(),
	NewAwsCodecommitTriggerInvalidRepositoryNameRule(),
	NewAwsCodedeployAppInvalidComputePlatformRule(),
	NewAwsCodedeployAppInvalidNameRule(),
	NewAwsCodedeployDeploymentConfigInvalidComputePlatformRule(),
	NewAwsCodedeployDeploymentConfigInvalidDeploymentConfigNameRule(),
	NewAwsCodedeployDeploymentGroupInvalidAppNameRule(),
	NewAwsCodedeployDeploymentGroupInvalidDeploymentConfigNameRule(),
	NewAwsCodedeployDeploymentGroupInvalidDeploymentGroupNameRule(),
	NewAwsCodepipelineInvalidNameRule(),
	NewAwsCodepipelineInvalidRoleArnRule(),
	NewAwsCodepipelineWebhookInvalidAuthenticationRule(),
	NewAwsCodepipelineWebhookInvalidNameRule(),
	NewAwsCodepipelineWebhookInvalidTargetActionRule(),
	NewAwsCodepipelineWebhookInvalidTargetPipelineRule(),
	NewAwsCognitoIdentityPoolInvalidDeveloperProviderNameRule(),
	NewAwsCognitoIdentityPoolInvalidIdentityPoolNameRule(),
	NewAwsCognitoIdentityPoolRolesAttachmentInvalidIdentityPoolIDRule(),
	NewAwsCognitoIdentityProviderInvalidProviderNameRule(),
	NewAwsCognitoIdentityProviderInvalidProviderTypeRule(),
	NewAwsCognitoIdentityProviderInvalidUserPoolIDRule(),
	NewAwsCognitoResourceServerInvalidIdentifierRule(),
	NewAwsCognitoResourceServerInvalidNameRule(),
	NewAwsCognitoUserGroupInvalidDescriptionRule(),
	NewAwsCognitoUserGroupInvalidNameRule(),
	NewAwsCognitoUserGroupInvalidRoleArnRule(),
	NewAwsCognitoUserGroupInvalidUserPoolIDRule(),
	NewAwsCognitoUserPoolClientInvalidDefaultRedirectURIRule(),
	NewAwsCognitoUserPoolClientInvalidNameRule(),
	NewAwsCognitoUserPoolClientInvalidUserPoolIDRule(),
	NewAwsCognitoUserPoolDomainInvalidCertificateArnRule(),
	NewAwsCognitoUserPoolDomainInvalidDomainRule(),
	NewAwsCognitoUserPoolDomainInvalidUserPoolIDRule(),
	NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule(),
	NewAwsCognitoUserPoolInvalidEmailVerificationSubjectRule(),
	NewAwsCognitoUserPoolInvalidMfaConfigurationRule(),
	NewAwsCognitoUserPoolInvalidNameRule(),
	NewAwsCognitoUserPoolInvalidSmsAuthenticationMessageRule(),
	NewAwsCognitoUserPoolInvalidSmsVerificationMessageRule(),
	NewAwsConfigAggregateAuthorizationInvalidAccountIDRule(),
	NewAwsConfigAggregateAuthorizationInvalidRegionRule(),
	NewAwsConfigConfigRuleInvalidDescriptionRule(),
	NewAwsConfigConfigRuleInvalidInputParametersRule(),
	NewAwsConfigConfigRuleInvalidMaximumExecutionFrequencyRule(),
	NewAwsConfigConfigRuleInvalidNameRule(),
	NewAwsConfigConfigurationAggregatorInvalidNameRule(),
	NewAwsConfigConfigurationRecorderInvalidNameRule(),
	NewAwsConfigConfigurationRecorderStatusInvalidNameRule(),
	NewAwsConfigDeliveryChannelInvalidNameRule(),
	NewAwsConfigOrganizationCustomRuleInvalidDescriptionRule(),
	NewAwsConfigOrganizationCustomRuleInvalidInputParametersRule(),
	NewAwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule(),
	NewAwsConfigOrganizationCustomRuleInvalidMaximumExecutionFrequencyRule(),
	NewAwsConfigOrganizationCustomRuleInvalidNameRule(),
	NewAwsConfigOrganizationCustomRuleInvalidResourceIDScopeRule(),
	NewAwsConfigOrganizationCustomRuleInvalidTagKeyScopeRule(),
	NewAwsConfigOrganizationCustomRuleInvalidTagValueScopeRule(),
	NewAwsConfigOrganizationManagedRuleInvalidDescriptionRule(),
	NewAwsConfigOrganizationManagedRuleInvalidInputParametersRule(),
	NewAwsConfigOrganizationManagedRuleInvalidMaximumExecutionFrequencyRule(),
	NewAwsConfigOrganizationManagedRuleInvalidNameRule(),
	NewAwsConfigOrganizationManagedRuleInvalidResourceIDScopeRule(),
	NewAwsConfigOrganizationManagedRuleInvalidRuleIdentifierRule(),
	NewAwsConfigOrganizationManagedRuleInvalidTagKeyScopeRule(),
	NewAwsConfigOrganizationManagedRuleInvalidTagValueScopeRule(),
	NewAwsCurReportDefinitionInvalidCompressionRule(),
	NewAwsCurReportDefinitionInvalidFormatRule(),
	NewAwsCurReportDefinitionInvalidReportNameRule(),
	NewAwsCurReportDefinitionInvalidS3BucketRule(),
	NewAwsCurReportDefinitionInvalidS3PrefixRule(),
	NewAwsCurReportDefinitionInvalidS3RegionRule(),
	NewAwsCurReportDefinitionInvalidTimeUnitRule(),
	NewAwsCustomerGatewayInvalidTypeRule(),
	NewAwsDatasyncAgentInvalidActivationKeyRule(),
	NewAwsDatasyncAgentInvalidNameRule(),
	NewAwsDatasyncLocationEfsInvalidEfsFileSystemArnRule(),
	NewAwsDatasyncLocationEfsInvalidSubdirectoryRule(),
	NewAwsDatasyncLocationNfsInvalidServerHostnameRule(),
	NewAwsDatasyncLocationNfsInvalidSubdirectoryRule(),
	NewAwsDatasyncLocationS3InvalidS3BucketArnRule(),
	NewAwsDatasyncLocationS3InvalidSubdirectoryRule(),
	NewAwsDatasyncTaskInvalidCloudwatchLogGroupArnRule(),
	NewAwsDatasyncTaskInvalidDestinationLocationArnRule(),
	NewAwsDatasyncTaskInvalidNameRule(),
	NewAwsDatasyncTaskInvalidSourceLocationArnRule(),
	NewAwsDevicefarmProjectInvalidNameRule(),
	NewAwsDirectoryServiceConditionalForwarderInvalidDirectoryIDRule(),
	NewAwsDirectoryServiceConditionalForwarderInvalidRemoteDomainNameRule(),
	NewAwsDirectoryServiceDirectoryInvalidDescriptionRule(),
	NewAwsDirectoryServiceDirectoryInvalidEditionRule(),
	NewAwsDirectoryServiceDirectoryInvalidNameRule(),
	NewAwsDirectoryServiceDirectoryInvalidPasswordRule(),
	NewAwsDirectoryServiceDirectoryInvalidShortNameRule(),
	NewAwsDirectoryServiceDirectoryInvalidSizeRule(),
	NewAwsDirectoryServiceDirectoryInvalidTypeRule(),
	NewAwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule(),
	NewAwsDirectoryServiceLogSubscriptionInvalidLogGroupNameRule(),
	NewAwsDlmLifecyclePolicyInvalidDescriptionRule(),
	NewAwsDlmLifecyclePolicyInvalidExecutionRoleArnRule(),
	NewAwsDlmLifecyclePolicyInvalidStateRule(),
	NewAwsDmsEndpointInvalidEndpointTypeRule(),
	NewAwsDmsEndpointInvalidSslModeRule(),
	NewAwsDmsReplicationTaskInvalidMigrationTypeRule(),
	NewAwsDxBgpPeerInvalidAddressFamilyRule(),
	NewAwsDxHostedPrivateVirtualInterfaceInvalidAddressFamilyRule(),
	NewAwsDxHostedPublicVirtualInterfaceInvalidAddressFamilyRule(),
	NewAwsDxPrivateVirtualInterfaceInvalidAddressFamilyRule(),
	NewAwsDxPublicVirtualInterfaceInvalidAddressFamilyRule(),
	NewAwsDynamoDBGlobalTableInvalidNameRule(),
	NewAwsDynamoDBTableInvalidBillingModeRule(),
	NewAwsDynamoDBTableInvalidHashKeyRule(),
	NewAwsDynamoDBTableInvalidNameRule(),
	NewAwsDynamoDBTableInvalidRangeKeyRule(),
	NewAwsDynamoDBTableItemInvalidHashKeyRule(),
	NewAwsDynamoDBTableItemInvalidRangeKeyRule(),
	NewAwsDynamoDBTableItemInvalidTableNameRule(),
	NewAwsEbsVolumeInvalidTypeRule(),
	NewAwsEc2CapacityReservationInvalidEndDateTypeRule(),
	NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule(),
	NewAwsEc2CapacityReservationInvalidInstancePlatformRule(),
	NewAwsEc2CapacityReservationInvalidTenancyRule(),
	NewAwsEc2ClientVpnEndpointInvalidTransportProtocolRule(),
	NewAwsEc2FleetInvalidExcessCapacityTerminationPolicyRule(),
	NewAwsEc2FleetInvalidTypeRule(),
	NewAwsEc2TransitGatewayInvalidAutoAcceptSharedAttachmentsRule(),
	NewAwsEc2TransitGatewayInvalidDefaultRouteTableAssociationRule(),
	NewAwsEc2TransitGatewayInvalidDefaultRouteTablePropagationRule(),
	NewAwsEc2TransitGatewayInvalidDNSSupportRule(),
	NewAwsEc2TransitGatewayVpcAttachmentInvalidDNSSupportRule(),
	NewAwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule(),
	NewAwsEcrLifecyclePolicyInvalidPolicyRule(),
	NewAwsEcrLifecyclePolicyInvalidRepositoryRule(),
	NewAwsEcrRepositoryInvalidNameRule(),
	NewAwsEcrRepositoryPolicyInvalidPolicyRule(),
	NewAwsEcrRepositoryPolicyInvalidRepositoryRule(),
	NewAwsEcsServiceInvalidLaunchTypeRule(),
	NewAwsEcsServiceInvalidPropagateTagsRule(),
	NewAwsEcsServiceInvalidSchedulingStrategyRule(),
	NewAwsEcsTaskDefinitionInvalidIpcModeRule(),
	NewAwsEcsTaskDefinitionInvalidNetworkModeRule(),
	NewAwsEcsTaskDefinitionInvalidPidModeRule(),
	NewAwsEfsFileSystemInvalidCreationTokenRule(),
	NewAwsEfsFileSystemInvalidKmsKeyIDRule(),
	NewAwsEfsFileSystemInvalidPerformanceModeRule(),
	NewAwsEfsFileSystemInvalidThroughputModeRule(),
	NewAwsEfsMountTargetInvalidFileSystemIDRule(),
	NewAwsEfsMountTargetInvalidIPAddressRule(),
	NewAwsEfsMountTargetInvalidSubnetIDRule(),
	NewAwsEksClusterInvalidNameRule(),
	NewAwsElasticBeanstalkApplicationInvalidDescriptionRule(),
	NewAwsElasticBeanstalkApplicationInvalidNameRule(),
	NewAwsElasticBeanstalkApplicationVersionInvalidApplicationRule(),
	NewAwsElasticBeanstalkApplicationVersionInvalidBucketRule(),
	NewAwsElasticBeanstalkApplicationVersionInvalidDescriptionRule(),
	NewAwsElasticBeanstalkApplicationVersionInvalidKeyRule(),
	NewAwsElasticBeanstalkApplicationVersionInvalidNameRule(),
	NewAwsElasticBeanstalkConfigurationTemplateInvalidApplicationRule(),
	NewAwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule(),
	NewAwsElasticBeanstalkConfigurationTemplateInvalidNameRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidApplicationRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidDescriptionRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidNameRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidTemplateNameRule(),
	NewAwsElasticBeanstalkEnvironmentInvalidVersionLabelRule(),
	NewAwsElastiCacheClusterInvalidAzModeRule(),
	NewAwsElasticsearchDomainInvalidDomainNameRule(),
	NewAwsElasticsearchDomainPolicyInvalidDomainNameRule(),
	NewAwsElastictranscoderPipelineInvalidAwsKmsKeyArnRule(),
	NewAwsElastictranscoderPipelineInvalidInputBucketRule(),
	NewAwsElastictranscoderPipelineInvalidNameRule(),
	NewAwsElastictranscoderPipelineInvalidOutputBucketRule(),
	NewAwsElastictranscoderPipelineInvalidRoleRule(),
	NewAwsElastictranscoderPresetInvalidContainerRule(),
	NewAwsElastictranscoderPresetInvalidDescriptionRule(),
	NewAwsElastictranscoderPresetInvalidNameRule(),
	NewAwsEmrClusterInvalidScaleDownBehaviorRule(),
	NewAwsFlowLogInvalidLogDestinationTypeRule(),
	NewAwsFlowLogInvalidTrafficTypeRule(),
	NewAwsFmsAdminAccountInvalidAccountIDRule(),
	NewAwsFsxLustreFileSystemInvalidWeeklyMaintenanceStartTimeRule(),
	NewAwsFsxWindowsFileSystemInvalidActiveDirectoryIDRule(),
	NewAwsFsxWindowsFileSystemInvalidDailyAutomaticBackupStartTimeRule(),
	NewAwsFsxWindowsFileSystemInvalidWeeklyMaintenanceStartTimeRule(),
	NewAwsGameliftAliasInvalidDescriptionRule(),
	NewAwsGameliftAliasInvalidNameRule(),
	NewAwsGameliftBuildInvalidNameRule(),
	NewAwsGameliftBuildInvalidOperatingSystemRule(),
	NewAwsGameliftBuildInvalidVersionRule(),
	NewAwsGameliftFleetInvalidBuildIDRule(),
	NewAwsGameliftFleetInvalidDescriptionRule(),
	NewAwsGameliftFleetInvalidEc2InstanceTypeRule(),
	NewAwsGameliftFleetInvalidNameRule(),
	NewAwsGameliftFleetInvalidNewGameSessionProtectionPolicyRule(),
	NewAwsGameliftGameSessionQueueInvalidNameRule(),
	NewAwsGlobalacceleratorAcceleratorInvalidIPAddressTypeRule(),
	NewAwsGlobalacceleratorAcceleratorInvalidNameRule(),
	NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckPathRule(),
	NewAwsGlobalacceleratorEndpointGroupInvalidHealthCheckProtocolRule(),
	NewAwsGlobalacceleratorEndpointGroupInvalidListenerArnRule(),
	NewAwsGlobalacceleratorListenerInvalidAcceleratorArnRule(),
	NewAwsGlobalacceleratorListenerInvalidClientAffinityRule(),
	NewAwsGlobalacceleratorListenerInvalidProtocolRule(),
	NewAwsGlueCatalogTableInvalidTableTypeRule(),
	NewAwsGlueCatalogTableInvalidViewExpandedTextRule(),
	NewAwsGlueCatalogTableInvalidViewOriginalTextRule(),
	NewAwsGlueConnectionInvalidConnectionTypeRule(),
	NewAwsGlueCrawlerInvalidSecurityConfigurationRule(),
	NewAwsGlueCrawlerInvalidTablePrefixRule(),
	NewAwsGlueTriggerInvalidTypeRule(),
	NewAwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule(),
	NewAwsGuarddutyInviteAccepterInvalidDetectorIDRule(),
	NewAwsGuarddutyIpsetInvalidDetectorIDRule(),
	NewAwsGuarddutyIpsetInvalidFormatRule(),
	NewAwsGuarddutyIpsetInvalidLocationRule(),
	NewAwsGuarddutyIpsetInvalidNameRule(),
	NewAwsGuarddutyMemberInvalidDetectorIDRule(),
	NewAwsGuarddutyMemberInvalidEmailRule(),
	NewAwsGuarddutyThreatintelsetInvalidDetectorIDRule(),
	NewAwsGuarddutyThreatintelsetInvalidFormatRule(),
	NewAwsGuarddutyThreatintelsetInvalidLocationRule(),
	NewAwsGuarddutyThreatintelsetInvalidNameRule(),
	NewAwsIAMAccessKeyInvalidStatusRule(),
	NewAwsIAMAccessKeyInvalidUserRule(),
	NewAwsIAMGroupInvalidNameRule(),
	NewAwsIAMGroupInvalidPathRule(),
	NewAwsIAMGroupMembershipInvalidGroupRule(),
	NewAwsIAMGroupPolicyAttachmentInvalidGroupRule(),
	NewAwsIAMGroupPolicyAttachmentInvalidPolicyArnRule(),
	NewAwsIAMGroupPolicyInvalidGroupRule(),
	NewAwsIAMGroupPolicyInvalidNameRule(),
	NewAwsIAMGroupPolicyInvalidPolicyRule(),
	NewAwsIAMInstanceProfileInvalidNameRule(),
	NewAwsIAMInstanceProfileInvalidPathRule(),
	NewAwsIAMInstanceProfileInvalidRoleRule(),
	NewAwsIAMOpenidConnectProviderInvalidURLRule(),
	NewAwsIAMPolicyAttachmentInvalidPolicyArnRule(),
	NewAwsIAMPolicyInvalidDescriptionRule(),
	NewAwsIAMPolicyInvalidNameRule(),
	NewAwsIAMPolicyInvalidPathRule(),
	NewAwsIAMPolicyInvalidPolicyRule(),
	NewAwsIAMRoleInvalidAssumeRolePolicyRule(),
	NewAwsIAMRoleInvalidDescriptionRule(),
	NewAwsIAMRoleInvalidNameRule(),
	NewAwsIAMRoleInvalidPathRule(),
	NewAwsIAMRoleInvalidPermissionsBoundaryRule(),
	NewAwsIAMRolePolicyAttachmentInvalidPolicyArnRule(),
	NewAwsIAMRolePolicyAttachmentInvalidRoleRule(),
	NewAwsIAMRolePolicyInvalidNameRule(),
	NewAwsIAMRolePolicyInvalidPolicyRule(),
	NewAwsIAMRolePolicyInvalidRoleRule(),
	NewAwsIAMSamlProviderInvalidNameRule(),
	NewAwsIAMSamlProviderInvalidSamlMetadataDocumentRule(),
	NewAwsIAMServerCertificateInvalidCertificateBodyRule(),
	NewAwsIAMServerCertificateInvalidCertificateChainRule(),
	NewAwsIAMServerCertificateInvalidNameRule(),
	NewAwsIAMServerCertificateInvalidPathRule(),
	NewAwsIAMServerCertificateInvalidPrivateKeyRule(),
	NewAwsIAMServiceLinkedRoleInvalidAwsServiceNameRule(),
	NewAwsIAMServiceLinkedRoleInvalidCustomSuffixRule(),
	NewAwsIAMServiceLinkedRoleInvalidDescriptionRule(),
	NewAwsIAMUserGroupMembershipInvalidUserRule(),
	NewAwsIAMUserInvalidNameRule(),
	NewAwsIAMUserInvalidPathRule(),
	NewAwsIAMUserInvalidPermissionsBoundaryRule(),
	NewAwsIAMUserLoginProfileInvalidUserRule(),
	NewAwsIAMUserPolicyAttachmentInvalidPolicyArnRule(),
	NewAwsIAMUserPolicyAttachmentInvalidUserRule(),
	NewAwsIAMUserPolicyInvalidNameRule(),
	NewAwsIAMUserPolicyInvalidPolicyRule(),
	NewAwsIAMUserPolicyInvalidUserRule(),
	NewAwsIAMUserSSHKeyInvalidEncodingRule(),
	NewAwsIAMUserSSHKeyInvalidPublicKeyRule(),
	NewAwsIAMUserSSHKeyInvalidStatusRule(),
	NewAwsIAMUserSSHKeyInvalidUsernameRule(),
	NewAwsInspectorAssessmentTargetInvalidNameRule(),
	NewAwsInspectorAssessmentTargetInvalidResourceGroupArnRule(),
	NewAwsInspectorAssessmentTemplateInvalidNameRule(),
	NewAwsInspectorAssessmentTemplateInvalidTargetArnRule(),
	NewAwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule(),
	NewAwsInstanceInvalidTenancyRule(),
	NewAwsInstanceInvalidTypeRule(),
	NewAwsIotPolicyAttachmentInvalidPolicyRule(),
	NewAwsIotPolicyInvalidNameRule(),
	NewAwsIotRoleAliasInvalidAliasRule(),
	NewAwsIotRoleAliasInvalidRoleArnRule(),
	NewAwsIotThingInvalidNameRule(),
	NewAwsIotThingInvalidThingTypeNameRule(),
	NewAwsIotThingPrincipalAttachmentInvalidThingRule(),
	NewAwsIotThingTypeInvalidNameRule(),
	NewAwsIotTopicRuleInvalidNameRule(),
	NewAwsKinesisAnalyticsApplicationInvalidCodeRule(),
	NewAwsKinesisAnalyticsApplicationInvalidDescriptionRule(),
	NewAwsKinesisAnalyticsApplicationInvalidNameRule(),
	NewAwsKinesisFirehoseDeliveryStreamInvalidNameRule(),
	NewAwsKinesisStreamInvalidEncryptionTypeRule(),
	NewAwsKinesisStreamInvalidKmsKeyIDRule(),
	NewAwsKinesisStreamInvalidNameRule(),
	NewAwsKmsAliasInvalidNameRule(),
	NewAwsKmsAliasInvalidTargetKeyIDRule(),
	NewAwsKmsCiphertextInvalidKeyIDRule(),
	NewAwsKmsExternalKeyInvalidDescriptionRule(),
	NewAwsKmsExternalKeyInvalidPolicyRule(),
	NewAwsKmsGrantInvalidGranteePrincipalRule(),
	NewAwsKmsGrantInvalidKeyIDRule(),
	NewAwsKmsGrantInvalidNameRule(),
	NewAwsKmsGrantInvalidRetiringPrincipalRule(),
	NewAwsKmsKeyInvalidDescriptionRule(),
	NewAwsKmsKeyInvalidKeyUsageRule(),
	NewAwsKmsKeyInvalidPolicyRule(),
	NewAwsLambdaAliasInvalidDescriptionRule(),
	NewAwsLambdaAliasInvalidFunctionNameRule(),
	NewAwsLambdaAliasInvalidFunctionVersionRule(),
	NewAwsLambdaEventSourceMappingInvalidEventSourceArnRule(),
	NewAwsLambdaEventSourceMappingInvalidFunctionNameRule(),
	NewAwsLambdaEventSourceMappingInvalidStartingPositionRule(),
	NewAwsLambdaFunctionInvalidDescriptionRule(),
	NewAwsLambdaFunctionInvalidFunctionNameRule(),
	NewAwsLambdaFunctionInvalidHandlerRule(),
	NewAwsLambdaFunctionInvalidKmsKeyArnRule(),
	NewAwsLambdaFunctionInvalidRoleRule(),
	NewAwsLambdaFunctionInvalidRuntimeRule(),
	NewAwsLambdaFunctionInvalidS3KeyRule(),
	NewAwsLambdaFunctionInvalidS3ObjectVersionRule(),
	NewAwsLambdaLayerVersionInvalidDescriptionRule(),
	NewAwsLambdaLayerVersionInvalidLayerNameRule(),
	NewAwsLambdaLayerVersionInvalidLicenseInfoRule(),
	NewAwsLambdaLayerVersionInvalidS3KeyRule(),
	NewAwsLambdaLayerVersionInvalidS3ObjectVersionRule(),
	NewAwsLambdaPermissionInvalidActionRule(),
	NewAwsLambdaPermissionInvalidEventSourceTokenRule(),
	NewAwsLambdaPermissionInvalidFunctionNameRule(),
	NewAwsLambdaPermissionInvalidPrincipalRule(),
	NewAwsLambdaPermissionInvalidQualifierRule(),
	NewAwsLambdaPermissionInvalidSourceAccountRule(),
	NewAwsLambdaPermissionInvalidSourceArnRule(),
	NewAwsLambdaPermissionInvalidStatementIDRule(),
	NewAwsLaunchConfigurationInvalidSpotPriceRule(),
	NewAwsLaunchConfigurationInvalidTypeRule(),
	NewAwsLaunchTemplateInvalidDescriptionRule(),
	NewAwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule(),
	NewAwsLaunchTemplateInvalidInstanceTypeRule(),
	NewAwsLaunchTemplateInvalidNameRule(),
	NewAwsLbInvalidIPAddressTypeRule(),
	NewAwsLbInvalidLoadBalancerTypeRule(),
	NewAwsLbListenerInvalidProtocolRule(),
	NewAwsLbTargetGroupInvalidProtocolRule(),
	NewAwsLbTargetGroupInvalidTargetTypeRule(),
	NewAwsLicensemanagerLicenseConfigurationInvalidLicenseCountingTypeRule(),
	NewAwsLightsailInstanceInvalidBlueprintIDRule(),
	NewAwsLightsailInstanceInvalidBundleIDRule(),
	NewAwsLightsailInstanceInvalidKeyPairNameRule(),
	NewAwsLightsailKeyPairInvalidNameRule(),
	NewAwsLightsailStaticIPAttachmentInvalidInstanceNameRule(),
	NewAwsLightsailStaticIPAttachmentInvalidStaticIPNameRule(),
	NewAwsLightsailStaticIPInvalidNameRule(),
	NewAwsMacieMemberAccountAssociationInvalidMemberAccountIDRule(),
	NewAwsMacieS3BucketAssociationInvalidBucketNameRule(),
	NewAwsMacieS3BucketAssociationInvalidMemberAccountIDRule(),
	NewAwsMacieS3BucketAssociationInvalidPrefixRule(),
	NewAwsMediaStoreContainerInvalidNameRule(),
	NewAwsMediaStoreContainerPolicyInvalidContainerNameRule(),
	NewAwsMqBrokerInvalidDeploymentModeRule(),
	NewAwsMskClusterInvalidClusterNameRule(),
	NewAwsMskClusterInvalidEnhancedMonitoringRule(),
	NewAwsMskClusterInvalidKafkaVersionRule(),
	NewAwsNetworkACLRuleInvalidRuleActionRule(),
	NewAwsOpsworksApplicationInvalidTypeRule(),
	NewAwsOpsworksInstanceInvalidArchitectureRule(),
	NewAwsOpsworksInstanceInvalidAutoScalingTypeRule(),
	NewAwsOpsworksInstanceInvalidRootDeviceTypeRule(),
	NewAwsOpsworksStackInvalidDefaultRootDeviceTypeRule(),
	NewAwsOrganizationsAccountInvalidEmailRule(),
	NewAwsOrganizationsAccountInvalidIAMUserAccessToBillingRule(),
	NewAwsOrganizationsAccountInvalidNameRule(),
	NewAwsOrganizationsAccountInvalidParentIDRule(),
	NewAwsOrganizationsAccountInvalidRoleNameRule(),
	NewAwsOrganizationsOrganizationInvalidFeatureSetRule(),
	NewAwsOrganizationsOrganizationalUnitInvalidNameRule(),
	NewAwsOrganizationsOrganizationalUnitInvalidParentIDRule(),
	NewAwsOrganizationsPolicyAttachmentInvalidPolicyIDRule(),
	NewAwsOrganizationsPolicyAttachmentInvalidTargetIDRule(),
	NewAwsOrganizationsPolicyInvalidContentRule(),
	NewAwsOrganizationsPolicyInvalidDescriptionRule(),
	NewAwsOrganizationsPolicyInvalidNameRule(),
	NewAwsOrganizationsPolicyInvalidTypeRule(),
	NewAwsPlacementGroupInvalidStrategyRule(),
	NewAwsQuicksightGroupInvalidAwsAccountIDRule(),
	NewAwsQuicksightGroupInvalidDescriptionRule(),
	NewAwsQuicksightGroupInvalidGroupNameRule(),
	NewAwsQuicksightGroupInvalidNamespaceRule(),
	NewAwsResourcegroupsGroupInvalidNameRule(),
	NewAwsRoute53DelegationSetInvalidReferenceNameRule(),
	NewAwsRoute53HealthCheckInvalidCloudwatchAlarmNameRule(),
	NewAwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule(),
	NewAwsRoute53HealthCheckInvalidFqdnRule(),
	NewAwsRoute53HealthCheckInvalidInsufficientDataHealthStatusRule(),
	NewAwsRoute53HealthCheckInvalidIPAddressRule(),
	NewAwsRoute53HealthCheckInvalidReferenceNameRule(),
	NewAwsRoute53HealthCheckInvalidResourcePathRule(),
	NewAwsRoute53HealthCheckInvalidSearchStringRule(),
	NewAwsRoute53HealthCheckInvalidTypeRule(),
	NewAwsRoute53QueryLogInvalidZoneIDRule(),
	NewAwsRoute53RecordInvalidHealthCheckIDRule(),
	NewAwsRoute53RecordInvalidNameRule(),
	NewAwsRoute53RecordInvalidSetIdentifierRule(),
	NewAwsRoute53RecordInvalidTypeRule(),
	NewAwsRoute53RecordInvalidZoneIDRule(),
	NewAwsRoute53ResolverEndpointInvalidDirectionRule(),
	NewAwsRoute53ResolverRuleAssociationInvalidResolverRuleIDRule(),
	NewAwsRoute53ResolverRuleAssociationInvalidVpcIDRule(),
	NewAwsRoute53ResolverRuleInvalidDomainNameRule(),
	NewAwsRoute53ResolverRuleInvalidResolverEndpointIDRule(),
	NewAwsRoute53ResolverRuleInvalidRuleTypeRule(),
	NewAwsRoute53ZoneAssociationInvalidVpcIDRule(),
	NewAwsRoute53ZoneAssociationInvalidVpcRegionRule(),
	NewAwsRoute53ZoneAssociationInvalidZoneIDRule(),
	NewAwsRoute53ZoneInvalidCommentRule(),
	NewAwsRoute53ZoneInvalidDelegationSetIDRule(),
	NewAwsRoute53ZoneInvalidNameRule(),
	NewAwsS3BucketInvalidAccelerationStatusRule(),
	NewAwsS3BucketInvalidRequestPayerRule(),
	NewAwsS3BucketInventoryInvalidIncludedObjectVersionsRule(),
	NewAwsS3BucketObjectInvalidACLRule(),
	NewAwsS3BucketObjectInvalidServerSideEncryptionRule(),
	NewAwsS3BucketObjectInvalidStorageClassRule(),
	NewAwsSagemakerEndpointConfigurationInvalidKmsKeyArnRule(),
	NewAwsSagemakerEndpointConfigurationInvalidNameRule(),
	NewAwsSagemakerEndpointInvalidEndpointConfigNameRule(),
	NewAwsSagemakerEndpointInvalidNameRule(),
	NewAwsSagemakerModelInvalidExecutionRoleArnRule(),
	NewAwsSagemakerModelInvalidNameRule(),
	NewAwsSagemakerNotebookInstanceInvalidInstanceTypeRule(),
	NewAwsSagemakerNotebookInstanceInvalidKmsKeyIDRule(),
	NewAwsSagemakerNotebookInstanceInvalidLifecycleConfigNameRule(),
	NewAwsSagemakerNotebookInstanceInvalidNameRule(),
	NewAwsSagemakerNotebookInstanceInvalidRoleArnRule(),
	NewAwsSagemakerNotebookInstanceInvalidSubnetIDRule(),
	NewAwsSagemakerNotebookInstanceLifecycleConfigurationInvalidNameRule(),
	NewAwsSecretsmanagerSecretInvalidDescriptionRule(),
	NewAwsSecretsmanagerSecretInvalidKmsKeyIDRule(),
	NewAwsSecretsmanagerSecretInvalidNameRule(),
	NewAwsSecretsmanagerSecretInvalidPolicyRule(),
	NewAwsSecretsmanagerSecretInvalidRotationLambdaArnRule(),
	NewAwsSecretsmanagerSecretVersionInvalidSecretIDRule(),
	NewAwsSecretsmanagerSecretVersionInvalidSecretStringRule(),
	NewAwsSecurityhubProductSubscriptionInvalidProductArnRule(),
	NewAwsSecurityhubStandardsSubscriptionInvalidStandardsArnRule(),
	NewAwsServiceDiscoveryHTTPNamespaceInvalidDescriptionRule(),
	NewAwsServiceDiscoveryHTTPNamespaceInvalidNameRule(),
	NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidDescriptionRule(),
	NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidNameRule(),
	NewAwsServiceDiscoveryPrivateDNSNamespaceInvalidVpcRule(),
	NewAwsServiceDiscoveryPublicDNSNamespaceInvalidDescriptionRule(),
	NewAwsServiceDiscoveryPublicDNSNamespaceInvalidNameRule(),
	NewAwsServiceDiscoveryServiceInvalidDescriptionRule(),
	NewAwsServicecatalogPortfolioInvalidDescriptionRule(),
	NewAwsServicecatalogPortfolioInvalidNameRule(),
	NewAwsServicecatalogPortfolioInvalidProviderNameRule(),
	NewAwsServicequotasServiceQuotaInvalidQuotaCodeRule(),
	NewAwsServicequotasServiceQuotaInvalidServiceCodeRule(),
	NewAwsSesDomainMailFromInvalidBehaviorOnMxFailureRule(),
	NewAwsSesIdentityNotificationTopicInvalidNotificationTypeRule(),
	NewAwsSesIdentityPolicyInvalidNameRule(),
	NewAwsSesReceiptFilterInvalidPolicyRule(),
	NewAwsSesReceiptRuleInvalidTLSPolicyRule(),
	NewAwsSfnActivityInvalidNameRule(),
	NewAwsSfnStateMachineInvalidDefinitionRule(),
	NewAwsSfnStateMachineInvalidNameRule(),
	NewAwsSfnStateMachineInvalidRoleArnRule(),
	NewAwsShieldProtectionInvalidNameRule(),
	NewAwsShieldProtectionInvalidResourceArnRule(),
	NewAwsSpotFleetRequestInvalidAllocationStrategyRule(),
	NewAwsSpotFleetRequestInvalidFleetTypeRule(),
	NewAwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule(),
	NewAwsSpotInstanceRequestInvalidInstanceInterruptionBehaviourRule(),
	NewAwsSsmActivationInvalidDescriptionRule(),
	NewAwsSsmActivationInvalidIAMRoleRule(),
	NewAwsSsmActivationInvalidNameRule(),
	NewAwsSsmAssociationInvalidAssociationNameRule(),
	NewAwsSsmAssociationInvalidComplianceSeverityRule(),
	NewAwsSsmAssociationInvalidDocumentVersionRule(),
	NewAwsSsmAssociationInvalidInstanceIDRule(),
	NewAwsSsmAssociationInvalidMaxConcurrencyRule(),
	NewAwsSsmAssociationInvalidMaxErrorsRule(),
	NewAwsSsmAssociationInvalidNameRule(),
	NewAwsSsmAssociationInvalidScheduleExpressionRule(),
	NewAwsSsmDocumentInvalidDocumentFormatRule(),
	NewAwsSsmDocumentInvalidDocumentTypeRule(),
	NewAwsSsmDocumentInvalidNameRule(),
	NewAwsSsmMaintenanceWindowInvalidNameRule(),
	NewAwsSsmMaintenanceWindowInvalidScheduleRule(),
	NewAwsSsmMaintenanceWindowTargetInvalidDescriptionRule(),
	NewAwsSsmMaintenanceWindowTargetInvalidNameRule(),
	NewAwsSsmMaintenanceWindowTargetInvalidOwnerInformationRule(),
	NewAwsSsmMaintenanceWindowTargetInvalidResourceTypeRule(),
	NewAwsSsmMaintenanceWindowTargetInvalidWindowIDRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidDescriptionRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidMaxErrorsRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidNameRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidTaskArnRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidTaskTypeRule(),
	NewAwsSsmMaintenanceWindowTaskInvalidWindowIDRule(),
	NewAwsSsmParameterInvalidAllowedPatternRule(),
	NewAwsSsmParameterInvalidDescriptionRule(),
	NewAwsSsmParameterInvalidKeyIDRule(),
	NewAwsSsmParameterInvalidNameRule(),
	NewAwsSsmParameterInvalidTierRule(),
	NewAwsSsmParameterInvalidTypeRule(),
	NewAwsSsmPatchBaselineInvalidApprovedPatchesComplianceLevelRule(),
	NewAwsSsmPatchBaselineInvalidDescriptionRule(),
	NewAwsSsmPatchBaselineInvalidNameRule(),
	NewAwsSsmPatchBaselineInvalidOperatingSystemRule(),
	NewAwsSsmPatchGroupInvalidBaselineIDRule(),
	NewAwsSsmPatchGroupInvalidPatchGroupRule(),
	NewAwsSsmResourceDataSyncInvalidNameRule(),
	NewAwsStoragegatewayCacheInvalidDiskIDRule(),
	NewAwsStoragegatewayCacheInvalidGatewayArnRule(),
	NewAwsStoragegatewayCachedIscsiVolumeInvalidGatewayArnRule(),
	NewAwsStoragegatewayCachedIscsiVolumeInvalidNetworkInterfaceIDRule(),
	NewAwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule(),
	NewAwsStoragegatewayCachedIscsiVolumeInvalidSourceVolumeArnRule(),
	NewAwsStoragegatewayCachedIscsiVolumeInvalidTargetNameRule(),
	NewAwsStoragegatewayGatewayInvalidActivationKeyRule(),
	NewAwsStoragegatewayGatewayInvalidGatewayNameRule(),
	NewAwsStoragegatewayGatewayInvalidGatewayTimezoneRule(),
	NewAwsStoragegatewayGatewayInvalidGatewayTypeRule(),
	NewAwsStoragegatewayGatewayInvalidMediumChangerTypeRule(),
	NewAwsStoragegatewayGatewayInvalidSmbGuestPasswordRule(),
	NewAwsStoragegatewayGatewayInvalidTapeDriveTypeRule(),
	NewAwsStoragegatewayNfsFileShareInvalidDefaultStorageClassRule(),
	NewAwsStoragegatewayNfsFileShareInvalidGatewayArnRule(),
	NewAwsStoragegatewayNfsFileShareInvalidKmsKeyArnRule(),
	NewAwsStoragegatewayNfsFileShareInvalidLocationArnRule(),
	NewAwsStoragegatewayNfsFileShareInvalidObjectACLRule(),
	NewAwsStoragegatewayNfsFileShareInvalidRoleArnRule(),
	NewAwsStoragegatewayNfsFileShareInvalidSquashRule(),
	NewAwsStoragegatewaySmbFileShareInvalidAuthenticationRule(),
	NewAwsStoragegatewaySmbFileShareInvalidDefaultStorageClassRule(),
	NewAwsStoragegatewaySmbFileShareInvalidGatewayArnRule(),
	NewAwsStoragegatewaySmbFileShareInvalidKmsKeyArnRule(),
	NewAwsStoragegatewaySmbFileShareInvalidLocationArnRule(),
	NewAwsStoragegatewaySmbFileShareInvalidObjectACLRule(),
	NewAwsStoragegatewaySmbFileShareInvalidRoleArnRule(),
	NewAwsStoragegatewayUploadBufferInvalidDiskIDRule(),
	NewAwsStoragegatewayUploadBufferInvalidGatewayArnRule(),
	NewAwsStoragegatewayWorkingStorageInvalidDiskIDRule(),
	NewAwsStoragegatewayWorkingStorageInvalidGatewayArnRule(),
	NewAwsSwfDomainInvalidDescriptionRule(),
	NewAwsSwfDomainInvalidNameRule(),
	NewAwsSwfDomainInvalidWorkflowExecutionRetentionPeriodInDaysRule(),
	NewAwsTransferServerInvalidEndpointTypeRule(),
	NewAwsTransferServerInvalidIdentityProviderTypeRule(),
	NewAwsTransferServerInvalidInvocationRoleRule(),
	NewAwsTransferServerInvalidLoggingRoleRule(),
	NewAwsTransferServerInvalidURLRule(),
	NewAwsTransferSSHKeyInvalidBodyRule(),
	NewAwsTransferSSHKeyInvalidServerIDRule(),
	NewAwsTransferSSHKeyInvalidUserNameRule(),
	NewAwsTransferUserInvalidHomeDirectoryRule(),
	NewAwsTransferUserInvalidPolicyRule(),
	NewAwsTransferUserInvalidRoleRule(),
	NewAwsTransferUserInvalidServerIDRule(),
	NewAwsTransferUserInvalidUserNameRule(),
	NewAwsVpcEndpointInvalidVpcEndpointTypeRule(),
	NewAwsVpcInvalidInstanceTenancyRule(),
	NewAwsWafByteMatchSetInvalidNameRule(),
	NewAwsWafGeoMatchSetInvalidNameRule(),
	NewAwsWafIpsetInvalidNameRule(),
	NewAwsWafRateBasedRuleInvalidMetricNameRule(),
	NewAwsWafRateBasedRuleInvalidNameRule(),
	NewAwsWafRateBasedRuleInvalidRateKeyRule(),
	NewAwsWafRegexMatchSetInvalidNameRule(),
	NewAwsWafRegexPatternSetInvalidNameRule(),
	NewAwsWafRuleGroupInvalidMetricNameRule(),
	NewAwsWafRuleGroupInvalidNameRule(),
	NewAwsWafRuleInvalidMetricNameRule(),
	NewAwsWafRuleInvalidNameRule(),
	NewAwsWafSizeConstraintSetInvalidNameRule(),
	NewAwsWafSQLInjectionMatchSetInvalidNameRule(),
	NewAwsWafWebACLInvalidMetricNameRule(),
	NewAwsWafWebACLInvalidNameRule(),
	NewAwsWafXSSMatchSetInvalidNameRule(),
	NewAwsWafregionalByteMatchSetInvalidNameRule(),
	NewAwsWafregionalGeoMatchSetInvalidNameRule(),
	NewAwsWafregionalIpsetInvalidNameRule(),
	NewAwsWafregionalRateBasedRuleInvalidMetricNameRule(),
	NewAwsWafregionalRateBasedRuleInvalidNameRule(),
	NewAwsWafregionalRateBasedRuleInvalidRateKeyRule(),
	NewAwsWafregionalRegexMatchSetInvalidNameRule(),
	NewAwsWafregionalRegexPatternSetInvalidNameRule(),
	NewAwsWafregionalRuleGroupInvalidMetricNameRule(),
	NewAwsWafregionalRuleGroupInvalidNameRule(),
	NewAwsWafregionalRuleInvalidMetricNameRule(),
	NewAwsWafregionalRuleInvalidNameRule(),
	NewAwsWafregionalSizeConstraintSetInvalidNameRule(),
	NewAwsWafregionalSQLInjectionMatchSetInvalidNameRule(),
	NewAwsWafregionalWebACLAssociationInvalidResourceArnRule(),
	NewAwsWafregionalWebACLAssociationInvalidWebACLIDRule(),
	NewAwsWafregionalWebACLInvalidMetricNameRule(),
	NewAwsWafregionalWebACLInvalidNameRule(),
	NewAwsWafregionalXSSMatchSetInvalidNameRule(),
	NewAwsWorklinkFleetInvalidAuditStreamArnRule(),
	NewAwsWorklinkFleetInvalidDeviceCaCertificateRule(),
	NewAwsWorklinkFleetInvalidDisplayNameRule(),
	NewAwsWorklinkFleetInvalidNameRule(),
	NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidCertificateRule(),
	NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule(),
	NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidFleetArnRule(),
	NewAwsXraySamplingRuleInvalidHostRule(),
	NewAwsXraySamplingRuleInvalidHTTPMethodRule(),
	NewAwsXraySamplingRuleInvalidResourceArnRule(),
	NewAwsXraySamplingRuleInvalidRuleNameRule(),
	NewAwsXraySamplingRuleInvalidServiceNameRule(),
	NewAwsXraySamplingRuleInvalidServiceTypeRule(),
	NewAwsXraySamplingRuleInvalidURLPathRule(),
}