package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tasansga/terraform-provider-kubefu/resourcegen"
	downloader "github.com/tasansga/terraform-provider-kubefu/resourcegen/downloader"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logger      = logrus.New()
	outputDir   = filepath.Join("kubefu", "generated")
	packageName = "generated"
	schemaDir   = "schemas"
)

type providerConfig struct {
	displayName string
	subdir      string
	factory     func(logrus.FieldLogger) downloader.SchemaDownloader
}

var providerConfigs = map[string]providerConfig{
	"k8s": {
		displayName: "kubernetes",
		subdir:      "k8s",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewKubernetesDownloader(logger)
		},
	},
	"cert-manager": {
		displayName: "cert-manager",
		subdir:      "cert-manager",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewCertManagerDownloader(logger)
		},
	},
	"flux": {
		displayName: "flux",
		subdir:      "flux",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewFluxDownloader(logger)
		},
	},
	"prometheus-operator": {
		displayName: "prometheus-operator",
		subdir:      "prometheus-operator",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewPrometheusOperatorDownloader(logger)
		},
	},
	"gateway-api": {
		displayName: "gateway-api",
		subdir:      "gateway-api",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewGatewayAPIDownloader(logger)
		},
	},
	"external-secrets": {
		displayName: "external-secrets",
		subdir:      "external-secrets",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewExternalSecretsDownloader(logger)
		},
	},
	"kustomize": {
		displayName: "kustomize",
		subdir:      "kustomize",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewKustomizeDownloader(logger)
		},
	},
	"karpenter-aws": {
		displayName: "karpenter-aws",
		subdir:      "karpenter-aws",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewKarpenterDownloader(logger)
		},
	},
	"karpenter-core": {
		displayName: "karpenter-core",
		subdir:      "karpenter-core",
		factory: func(logger logrus.FieldLogger) downloader.SchemaDownloader {
			return downloader.NewKarpenterCoreDownloader(logger)
		},
	},
}

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	rootCmd.PersistentFlags().StringVar(&outputDir, "out-dir", outputDir, "directory where generated data source files go")
	rootCmd.PersistentFlags().StringVar(&packageName, "package", packageName, "package used in created Go files")
	downloadSchemaCmd.Flags().StringVar(&schemaDir, "schema-dir", schemaDir, "directory where downloaded schemas land")
}

var isInit = false

var rootCmd = &cobra.Command{
	Use:           "resourcegen",
	Short:         "Generate terraform resource definitions in golang based on Kubernetes JSON resource definitions",
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.CalledAs() == "help" || cmd.Flags().Lookup("help").Changed {
			return nil
		}
		return nil
	},
}

// generateCmd is the entry point for future scaffolding that generates resources.
var generateCmd = &cobra.Command{
	Use:   "generate <schema-root>",
	Short: "Collect schema files and generate Terraform datasource Go sources",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		if err := resourcegen.GenerateFromSchemas(path, outputDir, packageName); err != nil {
			return err
		}
		logger.WithFields(logrus.Fields{
			"path":    path,
			"out-dir": outputDir,
		}).Info("generated resources from schemas")
		return nil
	},
}

var downloadSchemaCmd = &cobra.Command{
	Use:          "download-schema [provider...]",
	Short:        "Download missing OpenAPI specs for supported providers (defaults to all)",
	Args:         cobra.ArbitraryArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return downloadAllSchemas()
		}
		return downloadSchemasForProviders(args)
	},
}

func downloadSchemasForProviders(providers []string) error {
	seen := map[string]struct{}{}
	var downloadErrors []error

	for _, provider := range providers {
		cfg, ok := providerConfigs[provider]
		if !ok {
			downloadErrors = append(downloadErrors, fmt.Errorf("unsupported provider %q", provider))
			continue
		}
		if _, ok := seen[cfg.subdir]; ok {
			continue
		}
		seen[cfg.subdir] = struct{}{}
		if err := downloadSchemas(cfg); err != nil {
			downloadErrors = append(downloadErrors, err)
		}
	}

	if len(downloadErrors) == 0 {
		return nil
	}
	var errStrings []string
	for _, err := range downloadErrors {
		errStrings = append(errStrings, err.Error())
	}
	return fmt.Errorf("download errors: %s", strings.Join(errStrings, "; "))
}

func downloadSchemas(cfg providerConfig) error {
	targetDir := filepath.Join(schemaDir, cfg.subdir)
	logger.WithFields(logrus.Fields{
		"command":   "download-schema",
		"provider":  cfg.displayName,
		"schemaDir": targetDir,
	}).Info("starting schema download")

	schemaDownloader := cfg.factory(logger)
	count, err := schemaDownloader.DownloadMissingSchemas(targetDir)
	if err != nil {
		return err
	}

	if count == 0 {
		logger.WithFields(logrus.Fields{
			"dir":      targetDir,
			"provider": cfg.displayName,
		}).Info("no new schemas downloaded")
		return nil
	}

	logger.WithFields(logrus.Fields{
		"dir":      targetDir,
		"count":    count,
		"provider": cfg.displayName,
	}).Info("downloaded schemas for provider")
	return nil
}

func downloadAllSchemas() error {
	seen := map[string]struct{}{}
	var downloadErrors []error

	for _, cfg := range providerConfigs {
		if _, ok := seen[cfg.subdir]; ok {
			continue
		}
		seen[cfg.subdir] = struct{}{}
		if err := downloadSchemas(cfg); err != nil {
			downloadErrors = append(downloadErrors, err)
		}
	}

	if len(downloadErrors) == 0 {
		return nil
	}
	var errStrings []string
	for _, err := range downloadErrors {
		errStrings = append(errStrings, err.Error())
	}
	return fmt.Errorf("download errors: %s", strings.Join(errStrings, "; "))
}

var mainCmds = map[string]*cobra.Command{
	"generate":        generateCmd,
	"download-schema": downloadSchemaCmd,
}

func initCmds() *cobra.Command {
	if isInit {
		return rootCmd
	}
	for _, cmd := range mainCmds {
		rootCmd.AddCommand(cmd)
	}
	isInit = true
	return rootCmd
}

func main() {
	cmd := initCmds()
	logger.WithField("args", os.Args[1:]).Info("resourcegen starting")
	defer logger.Info("resourcegen shutting down")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
