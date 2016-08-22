// Code generated by protoc-gen-go.
// source: MediationMessage.proto
// DO NOT EDIT!

package communicator

import (
	"github.com/vmturbo/vmturbo-go-sdk/sdk"

	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Messages, sent from client to server
type MediationClientMessage struct {
	ValidationResponse *ValidationResponse `protobuf:"bytes,2,opt,name=validationResponse" json:"validationResponse,omitempty"`
	DiscoveryResponse  *DiscoveryResponse  `protobuf:"bytes,3,opt,name=discoveryResponse" json:"discoveryResponse,omitempty"`
	KeepAlive          *KeepAlive          `protobuf:"bytes,4,opt,name=keepAlive" json:"keepAlive,omitempty"`
	ActionProgress     *ActionProgress     `protobuf:"bytes,5,opt,name=actionProgress" json:"actionProgress,omitempty"`
	ActionResponse     *ActionResult       `protobuf:"bytes,6,opt,name=actionResponse" json:"actionResponse,omitempty"`
	// this is always required in reality. it's optional here because
	// we don't know if in the future, with embedded targets, we will
	// still use it or not
	MessageID        *int32 `protobuf:"varint,15,opt,name=messageID" json:"messageID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MediationClientMessage) Reset()         { *m = MediationClientMessage{} }
func (m *MediationClientMessage) String() string { return proto.CompactTextString(m) }
func (*MediationClientMessage) ProtoMessage()    {}

func (m *MediationClientMessage) GetValidationResponse() *ValidationResponse {
	if m != nil {
		return m.ValidationResponse
	}
	return nil
}

func (m *MediationClientMessage) GetDiscoveryResponse() *DiscoveryResponse {
	if m != nil {
		return m.DiscoveryResponse
	}
	return nil
}

func (m *MediationClientMessage) GetKeepAlive() *KeepAlive {
	if m != nil {
		return m.KeepAlive
	}
	return nil
}

func (m *MediationClientMessage) GetActionProgress() *ActionProgress {
	if m != nil {
		return m.ActionProgress
	}
	return nil
}

func (m *MediationClientMessage) GetActionResponse() *ActionResult {
	if m != nil {
		return m.ActionResponse
	}
	return nil
}

func (m *MediationClientMessage) GetMessageID() int32 {
	if m != nil && m.MessageID != nil {
		return *m.MessageID
	}
	return 0
}

// Messages, sent from server to client
type MediationServerMessage struct {
	ValidationRequest *ValidationRequest `protobuf:"bytes,2,opt,name=validationRequest" json:"validationRequest,omitempty"`
	DiscoveryRequest  *DiscoveryRequest  `protobuf:"bytes,3,opt,name=discoveryRequest" json:"discoveryRequest,omitempty"`
	ActionRequest     *ActionRequest     `protobuf:"bytes,4,opt,name=actionRequest" json:"actionRequest,omitempty"`
	// this is always required in reality. it's optional here because
	// we don't know if in the future, with embedded targets, we will
	// still use it or not
	MessageID        *int32 `protobuf:"varint,15,opt,name=messageID" json:"messageID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MediationServerMessage) Reset()         { *m = MediationServerMessage{} }
func (m *MediationServerMessage) String() string { return proto.CompactTextString(m) }
func (*MediationServerMessage) ProtoMessage()    {}

func (m *MediationServerMessage) GetValidationRequest() *ValidationRequest {
	if m != nil {
		return m.ValidationRequest
	}
	return nil
}

func (m *MediationServerMessage) GetDiscoveryRequest() *DiscoveryRequest {
	if m != nil {
		return m.DiscoveryRequest
	}
	return nil
}

func (m *MediationServerMessage) GetActionRequest() *ActionRequest {
	if m != nil {
		return m.ActionRequest
	}
	return nil
}

func (m *MediationServerMessage) GetMessageID() int32 {
	if m != nil && m.MessageID != nil {
		return *m.MessageID
	}
	return 0
}

// Structure to hold account parameters, passed to probe to connect and authenticate
// to target.
type AccountValue struct {
	// Name of the parameter. Should refer to the "name" field of AccountDefEntry message,
	// which is returned by the probe in registration phase, for example "userName",
	// "password" and so on.
	Key *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// String representation of the parameter value, for example "secretpassword",
	// "192.168.111.3" and so on.
	StringValue *string `protobuf:"bytes,2,opt,name=stringValue" json:"stringValue,omitempty"`
	// Set of property value lists
	GroupScopePropertyValues []*AccountValue_PropertyValueList `protobuf:"bytes,3,rep,name=groupScopePropertyValues" json:"groupScopePropertyValues,omitempty"`
	XXX_unrecognized         []byte                            `json:"-"`
}

func (m *AccountValue) Reset()         { *m = AccountValue{} }
func (m *AccountValue) String() string { return proto.CompactTextString(m) }
func (*AccountValue) ProtoMessage()    {}

func (m *AccountValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *AccountValue) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *AccountValue) GetGroupScopePropertyValues() []*AccountValue_PropertyValueList {
	if m != nil {
		return m.GroupScopePropertyValues
	}
	return nil
}

type AccountValue_PropertyValueList struct {
	Value            []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AccountValue_PropertyValueList) Reset()         { *m = AccountValue_PropertyValueList{} }
func (m *AccountValue_PropertyValueList) String() string { return proto.CompactTextString(m) }
func (*AccountValue_PropertyValueList) ProtoMessage()    {}

func (m *AccountValue_PropertyValueList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// Request for action to be performed in probe
type ActionRequest struct {
	ProbeType *string `protobuf:"bytes,1,req,name=probeType" json:"probeType,omitempty"`
	// Account values provide data to allow the probe to allow it to connect
	// to the probe
	AccountValue []*AccountValue `protobuf:"bytes,2,rep,name=accountValue" json:"accountValue,omitempty"`
	// An action execution DTO contains one or more action items
	ActionExecutionDTO *sdk.ActionExecutionDTO `protobuf:"bytes,3,req,name=actionExecutionDTO" json:"actionExecutionDTO,omitempty"`
	// For Cross Destination actions (from one target to another) 2 sets of account
	// values are needed
	SecondaryAccountValue []*AccountValue `protobuf:"bytes,4,rep,name=secondaryAccountValue" json:"secondaryAccountValue,omitempty"`
	XXX_unrecognized      []byte          `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}

func (m *ActionRequest) GetProbeType() string {
	if m != nil && m.ProbeType != nil {
		return *m.ProbeType
	}
	return ""
}

func (m *ActionRequest) GetAccountValue() []*AccountValue {
	if m != nil {
		return m.AccountValue
	}
	return nil
}

func (m *ActionRequest) GetActionExecutionDTO() *sdk.ActionExecutionDTO {
	if m != nil {
		return m.ActionExecutionDTO
	}
	return nil
}

func (m *ActionRequest) GetSecondaryAccountValue() []*AccountValue {
	if m != nil {
		return m.SecondaryAccountValue
	}
	return nil
}

// Result of the action execution. It is translated only once
// after action execution is either completed or failed
type ActionResult struct {
	Response         *ActionResponse `protobuf:"bytes,1,req,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ActionResult) Reset()         { *m = ActionResult{} }
func (m *ActionResult) String() string { return proto.CompactTextString(m) }
func (*ActionResult) ProtoMessage()    {}

func (m *ActionResult) GetResponse() *ActionResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// Progress of the currently executed action. Can be send multiple times
// for each action
type ActionProgress struct {
	Response         *ActionResponse `protobuf:"bytes,1,req,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ActionProgress) Reset()         { *m = ActionProgress{} }
func (m *ActionProgress) String() string { return proto.CompactTextString(m) }
func (*ActionProgress) ProtoMessage()    {}

func (m *ActionProgress) GetResponse() *ActionResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// This class holds response information about executing action. It contains: response: the
// ActionResponseState code representing the state of executing action responseDescription: the
// description message notifying detailed information about current status of executing action
type ActionResponse struct {
	// current action state
	ActionResponseState *sdk.ActionResponseState `protobuf:"varint,1,req,name=actionResponseState,enum=common_dto.ActionResponseState" json:"actionResponseState,omitempty"`
	// current action progress (0..100)
	Progress *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	// action state description, for example ("Moving VM...")
	ResponseDescription *string `protobuf:"bytes,3,req,name=responseDescription" json:"responseDescription,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}

func (m *ActionResponse) GetActionResponseState() sdk.ActionResponseState {
	if m != nil && m.ActionResponseState != nil {
		return *m.ActionResponseState
	}
	return sdk.ActionResponseState_PENDING_ACCEPT
}

func (m *ActionResponse) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

func (m *ActionResponse) GetResponseDescription() string {
	if m != nil && m.ResponseDescription != nil {
		return *m.ResponseDescription
	}
	return ""
}

// ContainerInfo message to the Operations Manager server.
// This message passes probe descriptions to the server.
type ContainerInfo struct {
	// Set of ProbeInfo objects, each one will carry information about one of the probe
	// that the container has loaded internally.
	Probes           []*ProbeInfo `protobuf:"bytes,1,rep,name=probes" json:"probes,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ContainerInfo) Reset()         { *m = ContainerInfo{} }
func (m *ContainerInfo) String() string { return proto.CompactTextString(m) }
func (*ContainerInfo) ProtoMessage()    {}

func (m *ContainerInfo) GetProbes() []*ProbeInfo {
	if m != nil {
		return m.Probes
	}
	return nil
}

type KeepAlive struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeepAlive) Reset()         { *m = KeepAlive{} }
func (m *KeepAlive) String() string { return proto.CompactTextString(m) }
func (*KeepAlive) ProtoMessage()    {}

type Ack struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}

type ValidationRequest struct {
	ProbeType        *string         `protobuf:"bytes,1,req,name=probeType" json:"probeType,omitempty"`
	AccountValue     []*AccountValue `protobuf:"bytes,2,rep,name=accountValue" json:"accountValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ValidationRequest) Reset()         { *m = ValidationRequest{} }
func (m *ValidationRequest) String() string { return proto.CompactTextString(m) }
func (*ValidationRequest) ProtoMessage()    {}

func (m *ValidationRequest) GetProbeType() string {
	if m != nil && m.ProbeType != nil {
		return *m.ProbeType
	}
	return ""
}

func (m *ValidationRequest) GetAccountValue() []*AccountValue {
	if m != nil {
		return m.AccountValue
	}
	return nil
}

type DiscoveryRequest struct {
	ProbeType        *string         `protobuf:"bytes,1,req,name=probeType" json:"probeType,omitempty"`
	AccountValue     []*AccountValue `protobuf:"bytes,2,rep,name=accountValue" json:"accountValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}

func (m *DiscoveryRequest) GetProbeType() string {
	if m != nil && m.ProbeType != nil {
		return *m.ProbeType
	}
	return ""
}

func (m *DiscoveryRequest) GetAccountValue() []*AccountValue {
	if m != nil {
		return m.AccountValue
	}
	return nil
}

// The ProbeInfo class provides a description of the probe that enables users to
// attach Operations Manager to a target, and enables the probe to add entities to the
// Operations Manager market as valid members of the supply chain.
//
// To enable users to use this probe, the ProbeInfo includes a probe type and
// the set of fields a user must give to provide credentials and other data necessary to
// attach to a target. The probe type is an arbitrary string, but REST API calls that
// invoke this probe must refer to it by the same type.
//
// To enable adding entities to the Operations Manager market, the ProbeInfo includes a
// set of {@link TemplateDTO} objects called the supplyChainDefinitionSet.
// Each template object describes an entity type that
// the probe can discover and add to the market. This description includes an EntityDTO object
// and its corresponding lists of bought and sold {@code CommodityDTO} objects. As the probe
// discovers entities, it must create instances that map to members of the supplyChainDefinitionSet.
type ProbeInfo struct {
	// probeType is a string identifier to define the type of the probe. You specify the value for this string in the
	// {@code probe-conf.xml} file for your probe. Note that for a given instance of Operations Manager,
	// every probe communicating with the server must have a unique type.
	//
	// The probe you create must include a type
	// to display in the user interface. The string you provide for the probe type appears
	// in the Target Configuration form as one of the choices for the category that you set for the probe.
	//
	// For example, in the standard targets, Hypervisor is a category. If your probe
	// category is also {@code Hypervisor} and you specify a type of 'MyProbe', then MyProbe
	// will appear in the user interface as an additional Hypervisor choice.
	// On the other hand, if the category you provide does not match one of the standard categories,
	// MyProbe will appear as a choice in the CUSTOM category.
	ProbeType *string `protobuf:"bytes,1,req,name=probeType" json:"probeType,omitempty"`
	// String identifier to define the category of the probe. You specify the value for this string in the
	// 'probe-conf.xml' file for your probe.
	//
	// The probe you create must include a category.
	// If the category you provide matches one of the standard categories, then your probe will appear
	// as a choice in the Target Configuration form alongside the other members of that category.
	// For example, in the standard targets, 'Hypervisor' is a category. If your probe
	// category is also 'Hypervisor' and you specify a type of 'MyProbe', then MyProbe
	// will appear in the user interface as an additional 'Hypervisor' choice.
	// On the other hand, if the category you provide does not match one of the standard categories,
	// MyProbe will appear as a choice in the 'CUSTOM' category.
	//
	// The set of standard categories is defined in the 'ProbeCategory' enumeration.
	ProbeCategory *string `protobuf:"bytes,2,req,name=probeCategory" json:"probeCategory,omitempty"`
	// Set of TemplateDTO objects that defines the types of entities the probe discovers, and
	// what their bought and sold commodities are. Any entity instances the probe creates must match
	// members of this set.
	SupplyChainDefinitionSet []*sdk.TemplateDTO `protobuf:"bytes,3,rep,name=supplyChainDefinitionSet" json:"supplyChainDefinitionSet,omitempty"`
	// List of AccountDefEntry objects that describe the fields users provide as
	// input (i.e. ip, user, pass, ...). These fields appear in the Operations Manager user interface
	// when users add targets of this probe's type. REST API calls to add targets also provide data
	// for these fields (i.e. ip, user, password, ...).
	//
	// Order of elements in the list specifyes the order they appear in the UI.
	// List must not contain entries with equal "name" field. This is up to client to ensure this.
	AccountDefinition []*AccountDefEntry `protobuf:"bytes,4,rep,name=accountDefinition" json:"accountDefinition,omitempty"`
	// Specifies the interval at which discoveries will be executed for this probe.
	// The value is specified in seconds. If no value is provided for rediscoveryIntervalSeconds
	// a default of 600 seconds (10 minutes) will be used. The minimum value allowed for this
	// field is 60 seconds (1 minute).
	RediscoveryIntervalSeconds *int32 `protobuf:"varint,5,opt,name=rediscoveryIntervalSeconds" json:"rediscoveryIntervalSeconds,omitempty"`
	XXX_unrecognized           []byte `json:"-"`
}

func (m *ProbeInfo) Reset()         { *m = ProbeInfo{} }
func (m *ProbeInfo) String() string { return proto.CompactTextString(m) }
func (*ProbeInfo) ProtoMessage()    {}

func (m *ProbeInfo) GetProbeType() string {
	if m != nil && m.ProbeType != nil {
		return *m.ProbeType
	}
	return ""
}

func (m *ProbeInfo) GetProbeCategory() string {
	if m != nil && m.ProbeCategory != nil {
		return *m.ProbeCategory
	}
	return ""
}

func (m *ProbeInfo) GetSupplyChainDefinitionSet() []*sdk.TemplateDTO {
	if m != nil {
		return m.SupplyChainDefinitionSet
	}
	return nil
}

func (m *ProbeInfo) GetAccountDefinition() []*AccountDefEntry {
	if m != nil {
		return m.AccountDefinition
	}
	return nil
}

func (m *ProbeInfo) GetRediscoveryIntervalSeconds() int32 {
	if m != nil && m.RediscoveryIntervalSeconds != nil {
		return *m.RediscoveryIntervalSeconds
	}
	return 0
}
