package probe

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
)

type ProbeBuilder struct {
	probeConf              *ProbeConfig
	registrationClient     TurboRegistrationClient
	actionClient           TurboActionExecutorClient
	builderError           error
	supplyChainProvider    ISupplyChainProvider
	accountDefProvider     IAccountDefinitionProvider
	actionPolicyProvider   IActionPolicyProvider
	entityMetadataProvider IEntityMetadataProvider

	discoveryClient *TargetDiscoveryAgent
	discoveryTarget string
}

func ErrorInvalidTargetIdentifier() error {
	return errors.New("Null Target Identifier")
}

func ErrorInvalidProbeConfig() error {
	return errors.New("Null Probe configuration")
}

func ErrorInvalidProbeType() error {
	return errors.New("Null Probe type")
}

func ErrorInvalidProbeCategory() error {
	return errors.New("Null Probe category")
}

func ErrorInvalidRegistrationClient() error {
	return errors.New("Null registration client")
}

func ErrorInvalidActionClient() error {
	return errors.New("Null action client")
}

func ErrorInvalidDiscoveryClient(targetId string, discoveryType string) error {
	return errors.New(fmt.Sprintf("Invalid %s discovery client for target [%s]", discoveryType, targetId))
}

func ErrorDiscoveryClient(probeCategory, probeType string, discoveryType string, errMsg string) error {
	return errors.New(fmt.Sprintf("Invalid %s discovery client for [%s:%s] : %s", discoveryType, probeCategory, probeType, errMsg))
}

func ErrorUndefinedDiscoveryClient() error {
	return errors.New("No discovery clients defined")
}

func ErrorCreatingProbe(probeCategory, probeType string) error {
	return errors.New("Error creating probe for " + probeCategory + "::" + probeType)
}

// Create an instance of ProbeBuilder using the given probe configuration
func NewProbeBuilderWithConfig(probeConf *ProbeConfig) (*ProbeBuilder, error) {
	if probeConf == nil {
		return nil, ErrorInvalidProbeConfig()
	}

	if probeConf.ProbeType == "" {
		return nil, ErrorInvalidProbeType()
	}

	if probeConf.ProbeCategory == "" {
		return nil, ErrorInvalidProbeCategory()
	}

	return &ProbeBuilder{
		probeConf: probeConf,
	}, nil
}

// Get an instance of ProbeBuilder
func NewProbeBuilder(probeType string, probeCategory string) *ProbeBuilder {
	// Validate probe type and category
	probeBuilder := &ProbeBuilder{}
	if probeType == "" {
		probeBuilder.builderError = ErrorInvalidProbeType()
		return probeBuilder
	}

	if probeCategory == "" {
		probeBuilder.builderError = ErrorInvalidProbeCategory()
		return probeBuilder
	}

	probeConf := &ProbeConfig{
		ProbeCategory:        probeCategory,
		ProbeType:            probeType,
		FullDiscovery:        DEFAULT_FULL_DISCOVERY_IN_SECS,
		IncrementalDiscovery: DISCOVERY_NOT_SUPPORTED,
		PerformanceDiscovery: DISCOVERY_NOT_SUPPORTED,
	}

	return &ProbeBuilder{
		probeConf: probeConf,
	}
}

