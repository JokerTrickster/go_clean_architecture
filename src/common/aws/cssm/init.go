package cssm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"main/common"
	"strings"
)

var AwsClientSsm *ssm.Client

func InitAwsSsm() error {
	var awsConfig aws.Config
	var err error

	awsConfig, err = AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("init aws - region : %s", "us-east-1")
	}

	AwsClientSsm = ssm.NewFromConfig(awsConfig)

	if err != nil {
		return err
	}
	return nil
}

func AwsGetParam(ctx context.Context, path string) (string, error) {
	param, err := AwsClientSsm.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: common.BoolToPointer(true),
	})
	if err != nil {
		return "", fmt.Errorf("get cssm param : %s, %v", path, err)
	}
	return aws.ToString(param.Parameter.Value), nil
}

func AwsGetParams(paths []string) ([]string, error) {
	ctx := context.TODO()
	params, err := AwsClientSsm.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          paths,
		WithDecryption: common.BoolToPointer(true),
	})
	if err != nil {
		return nil, fmt.Errorf("get cssm params : %s", strings.Join(paths, ","))
	}
	result := make([]string, len(paths))
	for i, path := range paths {
		val := ""
		for _, parameter := range params.Parameters {
			if strings.Contains(aws.ToString(parameter.ARN), path) {
				val = aws.ToString(parameter.Value)
				break
			}
		}
		result[i] = val
	}
	return result, nil
}
