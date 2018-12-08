/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	internalinterfaces "github.com/oracle/oci-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DhcpOptions returns a DhcpOptionInformer.
	DhcpOptions() DhcpOptionInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// InternetGatewaies returns a InternetGatewayInformer.
	InternetGatewaies() InternetGatewayInformer
	// RouteTables returns a RouteTableInformer.
	RouteTables() RouteTableInformer
	// SecurityRuleSets returns a SecurityRuleSetInformer.
	SecurityRuleSets() SecurityRuleSetInformer
	// Subnets returns a SubnetInformer.
	Subnets() SubnetInformer
	// Vcns returns a VcnInformer.
	Vcns() VcnInformer
	// Volumes returns a VolumeInformer.
	Volumes() VolumeInformer
	// VolumeBackups returns a VolumeBackupInformer.
	VolumeBackups() VolumeBackupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// DhcpOptions returns a DhcpOptionInformer.
func (v *version) DhcpOptions() DhcpOptionInformer {
	return &dhcpOptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InternetGatewaies returns a InternetGatewayInformer.
func (v *version) InternetGatewaies() InternetGatewayInformer {
	return &internetGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteTables returns a RouteTableInformer.
func (v *version) RouteTables() RouteTableInformer {
	return &routeTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityRuleSets returns a SecurityRuleSetInformer.
func (v *version) SecurityRuleSets() SecurityRuleSetInformer {
	return &securityRuleSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Subnets returns a SubnetInformer.
func (v *version) Subnets() SubnetInformer {
	return &subnetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Vcns returns a VcnInformer.
func (v *version) Vcns() VcnInformer {
	return &vcnInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Volumes returns a VolumeInformer.
func (v *version) Volumes() VolumeInformer {
	return &volumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeBackups returns a VolumeBackupInformer.
func (v *version) VolumeBackups() VolumeBackupInformer {
	return &volumeBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
