package probe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewProbeBuilder(t *testing.T) {
	probeBuilder, err := NewProbeBuilderWithConfig(nil)
	assert.Nil(t, probeBuilder)
	assert.EqualValues(t, ErrorInvalidProbeConfig(), err)

	probeConfig := &ProbeConfig{
		ProbeCategory: "",
		ProbeType:     "TYPE1",
	}
	probeBuilder, err = NewProbeBuilderWithConfig(probeConfig)
	assert.Nil(t, probeBuilder)
	assert.EqualValues(t, ErrorInvalidProbeCategory(), err)

	probeConfig = &ProbeConfig{
		ProbeCategory: "CAT1",
		ProbeType:     "",
	}
	probeBuilder, err = NewProbeBuilderWithConfig(probeConfig)
	assert.Nil(t, probeBuilder)
	assert.EqualValues(t, ErrorInvalidProbeType(), err)
}

func TestProbeBuilderCreate(t *testing.T) {
	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(0, 0))
	assert.Nil(t, err)

	discoveryClient := testDiscoveryClient()
	targetId := discoveryClient.TargetId
	probeBuilder.DiscoveryBy(discoveryClient)

	probe, err := probeBuilder.Create()
	assert.Nil(t, err)
	assert.NotNil(t, probe)

	assert.NotNil(t, probe.RegistrationClient.discoveryMetadata)
	discoveryMetadata := probe.RegistrationClient.discoveryMetadata
	assert.EqualValues(t, DEFAULT_FULL_DISCOVERY_IN_SECS, discoveryMetadata.GetFullRediscoveryIntervalSeconds())
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds())
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds())

	assert.NotNil(t, probe.RegistrationClient)
	assert.Nil(t, probe.RegistrationClient.ISupplyChainProvider)
	assert.Nil(t, probe.RegistrationClient.IAccountDefinitionProvider)
	assert.NotNil(t, probe.RegistrationClient.IActionPolicyProvider)
	assert.NotEmpty(t, probe.RegistrationClient.IActionPolicyProvider.GetActionPolicy())

	discoveryAgent := probe.DiscoveryClientMap[targetId]
	assert.NotNil(t, discoveryAgent)
	assert.NotNil(t, discoveryAgent.TurboDiscoveryClient)
	assert.Nil(t, discoveryAgent.IIncrementalDiscovery)
	assert.Nil(t, discoveryAgent.IPerformanceDiscovery)
}

func TestProbeBuilderCreateInvalidDiscoveryConfig(t *testing.T) {
	discoveryClient := testDiscoveryClient()

	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(100, 0))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probe, err := probeBuilder.Create()
	assert.NotNil(t, err)
	assert.Nil(t, probe)

	probeBuilder, err = NewProbeBuilderWithConfig(testProbeConfig(0, 100))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probe, err = probeBuilder.Create()
	assert.NotNil(t, err)
	assert.Nil(t, probe)
}

func TestProbeBuilderCreateWithIncrementalDiscovery(t *testing.T) {
	discoveryClient := testDiscoveryClientComplete()
	targetId := discoveryClient.TargetId

	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(100, 0))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probe, err := probeBuilder.Create()

	assert.Nil(t, err)
	assert.NotNil(t, probe)

	discoveryAgent := probe.DiscoveryClientMap[targetId]
	assert.NotNil(t, discoveryAgent)
	assert.NotNil(t, discoveryAgent.IIncrementalDiscovery)
	assert.NotNil(t, discoveryAgent.TurboDiscoveryClient)
}

func TestProbeBuilderCreateWithPerformanceDiscovery(t *testing.T) {
	discoveryClient := testDiscoveryClientComplete()
	targetId := discoveryClient.TargetId

	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(0, 100))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)

	probe, err := probeBuilder.Create()

	assert.Nil(t, err)
	assert.NotNil(t, probe)

	discoveryAgent := probe.DiscoveryClientMap[targetId]
	assert.NotNil(t, discoveryAgent)
	assert.NotNil(t, discoveryAgent.TurboDiscoveryClient)
	assert.NotNil(t, discoveryAgent.IPerformanceDiscovery)
}