// Build an instance of TurboProbe.
func (pb *ProbeBuilder) Create() (*TurboProbe, error) {
	if pb.builderError != nil {
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	if pb.discoveryClient == nil {
		pb.builderError = ErrorUndefinedDiscoveryClient()
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}
	targetDiscoveryAgent := pb.discoveryClient
	probeConf := pb.probeConf

	// validate the discovery intervals
	discoveryMetadata := NewDiscoveryMetadata()
	discoveryMetadata.SetFullRediscoveryIntervalSeconds(checkFullRediscoveryInterval(probeConf.FullDiscovery))
	discoveryMetadata.SetIncrementalRediscoveryIntervalSeconds(checkIncrementalRediscoveryInterval(probeConf.IncrementalDiscovery))
	discoveryMetadata.SetPerformanceRediscoveryIntervalSeconds(checkPerformanceRediscoveryInterval(probeConf.PerformanceDiscovery))

	if (discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds() >= DEFAULT_MIN_DISCOVERY_IN_SECS) && (targetDiscoveryAgent.IIncrementalDiscovery == nil) {
		pb.builderError = ErrorDiscoveryClient(probeConf.ProbeCategory, probeConf.ProbeType,
			"incremental", "Incremental rediscovery interval is "+
				"declared but IIncrementalDiscoveryProbe interface is not implemented.")
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	if (discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds() >= DEFAULT_MIN_DISCOVERY_IN_SECS) && (targetDiscoveryAgent.IPerformanceDiscovery == nil) {
		pb.builderError = ErrorDiscoveryClient(probeConf.ProbeCategory, probeConf.ProbeType,
			"performance", "Performance rediscovery interval is "+
				"declared but IPerformanceDiscovery interface is not implemented.")
		glog.Errorf(pb.builderError.Error())
		return nil, pb.builderError
	}

	checkRediscoveryIntervalValidity(pb.probeConf.FullDiscovery,
		pb.probeConf.IncrementalDiscovery,
		pb.probeConf.PerformanceDiscovery)

	turboProbe := turboProbe() //probeConf)
	turboProbe.ProbeConfiguration = probeConf
	turboProbe.RegistrationClient.discoveryMetadata = discoveryMetadata
	turboProbe.RegistrationClient.ISupplyChainProvider = pb.registrationClient
	turboProbe.RegistrationClient.IAccountDefinitionProvider = pb.registrationClient

	if pb.supplyChainProvider != nil {
		turboProbe.RegistrationClient.ISupplyChainProvider = pb.supplyChainProvider
	}

	if pb.accountDefProvider != nil {
		turboProbe.RegistrationClient.IAccountDefinitionProvider = pb.accountDefProvider
	}

	if pb.actionPolicyProvider != nil {
		turboProbe.RegistrationClient.IActionPolicyProvider = pb.actionPolicyProvider
	}

	turboProbe.DiscoveryClientMap[pb.discoveryTarget] = targetDiscoveryAgent //discoveryClient
	turboProbe.ActionClient = pb.actionClient

	return turboProbe, nil
}

func checkFullRediscoveryInterval(rediscoveryIntervalSec int32) int32 {
	if rediscoveryIntervalSec <= 0 {
		glog.V(3).Infof("No rediscovery interval specified. Using a default value of %d seconds", DEFAULT_FULL_DISCOVERY_IN_SECS)
		return DEFAULT_FULL_DISCOVERY_IN_SECS
	}

	if rediscoveryIntervalSec < DEFAULT_MIN_DISCOVERY_IN_SECS {
		glog.Warning("Rediscovery interval value of %d is below minimum value allowed."+
			" Setting full rediscovery interval to minimum allowed value of %d seconds.",
			rediscoveryIntervalSec, DEFAULT_MIN_DISCOVERY_IN_SECS)
		return DEFAULT_MIN_DISCOVERY_IN_SECS
	}

	return rediscoveryIntervalSec
}

func checkIncrementalRediscoveryInterval(incrementalDiscoverySec int32) int32 {
	if incrementalDiscoverySec <= 0 {
		glog.V(3).Infof("Incremental discovery is not supported.")
		return DISCOVERY_NOT_SUPPORTED
	}

	if incrementalDiscoverySec < DEFAULT_MIN_DISCOVERY_IN_SECS {
		glog.Warning("Incremental interval value of %d is below minimum value allowed."+
			" Setting incremental rediscovery interval to minimum allowed value of %d seconds.",
			incrementalDiscoverySec, DEFAULT_MIN_DISCOVERY_IN_SECS)
		return DEFAULT_MIN_DISCOVERY_IN_SECS
	}

	return incrementalDiscoverySec
}

func checkPerformanceRediscoveryInterval(performanceDiscoverySec int32) int32 {
	if performanceDiscoverySec <= 0 {
		glog.V(3).Infof("Performance discovery is not supported.")
		return DISCOVERY_NOT_SUPPORTED
	}

	if performanceDiscoverySec < DEFAULT_MIN_DISCOVERY_IN_SECS {
		glog.Warning("Incremental interval value of %d is below minimum value allowed."+
			" Setting performance rediscovery interval to minimum allowed value of %d seconds.",
			performanceDiscoverySec, DEFAULT_MIN_DISCOVERY_IN_SECS)
		return DEFAULT_MIN_DISCOVERY_IN_SECS
	}

	return performanceDiscoverySec
}

func checkRediscoveryIntervalValidity(rediscoveryIntervalSec, incrementalDiscoverySec, performanceDiscoverySec int32) {

	if performanceDiscoverySec >= rediscoveryIntervalSec {
		glog.Warning(discoveryConfigError("performance", "full"))
	}

	if incrementalDiscoverySec >= rediscoveryIntervalSec {
		glog.Warning(discoveryConfigError("incremental", "full"))
	}

	if incrementalDiscoverySec >= performanceDiscoverySec && performanceDiscoverySec != DISCOVERY_NOT_SUPPORTED {
		glog.Warning(discoveryConfigError("incremental", "performance"))
	}
}

func discoveryConfigError(discoveryType1, discoveryType2 string) string {
	return fmt.Sprintf("%s rediscovery interval is greater than %s rediscovery interval, will be skipped!",
		discoveryType1, discoveryType2)
}

// Set the supply chain provider for the probe
func (pb *ProbeBuilder) WithSupplyChain(supplyChainProvider ISupplyChainProvider) *ProbeBuilder {
	if supplyChainProvider == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.supplyChainProvider = supplyChainProvider

	return pb
}

// Set the provider that for the account definition for discovering the probe targets
func (pb *ProbeBuilder) WithAccountDef(accountDefProvider IAccountDefinitionProvider) *ProbeBuilder {
	if accountDefProvider == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.accountDefProvider = accountDefProvider

	return pb
}

// Set the provider for the policies regarding the supported action types
func (pb *ProbeBuilder) WithActionPolicies(actionPolicyProvider IActionPolicyProvider) *ProbeBuilder {
	if actionPolicyProvider == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.actionPolicyProvider = actionPolicyProvider

	return pb
}

// Set the provider for the metadata for generating unique identifiers for the probe entities
func (pb *ProbeBuilder) WithEntityMetadata(entityMetadataProvider IEntityMetadataProvider) *ProbeBuilder {
	if entityMetadataProvider == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.entityMetadataProvider = entityMetadataProvider

	return pb
}

// Set the registration client for the probe
func (pb *ProbeBuilder) RegisteredBy(registrationClient TurboRegistrationClient) *ProbeBuilder {
	if registrationClient == nil {
		pb.builderError = ErrorInvalidRegistrationClient()
		return pb
	}
	pb.registrationClient = registrationClient

	return pb
}

// Set a probe target and its discovery client //TODO: to remove
func (pb *ProbeBuilder) DiscoversTarget(targetId string, discoveryClient TurboDiscoveryClient) *ProbeBuilder {
	if targetId == "" {
		pb.builderError = ErrorInvalidTargetIdentifier()
		return pb
	}
	if discoveryClient == nil {
		pb.builderError = ErrorInvalidDiscoveryClient(targetId, "full")
		return pb
	}

	//pb.discoveryClientMap[targetId] = discoveryClient

	pb.discoveryTarget = targetId
	targetDiscoveryAgent := NewTargetDiscoveryAgent(targetId)
	targetDiscoveryAgent.TurboDiscoveryClient = discoveryClient
	pb.discoveryClient = targetDiscoveryAgent

	return pb
}

// Set the incremental discovery client for the probe target
func (pb *ProbeBuilder) DiscoveryBy(discoveryClient *TargetDiscoveryAgent) *ProbeBuilder {
	if discoveryClient.TargetId == "" {
		pb.builderError = ErrorInvalidTargetIdentifier()
		return pb
	}
	if discoveryClient.TurboDiscoveryClient == nil {
		pb.builderError = ErrorInvalidDiscoveryClient(discoveryClient.TargetId, "full")
		return pb
	}

	pb.discoveryTarget = discoveryClient.TargetId
	pb.discoveryClient = discoveryClient
	return pb
}

// Set the action client for the probe
func (pb *ProbeBuilder) ExecutesActionsBy(actionClient TurboActionExecutorClient) *ProbeBuilder {
	if actionClient == nil {
		pb.builderError = ErrorInvalidActionClient()
		return pb
	}
	pb.actionClient = actionClient

	return pb
}
