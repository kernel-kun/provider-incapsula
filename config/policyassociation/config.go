package policyassociation

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("incapsula_account_policy_association", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "account"

		r.References["default_non_mandatory_policy_ids"] = config.Reference{
			Type: "github.com/kernel-kun/provider-incapsula/apis/policy/v1alpha1.Policy",
		}
	})
}
