// Code generated by github.com/aws/aws-sdk-go-v2/config. DO NOT EDIT.

package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

// CredentialsProviderProvider implementor assertions
var (
	_ CredentialsProviderProvider = WithCredentialsProvider{aws.NewStaticCredentialsProvider("", "", "")}
)

// CustomCABundleProvider implementor assertions
var (
	_ CustomCABundleProvider = &EnvConfig{}
	_ CustomCABundleProvider = WithCustomCABundle([]byte{})
)

// DefaultRegionProvider implementor assertions
var (
	_ DefaultRegionProvider = WithDefaultRegion("")
)

// MFATokenFuncProvider implementor assertions
var (
	_ MFATokenFuncProvider = WithMFATokenFunc(func() (string, error) { return "", nil })
)

// RegionProvider implementor assertions
var (
	_ RegionProvider = &EnvConfig{}
	_ RegionProvider = &SharedConfig{}
	_ RegionProvider = WithRegion("")
)

// SharedConfigFilesProvider implementor assertions
var (
	_ SharedConfigFilesProvider = &EnvConfig{}
	_ SharedConfigFilesProvider = WithSharedConfigFiles([]string{})
)

// SharedConfigProfileProvider implementor assertions
var (
	_ SharedConfigProfileProvider = &EnvConfig{}
	_ SharedConfigProfileProvider = WithSharedConfigProfile("")
)
