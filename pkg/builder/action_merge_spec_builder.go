package builder

import "github.com/turbonomic/turbo-go-sdk/pkg/proto"

type ActionMergeSpecBuilder interface {

}

type MoveMergeSpecBuilder struct {
	EntityIDs []string
	TargetID string
	ProviderType proto.EntityDTO_EntityType
}

func NewMoveMergeSpecBuilder() *MoveMergeSpecBuilder {
	return &MoveMergeSpecBuilder{}
}

func (m *MoveMergeSpecBuilder) ForEntities(entityIDs []string) *MoveMergeSpecBuilder{
	m.EntityIDs = entityIDs
	return m
}

func (m *MoveMergeSpecBuilder) MergedTo(targetID string) *MoveMergeSpecBuilder{
	m.TargetID = targetID
	return m
}

func (m *MoveMergeSpecBuilder) MovingOn(providerType proto.EntityDTO_EntityType) *MoveMergeSpecBuilder{
	m.ProviderType = providerType
	return m
}

type ResizeMergeSpecBuilder struct {
	EntityIDs []string
	TargetID string
	CommTypes []proto.CommodityDTO_CommodityType
	ChangedAttr      []proto.ActionItemDTO_CommodityAttribute
}

func (r *ResizeMergeSpecBuilder) ForEntities(entityIDs []string) *ResizeMergeSpecBuilder{
	r.EntityIDs = entityIDs
	return r
}

func (r *ResizeMergeSpecBuilder) MergedTo(targetID string) *ResizeMergeSpecBuilder{
	r.TargetID = targetID
	return r
}

func (r *ResizeMergeSpecBuilder) ForCommodities(commTypes []proto.CommodityDTO_CommodityType) *ResizeMergeSpecBuilder{
	r.CommTypes = commTypes
	return r
}

func (r *ResizeMergeSpecBuilder) ForCommoditiesAndAttributes(commTypes []proto.CommodityDTO_CommodityType,
	changedAttr      []proto.ActionItemDTO_CommodityAttribute) *ResizeMergeSpecBuilder{
	r.CommTypes = commTypes
	r.ChangedAttr = changedAttr
	return r
}

type ProvisionMergeSpecBuilder struct {
	EntityIDs []string
	TargetID string
}

func (pb *ProvisionMergeSpecBuilder) ForEntities(entityIDs []string) *ProvisionMergeSpecBuilder{
	pb.EntityIDs = entityIDs
	return pb
}

func (pb *ProvisionMergeSpecBuilder) MergedTo(targetID string) *ProvisionMergeSpecBuilder{
	pb.TargetID = targetID
	return pb
}