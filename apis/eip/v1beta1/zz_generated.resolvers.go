/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta1 "github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EIPAssociate.
func (mg *EIPAssociate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1beta1.NetworkList{},
			Managed: &v1beta1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortIDRef,
		Selector:     mg.Spec.ForProvider.PortIDSelector,
		To: reference.To{
			List:    &v1beta1.PortList{},
			Managed: &v1beta1.Port{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortID")
	}
	mg.Spec.ForProvider.PortID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortIDRef = rsp.ResolvedReference

	return nil
}
