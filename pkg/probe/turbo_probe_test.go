package probe

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTurboProbe(t *testing.T) {
	newProbe := turboProbe()

	fmt.Printf("%++v\n", newProbe.RegistrationClient.GetActionPolicy())
	assert.NotNil(t, newProbe.RegistrationClient)
	assert.Nil(t, newProbe.RegistrationClient.ISupplyChainProvider)
	assert.Nil(t, newProbe.RegistrationClient.IAccountDefinitionProvider)
	assert.NotNil(t, newProbe.RegistrationClient.IActionPolicyProvider)
	assert.NotEmpty(t, newProbe.RegistrationClient.IActionPolicyProvider.GetActionPolicy())

	assert.NotNil(t, newProbe.RegistrationClient.discoveryMetadata)
	assert.Equal(t, DEFAULT_FULL_DISCOVERY_IN_SECS,
		newProbe.RegistrationClient.discoveryMetadata.GetFullRediscoveryIntervalSeconds())
	assert.Equal(t, DISCOVERY_NOT_SUPPORTED,
		newProbe.RegistrationClient.discoveryMetadata.GetIncrementalRediscoveryIntervalSeconds())
	assert.Equal(t, DISCOVERY_NOT_SUPPORTED,
		newProbe.RegistrationClient.discoveryMetadata.GetPerformanceRediscoveryIntervalSeconds())
}
