module github.com/sqooba/go-common

go 1.15

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.3.1-0.20170228224354-599cba5e7b61 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200916030750-2334cc1a136f // indirect
)

replace (
	github.com/sqooba/go-common => ../sqooba-go-common
)
