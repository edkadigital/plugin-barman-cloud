/*
Copyright © contributors to CloudNativePG, established as
CloudNativePG a Series of LF Projects, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package operator

import (
	"context"

	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/http"
	"github.com/cloudnative-pg/cnpg-i/pkg/lifecycle"
	"github.com/cloudnative-pg/cnpg-i/pkg/reconciler"
	"google.golang.org/grpc"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CNPGI is the implementation of the CNPG-i server
type CNPGI struct {
	Client         client.Client
	PluginPath     string
	ServerCertPath string
	ServerKeyPath  string
	ClientCertPath string
	ServerAddress  string
}

// NeedLeaderElection lets every replica serve the CNPG-I endpoints, while the
// ObjectStore reconciler keeps running on the leader only.
//
// The operator reaches the plugin through its Service and loads it before
// every Cluster reconcile, failover included. When only the leader served,
// a standby replica never became ready, so losing the leader's node left the
// Service without an endpoint and the operator without failover until a new
// leader started. The hooks only read the API, return patches, or create and
// patch the Role and RoleBinding a Cluster needs, so any replica can answer.
func (c *CNPGI) NeedLeaderElection() bool {
	return false
}

// Start starts the GRPC server
// of the operator plugin
func (c *CNPGI) Start(ctx context.Context) error {
	enrich := func(server *grpc.Server) error {
		reconciler.RegisterReconcilerHooksServer(server, ReconcilerImplementation{
			Client: c.Client,
		})
		lifecycle.RegisterOperatorLifecycleServer(server, LifecycleImplementation{
			Client: c.Client,
		})
		return nil
	}

	srv := http.Server{
		IdentityImpl:   IdentityImplementation{},
		Enrichers:      []http.ServerEnricher{enrich},
		PluginPath:     c.PluginPath,
		ServerCertPath: c.ServerCertPath,
		ServerKeyPath:  c.ServerKeyPath,
		ClientCertPath: c.ClientCertPath,
		ServerAddress:  c.ServerAddress,
	}

	return srv.Start(ctx)
}
