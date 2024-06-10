// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"

	azurenode "github.com/gardener/remedy-controller/pkg/controller/azure/node"
	azurepublicipaddress "github.com/gardener/remedy-controller/pkg/controller/azure/publicipaddress"
	azureservice "github.com/gardener/remedy-controller/pkg/controller/azure/service"
	azurevirtualmachine "github.com/gardener/remedy-controller/pkg/controller/azure/virtualmachine"
)

// ControllerSwitchOptions are the controllercmd.SwitchOptions for the manager controllers.
func ControllerSwitchOptions() *controllercmd.SwitchOptions {
	return controllercmd.NewSwitchOptions(
		controllercmd.Switch(azurepublicipaddress.ControllerName, azurepublicipaddress.AddToManager),
		controllercmd.Switch(azurevirtualmachine.ControllerName, azurevirtualmachine.AddToManager),
	)
}

// TargetControllerSwitchOptions are the controllercmd.SwitchOptions for the target cluster manager controllers.
func TargetControllerSwitchOptions() *controllercmd.SwitchOptions {
	return controllercmd.NewSwitchOptions(
		controllercmd.Switch(azureservice.ControllerName, azureservice.AddToManager),
		controllercmd.Switch(azurenode.ControllerName, azurenode.AddToManager),
	)
}
