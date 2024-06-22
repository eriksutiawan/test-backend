# test-backend
test backend developer

change environtment (.env) :
- JWT_SECRET_KEY=your_jwt_secret_key
- GOOGLE_CLIENT_ID=your_google_client_id
- GOOGLE_CLIENT_SECRET=your_google_client_secret
- GOOGLE_REDIRECT_URL="http://localhost:8080/auth/google/callback"
- MICROSOFT_CLIENT_ID=your_microsoft_client_id
- MICROSOFT_CLIENT_SECRET=your_microsoft_client_secret
- MICROSOFT_REDIRECT_URL="http://localhost:8080/auth/microsoft/callback"
- MONGO_URI="mongodb://localhost:27017"
- COLLECTION_NAME="test"

run terminal:
- go mod tidy
- go mod vendor
- go run main.go

tools: 
- mongoDb
  
swagger:
- http://localhost:8080/swagger/index.html

base url api :
- http://localhost:8080/api

add model best_plant_disease.h5 in folder internal/pkg/model
