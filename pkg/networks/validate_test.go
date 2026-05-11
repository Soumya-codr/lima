// SPDX-FileCopyrightText: Copyright The Lima Authors
// SPDX-License-Identifier: Apache-2.0

package networks

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Paths    Paths
		Group    string
		Networks map[string]Network
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid user-v2",
			fields: fields{
				Paths: Paths{
					SocketVMNet: "/tmp/socket_vmnet",
					VarRun:      "/tmp/lima",
				},
				Networks: map[string]Network{
					"test": {Mode: ModeUserV2},
				},
			},
			wantErr: false,
		},
		{
			name: "valid bridged",
			fields: fields{
				Paths: Paths{
					SocketVMNet: "/tmp/socket_vmnet",
					VarRun:      "/tmp/lima",
				},
				Networks: map[string]Network{
					"test": {Mode: ModeBridged, Interface: "en0"},
				},
			},
			wantErr: false,
		},
		{
			name: "bridged missing interface",
			fields: fields{
				Paths: Paths{
					SocketVMNet: "/tmp/socket_vmnet",
					VarRun:      "/tmp/lima",
				},
				Networks: map[string]Network{
					"test": {Mode: ModeBridged},
				},
			},
			wantErr: true,
		},
		{
			name: "unknown mode",
			fields: fields{
				Paths: Paths{
					SocketVMNet: "/tmp/socket_vmnet",
					VarRun:      "/tmp/lima",
				},
				Networks: map[string]Network{
					"test": {Mode: "unknown"},
				},
			},
			wantErr: true,
		},
		{
			name: "empty mode",
			fields: fields{
				Paths: Paths{
					SocketVMNet: "/tmp/socket_vmnet",
					VarRun:      "/tmp/lima",
				},
				Networks: map[string]Network{
					"test": {Mode: ""},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Paths:    tt.fields.Paths,
				Group:    tt.fields.Group,
				Networks: tt.fields.Networks,
			}
			err := c.Validate()
			if tt.wantErr {
				assert.Assert(t, err != nil, "Validate() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
