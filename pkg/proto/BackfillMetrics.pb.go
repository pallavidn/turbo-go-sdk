// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: BackfillMetrics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a transport message from a mediation probe to services consuming
// entity metrics data. It contains the smallest amount of data needed to backfill
// percentiles, historical values, and stats for previously-discovered entities.
// This message is intended to be chunked/streamed over time.
type BackfilledMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Essentially a session ID for the overall backfill operation, in case there are
	// issues like orphaned messages, or somehow there are two different operations at
	// once.
	BackfillOperationId *int64 `protobuf:"varint,1,req,name=backfill_operation_id,json=backfillOperationId" json:"backfill_operation_id,omitempty"`
	// The segment of the message stream. Its 3 modes signify whether the stream is
	// starting, in progress, or at the end.
	//
	// Types that are assignable to Segment:
	//
	//	*BackfilledMetrics_Metadata_
	//	*BackfilledMetrics_Data_
	//	*BackfilledMetrics_End_
	Segment isBackfilledMetrics_Segment `protobuf_oneof:"segment"`
}

func (x *BackfilledMetrics) Reset() {
	*x = BackfilledMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics) ProtoMessage() {}

func (x *BackfilledMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0}
}

func (x *BackfilledMetrics) GetBackfillOperationId() int64 {
	if x != nil && x.BackfillOperationId != nil {
		return *x.BackfillOperationId
	}
	return 0
}

func (m *BackfilledMetrics) GetSegment() isBackfilledMetrics_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *BackfilledMetrics) GetMetadata() *BackfilledMetrics_Metadata {
	if x, ok := x.GetSegment().(*BackfilledMetrics_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

func (x *BackfilledMetrics) GetData() *BackfilledMetrics_Data {
	if x, ok := x.GetSegment().(*BackfilledMetrics_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *BackfilledMetrics) GetEnd() *BackfilledMetrics_End {
	if x, ok := x.GetSegment().(*BackfilledMetrics_End_); ok {
		return x.End
	}
	return nil
}

type isBackfilledMetrics_Segment interface {
	isBackfilledMetrics_Segment()
}

type BackfilledMetrics_Metadata_ struct {
	// If the segment is this type, the stream is at the beginning, and it contains
	// metadata about the metrics that will come in later messages.
	Metadata *BackfilledMetrics_Metadata `protobuf:"bytes,2,opt,name=metadata,oneof"`
}

type BackfilledMetrics_Data_ struct {
	// A segment of the message that contains actual entity metric data.
	Data *BackfilledMetrics_Data `protobuf:"bytes,3,opt,name=data,oneof"`
}

type BackfilledMetrics_End_ struct {
	// If the segment is this type, it signifies that the stream is finished,
	// and the backfill operation can continue downstream.
	End *BackfilledMetrics_End `protobuf:"bytes,4,opt,name=end,oneof"`
}

func (*BackfilledMetrics_Metadata_) isBackfilledMetrics_Segment() {}

func (*BackfilledMetrics_Data_) isBackfilledMetrics_Segment() {}

func (*BackfilledMetrics_End_) isBackfilledMetrics_Segment() {}

// Using shared window across entities since currently percentile blobs are calculated
// for LARGEST observation period across all entities. We may be able to change this.
type BackfilledMetrics_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Beginning of time window for backfill period.
	MetricsStartTimestamp *int64 `protobuf:"varint,1,req,name=metrics_start_timestamp,json=metricsStartTimestamp" json:"metrics_start_timestamp,omitempty"`
	// End of time window for backfill period. I assume this would be up to the
	// the start of the first discovery.
	MetricsEndTimestamp *int64 `protobuf:"varint,2,req,name=metrics_end_timestamp,json=metricsEndTimestamp" json:"metrics_end_timestamp,omitempty"`
}

func (x *BackfilledMetrics_Metadata) Reset() {
	*x = BackfilledMetrics_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_Metadata) ProtoMessage() {}

func (x *BackfilledMetrics_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_Metadata.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_Metadata) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BackfilledMetrics_Metadata) GetMetricsStartTimestamp() int64 {
	if x != nil && x.MetricsStartTimestamp != nil {
		return *x.MetricsStartTimestamp
	}
	return 0
}

func (x *BackfilledMetrics_Metadata) GetMetricsEndTimestamp() int64 {
	if x != nil && x.MetricsEndTimestamp != nil {
		return *x.MetricsEndTimestamp
	}
	return 0
}

// Actual metrics data being streamed.
// Represents a chunk of data. Each chunk can have as many
// EntityInfo/metric data pairs as makes sense.
type BackfilledMetrics_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityInfo []*BackfilledMetrics_EntityInfo `protobuf:"bytes,1,rep,name=entity_info,json=entityInfo" json:"entity_info,omitempty"`
}

