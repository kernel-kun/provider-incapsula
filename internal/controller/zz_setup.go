/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	policyassociation "github.com/kernel-kun/provider-incapsula/internal/controller/account/policyassociation"
	assetassociation "github.com/kernel-kun/provider-incapsula/internal/controller/policy/assetassociation"
	policy "github.com/kernel-kun/provider-incapsula/internal/controller/policy/policy"
	providerconfig "github.com/kernel-kun/provider-incapsula/internal/controller/providerconfig"
	site "github.com/kernel-kun/provider-incapsula/internal/controller/site/site"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policyassociation.Setup,
		assetassociation.Setup,
		policy.Setup,
		providerconfig.Setup,
		site.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
