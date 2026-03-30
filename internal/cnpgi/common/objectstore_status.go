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

package common

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	barmancloudv1 "github.com/cloudnative-pg/plugin-barman-cloud/api/v1"
)

// SetWALSubmissionTimes stores the earliest and latest WAL archive
// submission timestamps observed for a server.
func SetWALSubmissionTimes(
	ctx context.Context,
	c client.Client,
	objectStoreKey client.ObjectKey,
	serverName string,
	submissionTimes []time.Time,
) error {
	var earliest time.Time
	var latest time.Time
	var hasSubmissionTime bool

	for _, submissionTime := range submissionTimes {
		if submissionTime.IsZero() {
			continue
		}

		if !hasSubmissionTime {
			earliest = submissionTime
			latest = submissionTime
			hasSubmissionTime = true
			continue
		}

		if submissionTime.Before(earliest) {
			earliest = submissionTime
		}
		if submissionTime.After(latest) {
			latest = submissionTime
		}
	}

	if !hasSubmissionTime {
		return nil
	}

	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		var objectStore barmancloudv1.ObjectStore

		if err := c.Get(ctx, objectStoreKey, &objectStore); err != nil {
			return err
		}

		recoveryWindow := objectStore.Status.ServerRecoveryWindow[serverName]
		if recoveryWindow.FirstWALSubmissionTime == nil ||
			earliest.Before(recoveryWindow.FirstWALSubmissionTime.Time) {
			recoveryWindow.FirstWALSubmissionTime = ptr.To(metav1.NewTime(earliest))
		}
		if recoveryWindow.LastWALSubmissionTime == nil ||
			latest.After(recoveryWindow.LastWALSubmissionTime.Time) {
			recoveryWindow.LastWALSubmissionTime = ptr.To(metav1.NewTime(latest))
		}

		if objectStore.Status.ServerRecoveryWindow == nil {
			objectStore.Status.ServerRecoveryWindow = make(map[string]barmancloudv1.RecoveryWindow)
		}
		objectStore.Status.ServerRecoveryWindow[serverName] = recoveryWindow

		return c.Status().Update(ctx, &objectStore)
	})
}
