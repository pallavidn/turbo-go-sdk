package probe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.ibm.com/turbonomic/turbo-go-sdk/pkg"
)

func TestGetProbeInfoWithDefaultDiscoveryIntervals(t *testing.T) {
	pc, _ := NewProbeConfig("Type1", "Category1", "UICategory1")
	theProbe, _ := newTurboProbe(pc)

	probeInfo, _ := theProbe.GetProbeInfo()
	assert.EqualValues(t, pkg.DEFAULT_FULL_DISCOVERY_IN_SECS, probeInfo.GetFullRediscoveryIntervalSeconds())
	assert.EqualValues(t, 0, probeInfo.GetIncrementalRediscoveryIntervalSeconds())
	assert.EqualValues(t, 0, probeInfo.GetPerformanceRediscoveryIntervalSeconds())
}

func TestGetProbeInfoWithIllegalDiscoveryIntervals(t *testing.T) {
	discoveryMetadata := NewDiscoveryMetadata()
	discoveryMetadata.SetFullRediscoveryIntervalSeconds(10)
	discoveryMetadata.SetIncrementalRediscoveryIntervalSeconds(300)
	discoveryMetadata.SetPerformanceRediscoveryIntervalSeconds(-1)
	pc, _ := NewProbeConfig("Type1", "Category1", "UICategory1")
	pc.SetDiscoveryMetadata(discoveryMetadata)
	version := "foo"
	displayName := "bar"
	pc.WithVersion(version)
	pc.WithDisplayName(displayName)
	theProbe, _ := newTurboProbe(pc)

	probeInfo, _ := theProbe.GetProbeInfo()
	assert.EqualValues(t, pkg.DEFAULT_MIN_DISCOVERY_IN_SECS, probeInfo.GetFullRediscoveryIntervalSeconds())
	assert.EqualValues(t, 300.0, probeInfo.GetIncrementalRediscoveryIntervalSeconds())
	assert.EqualValues(t, 0, probeInfo.GetPerformanceRediscoveryIntervalSeconds())
	assert.EqualValues(t, version, pc.Version)
	assert.EqualValues(t, displayName, pc.DisplayName)
}

func TestNewProbe(t *testing.T) {
	pcTable := []struct {
		probeType string
		category  string
	}{
		{"type1", "category1"},
		{"", "category1"},
		{"type1", ""},
		{"", ""},
	}

	for _, item := range pcTable {
		pc := &ProbeConfig{
			ProbeCategory: item.category,
			ProbeType:     item.probeType,
		}
		theProbe, err := newTurboProbe(pc)

		if pc.Validate() == nil {
			assert.Nil(t, err)
			assert.NotNil(t, theProbe)
			assert.NotNil(t, theProbe.TargetsToAdd)
			assert.NotNil(t, theProbe.ProbeConfiguration)
			assert.NotNil(t, theProbe.RegistrationClient)
			discoveryMetadata := theProbe.ProbeConfiguration.discoveryMetadata
			assert.EqualValues(t, pkg.DEFAULT_FULL_DISCOVERY_IN_SECS, discoveryMetadata.GetFullRediscoveryIntervalSeconds())
			assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED, discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds())
			assert.EqualValues(t, pkg.DISCOVERY_NOT_SUPPORTED, discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds())
		} else {
			assert.NotNil(t, err)
			assert.Nil(t, theProbe)
		}
	}
}
