/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-gcp/apis/compute/v1beta1"
	common "github.com/upbound/provider-gcp/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ManagedZone.
func (mg *ManagedZone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PeeringConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURL),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURLRef,
				Selector:     mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURLSelector,
				To: reference.To{
					List:    &v1beta1.NetworkList{},
					Managed: &v1beta1.Network{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURL")
			}
			mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PeeringConfig[i3].TargetNetwork[i4].NetworkURLRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateVisibilityConfig); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURL),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURLRef,
				Selector:     mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURLSelector,
				To: reference.To{
					List:    &v1beta1.NetworkList{},
					Managed: &v1beta1.Network{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURL")
			}
			mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PrivateVisibilityConfig[i3].Networks[i4].NetworkURLRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Networks); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Networks[i3].NetworkURL),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.Networks[i3].NetworkRef,
			Selector:     mg.Spec.ForProvider.Networks[i3].NetworkSelector,
			To: reference.To{
				List:    &v1beta1.NetworkList{},
				Managed: &v1beta1.Network{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Networks[i3].NetworkURL")
		}
		mg.Spec.ForProvider.Networks[i3].NetworkURL = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Networks[i3].NetworkRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this RecordSet.
func (mg *RecordSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ManagedZone),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ManagedZoneRef,
		Selector:     mg.Spec.ForProvider.ManagedZoneSelector,
		To: reference.To{
			List:    &ManagedZoneList{},
			Managed: &ManagedZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ManagedZone")
	}
	mg.Spec.ForProvider.ManagedZone = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ManagedZoneRef = rsp.ResolvedReference

	return nil
}