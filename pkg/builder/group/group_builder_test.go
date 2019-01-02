package group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"testing"
)

func TestGroupBuilderInvalid(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	selectionSpec := StringProperty().
		Name("Property1").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("Hello")

	// valid entity type, no members
	_, err := StaticGroup(id).OfType(eType).Build()
	fmt.Printf("******** StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	_, err = DynamicGroup(id).OfType(eType).Build()
	fmt.Printf("******** DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)

	// no entity type
	_, err = StaticGroup(id).WithEntities([]string{"abc", "xyz"}).Build()
	fmt.Printf("******** StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	_, err = DynamicGroup(id).MatchingEntities(SelectedBy(selectionSpec)).Build()
	fmt.Printf("******** DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)

	// matching criterion for static group
	_, err = StaticGroup(id).OfType(eType).MatchingEntities(SelectedBy(selectionSpec)).WithEntities([]string{"abc", "xyz"}).Build()
	fmt.Printf("####### StaticGroup Error %s\n", err)
	assert.NotNil(t, err)

	// member list for dynamic group
	_, err = DynamicGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"}).MatchingEntities(SelectedBy(selectionSpec)).Build()
	fmt.Printf("####### DynamicGroup Error %s\n", err)
	assert.NotNil(t, err)
}

func TestGroupBuilderEntityType(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	// valid entity type
	groupBuilder := StaticGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"})
	assert.NotNil(t, groupBuilder)

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, groupDTO.EntityType)
	assert.EqualValues(t, eType, groupDTO.GetEntityType())

	// unknown entity type
	var fakeType proto.EntityDTO_EntityType = 70
	groupBuilder = StaticGroup(id).OfType(fakeType).WithEntities([]string{"abc", "xyz"})

	groupDTO, err = groupBuilder.Build()
	assert.NotNil(t, err)
	//
	// default entity type
	eType = 0
	groupBuilder3 := StaticGroup(id).OfType(eType).WithEntities([]string{"abc", "xyz"})

	groupDTO3, err := groupBuilder3.Build()
	assert.Nil(t, err)
	assert.NotNil(t, groupDTO3.EntityType)
	assert.EqualValues(t, proto.EntityDTO_SWITCH, groupDTO3.GetEntityType())

	// overwrite existing entity type
	groupBuilder3 = groupBuilder3.OfType(30).WithEntities([]string{"abc", "xyz"})

	groupDTO3, err = groupBuilder3.Build()
	assert.NotNil(t, err)
}

func TestGroupBuilderDynamic(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	selectionSpec := StringProperty().
		Name("Property1").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("Hello")

	selectionSpec2 := StringProperty().
		Name("Property2").
		Expression(proto.GroupDTO_SelectionSpec_EQUAL_TO).
		SetProperty("World")

	groupBuilder := DynamicGroup(id).
		OfType(eType).
		MatchingEntities(SelectedBy(selectionSpec).and(selectionSpec2))

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)

	assert.EqualValues(t, eType, groupDTO.GetEntityType())
	assert.EqualValues(t, id, groupDTO.GetDisplayName())

	assert.EqualValues(t, 2, len(groupDTO.GetSelectionSpecList().SelectionSpec))
	assert.Nil(t, groupDTO.GetMemberList())
}

func TestGroupBuilderStatic(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	groupBuilder := StaticGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	assert.EqualValues(t, eType, groupDTO.GetEntityType())
	assert.EqualValues(t, id, groupDTO.GetDisplayName())
	assert.EqualValues(t, 2, len(groupDTO.GetMemberList().GetMember()))
	assert.Nil(t, groupDTO.GetSelectionSpecList())
}

func TestGroupBuilderSetConsistentResize(t *testing.T) {
	id := "group1"
	eType := proto.EntityDTO_CONTAINER

	groupBuilder := StaticGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"})

	groupDTO, err := groupBuilder.Build()
	assert.Nil(t, err)
	consistentResize := groupDTO.GetIsConsistentResizing()
	assert.False(t, consistentResize)

	fmt.Printf("%++v\n", groupDTO)

	groupBuilder = StaticGroup(id).
		OfType(eType).
		WithEntities([]string{"abc", "xyz"}).
		ResizeConsistently()

	groupDTO, err = groupBuilder.Build()
	assert.Nil(t, err)
	consistentResize = groupDTO.GetIsConsistentResizing()
	assert.True(t, consistentResize)

	fmt.Printf("%++v\n", groupDTO)
}
