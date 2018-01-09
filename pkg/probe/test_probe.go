package probe

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type TestProbe struct{}
type TestProbeDiscoveryClient struct{}
type TestProbeRegistrationClient struct{}
type TestProbeActionClient struct{}

type TestProbeDiscoveryClientComplete struct{}

func (handler *TestProbeDiscoveryClient) GetAccountValues() *TurboTargetInfo {
	return nil
}
func (handler *TestProbeDiscoveryClient) Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error) {
	return nil, fmt.Errorf("TestProbeDiscoveryClient Validate not implemented")
}

func (handler *TestProbeDiscoveryClient) Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	return fakeDiscoveryResponse(), nil //.Errorf("TestProbeDiscoveryClient Discover not implemented")
}

func (handler *TestProbeDiscoveryClientComplete) GetAccountValues() *TurboTargetInfo {
	return nil
}
func (handler *TestProbeDiscoveryClientComplete) Validate(accountValues []*proto.AccountValue) (*proto.ValidationResponse, error) {
	return nil, fmt.Errorf("TestProbeDiscoveryClient Validate not implemented")
}
func (handler *TestProbeDiscoveryClientComplete) Discover(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	return fakeDiscoveryResponse(), nil
}

func (handler *TestProbeDiscoveryClientComplete) DiscoverIncremental(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	return fakeDiscoveryResponse(), nil
}

func (handler *TestProbeDiscoveryClientComplete) DiscoverPerformance(accountValues []*proto.AccountValue) (*proto.DiscoveryResponse, error) {
	return fakeDiscoveryResponse(), nil
}

func fakeDiscoveryResponse() *proto.DiscoveryResponse {
	var entityDtos []*proto.EntityDTO
	eType := proto.EntityDTO_VIRTUAL_MACHINE
	id := "Id1"
	dName := "Entity1"
	entityDto := &proto.EntityDTO{
		EntityType:  &eType,
		Id:          &id,
		DisplayName: &dName,
	}
	entityDtos = append(entityDtos, entityDto)
	discoveryResponse := &proto.DiscoveryResponse{
		EntityDTO: entityDtos,
	}
	return discoveryResponse
}

func (registrationClient *TestProbeRegistrationClient) GetSupplyChainDefinition() []*proto.TemplateDTO {
	return nil
}
func (registrationClient *TestProbeRegistrationClient) GetAccountDefinition() []*proto.AccountDefEntry {
	return nil
}
func (registrationClient *TestProbeRegistrationClient) GetIdentifyingFields() string {
	return ""
}

func (actionClient *TestProbeActionClient) ExecuteAction(actionExecutionDTO *proto.ActionExecutionDTO,
	accountValues []*proto.AccountValue,
	progressTracker ActionProgressTracker) (*proto.ActionResult, error) {

	return nil, fmt.Errorf("TestProbeDiscoveryClient ExecuteAction not implemented")
}
