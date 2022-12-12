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
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/providers/awslib"
	"github.com/elastic/elastic-agent-libs/logp"
)

type S3Fetcher struct {
	log        *logp.Logger
	cfg        S3FetcherConfig
	s3Provider awslib.S3BucketDescriber
	resourceCh chan fetching.ResourceInfo
	awsConfig  aws.Config
}

type S3FetcherConfig struct {
	fetching.AwsBaseFetcherConfig `config:",inline"`
}

type S3Resource struct {
	bucket awslib.S3BucketDescription
}

func (f *S3Fetcher) Fetch(ctx context.Context, cMetadata fetching.CycleMetadata) error {
	f.log.Info("Starting S3Fetcher.Fetch")

	buckets, err := f.s3Provider.DescribeS3Buckets(ctx, f.awsConfig, f.log)
	if err != nil {
		return fmt.Errorf("failed to load buckets from S3: %w", err)
	}

	for _, bucket := range buckets {
		resource := S3Resource{bucket}
		f.log.Debugf("Fetched bucket: %s", *bucket.Bucket.Name)
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
	return r.bucket
}

func (r S3Resource) GetMetadata() (fetching.ResourceMetadata, error) {
	return fetching.ResourceMetadata{
		// TODO: use proper ARN
		ID:      fmt.Sprintf("%s", *r.bucket.Bucket.Name),
		Type:    fetching.CloudStorage,
		SubType: fetching.S3Type,
		Name:    *r.bucket.Bucket.Name,
	}, nil
}

func (r S3Resource) GetElasticCommonData() any { return nil }
