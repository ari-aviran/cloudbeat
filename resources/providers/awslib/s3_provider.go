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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/elastic/elastic-agent-libs/logp"
)

type S3BucketDescription struct {
	Bucket     s3Types.Bucket
	Encryption *s3Types.ServerSideEncryptionRule
}

type S3BucketDescriber interface {
	DescribeS3Buckets(ctx context.Context, cfg aws.Config, log *logp.Logger) ([]S3BucketDescription, error)
}

type S3Provider struct {
}

func NewS3Provider() *S3Provider {
	return &S3Provider{}
}

func (provider S3Provider) DescribeS3Buckets(ctx context.Context, cfg aws.Config, log *logp.Logger) ([]S3BucketDescription, error) {
	client := s3.NewFromConfig(cfg)

	clientBuckets, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Errorf("Could not list s3 buckets: %v", err)
		return nil, err
	}

	var result []S3BucketDescription

	for _, clientBucket := range clientBuckets.Buckets {
		if *clientBucket.Name != "cis-aws-ari-test" && *clientBucket.Name != "cis-aws-ari-test-2" {
			continue
		}
		// TODO: Split functions
		encryption, err := client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{Bucket: clientBucket.Name})
		var encryptionRule *s3Types.ServerSideEncryptionRule = nil

		if err != nil {
			// TODO: This does not work, why?
			if awsError, ok := err.(awserr.Error); ok {
				log.Infof("ARI ARI ARI ARI ARI!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
				if awsError.Code() != "ServerSideEncryptionConfigurationNotFoundError" {
					log.Errorf("Could not get encryption for bucket %s. Error: %v", *clientBucket.Name, err)
				}
			} else {
				log.Errorf("Could not get encryption for bucket %s. Error: %v", *clientBucket.Name, err)
			}
		} else {
			if len(encryption.ServerSideEncryptionConfiguration.Rules) > 0 {
				encryptionRule = &encryption.ServerSideEncryptionConfiguration.Rules[0]
			}
		}

		result = append(result, S3BucketDescription{clientBucket, encryptionRule})
	}

	return result, nil
}
