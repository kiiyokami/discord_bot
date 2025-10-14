package main

import (
	"context"
	"discord_bot/bot"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("AWS config error: %v", err)
	} else {
		log.Println("AWS config loaded successfully")
		svc := ssm.NewFromConfig(cfg)
		result, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{
			Name:           aws.String("/discord/token"),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			log.Printf("Parameter Store error: %v", err)
		} else {
			log.Println("Successfully retrieved token from Parameter Store")
			bot.Run(*result.Parameter.Value)
			return
		}
	}
	
	log.Println("Falling back to environment variable")
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("No Discord token found in Parameter Store or environment")
	}
	bot.Run(token)
}

