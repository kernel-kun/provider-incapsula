/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/kernel-kun/provider-incapsula/config/assetassociation"
	"github.com/kernel-kun/provider-incapsula/config/policy"
	"github.com/kernel-kun/provider-incapsula/config/policyassociation"
	"github.com/kernel-kun/provider-incapsula/config/site"
)

const (
	resourcePrefix = "incapsula"
	modulePath     = "github.com/kernel-kun/provider-incapsula"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		site.Configure,
		policy.Configure,
		policyassociation.Configure,
		assetassociation.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
