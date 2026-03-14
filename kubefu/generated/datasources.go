package generated

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	versionpkg "github.com/tasansga/terraform-provider-kubefu/resourcegen/version"
)

type Versions struct {
	K8sVersions                []string
	FluxVersions               []string
	CertManagerVersions        []string
	PrometheusOperatorVersions []string
	GatewayAPIVersions         []string
	ExternalSecretsVersions    []string
	KustomizeVersions          []string
	KarpenterAWSVersions       []string
	KarpenterCoreVersions      []string
}

func DataSources(versions Versions) map[string]*schema.Resource {
	result := make(map[string]*schema.Resource, 298)
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoChallengeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_challenge_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoChallengeV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_challenge_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoOrderV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_acme_cert_manager_io_order_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerAcmeCertManagerIoOrderV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_acme_cert_manager_io_order_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateRequestV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_request_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateRequestV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_request_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoCertificateV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoCertificateV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_certificate_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoCertificateV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_certificate_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoClusterIssuerV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_cluster_issuer_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoClusterIssuerV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_cluster_issuer_v1beta1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoIssuerV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Alpha2()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoIssuerV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1alpha2",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1alpha2"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Alpha3()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoIssuerV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1alpha3",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1alpha3"] = ds
	}
	{
		ds := dataSourceCertManagerCertManagerIoIssuerV1Beta1()
		configured := versions.versionFor("cert_manager")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceCertManagerCertManagerIoIssuerV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_cert_manager_cert_manager_io_issuer_v1beta1",
				"cert_manager",
				strings.Join(dataSourceCertManagerCertManagerIoIssuerV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_cert_manager_cert_manager_io_issuer_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterExternalSecretV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_external_secret_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_push_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterPushSecretV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_push_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoClusterSecretStoreV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_cluster_secret_store_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoExternalSecretV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_external_secret_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoExternalSecretV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_external_secret_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_push_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoPushSecretV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_push_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoSecretStoreV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_external_secrets_io_secret_store_v1beta1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsExternalSecretsIoSecretStoreV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_external_secrets_io_secret_store_v1beta1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_acr_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoACRAccessTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_acr_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_cloudsmith_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoCloudsmithAccessTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_cloudsmith_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_cluster_generator_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoClusterGeneratorV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_cluster_generator_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_ecr_authorization_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoECRAuthorizationTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_ecr_authorization_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_fake_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoFakeV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_fake_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_gcr_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGCRAccessTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_gcr_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_generator_state_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGeneratorStateV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_generator_state_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_github_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGithubAccessTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_github_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_grafana_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoGrafanaV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_grafana_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_mfa_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoMFAV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_mfa_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_password_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoPasswordV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_password_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_quay_access_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoQuayAccessTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_quay_access_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_ssh_key_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoSSHKeyV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_ssh_key_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_sts_session_token_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoSTSSessionTokenV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_sts_session_token_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_uuid_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoUUIDV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_uuid_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_vault_dynamic_secret_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_vault_dynamic_secret_v1alpha1"] = ds
	}
	{
		ds := dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1()
		configured := versions.versionFor("external_secrets")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_external_secrets_generators_external_secrets_io_webhook_v1alpha1",
				"external_secrets",
				strings.Join(dataSourceExternalSecretsGeneratorsExternalSecretsIoWebhookV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_external_secrets_generators_external_secrets_io_webhook_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta1",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta1"] = ds
	}
	{
		ds := dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta2",
				"flux",
				strings.Join(dataSourceFluxHelmToolkitFluxcdIoHelmReleaseV2Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_helm_toolkit_fluxcd_io_helm_release_v2beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImagePolicyV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImagePolicyV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_policy_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageRepositoryV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1alpha2"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta1",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta2",
				"flux",
				strings.Join(dataSourceFluxImageToolkitFluxcdIoImageUpdateAutomationV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_image_toolkit_fluxcd_io_image_update_automation_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta1",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta2",
				"flux",
				strings.Join(dataSourceFluxKustomizeToolkitFluxcdIoKustomizationV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_kustomize_toolkit_fluxcd_io_kustomization_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta3",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoAlertV1Beta3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_alert_v1beta3"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta3",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoProviderV1Beta3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_provider_v1beta3"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoReceiverV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta1",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta2",
				"flux",
				strings.Join(dataSourceFluxNotificationToolkitFluxcdIoReceiverV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_notification_toolkit_fluxcd_io_receiver_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_extensions_fluxcd_io_artifact_generator_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceExtensionsFluxcdIoArtifactGeneratorV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_extensions_fluxcd_io_artifact_generator_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoBucketV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoBucketV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_bucket_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_external_artifact_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoExternalArtifactV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_external_artifact_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoGitRepositoryV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_git_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmChartV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmChartV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_chart_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoHelmRepositoryV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_helm_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1"] = ds
	}
	{
		ds := dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2()
		configured := versions.versionFor("flux")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1beta2",
				"flux",
				strings.Join(dataSourceFluxSourceToolkitFluxcdIoOCIRepositoryV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_flux_source_toolkit_fluxcd_io_oci_repository_v1beta2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1alpha3",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoBackendTLSPolicyV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_backend_tls_policy_v1alpha3"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_class_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGatewayV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_gateway_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoGRPCRouteV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_grpc_route_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoHTTPRouteV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_http_route_v1beta1"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1alpha2",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1alpha2"] = ds
	}
	{
		ds := dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1()
		configured := versions.versionFor("gateway_api")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1beta1",
				"gateway_api",
				strings.Join(dataSourceGatewayApiGatewayNetworkingK8sIoReferenceGrantV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_gateway_api_gateway_networking_k8s_io_reference_grant_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_external_admission_hook_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoExternalAdmissionHookConfigurationV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_external_admission_hook_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_initializer_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoInitializerConfigurationV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_initializer_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyBindingV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingAdmissionPolicyV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_admission_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoMutatingWebhookConfigurationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyBindingV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingAdmissionPolicyV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_admission_policy_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAdmissionregistrationK8sIoValidatingWebhookConfigurationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1",
				"k8s",
				strings.Join(dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1"] = ds
	}
	{
		ds := dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sApiextensionsK8sIoCustomResourceDefinitionV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_apiextensions_k8s_io_custom_resource_definition_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sApiregistrationK8sIoAPIServiceV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sApiregistrationK8sIoAPIServiceV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_apiregistration_k8s_io_api_service_v1",
				"k8s",
				strings.Join(dataSourceK8sApiregistrationK8sIoAPIServiceV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_apiregistration_k8s_io_api_service_v1"] = ds
	}
	{
		ds := dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_apiregistration_k8s_io_api_service_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sApiregistrationK8sIoAPIServiceV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_apiregistration_k8s_io_api_service_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_auditregistration_k8s_io_audit_sink_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAuditregistrationK8sIoAuditSinkV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_auditregistration_k8s_io_audit_sink_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authentication_k8s_io_self_subject_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoSelfSubjectReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authentication_k8s_io_self_subject_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoTokenReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthenticationK8sIoTokenReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authentication_k8s_io_token_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoTokenReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authentication_k8s_io_token_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authentication_k8s_io_token_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthenticationK8sIoTokenReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authentication_k8s_io_token_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoLocalSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_local_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSelfSubjectRulesReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_self_subject_rules_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_subject_access_review_v1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_subject_access_review_v1"] = ds
	}
	{
		ds := dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_authorization_k8s_io_subject_access_review_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sAuthorizationK8sIoSubjectAccessReviewV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_authorization_k8s_io_subject_access_review_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoCertificateSigningRequestV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_certificate_signing_request_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoClusterTrustBundleV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_cluster_trust_bundle_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCertificatesK8sIoPodCertificateRequestV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_certificates_k8s_io_pod_certificate_request_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_candidate_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseCandidateV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_coordination_k8s_io_lease_candidate_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoordinationK8sIoLeaseV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_v1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_coordination_k8s_io_lease_v1"] = ds
	}
	{
		ds := dataSourceK8sCoordinationK8sIoLeaseV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoordinationK8sIoLeaseV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_coordination_k8s_io_lease_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sCoordinationK8sIoLeaseV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_coordination_k8s_io_lease_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sCoreBindingV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreBindingV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreBindingV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreComponentStatusV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreComponentStatusV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_component_status_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreComponentStatusV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_component_status_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreConfigMapV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreConfigMapV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_config_map_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreConfigMapV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_config_map_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreEndpointsV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreEndpointsV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_endpoints_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreEndpointsV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_endpoints_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreEventV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreEventV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_event_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreEventV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_event_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreLimitRangeV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreLimitRangeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_limit_range_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreLimitRangeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_limit_range_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreNamespaceV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreNamespaceV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_namespace_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreNamespaceV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_namespace_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreNodeV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreNodeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_node_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreNodeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_node_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePersistentVolumeClaimV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCorePersistentVolumeClaimV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_persistent_volume_claim_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePersistentVolumeClaimV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_persistent_volume_claim_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePersistentVolumeV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCorePersistentVolumeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_persistent_volume_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePersistentVolumeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_persistent_volume_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePodTemplateV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCorePodTemplateV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_pod_template_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePodTemplateV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_pod_template_v1"] = ds
	}
	{
		ds := dataSourceK8sCorePodV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCorePodV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_pod_v1",
				"k8s",
				strings.Join(dataSourceK8sCorePodV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_pod_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreReplicationControllerV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreReplicationControllerV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_replication_controller_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreReplicationControllerV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_replication_controller_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreResourceQuotaV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreResourceQuotaV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_resource_quota_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreResourceQuotaV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_resource_quota_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreSecretV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreSecretV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_secret_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreSecretV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_secret_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreServiceAccountV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreServiceAccountV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_service_account_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreServiceAccountV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_service_account_v1"] = ds
	}
	{
		ds := dataSourceK8sCoreServiceV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sCoreServiceV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_core_service_v1",
				"k8s",
				strings.Join(dataSourceK8sCoreServiceV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_core_service_v1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sDiscoveryK8sIoEndpointSliceV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_discovery_k8s_io_endpoint_slice_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sDiscoveryK8sIoEndpointSliceV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_discovery_k8s_io_endpoint_slice_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sEventsK8sIoEventV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sEventsK8sIoEventV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_events_k8s_io_event_v1",
				"k8s",
				strings.Join(dataSourceK8sEventsK8sIoEventV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_events_k8s_io_event_v1"] = ds
	}
	{
		ds := dataSourceK8sEventsK8sIoEventV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sEventsK8sIoEventV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_events_k8s_io_event_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sEventsK8sIoEventV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_events_k8s_io_event_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoFlowSchemaV1Beta3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta3"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3",
				"k8s",
				strings.Join(dataSourceK8sFlowcontrolApiserverK8sIoPriorityLevelConfigurationV1Beta3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta3"] = ds
	}
	{
		ds := dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_internal_apiserver_k8s_io_storage_version_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sInternalApiserverK8sIoStorageVersionV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_internal_apiserver_k8s_io_storage_version_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_cluster_cidr_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoClusterCIDRV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_cluster_cidr_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIngressClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_class_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ingress_class_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIngressClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ingress_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIngressV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ingress_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIngressV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIngressV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ingress_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIngressV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ingress_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIPAddressV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoIPAddressV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoIPAddressV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_ip_address_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoIPAddressV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_ip_address_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoNetworkPolicyV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoNetworkPolicyV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_network_policy_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoNetworkPolicyV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_network_policy_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoServiceCIDRV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_networking_k8s_io_service_cidr_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNetworkingK8sIoServiceCIDRV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_networking_k8s_io_service_cidr_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNodeK8sIoRuntimeClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sNodeK8sIoRuntimeClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sNodeK8sIoRuntimeClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_node_k8s_io_runtime_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sNodeK8sIoRuntimeClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_node_k8s_io_runtime_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleBindingV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoClusterRoleV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_cluster_role_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleBindingV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_binding_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_rbac_authorization_k8s_io_role_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sRbacAuthorizationK8sIoRoleV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_rbac_authorization_k8s_io_role_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoDeviceClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoDeviceClassV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoDeviceClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceClassV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoDeviceClassV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_class_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceClassV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_device_class_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_device_taint_rule_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoDeviceTaintRuleV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_device_taint_rule_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingContextV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_context_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_pod_scheduling_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoPodSchedulingV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_pod_scheduling_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_parameters_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimParametersV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_parameters_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimTemplateV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_template_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClaimV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClaimV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_claim_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClaimV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_claim_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_parameters_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassParametersV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_parameters_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClassV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceClassV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceClassV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_class_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceClassV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_class_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceSliceV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Alpha2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceSliceV1Alpha2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1alpha2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Alpha2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1alpha2"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Alpha3()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceSliceV1Alpha3CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1alpha3",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Alpha3CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1alpha3"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceSliceV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sResourceK8sIoResourceSliceV1Beta2()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sResourceK8sIoResourceSliceV1Beta2CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_resource_k8s_io_resource_slice_v1beta2",
				"k8s",
				strings.Join(dataSourceK8sResourceK8sIoResourceSliceV1Beta2CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_resource_k8s_io_resource_slice_v1beta2"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sSchedulingK8sIoPriorityClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_priority_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoPriorityClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_scheduling_k8s_io_priority_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_scheduling_k8s_io_workload_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSchedulingK8sIoWorkloadV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_scheduling_k8s_io_workload_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sSettingsK8sIoPodPresetV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sSettingsK8sIoPodPresetV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_settings_k8s_io_pod_preset_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sSettingsK8sIoPodPresetV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_settings_k8s_io_pod_preset_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIDriverV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSIDriverV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_driver_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIDriverV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_driver_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIDriverV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSIDriverV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_driver_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIDriverV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_driver_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSINodeV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSINodeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_node_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSINodeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_node_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSINodeV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSINodeV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_node_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSINodeV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_node_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSIStorageCapacityV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoCSIStorageCapacityV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_csi_storage_capacity_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoStorageClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoStorageClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_storage_class_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoStorageClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_storage_class_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoStorageClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoStorageClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_storage_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoStorageClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_storage_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttachmentV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attachment_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttachmentV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attachment_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttributesClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storage_k8s_io_volume_attributes_class_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStorageK8sIoVolumeAttributesClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storage_k8s_io_volume_attributes_class_v1beta1"] = ds
	}
	{
		ds := dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1alpha1",
				"k8s",
				strings.Join(dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1alpha1"] = ds
	}
	{
		ds := dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1()
		configured := versions.versionFor("k8s")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1beta1",
				"k8s",
				strings.Join(dataSourceK8sStoragemigrationK8sIoStorageVersionMigrationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_k8s_storagemigration_k8s_io_storage_version_migration_v1beta1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_k8s_aws_ec2_node_class_v1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_k8s_aws_ec2_node_class_v1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1Beta1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_k8s_aws_ec2_node_class_v1beta1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterK8sAwsEC2NodeClassV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_k8s_aws_ec2_node_class_v1beta1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterShNodeClaimV1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterShNodeClaimV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_sh_node_claim_v1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterShNodeClaimV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_sh_node_claim_v1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_sh_node_claim_v1beta1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterShNodeClaimV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_sh_node_claim_v1beta1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_sh_node_overlay_v1alpha1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterShNodeOverlayV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_sh_node_overlay_v1alpha1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterShNodePoolV1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterShNodePoolV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_sh_node_pool_v1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterShNodePoolV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_sh_node_pool_v1"] = ds
	}
	{
		ds := dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1()
		configured := versions.versionFor("karpenter_aws")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_aws_karpenter_sh_node_pool_v1beta1",
				"karpenter_aws",
				strings.Join(dataSourceKarpenterAwsKarpenterShNodePoolV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_aws_karpenter_sh_node_pool_v1beta1"] = ds
	}
	{
		ds := dataSourceKarpenterCoreKarpenterShNodeClaimV1()
		configured := versions.versionFor("karpenter_core")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterCoreKarpenterShNodeClaimV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_core_karpenter_sh_node_claim_v1",
				"karpenter_core",
				strings.Join(dataSourceKarpenterCoreKarpenterShNodeClaimV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_core_karpenter_sh_node_claim_v1"] = ds
	}
	{
		ds := dataSourceKarpenterCoreKarpenterShNodeClaimV1Beta1()
		configured := versions.versionFor("karpenter_core")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterCoreKarpenterShNodeClaimV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_core_karpenter_sh_node_claim_v1beta1",
				"karpenter_core",
				strings.Join(dataSourceKarpenterCoreKarpenterShNodeClaimV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_core_karpenter_sh_node_claim_v1beta1"] = ds
	}
	{
		ds := dataSourceKarpenterCoreKarpenterShNodeOverlayV1Alpha1()
		configured := versions.versionFor("karpenter_core")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterCoreKarpenterShNodeOverlayV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_core_karpenter_sh_node_overlay_v1alpha1",
				"karpenter_core",
				strings.Join(dataSourceKarpenterCoreKarpenterShNodeOverlayV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_core_karpenter_sh_node_overlay_v1alpha1"] = ds
	}
	{
		ds := dataSourceKarpenterCoreKarpenterShNodePoolV1()
		configured := versions.versionFor("karpenter_core")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterCoreKarpenterShNodePoolV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_core_karpenter_sh_node_pool_v1",
				"karpenter_core",
				strings.Join(dataSourceKarpenterCoreKarpenterShNodePoolV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_core_karpenter_sh_node_pool_v1"] = ds
	}
	{
		ds := dataSourceKarpenterCoreKarpenterShNodePoolV1Beta1()
		configured := versions.versionFor("karpenter_core")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKarpenterCoreKarpenterShNodePoolV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_karpenter_core_karpenter_sh_node_pool_v1beta1",
				"karpenter_core",
				strings.Join(dataSourceKarpenterCoreKarpenterShNodePoolV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_karpenter_core_karpenter_sh_node_pool_v1beta1"] = ds
	}
	{
		ds := dataSourceKustomizeKustomizeConfigK8sIoConfigMapArgsV1Beta1()
		configured := versions.versionFor("kustomize")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKustomizeKustomizeConfigK8sIoConfigMapArgsV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_kustomize_kustomize_config_k8s_io_config_map_args_v1beta1",
				"kustomize",
				strings.Join(dataSourceKustomizeKustomizeConfigK8sIoConfigMapArgsV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_kustomize_kustomize_config_k8s_io_config_map_args_v1beta1"] = ds
	}
	{
		ds := dataSourceKustomizeKustomizeConfigK8sIoGeneratorArgsV1Beta1()
		configured := versions.versionFor("kustomize")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKustomizeKustomizeConfigK8sIoGeneratorArgsV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_kustomize_kustomize_config_k8s_io_generator_args_v1beta1",
				"kustomize",
				strings.Join(dataSourceKustomizeKustomizeConfigK8sIoGeneratorArgsV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_kustomize_kustomize_config_k8s_io_generator_args_v1beta1"] = ds
	}
	{
		ds := dataSourceKustomizeKustomizeConfigK8sIoKustomizationV1Beta1()
		configured := versions.versionFor("kustomize")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKustomizeKustomizeConfigK8sIoKustomizationV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_kustomize_kustomize_config_k8s_io_kustomization_v1beta1",
				"kustomize",
				strings.Join(dataSourceKustomizeKustomizeConfigK8sIoKustomizationV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_kustomize_kustomize_config_k8s_io_kustomization_v1beta1"] = ds
	}
	{
		ds := dataSourceKustomizeKustomizeConfigK8sIoKvPairSourcesV1Beta1()
		configured := versions.versionFor("kustomize")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKustomizeKustomizeConfigK8sIoKvPairSourcesV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_kustomize_kustomize_config_k8s_io_kv_pair_sources_v1beta1",
				"kustomize",
				strings.Join(dataSourceKustomizeKustomizeConfigK8sIoKvPairSourcesV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_kustomize_kustomize_config_k8s_io_kv_pair_sources_v1beta1"] = ds
	}
	{
		ds := dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1()
		configured := versions.versionFor("kustomize")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_kustomize_kustomize_config_k8s_io_secret_args_v1beta1",
				"kustomize",
				strings.Join(dataSourceKustomizeKustomizeConfigK8sIoSecretArgsV1Beta1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_kustomize_kustomize_config_k8s_io_secret_args_v1beta1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_config_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerConfigV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_config_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComAlertmanagerV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_alertmanager_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_pod_monitor_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPodMonitorV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_pod_monitor_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComProbeV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComProbeV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_probe_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComProbeV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_probe_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_agent_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusAgentV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_agent_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_rule_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusRuleV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_rule_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_prometheus_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComPrometheusV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_prometheus_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_scrape_config_v1alpha1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComScrapeConfigV1Alpha1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_scrape_config_v1alpha1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_service_monitor_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComServiceMonitorV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_service_monitor_v1"] = ds
	}
	{
		ds := dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1()
		configured := versions.versionFor("prometheus_operator")
		if len(configured) > 0 {
			incompatible := versionpkg.FilterIncompatible(configured, dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1CompatibleVersions)
			if len(incompatible) > 0 {
			ds.DeprecationMessage = fmt.Sprintf(
				"%s is only guaranteed to work with %s versions %s; configured versions %s may be incompatible",
				"kubefu_prometheus_operator_monitoring_coreos_com_thanos_ruler_v1",
				"prometheus_operator",
				strings.Join(dataSourcePrometheusOperatorMonitoringCoreosComThanosRulerV1CompatibleVersions, ", "),
				strings.Join(incompatible, ", "),
			)
			}
		}
		result["kubefu_prometheus_operator_monitoring_coreos_com_thanos_ruler_v1"] = ds
	}
	return result
}

func (v Versions) versionFor(provider string) []string {
	switch provider {
	case "k8s":
		return v.K8sVersions
	case "flux":
		return v.FluxVersions
	case "cert_manager", "cert-manager":
		return v.CertManagerVersions
	case "prometheus_operator", "prometheus-operator", "prometheus":
		return v.PrometheusOperatorVersions
	case "gateway_api", "gateway-api", "gateway":
		return v.GatewayAPIVersions
	case "external_secrets", "external-secrets", "externalsecrets":
		return v.ExternalSecretsVersions
	case "kustomize":
		return v.KustomizeVersions
	case "karpenter_aws", "karpenter-aws":
		return v.KarpenterAWSVersions
	case "karpenter_core", "karpenter-core":
		return v.KarpenterCoreVersions
	default:
		return nil
	}
}
