package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
    // Using the SDK's default configuration, loading additional config
    // and credentials values from the environment variables, shared
    // credentials, and shared configuration files
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }
	
	log.Print(cfg)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func getRdsClient() {
	// Create an RDS service client
	rdsClient := rds.New(sess)

	// Describe the Aurora DB instance
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String("dbInstanceIdentifier"),
	}

	result, err := rdsClient.DescribeDBInstances(input)
	if err != nil {
		log.Fatal("Error describing DB instance:", err)
	}

	// Print information about the Aurora DB instance
	if len(result.DBInstances) > 0 {
		dbInstance := result.DBInstances[0]
		fmt.Printf("DB Instance ID: %s\n", *dbInstance.DBInstanceIdentifier)
		fmt.Printf("Engine: %s\n", *dbInstance.Engine)
		fmt.Printf("Endpoint: %s\n", *dbInstance.Endpoint.Address)
		fmt.Printf("Port: %d\n", *dbInstance.Endpoint.Port)
	} else {
		fmt.Println("DB instance not found")
	}
}