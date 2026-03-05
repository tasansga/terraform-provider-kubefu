package generated

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Versions struct {
	K8sVersion         string
	FluxVersion        string
	CertManagerVersion string
	PrometheusOperatorVersion string
	GatewayAPIVersion         string
	ExternalSecretsVersion    string
}

func DataSources(versions Versions) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource, 281)
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoChallengeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoOrderV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateRequestV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoCertificateV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoClusterIssuerV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoIssuerV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Alpha2()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoIssuerV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Alpha3()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoIssuerV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Beta1()
		version := versions.versionFor("cert_manager")
		if version != "" && !dataSourceCertManagerCertManagerIoIssuerV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_push_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_push_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoExternalSecretV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_push_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_push_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoSecretStoreV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_acr_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_acr_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_cloudsmith_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_cloudsmith_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_cluster_generator_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_cluster_generator_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_ecr_authorization_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_ecr_authorization_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_fake_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_fake_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_gcr_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_gcr_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_generator_state_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_generator_state_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_github_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_github_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_grafana_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_grafana_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_mfa_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_mfa_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_password_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_password_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_quay_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_quay_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_ssh_key_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_ssh_key_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_sts_session_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_sts_session_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_uuid_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_uuid_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_vault_dynamic_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_vault_dynamic_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1()
		version := versions.versionFor("external_secrets")
		if version != "" && !dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_webhook_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_external_secrets_generators_external_secrets_io_webhook_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta1",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta1"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta2",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImagePolicyV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta1",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta2",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta3",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta3"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta3",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta3"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoReceiverV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_extensions_fluxcd_io_artifact_generator_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_extensions_fluxcd_io_artifact_generator_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoBucketV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_external_artifact_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_external_artifact_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmChartV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2()
		version := versions.versionFor("flux")
		if version != "" && !dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1alpha3",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1alpha3"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1()
		version := versions.versionFor("gateway_api")
		if version != "" && !dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_external_admission_hook_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_external_admission_hook_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_initializer_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_initializer_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1",
				"k8s",
				strings.Join(dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1"] = ds
	}
	{
		ds := dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sApiregistrationK8sIoAPIServiceV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sApiregistrationK8sIoAPIServiceV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_apiregistration_k8s_io_api_service_v1",
				"k8s",
				strings.Join(dataSourceK8sApiregistrationK8sIoAPIServiceV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_apiregistration_k8s_io_api_service_v1"] = ds
	}
	{
		ds := dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_apiregistration_k8s_io_api_service_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_apiregistration_k8s_io_api_service_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_auditregistration_k8s_io_audit_sink_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_auditregistration_k8s_io_audit_sink_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoTokenReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthenticationK8sIoTokenReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authentication_k8s_io_token_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoTokenReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authentication_k8s_io_token_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authentication_k8s_io_token_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authentication_k8s_io_token_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_authorization_k8s_io_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_authorization_k8s_io_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoordinationK8sIoLeaseV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_v1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_coordination_k8s_io_lease_v1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoordinationK8sIoLeaseV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_coordination_k8s_io_lease_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoreBindingV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreBindingV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreBindingV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreComponentStatusV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreComponentStatusV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_component_status_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreComponentStatusV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_component_status_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreConfigMapV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreConfigMapV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_config_map_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreConfigMapV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_config_map_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreEndpointsV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreEndpointsV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_endpoints_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreEndpointsV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_endpoints_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreEventV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreEventV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_event_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreEventV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_event_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreLimitRangeV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreLimitRangeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_limit_range_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreLimitRangeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_limit_range_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreNamespaceV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreNamespaceV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_namespace_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreNamespaceV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_namespace_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreNodeV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreNodeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_node_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreNodeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_node_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePersistentVolumeClaimV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCorePersistentVolumeClaimV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_persistent_volume_claim_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePersistentVolumeClaimV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_persistent_volume_claim_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePersistentVolumeV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCorePersistentVolumeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_persistent_volume_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePersistentVolumeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_persistent_volume_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePodTemplateV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCorePodTemplateV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_pod_template_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePodTemplateV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_pod_template_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePodV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCorePodV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_pod_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePodV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_pod_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreReplicationControllerV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreReplicationControllerV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_replication_controller_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreReplicationControllerV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_replication_controller_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreResourceQuotaV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreResourceQuotaV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_resource_quota_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreResourceQuotaV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_resource_quota_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreSecretV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreSecretV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_secret_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreSecretV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_secret_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreServiceAccountV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreServiceAccountV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_service_account_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreServiceAccountV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_service_account_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreServiceV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sCoreServiceV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_core_service_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreServiceV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_core_service_v1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sDiscoveryK8sIoEndpointSliceV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sEventsK8sIoEventV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sEventsK8sIoEventV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_events_k8s_io_event_v1",
				"k8s",
				strings.Join(dataSourceK8sEventsK8sIoEventV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_events_k8s_io_event_v1"] = ds
	}
	{
		ds := dataSourceK8sEventsK8sIoEventV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sEventsK8sIoEventV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_events_k8s_io_event_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sEventsK8sIoEventV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_events_k8s_io_event_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3"] = ds
	}
	{
		ds := dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_internal_apiserver_k8s_io_storage_version_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_internal_apiserver_k8s_io_storage_version_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_cluster_cidr_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_cluster_cidr_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIngressClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_class_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ingress_class_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIngressClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ingress_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIngressV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ingress_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIngressV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ingress_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIPAddressV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoIPAddressV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoNetworkPolicyV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoNetworkPolicyV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_network_policy_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoNetworkPolicyV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_network_policy_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoServiceCIDRV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNodeK8sIoRuntimeClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sNodeK8sIoRuntimeClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoDeviceClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoDeviceClassV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoDeviceClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoDeviceClassV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_taint_rule_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_device_taint_rule_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_parameters_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_parameters_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClaimV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_parameters_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_parameters_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClassV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceClassV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceSliceV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Alpha2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceSliceV1Alpha2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Alpha2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Alpha3()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceSliceV1Alpha3IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Alpha3CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceSliceV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Beta2()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sResourceK8sIoResourceSliceV1Beta2IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Beta2CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sSchedulingK8sIoPriorityClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_workload_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_scheduling_k8s_io_workload_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sSettingsK8sIoPodPresetV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sSettingsK8sIoPodPresetV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_settings_k8s_io_pod_preset_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSettingsK8sIoPodPresetV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_settings_k8s_io_pod_preset_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIDriverV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSIDriverV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_driver_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIDriverV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_driver_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIDriverV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSIDriverV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_driver_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIDriverV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_driver_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSINodeV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSINodeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_node_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSINodeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_node_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSINodeV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSINodeV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_node_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSINodeV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_node_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSIStorageCapacityV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoStorageClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoStorageClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_storage_class_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoStorageClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_storage_class_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoStorageClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoStorageClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_storage_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoStorageClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_storage_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttachmentV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttributesClassV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1()
		version := versions.versionFor("k8s")
		if version != "" && !dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1beta1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_config_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_config_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_pod_monitor_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_pod_monitor_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComProbeV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComProbeV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_probe_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComProbeV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_probe_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_agent_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_agent_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_rule_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_rule_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_scrape_config_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_scrape_config_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_service_monitor_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_service_monitor_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1()
		version := versions.versionFor("prometheus_operator")
		if version != "" && !dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1IsCompatibleWith(version) {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured version %q may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_thanos_ruler_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1CompatibleVersions, ", "),
				version,
			)
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_thanos_ruler_v1"] = ds
	}
	return result
}

func (v Versions) versionFor(provider string) string {
	switch provider {
	case "k8s":
		return v.K8sVersion
	case "flux":
		return v.FluxVersion
	case "cert_manager", "cert-manager":
		return v.CertManagerVersion
	case "prometheus_operator", "prometheus-operator", "prometheus":
		return v.PrometheusOperatorVersion
	case "gateway_api", "gateway-api", "gateway":
		return v.GatewayAPIVersion
	case "external_secrets", "external-secrets", "externalsecrets":
		return v.ExternalSecretsVersion
	default:
		return ""
	}
}