func (x *BackfilledMetrics_Data) Reset() {
	*x = BackfilledMetrics_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_Data) ProtoMessage() {}

func (x *BackfilledMetrics_Data) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_Data.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_Data) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *BackfilledMetrics_Data) GetEntityInfo() []*BackfilledMetrics_EntityInfo {
	if x != nil {
		return x.EntityInfo
	}
	return nil
}

type BackfilledMetrics_EntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimal set of information needed to get the OID of the entity, TP-side.
	TargetEntityIdSet *BackfilledMetrics_EntityIdSet `protobuf:"bytes,1,opt,name=target_entity_id_set,json=targetEntityIdSet" json:"target_entity_id_set,omitempty"`
	// Temp ID for the entity sot that we do not have to repeat the entity id set.
	EntityInfoId *uint64 `protobuf:"varint,2,opt,name=entity_info_id,json=entityInfoId" json:"entity_info_id,omitempty"`
	// Simulates a "virtual" discovery number had this been discovered
	// under the current live discovery workflow.
	DiscoveryNumber *uint64 `protobuf:"varint,3,req,name=discovery_number,json=discoveryNumber" json:"discovery_number,omitempty"`
	// Commodities sold and bought ala regular discovery.
	CommoditiesSold   []*BackfilledMetrics_CommoditySoldMetrics `protobuf:"bytes,4,rep,name=commodities_sold,json=commoditiesSold" json:"commodities_sold,omitempty"`
	CommoditiesBought []*BackfilledMetrics_CommodityBought      `protobuf:"bytes,5,rep,name=commodities_bought,json=commoditiesBought" json:"commodities_bought,omitempty"`
}

func (x *BackfilledMetrics_EntityInfo) Reset() {
	*x = BackfilledMetrics_EntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_EntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_EntityInfo) ProtoMessage() {}

func (x *BackfilledMetrics_EntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_EntityInfo.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_EntityInfo) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 2}
}

func (x *BackfilledMetrics_EntityInfo) GetTargetEntityIdSet() *BackfilledMetrics_EntityIdSet {
	if x != nil {
		return x.TargetEntityIdSet
	}
	return nil
}

func (x *BackfilledMetrics_EntityInfo) GetEntityInfoId() uint64 {
	if x != nil && x.EntityInfoId != nil {
		return *x.EntityInfoId
	}
	return 0
}

func (x *BackfilledMetrics_EntityInfo) GetDiscoveryNumber() uint64 {
	if x != nil && x.DiscoveryNumber != nil {
		return *x.DiscoveryNumber
	}
	return 0
}

func (x *BackfilledMetrics_EntityInfo) GetCommoditiesSold() []*BackfilledMetrics_CommoditySoldMetrics {
	if x != nil {
		return x.CommoditiesSold
	}
	return nil
}

func (x *BackfilledMetrics_EntityInfo) GetCommoditiesBought() []*BackfilledMetrics_CommodityBought {
	if x != nil {
		return x.CommoditiesBought
	}
	return nil
}

