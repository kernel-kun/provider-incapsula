package assetassociation

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("incapsula_policy_asset_association", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		// r.Kind = "AssetAssociation"
		r.ShortGroup = "policy"

		r.References["policy_id"] = config.Reference{
			Type: "Policy",
		}

		r.References["asset_id"] = config.Reference{
			Type: "github.com/kernel-kun/provider-incapsula/apis/site/v1alpha1.Site",
		}
	})
}
