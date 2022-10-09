1. [Main Repo Link](https://github.com/plutov/paypal)

- replace directive in go.mod file to use forked repository as package import
    - [Reference post](https://stackoverflow.com/questions/14323872/using-forked-package-import-in-go)
    - new go.mod file: [forked_import_go.mod](#forked_import_go.mod) - this file in itself won't work, you have to use the filename=`go.mod`.

- Separate unit tests with integration tests : [separating tests using build tags](https://mickey.dev/posts/go-build-tags-testing/)
- To run tagged tests, use `go test -tags=integration`
    - in case of paypal, cannot run the integration test even after specifying tags because of variables `testClientID` and `testSecret` being duplicate in `order_test.go` and `integration_test.go`
    - **cheap solution**: comment out the `clientID` and `secret` declared in `order_test.go`.
- all structs/interfaces are housed in a **single file**: `types.go`. **= Convention??**

# Our Contribution - Disputes
1. [Paypal Developer API Link- Disputes](https://developer.paypal.com/docs/api/customer-disputes/v1/)