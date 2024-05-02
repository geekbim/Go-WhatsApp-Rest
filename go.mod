// +heroku goVersion go1.17
module go_wa_rest

go 1.2

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-openapi/runtime v0.28.0
	github.com/go-playground/assert v1.2.1
	github.com/gorilla/mux v1.8.1
	github.com/hashicorp/go-multierror v1.1.1
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/sirupsen/logrus v1.9.3
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/stretchr/testify v1.9.0
	go.mau.fi/whatsmeow v0.0.0-20240327124018-350073db195c
	google.golang.org/protobuf v1.34.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.mau.fi/util v0.4.2 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
)
