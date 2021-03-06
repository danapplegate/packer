// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in builder/azure for license information.

package arm

import (
	"github.com/hashicorp/packer/builder/azure/common"
	"github.com/hashicorp/packer/builder/azure/common/constants"
	"github.com/mitchellh/multistep"
)

func processInterruptibleResult(
	result common.InterruptibleTaskResult, sayError func(error), state multistep.StateBag) multistep.StepAction {
	if result.IsCancelled {
		return multistep.ActionHalt
	}

	return processStepResult(result.Err, sayError, state)
}

func processStepResult(
	err error, sayError func(error), state multistep.StateBag) multistep.StepAction {

	if err != nil {
		state.Put(constants.Error, err)
		sayError(err)

		return multistep.ActionHalt
	}

	return multistep.ActionContinue

}
