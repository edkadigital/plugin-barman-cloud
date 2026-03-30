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
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	barmancloudv1 "github.com/cloudnative-pg/plugin-barman-cloud/api/v1"
)

func TestSetWALSubmissionTimes(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()
	if err := barmancloudv1.AddToScheme(scheme); err != nil {
		t.Fatalf("add scheme: %v", err)
	}

	firstRecoverabilityPoint := metav1.NewTime(time.Date(2026, 3, 30, 0, 0, 0, 0, time.UTC))
	objectStore := &barmancloudv1.ObjectStore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-store",
			Namespace: "postgres",
		},
		Status: barmancloudv1.ObjectStoreStatus{
			ServerRecoveryWindow: map[string]barmancloudv1.RecoveryWindow{
				"postgres": {
					FirstRecoverabilityPoint: &firstRecoverabilityPoint,
				},
			},
		},
	}

	cli := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(&barmancloudv1.ObjectStore{}).
		WithObjects(objectStore).
		Build()

	firstBatch := []time.Time{
		time.Date(2026, 3, 30, 10, 0, 5, 0, time.UTC),
		time.Date(2026, 3, 30, 10, 0, 1, 0, time.UTC),
		time.Date(2026, 3, 30, 10, 0, 9, 0, time.UTC),
	}
	if err := SetWALSubmissionTimes(
		context.Background(),
		cli,
		client.ObjectKeyFromObject(objectStore),
		"postgres",
		firstBatch,
	); err != nil {
		t.Fatalf("set first batch: %v", err)
	}

	secondBatch := []time.Time{
		time.Date(2026, 3, 30, 9, 59, 58, 0, time.UTC),
		time.Date(2026, 3, 30, 10, 0, 15, 0, time.UTC),
	}
	if err := SetWALSubmissionTimes(
		context.Background(),
		cli,
		client.ObjectKeyFromObject(objectStore),
		"postgres",
		secondBatch,
	); err != nil {
		t.Fatalf("set second batch: %v", err)
	}

	var updated barmancloudv1.ObjectStore
	if err := cli.Get(context.Background(), client.ObjectKeyFromObject(objectStore), &updated); err != nil {
		t.Fatalf("get updated object store: %v", err)
	}

	window := updated.Status.ServerRecoveryWindow["postgres"]
	if window.FirstRecoverabilityPoint == nil || !window.FirstRecoverabilityPoint.Equal(&firstRecoverabilityPoint) {
		t.Fatalf("expected first recoverability point to be preserved, got %#v", window.FirstRecoverabilityPoint)
	}
	if window.FirstWALSubmissionTime == nil ||
		!window.FirstWALSubmissionTime.Time.Equal(time.Date(2026, 3, 30, 9, 59, 58, 0, time.UTC)) {
		t.Fatalf("unexpected first WAL submission time: %#v", window.FirstWALSubmissionTime)
	}
	if window.LastWALSubmissionTime == nil ||
		!window.LastWALSubmissionTime.Time.Equal(time.Date(2026, 3, 30, 10, 0, 15, 0, time.UTC)) {
		t.Fatalf("unexpected last WAL submission time: %#v", window.LastWALSubmissionTime)
	}
}
