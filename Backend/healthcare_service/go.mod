module healthcare_service

go 1.18

require (
	authorization v0.0.0-00010101000000-000000000000
	github.com/casbin/casbin v1.9.1
	github.com/cristalhq/jwt/v4 v4.0.2
	github.com/gorilla/mux v1.8.0
	github.com/nats-io/nats.go v1.31.0
	go.mongodb.org/mongo-driver v1.11.4
	nats v0.0.0-00010101000000-000000000000
)

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace authorization => ../authorization

replace nats => ../nats