func TestValidateDiscoveryIntervals(t *testing.T) {
	fullDiscoveryInterval := checkFullRediscoveryInterval(0)
	assert.EqualValues(t, DEFAULT_FULL_DISCOVERY_IN_SECS, fullDiscoveryInterval)

	fullDiscoveryInterval = checkFullRediscoveryInterval(-1)
	assert.EqualValues(t, DEFAULT_FULL_DISCOVERY_IN_SECS, fullDiscoveryInterval)

	fullDiscoveryInterval = checkFullRediscoveryInterval(10)
	assert.EqualValues(t, DEFAULT_MIN_DISCOVERY_IN_SECS, fullDiscoveryInterval)

	fullDiscoveryInterval = checkFullRediscoveryInterval(100)
	assert.EqualValues(t, 100, fullDiscoveryInterval)

	fullDiscoveryInterval = checkFullRediscoveryInterval(900)
	assert.EqualValues(t, 900, fullDiscoveryInterval)

	incrDiscoveryInterval := checkIncrementalRediscoveryInterval(0)
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, incrDiscoveryInterval)

	incrDiscoveryInterval = checkIncrementalRediscoveryInterval(-1)
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, incrDiscoveryInterval)

	incrDiscoveryInterval = checkIncrementalRediscoveryInterval(10)
	assert.EqualValues(t, DEFAULT_MIN_DISCOVERY_IN_SECS, incrDiscoveryInterval)

	incrDiscoveryInterval = checkIncrementalRediscoveryInterval(100)
	assert.EqualValues(t, 100, incrDiscoveryInterval)

	incrDiscoveryInterval = checkIncrementalRediscoveryInterval(900)
	assert.EqualValues(t, 900, incrDiscoveryInterval)

	perfDiscoveryInterval := checkPerformanceRediscoveryInterval(0)
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, perfDiscoveryInterval)

	perfDiscoveryInterval = checkPerformanceRediscoveryInterval(-1)
	assert.EqualValues(t, DISCOVERY_NOT_SUPPORTED, perfDiscoveryInterval)

	perfDiscoveryInterval = checkPerformanceRediscoveryInterval(10)
	assert.EqualValues(t, DEFAULT_MIN_DISCOVERY_IN_SECS, perfDiscoveryInterval)

	perfDiscoveryInterval = checkPerformanceRediscoveryInterval(100)
	assert.EqualValues(t, 100, perfDiscoveryInterval)

	perfDiscoveryInterval = checkPerformanceRediscoveryInterval(900)
	assert.EqualValues(t, 900, perfDiscoveryInterval)
}

func TestProbeBuilderRegistrationClient(t *testing.T) {
	discoveryClient := testDiscoveryClient()

	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(-1, 0))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probeBuilder.WithSupplyChain(&TestProbeRegistrationClient{})
	probeBuilder.WithAccountDef(&TestProbeRegistrationClient{})

	probe, err := probeBuilder.Create()

	assert.Nil(t, err)
	assert.NotNil(t, probe)

	assert.NotNil(t, probe.RegistrationClient)
	assert.NotNil(t, probe.RegistrationClient.discoveryMetadata)

	assert.NotNil(t, probe.RegistrationClient.ISupplyChainProvider)
	assert.NotNil(t, probe.RegistrationClient.IAccountDefinitionProvider)
	assert.NotNil(t, probe.RegistrationClient.IActionPolicyProvider)
	assert.NotEmpty(t, probe.RegistrationClient.IActionPolicyProvider.GetActionPolicy())

	probeBuilder, err = NewProbeBuilderWithConfig(testProbeConfig(-1, 0))
	assert.Nil(t, err)

	probeBuilder.WithSupplyChain(nil)
	probe, err = probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)

	probeBuilder.WithAccountDef(nil)
	probe, err = probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)

	probeBuilder.WithActionPolicies(nil)
	probe, err = probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)

	probeBuilder.WithEntityMetadata(nil)
	probe, err = probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)
}

func TestProbeBuilderDiscoveryClient(t *testing.T) {
	discoveryClient := NewTargetDiscoveryAgent("")
	discoveryClient.TurboDiscoveryClient = &TestProbeDiscoveryClient{}

	probeBuilder, err := NewProbeBuilderWithConfig(testProbeConfig(-1, 0))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probe, err := probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)

	discoveryClient = NewTargetDiscoveryAgent("target1")

	probeBuilder, err = NewProbeBuilderWithConfig(testProbeConfig(-1, 0))
	assert.Nil(t, err)
	probeBuilder.DiscoveryBy(discoveryClient)
	probe, err = probeBuilder.Create()
	assert.Nil(t, probe)
	assert.NotNil(t, err)
}