// Slimmed down version of CommodityDTO. Contains only the values
// needed for backfilling relevant metrics.
type BackfilledMetrics_CommoditySoldMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of commodity being sold/type of turbo metric.
	CommodityType *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=commodity_type,json=commodityType,enum=common_dto.CommodityDTO_CommodityType" json:"commodity_type,omitempty"`
	// Identifier for selling provider. Can be referenced by
	// bought relations.
	Key *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	// Used value for the virtual discovery.
	Used *float64 `protobuf:"fixed64,3,opt,name=used" json:"used,omitempty"`
	// Peak value for the virtual discovery.
	Peak *float64 `protobuf:"fixed64,4,opt,name=peak" json:"peak,omitempty"`
	// Metrics parsed into utilization data.
	Data *CommodityDTO_UtilizationData `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (x *BackfilledMetrics_CommoditySoldMetrics) Reset() {
	*x = BackfilledMetrics_CommoditySoldMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_CommoditySoldMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_CommoditySoldMetrics) ProtoMessage() {}

func (x *BackfilledMetrics_CommoditySoldMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_CommoditySoldMetrics.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_CommoditySoldMetrics) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 3}
}

func (x *BackfilledMetrics_CommoditySoldMetrics) GetCommodityType() CommodityDTO_CommodityType {
	if x != nil && x.CommodityType != nil {
		return *x.CommodityType
	}
	return CommodityDTO_CLUSTER
}

func (x *BackfilledMetrics_CommoditySoldMetrics) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *BackfilledMetrics_CommoditySoldMetrics) GetUsed() float64 {
	if x != nil && x.Used != nil {
		return *x.Used
	}
	return 0
}

func (x *BackfilledMetrics_CommoditySoldMetrics) GetPeak() float64 {
	if x != nil && x.Peak != nil {
		return *x.Peak
	}
	return 0
}

func (x *BackfilledMetrics_CommoditySoldMetrics) GetData() *CommodityDTO_UtilizationData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Useful for modifying bought relations, since the interplay between
// bought/sold can be used downstream.
type BackfilledMetrics_CommodityBought struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provider key as above.
	ProviderKey *string `protobuf:"bytes,1,req,name=providerKey" json:"providerKey,omitempty"`
	// Amounts bought from provider.
	Bought []*BackfilledMetrics_CommoditySoldMetrics `protobuf:"bytes,2,rep,name=bought" json:"bought,omitempty"`
}

func (x *BackfilledMetrics_CommodityBought) Reset() {
	*x = BackfilledMetrics_CommodityBought{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_CommodityBought) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_CommodityBought) ProtoMessage() {}

func (x *BackfilledMetrics_CommodityBought) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_CommodityBought.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_CommodityBought) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 4}
}

func (x *BackfilledMetrics_CommodityBought) GetProviderKey() string {
	if x != nil && x.ProviderKey != nil {
		return *x.ProviderKey
	}
	return ""
}

func (x *BackfilledMetrics_CommodityBought) GetBought() []*BackfilledMetrics_CommoditySoldMetrics {
	if x != nil {
		return x.Bought
	}
	return nil
}

// Signifies the end of the backfill metrics stream.
type BackfilledMetrics_End struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of distinct entities that the stream has metrics for.
	EntitiesAffected *uint64 `protobuf:"varint,1,req,name=entities_affected,json=entitiesAffected" json:"entities_affected,omitempty"`
	// Total utilization points loaded.
	TotalPoints *uint64 `protobuf:"varint,2,req,name=total_points,json=totalPoints" json:"total_points,omitempty"`
	// Total number of virtual discoveries returned.
	VirtualDiscoveryCount *uint64 `protobuf:"varint,3,req,name=virtual_discovery_count,json=virtualDiscoveryCount" json:"virtual_discovery_count,omitempty"`
}

func (x *BackfilledMetrics_End) Reset() {
	*x = BackfilledMetrics_End{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_End) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_End) ProtoMessage() {}

func (x *BackfilledMetrics_End) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_End.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_End) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 5}
}

func (x *BackfilledMetrics_End) GetEntitiesAffected() uint64 {
	if x != nil && x.EntitiesAffected != nil {
		return *x.EntitiesAffected
	}
	return 0
}

func (x *BackfilledMetrics_End) GetTotalPoints() uint64 {
	if x != nil && x.TotalPoints != nil {
		return *x.TotalPoints
	}
	return 0
}

func (x *BackfilledMetrics_End) GetVirtualDiscoveryCount() uint64 {
	if x != nil && x.VirtualDiscoveryCount != nil {
		return *x.VirtualDiscoveryCount
	}
	return 0
}

// Models the minimum amount of data about an entity needed to resolve
// its OID by TP. Each of these is defined by probes in
// `probe-conf/identity-metadata.yml`.
type BackfilledMetrics_EntityIdSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NonvolatileProperties map[string]string `protobuf:"bytes,1,rep,name=nonvolatile_properties,json=nonvolatileProperties" json:"nonvolatile_properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	VolatileProperties    map[string]string `protobuf:"bytes,2,rep,name=volatile_properties,json=volatileProperties" json:"volatile_properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HeuristicProperties   map[string]string `protobuf:"bytes,3,rep,name=heuristic_properties,json=heuristicProperties" json:"heuristic_properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *BackfilledMetrics_EntityIdSet) Reset() {
	*x = BackfilledMetrics_EntityIdSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackfillMetrics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackfilledMetrics_EntityIdSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackfilledMetrics_EntityIdSet) ProtoMessage() {}

func (x *BackfilledMetrics_EntityIdSet) ProtoReflect() protoreflect.Message {
	mi := &file_BackfillMetrics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackfilledMetrics_EntityIdSet.ProtoReflect.Descriptor instead.
func (*BackfilledMetrics_EntityIdSet) Descriptor() ([]byte, []int) {
	return file_BackfillMetrics_proto_rawDescGZIP(), []int{0, 6}
}

func (x *BackfilledMetrics_EntityIdSet) GetNonvolatileProperties() map[string]string {
	if x != nil {
		return x.NonvolatileProperties
	}
	return nil
}

func (x *BackfilledMetrics_EntityIdSet) GetVolatileProperties() map[string]string {
	if x != nil {
		return x.VolatileProperties
	}
	return nil
}

func (x *BackfilledMetrics_EntityIdSet) GetHeuristicProperties() map[string]string {
	if x != nil {
		return x.HeuristicProperties
	}
	return nil
}

var File_BackfillMetrics_proto protoreflect.FileDescriptor

var file_BackfillMetrics_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x64, 0x74, 0x6f, 0x1a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x54, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x0f, 0x0a, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x62, 0x61,
	0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x13, 0x62, 0x61, 0x63, 0x6b, 0x66,
	0x69, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x44,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c,
	0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x64, 0x48, 0x00,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x1a, 0x76, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x03, 0x52, 0x15, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x13, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x51, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0xf6, 0x02, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x5a, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x53, 0x65, 0x74, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x53, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53,
	0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x6f, 0x6c, 0x64, 0x12, 0x5c, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x62, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x1a, 0xdd, 0x01, 0x0a, 0x14, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x79, 0x44, 0x54, 0x4f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x61, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x70, 0x65, 0x61, 0x6b, 0x12, 0x3c, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x44, 0x54, 0x4f, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x7f, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x4a,
	0x0a, 0x06, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x1a, 0x8d, 0x01, 0x0a, 0x03, 0x45,
	0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x61,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x10, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x15, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0xce, 0x04, 0x0a, 0x0b, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x53, 0x65, 0x74, 0x12, 0x7b, 0x0a, 0x16, 0x6e, 0x6f,
	0x6e, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x53, 0x65, 0x74, 0x2e, 0x4e, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x15, 0x6e, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x13, 0x76, 0x6f, 0x6c, 0x61, 0x74,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x53, 0x65, 0x74, 0x2e,
	0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x75, 0x0a, 0x14, 0x68,
	0x65, 0x75, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x53, 0x65, 0x74, 0x2e, 0x48, 0x65, 0x75, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x68,
	0x65, 0x75, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x1a, 0x48, 0x0a, 0x1a, 0x4e, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x45, 0x0a, 0x17,
	0x56, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x48, 0x65, 0x75, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x5d, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6d,
	0x74, 0x75, 0x72, 0x62, 0x6f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x74, 0x6f, 0x42, 0x08, 0x42, 0x61, 0x63, 0x6b, 0x66,
	0x69, 0x6c, 0x6c, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x62, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x2f, 0x74,
	0x75, 0x72, 0x62, 0x6f, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_BackfillMetrics_proto_rawDescOnce sync.Once
	file_BackfillMetrics_proto_rawDescData = file_BackfillMetrics_proto_rawDesc
)

func file_BackfillMetrics_proto_rawDescGZIP() []byte {
	file_BackfillMetrics_proto_rawDescOnce.Do(func() {
		file_BackfillMetrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_BackfillMetrics_proto_rawDescData)
	})
	return file_BackfillMetrics_proto_rawDescData
}

var file_BackfillMetrics_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_BackfillMetrics_proto_goTypes = []interface{}{
	(*BackfilledMetrics)(nil),                      // 0: common_dto.BackfilledMetrics
	(*BackfilledMetrics_Metadata)(nil),             // 1: common_dto.BackfilledMetrics.Metadata
	(*BackfilledMetrics_Data)(nil),                 // 2: common_dto.BackfilledMetrics.Data
	(*BackfilledMetrics_EntityInfo)(nil),           // 3: common_dto.BackfilledMetrics.EntityInfo
	(*BackfilledMetrics_CommoditySoldMetrics)(nil), // 4: common_dto.BackfilledMetrics.CommoditySoldMetrics
	(*BackfilledMetrics_CommodityBought)(nil),      // 5: common_dto.BackfilledMetrics.CommodityBought
	(*BackfilledMetrics_End)(nil),                  // 6: common_dto.BackfilledMetrics.End
	(*BackfilledMetrics_EntityIdSet)(nil),          // 7: common_dto.BackfilledMetrics.EntityIdSet
	nil,                                            // 8: common_dto.BackfilledMetrics.EntityIdSet.NonvolatilePropertiesEntry
	nil,                                            // 9: common_dto.BackfilledMetrics.EntityIdSet.VolatilePropertiesEntry
	nil,                                            // 10: common_dto.BackfilledMetrics.EntityIdSet.HeuristicPropertiesEntry
	(CommodityDTO_CommodityType)(0),                // 11: common_dto.CommodityDTO.CommodityType
	(*CommodityDTO_UtilizationData)(nil),           // 12: common_dto.CommodityDTO.UtilizationData
}
var file_BackfillMetrics_proto_depIdxs = []int32{
	1,  // 0: common_dto.BackfilledMetrics.metadata:type_name -> common_dto.BackfilledMetrics.Metadata
	2,  // 1: common_dto.BackfilledMetrics.data:type_name -> common_dto.BackfilledMetrics.Data
	6,  // 2: common_dto.BackfilledMetrics.end:type_name -> common_dto.BackfilledMetrics.End
	3,  // 3: common_dto.BackfilledMetrics.Data.entity_info:type_name -> common_dto.BackfilledMetrics.EntityInfo
	7,  // 4: common_dto.BackfilledMetrics.EntityInfo.target_entity_id_set:type_name -> common_dto.BackfilledMetrics.EntityIdSet
	4,  // 5: common_dto.BackfilledMetrics.EntityInfo.commodities_sold:type_name -> common_dto.BackfilledMetrics.CommoditySoldMetrics
	5,  // 6: common_dto.BackfilledMetrics.EntityInfo.commodities_bought:type_name -> common_dto.BackfilledMetrics.CommodityBought
	11, // 7: common_dto.BackfilledMetrics.CommoditySoldMetrics.commodity_type:type_name -> common_dto.CommodityDTO.CommodityType
	12, // 8: common_dto.BackfilledMetrics.CommoditySoldMetrics.data:type_name -> common_dto.CommodityDTO.UtilizationData
	4,  // 9: common_dto.BackfilledMetrics.CommodityBought.bought:type_name -> common_dto.BackfilledMetrics.CommoditySoldMetrics
	8,  // 10: common_dto.BackfilledMetrics.EntityIdSet.nonvolatile_properties:type_name -> common_dto.BackfilledMetrics.EntityIdSet.NonvolatilePropertiesEntry
	9,  // 11: common_dto.BackfilledMetrics.EntityIdSet.volatile_properties:type_name -> common_dto.BackfilledMetrics.EntityIdSet.VolatilePropertiesEntry
	10, // 12: common_dto.BackfilledMetrics.EntityIdSet.heuristic_properties:type_name -> common_dto.BackfilledMetrics.EntityIdSet.HeuristicPropertiesEntry
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_BackfillMetrics_proto_init() }
func file_BackfillMetrics_proto_init() {
	if File_BackfillMetrics_proto != nil {
		return
	}
	file_CommonDTO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BackfillMetrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_EntityInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_CommoditySoldMetrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_CommodityBought); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_End); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_BackfillMetrics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackfilledMetrics_EntityIdSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_BackfillMetrics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BackfilledMetrics_Metadata_)(nil),
		(*BackfilledMetrics_Data_)(nil),
		(*BackfilledMetrics_End_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_BackfillMetrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BackfillMetrics_proto_goTypes,
		DependencyIndexes: file_BackfillMetrics_proto_depIdxs,
		MessageInfos:      file_BackfillMetrics_proto_msgTypes,
	}.Build()
	File_BackfillMetrics_proto = out.File
	file_BackfillMetrics_proto_rawDesc = nil
	file_BackfillMetrics_proto_goTypes = nil
	file_BackfillMetrics_proto_depIdxs = nil
}
