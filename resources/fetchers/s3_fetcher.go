// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aquasecurity/defsec/pkg/rego/convert"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/providers/awslib"
	"github.com/elastic/elastic-agent-libs/logp"
	"reflect"
)

type S3Fetcher struct {
	log        *logp.Logger
	cfg        S3FetcherConfig
	s3Provider awslib.S3BucketDescriber
	resourceCh chan fetching.ResourceInfo
}

type S3FetcherConfig struct {
	fetching.AwsBaseFetcherConfig `config:",inline"`
}

type S3Resource struct {
	bucket awslib.S3BucketDescription
}

func (f *S3Fetcher) Fetch(ctx context.Context, cMetadata fetching.CycleMetadata) error {
	f.log.Info("Starting S3Fetcher.Fetch")

	buckets, err := f.s3Provider.DescribeS3Buckets(ctx, f.log)
	if err != nil {
		return fmt.Errorf("failed to load buckets from S3 %w", err)
	}

	for _, bucket := range buckets {
		resource := S3Resource{awslib.S3BucketDescription(bucket)}
		f.log.Infof("Fetched bucket: %s", bucket.Name.Value())
		f.resourceCh <- fetching.ResourceInfo{
			Resource:      resource,
			CycleMetadata: cMetadata,
		}
	}
	return nil
}

func (f *S3Fetcher) Stop() {
}

func (r S3Resource) GetData() interface{} {
	converted := convert.StructToRego(reflect.ValueOf(r.bucket))
	for i, policy := range r.bucket.BucketPolicies {
		var output interface{}
		marshalled, _ := policy.Document.Parsed.MarshalJSON()
		json.Unmarshal(marshalled, &output)
		a, _ := converted["bucketpolicies"].([]interface{})
		b, _ := a[i].(map[string]interface{})
		c, _ := b["document"].(map[string]interface{})
		c["value"] = output
	}
	return converted
}

func (r S3Resource) GetMetadata() (fetching.ResourceMetadata, error) {
	return fetching.ResourceMetadata{
		// A compromise because aws-sdk do not return an arn for an Elb
		ID:      fmt.Sprintf("%s", r.bucket.Metadata.Reference()),
		Type:    fetching.CloudStorage,
		SubType: "S3 bucket",
		Name:    r.bucket.Name.Value(),
	}, nil
}

func (r S3Resource) GetElasticCommonData() any { return nil }
