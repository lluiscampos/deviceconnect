// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package store

import (
	"context"

	"github.com/mendersoftware/deviceconnect/model"
)

// DataStore interface for DataStore services
type DataStore interface {
	Ping(ctx context.Context) error
	ProvisionTenant(ctx context.Context, tenantID string) error
	ProvisionDevice(ctx context.Context, tenantID string, deviceID string) error
	DeleteDevice(ctx context.Context, tenantID, deviceID string) error
	GetDevice(ctx context.Context, tenantID, deviceID string) (*model.Device, error)
	UpdateDeviceStatus(ctx context.Context, tenantID, deviceID, status string) error
	UpsertSession(ctx context.Context, tenantID, userID, devID string) (*model.Session, error)
	GetSession(ctx context.Context, tenantID, sessionID string) (*model.Session, error)
	UpdateSessionStatus(ctx context.Context, tenantID, sessionID, status string) error
}
