package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestResizeActionMergeSpec_Create(t *testing.T) {

	targetId := "test-deployment"
	entities := []string{"e1", "e2", "e3"}
	comms := []proto.CommodityDTO_CommodityType{proto.CommodityDTO_VCPU, proto.CommodityDTO_VMEM}
	spec := NewResizeMergeSpecBuilder()
	spec.ForEntities(entities).
		ForCommodities(comms).
		MergedTo(targetId)

	resizeSpec, err := spec.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, *resizeSpec.TargetEntity, targetId)
	assert.ElementsMatch(t, resizeSpec.EntityIds, entities)
	assert.ElementsMatch(t, comms, resizeSpec.GetResizeSpec().CommodityType)

	assert.True(t, resizeSpec.GetProvisionSpec() == nil)
	assert.True(t, resizeSpec.GetMoveSpec() == nil)
	assert.True(t, resizeSpec.GetResizeSpec() != nil)
}

func TestMoveActionMergeSpec_Create(t *testing.T) {

	targetId := "test-deployment"
	entities := []string{"e1", "e2", "e3"}
	provider := proto.EntityDTO_VIRTUAL_MACHINE
	spec := NewMoveMergeSpecBuilder()
	spec.ForEntities(entities).MovingOn(&provider).MergedTo(targetId)

	moveSpec, err := spec.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, *moveSpec.TargetEntity, targetId)
	assert.ElementsMatch(t, moveSpec.EntityIds, entities)
	assert.EqualValues(t, provider, moveSpec.GetMoveSpec().GetProviderType())

	assert.True(t, moveSpec.GetProvisionSpec() == nil)
	assert.True(t, moveSpec.GetMoveSpec() != nil)
	assert.True(t, moveSpec.GetResizeSpec() == nil)
}

func TestProvisionActionMergeSpec_Create(t *testing.T) {

	targetId := "test-deployment"
	entities := []string{"e1", "e2", "e3"}
	spec := NewProvisionMergeSpecBuilder()
	spec.ForEntities(entities).MergedTo(targetId)

	provisionSpec, err := spec.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, *provisionSpec.TargetEntity, targetId)
	assert.ElementsMatch(t, provisionSpec.EntityIds, entities)

	assert.True(t, provisionSpec.GetProvisionSpec() != nil)
	assert.True(t, provisionSpec.GetMoveSpec() == nil)
	assert.True(t, provisionSpec.GetResizeSpec() == nil)
}
