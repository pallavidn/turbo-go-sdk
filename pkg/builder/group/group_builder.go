package group

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/pkg/builder"
)

type GroupType string

const (
	// Static group contains a fixed list of entity id's
	STATIC_GROUP      GroupType = "Static"
	// Dynamic group contains selection criteria to select entity id's
	DYNAMIC_GROUP     GroupType = "Dynamic"
)
//
//type GroupBuilder interface {
//	OfType(eType proto.EntityDTO_EntityType) *GroupBuilder
//	WithEntities(entities []string) *GroupBuilder
//	MatchingEntities(matching *Matching) *GroupBuilder
//	UpdateType(updateType proto.UpdateType) *GroupBuilder
//	GroupType(updateType proto.GroupDTO_GroupType) *GroupBuilder
//}

// Builder for creating a GroupDTO
type AbstractBuilder struct {
	groupId string
	entityTypePtr *proto.EntityDTO_EntityType
	memberList []string
	matching *Matching
	//groupDTO *proto.GroupDTO
	ec *builder.ErrorCollector
	groupType GroupType
}

// Create a new instance of AbstractBuilder.
// Specify the group id and if the group is static or dynamic.
func newAbstractBuilder(id string, groupType GroupType) *AbstractBuilder {
	// group id and display name
	//groupId := &proto.GroupDTO_GroupName{
	//	GroupName: id,
	//}
	//groupDTO := &proto.GroupDTO{
	//	DisplayName: &id,
	//	Info: groupId,
	//}

	groupBuilder := &AbstractBuilder{
		//groupDTO: groupDTO,
		groupType: groupType,
		groupId: id,
		ec:  new(builder.ErrorCollector),
	}
	return groupBuilder
}

// Create a new instance of builder for creating Static groups.
// Static group contains a fixed list of entity id's
func StaticGroup(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, STATIC_GROUP)
	return groupBuilder
}

// Create a new instance of builder for creating Dynamic groups.
// Dynamic group contains selection criteria using entity properties to select entities.
func DynamicGroup(id string) *AbstractBuilder {
	groupBuilder := newAbstractBuilder(id, DYNAMIC_GROUP)
	return groupBuilder
}

// Return the Protobuf GroupDTO object. There is no constraint object with this group.
// Return error if errors were collected during the building of the group properties.
func (groupBuilder *AbstractBuilder) Build() (*proto.GroupDTO, error) {

	groupId := &proto.GroupDTO_GroupName{
		GroupName: groupBuilder.groupId,
	}
	groupDTO := &proto.GroupDTO{
		DisplayName: &groupBuilder.groupId,
		Info: groupId,
	}

	err := groupBuilder.setupEntityType(groupDTO)
	if err != nil {
		fmt.Printf("[groupBuilder] %s\n", err)
		groupBuilder.ec.Collect(err)
	}

	if groupBuilder.groupType == STATIC_GROUP {
		err := groupBuilder.setUpStaticMembers(groupDTO)
		if err != nil {
			fmt.Printf("[groupBuilder] %s\n", err)
			groupBuilder.ec.Collect(err)
		}
	} else {
		err := groupBuilder.setUpDynamicGroup(groupDTO)
		if err != nil {
			fmt.Printf("[groupBuilder] %s\n", err)
			groupBuilder.ec.Collect(err)
		}
	}

	if groupBuilder.ec.Count() > 0 {
		fmt.Printf("Error : %s\n", groupBuilder.ec.Error())
		return nil, fmt.Errorf(groupBuilder.ec.Error())
	}

	//members := groupBuilder.groupDTO.Members
	//if members == nil {
	//	fmt.Printf("NULL MEMBERS\n")
	//}
	//if groupBuilder.groupDTO == nil {
	//	return nil, fmt.Errorf("Group dto is null\n")
	//}

	return groupDTO, nil
}

// Set the entity type for the members of the group.
// All the entities in a group belong to the same entity type.
func (groupBuilder *AbstractBuilder) OfType(eType proto.EntityDTO_EntityType) *AbstractBuilder {

	// Check entity type
	if groupBuilder.entityTypePtr != nil && *groupBuilder.entityTypePtr !=  eType {
		groupBuilder.ec.Collect(fmt.Errorf("Cannot add members, input EntityType - %s is not consistent with existing EntityType %s ",
			eType, *groupBuilder.entityTypePtr))
		return groupBuilder
	}

	// Setup entity type
	groupBuilder.entityTypePtr = &eType

	return groupBuilder
}


func (groupBuilder *AbstractBuilder) setupEntityType(groupDTO *proto.GroupDTO) error {
	if groupBuilder.entityTypePtr == nil {
		return fmt.Errorf("Entity type is not set")
	}
	// Validate entity type
	entityType := *groupBuilder.entityTypePtr
	_, valid := proto.EntityDTO_EntityType_name[int32(entityType)]

	if !valid {
		return fmt.Errorf("invalid entity type %v\n", entityType)
	}

	// Setup entity type
	groupDTO.EntityType = &entityType
	return nil
}

// Set the members for a static group. Input is a list of UUIDs for the entities that belong to the group.
func (groupBuilder *AbstractBuilder) WithEntities(entities []string) *AbstractBuilder {

	// Assert that the group is a static group
	if groupBuilder.groupType != STATIC_GROUP {
		groupBuilder.ec.Collect(fmt.Errorf("Cannot set member uuid list for dynamic group"))
		return groupBuilder
	}
	groupBuilder.memberList = entities

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) setUpStaticMembers(groupDTO *proto.GroupDTO) error {

	if len(groupBuilder.memberList) == 0 {
		return fmt.Errorf("Empty member list")
	}

	// Set the Group DTO member field
	memberList := &proto.GroupDTO_MemberList{
		MemberList: &proto.GroupDTO_MembersList{
			Member: groupBuilder.memberList,
		},
	}
	groupDTO.Members = memberList
	return nil
}

// Set the members matching criteria for a dynamic group.
func (groupBuilder *AbstractBuilder) MatchingEntities(matching *Matching) *AbstractBuilder {

	// Assert that the group is a dynamci group
	if groupBuilder.groupType != DYNAMIC_GROUP {
		groupBuilder.ec.Collect(fmt.Errorf("Cannot set matching criteria for static group"))
		return groupBuilder
	}
	groupBuilder.matching = matching

	return groupBuilder
}

func (groupBuilder *AbstractBuilder) setUpDynamicGroup(groupDTO *proto.GroupDTO) error {

	if groupBuilder.matching == nil {
		return fmt.Errorf("Null matching criteria for member selection")
	}

	var selectionSpecList []*proto.GroupDTO_SelectionSpec

	// Build the selection spec list from the matching criteria
	selectionSpecBuilderList := groupBuilder.matching.selectionSpecBuilderList
	for _, specBuilder := range selectionSpecBuilderList {
		selectionSpec := specBuilder.Build()
		selectionSpecList = append(selectionSpecList, selectionSpec)
	}

	// Set the Group DTO member field
	selectionSpecList_ := &proto.GroupDTO_SelectionSpecList_{
		SelectionSpecList: &proto.GroupDTO_SelectionSpecList{
			SelectionSpec: selectionSpecList,
		},
	}
	groupDTO.Members = selectionSpecList_
	return nil
}
