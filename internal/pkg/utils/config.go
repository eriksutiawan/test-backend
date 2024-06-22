package utils

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
)

func ConnectToMongo() *mongo.Database {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	var err error
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(os.Getenv("COLLECTION_NAME"))
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error : ", err)
	}
}

var (
	GoogleOAuthConfig    *oauth2.Config
	MicrosoftOAuthConfig *oauth2.Config
)

func InitOAuthConfig() {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}

	MicrosoftOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("MICROSOFT_REDIRECT_URL"),
		Endpoint:     microsoft.AzureADEndpoint("common"),
		Scopes:       []string{"https://graph.microsoft.com/User.Read"},
	}
}
