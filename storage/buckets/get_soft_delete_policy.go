// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buckets

// [START storage_get_soft_delete_policy]
import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

func getSoftDeletePolicy(w io.Writer, bucketName string) error {
	// bucketName := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucket := client.Bucket(bucketName)

	// Get the current bucket attributes.
	attrs, err := bucket.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("bucket.Attrs: %w", err)
	}

	// Check if soft delete policy is set.
	softDeletePolicy := attrs.SoftDeletePolicy
	if softDeletePolicy == nil {
		fmt.Printf("Bucket %s does not have a soft delete policy set.\n", bucketName)
		return nil
	}
	if softDeletePolicy.RetentionDuration == 0 {
		fmt.Printf("Soft delete is disabled for bucket %s.\n", bucketName)
		return nil
	}

	fmt.Fprintf(w, "Soft delete policy for bucket %s is:\n EffectiveTime: %s\n RetentionDuration: %s\n", bucketName, softDeletePolicy.EffectiveTime.String(), softDeletePolicy.RetentionDuration.String())
	return nil
}

// [START storage_get_soft_delete_policy]
