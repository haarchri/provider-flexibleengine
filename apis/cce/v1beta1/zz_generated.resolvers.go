/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta13 "github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1"
	v1beta1 "github.com/FrangipaneTeam/provider-flexibleengine/apis/eip/v1beta1"
	v1beta12 "github.com/FrangipaneTeam/provider-flexibleengine/apis/kms/v1beta1"
	v1beta11 "github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1"
	common "github.com/FrangipaneTeam/provider-flexibleengine/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Addon.
func (mg *Addon) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CCENameSpace.
func (mg *CCENameSpace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EIP),
		Extract:      common.AddressExtractor(),
		Reference:    mg.Spec.ForProvider.EIPRef,
		Selector:     mg.Spec.ForProvider.EIPSelector,
		To: reference.To{
			List:    &v1beta1.EIPList{},
			Managed: &v1beta1.EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIP")
	}
	mg.Spec.ForProvider.EIP = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EIPRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HighwaySubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HighwaySubnetIDRef,
		Selector:     mg.Spec.ForProvider.HighwaySubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HighwaySubnetID")
	}
	mg.Spec.ForProvider.HighwaySubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HighwaySubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      common.IDExtractor(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VPCIDRef,
		Selector:     mg.Spec.ForProvider.VPCIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCList{},
			Managed: &v1beta11.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Node.
func (mg *Node) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDRef,
			Selector:     mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDSelector,
			To: reference.To{
				List:    &v1beta12.KeyList{},
				Managed: &v1beta12.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID")
		}
		mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.EIPIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.EIPIdsRefs,
		Selector:      mg.Spec.ForProvider.EIPIdsSelector,
		To: reference.To{
			List:    &v1beta1.EIPList{},
			Managed: &v1beta1.EIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EIPIds")
	}
	mg.Spec.ForProvider.EIPIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.EIPIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1beta13.KeyPairList{},
			Managed: &v1beta13.KeyPair{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodePool.
func (mg *NodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataVolumes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDRef,
			Selector:     mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDSelector,
			To: reference.To{
				List:    &v1beta12.KeyList{},
				Managed: &v1beta12.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID")
		}
		mg.Spec.ForProvider.DataVolumes[i3].KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataVolumes[i3].KMSKeyIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyPair),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KeyPairRef,
		Selector:     mg.Spec.ForProvider.KeyPairSelector,
		To: reference.To{
			List:    &v1beta13.KeyPairList{},
			Managed: &v1beta13.KeyPair{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyPair")
	}
	mg.Spec.ForProvider.KeyPair = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyPairRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      common.IDExtractor(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.VPCSubnetList{},
			Managed: &v1beta11.VPCSubnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pvc.
func (mg *Pvc) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &CCENameSpaceList{},
			Managed: &CCENameSpace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}
