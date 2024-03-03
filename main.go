package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func main() {
	// Test parameter group name
	parameterGroupName := "rds-pg"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	rdsClient := rds.NewFromConfig(cfg)

	// Describe before modify
	describeInput := &rds.DescribeDBParameterGroupsInput{
		DBParameterGroupName: &parameterGroupName,
	}
	var meta *rds.DescribeDBParameterGroupsOutput
	meta, err = rdsClient.DescribeDBParameterGroups(context.TODO(), describeInput)
	if err != nil {
		fmt.Printf("Describe DB Parameter Group: %s", err)
		os.Exit(1)
	}

	fmt.Printf("ApplyMethod = (%+v) Before modify ", meta.DBParameterGroups)
	fmt.Printf("ApplyMethod = (%+v) Before modify ", meta.DBParameterGroups)
	os.Exit(0)

	// Modify
	var params = []types.Parameter{
		types.Parameter{
			ParameterName:  aws.String("log_statement"),
			ApplyMethod:    "immediate",
			ParameterValue: aws.String("mod"),
		},
	}

	input := &rds.ModifyDBParameterGroupInput{
		DBParameterGroupName: &parameterGroupName,
		Parameters:           params,
	}

	_, err = rdsClient.ModifyDBParameterGroup(context.TODO(), input)
	if err != nil {
		fmt.Printf("modifying DB Parameter Group: %s", err)
		os.Exit(1)
	}
	fmt.Printf("modifying DB Parameter Group are succeeded.")
}
