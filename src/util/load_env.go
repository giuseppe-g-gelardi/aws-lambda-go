package util

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Env struct {
	Url          string
	Key          string
	ClientId     string
	ClientSecret string
	Environment  string
	DevUrl       string
	ProdUrl      string
}

func LoadEnv() Env {
	godotenv.Load()

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_ANON_KEY")
	client_id := os.Getenv("GITHUB_CLIENT_ID")
	client_secret := os.Getenv("GITHUB_CLIENT_SECRET")
	environment := os.Getenv("VERCEL_ENV")
	dev_url := os.Getenv("DEV_BASE_URL")
	prod_url := os.Getenv("PROD_BASE_URL")

	return Env{
		Url:          url,
		Key:          key,
		ClientId:     client_id,
		ClientSecret: client_secret,
		Environment:  environment,
		DevUrl:       dev_url,
		ProdUrl:      prod_url,
	}
}
