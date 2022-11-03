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

package awslib

import (
	"context"
	"github.com/aquasecurity/defsec/pkg/providers/aws/s3"
	awsScanner "github.com/aquasecurity/defsec/pkg/scanners/cloud/aws"
	"github.com/aquasecurity/defsec/pkg/scanners/options"
	"github.com/elastic/elastic-agent-libs/logp"
)

type S3BucketDescription s3.Bucket
type S3BucketDescriptions []s3.Bucket

type S3BucketDescriber interface {
	DescribeS3Buckets(ctx context.Context, log *logp.Logger) (S3BucketDescriptions, error)
}

type S3Provider struct {
	scanner *awsScanner.Scanner
}

func NewS3Provider() *S3Provider {
	// TODO:(ARI) Credentials are self-obtained in the adaptor code, is this fine?
	opts := []options.ScannerOption{
		awsScanner.ScannerWithAWSRegion("us-east-1"),
		awsScanner.ScannerWithAWSServices("s3"),
	}
	scanner := awsScanner.New(opts...)

	return &S3Provider{
		scanner: scanner,
	}
}

func (provider S3Provider) DescribeS3Buckets(ctx context.Context, log *logp.Logger) (S3BucketDescriptions, error) {
	cloudState, err := provider.scanner.CreateState(ctx)
	if err != nil {
		log.Errorf("ARI Error %v", err)
	}
	return cloudState.AWS.S3.Buckets, err
}