func TestNewProbeBuilderWithoutRegistrationClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"

	probe, err := createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidRegistrationClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderWithoutDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetID := "T1"

	registrationClient := &TestProbeRegistrationClient{}

	probe, err := createProbe(probeType, probeCat, registrationClient, targetID, nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorUndefinedDiscoveryClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderWithRegistrationAndDiscoveryClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := "T1"

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	builder := NewProbeBuilder(probeType, probeCat)
	builder.RegisteredBy(registrationClient)
	builder.DiscoversTarget(targetId, discoveryClient)
	probe, err := builder.Create()

	if !reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, err)
	}

	if !reflect.DeepEqual(registrationClient.GetSupplyChainDefinition(), probe.RegistrationClient.GetSupplyChainDefinition()) {
		t.Errorf("\nExpected %+v, \ngot      %+v", registrationClient, probe.RegistrationClient)
	}
	if !reflect.DeepEqual(registrationClient.GetAccountDefinition(), probe.RegistrationClient.GetAccountDefinition()) {
		t.Errorf("\nExpected %+v, \ngot      %+v", registrationClient, probe.RegistrationClient)
	}

	dc := probe.getDiscoveryClient(targetId)
	if !reflect.DeepEqual(discoveryClient, dc) {
		t.Errorf("\nExpected %+v, \ngot      %+v", discoveryClient, dc)
	}
}

func TestNewProbeBuilderWithActionClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	actionClient := &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}
	builder.ExecutesActionsBy(actionClient)
	probe, err := builder.Create()

	if !reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, err)
	}
	if !reflect.DeepEqual(actionClient, probe.ActionClient) {
		t.Errorf("\nExpected %+v, \ngot      %+v", actionClient, probe.ActionClient)
	}
}

func TestNewProbeBuilderWithInvalidActionClient(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"
	targetId := "T1"
	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	var actionClient TurboActionExecutorClient //:= &TestProbeActionClient{}
	builder := NewProbeBuilder(probeType, probeCat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}
	builder.ExecutesActionsBy(actionClient)
	probe, err := builder.Create()

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidActionClient(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func TestNewProbeBuilderInvalidTargetId(t *testing.T) {
	probeType := "Type1"
	probeCat := "Cloud"

	registrationClient := &TestProbeRegistrationClient{}
	discoveryClient := &TestProbeDiscoveryClient{}
	probe, err := createProbe(probeType, probeCat, registrationClient, "", discoveryClient)
	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidTargetIdentifier(), err)
	}
	var expected *TurboProbe
	if !reflect.DeepEqual(expected, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

}

func TestNewProbeBuilderInvalidProbeType(t *testing.T) {
	probeType := ""
	probeCat := "Cloud"

	probe, err := createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeType(), err)
	}
	var expected1 *TurboProbe
	if !reflect.DeepEqual(expected1, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}

	probeType = "Type1"
	probeCat = ""

	probe, err = createProbe(probeType, probeCat, nil, "", nil)

	if reflect.DeepEqual(nil, err) {
		t.Errorf("\nExpected %+v, \ngot      %+v", ErrorInvalidProbeCategory(), err)
	}
	var expected2 *TurboProbe
	if !reflect.DeepEqual(expected2, probe) {
		t.Errorf("\nExpected %+v, \ngot      %+v", nil, probe)
	}
}

func createProbe(probeType, probeCat string,
	registrationClient TurboRegistrationClient,
	targetId string, discoveryClient TurboDiscoveryClient) (*TurboProbe, error) {

	builder := NewProbeBuilder(probeType, probeCat)

	if registrationClient != nil {
		builder.RegisteredBy(registrationClient)
	}

	if targetId != "" || discoveryClient != nil {
		builder.DiscoversTarget(targetId, discoveryClient)
	}

	return builder.Create()
}

func testDiscoveryClient() *TargetDiscoveryAgent {
	targetId := "Target1"
	discoveryClient := NewTargetDiscoveryAgent(targetId) //&TestProbeDiscoveryClient{}
	discoveryClient.TurboDiscoveryClient = &TestProbeDiscoveryClient{}
	return discoveryClient
}

func testDiscoveryClientComplete() *TargetDiscoveryAgent {
	targetId := "Target1"
	discoveryClient := NewTargetDiscoveryAgent(targetId) //&TestProbeDiscoveryClient{}
	discoveryClient.TurboDiscoveryClient = &TestProbeDiscoveryClientComplete{}
	discoveryClient.IIncrementalDiscovery = &TestProbeDiscoveryClientComplete{}
	discoveryClient.IPerformanceDiscovery = &TestProbeDiscoveryClientComplete{}
	return discoveryClient
}

func testProbeConfig(IncrementalDiscovery, PerformanceDiscovery int32) *ProbeConfig {
	probeConfig := &ProbeConfig{
		ProbeCategory:        "CAT1",
		ProbeType:            "TYPE1",
		FullDiscovery:        DEFAULT_FULL_DISCOVERY_IN_SECS,
		IncrementalDiscovery: IncrementalDiscovery,
		PerformanceDiscovery: PerformanceDiscovery,
	}
	return probeConfig
}
