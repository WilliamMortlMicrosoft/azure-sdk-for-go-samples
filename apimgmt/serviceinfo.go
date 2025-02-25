// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import "context"

// ServiceInfo is used to pass API Management Service configuration information to support functions.
type ServiceInfo struct {
	Ctx               context.Context
	ResourceGroupName string
	ServiceName       string
	Email             string
	Name              string
}
