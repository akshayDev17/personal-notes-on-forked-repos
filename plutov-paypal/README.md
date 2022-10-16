1. [Main Repo Link](https://github.com/plutov/paypal)

- replace directive in go.mod file to use forked repository as package import
    - [Reference post](https://stackoverflow.com/questions/14323872/using-forked-package-import-in-go)
    - new go.mod file: [forked_import_go.mod](#forked_import_go.mod) - this file in itself won't work, you have to use the filename=`go.mod`.

# Integration Testing
- Separate unit tests with integration tests : [separating tests using build tags](https://mickey.dev/posts/go-build-tags-testing/)
- To run tagged tests, use `go test -tags=integration`
    - in case of paypal, cannot run the integration test even after specifying tags because of variables `testClientID` and `testSecret` being duplicate in `order_test.go` and `integration_test.go`
    - **cheap solution**: comment out the `clientID` and `secret` declared in `order_test.go`.
- all structs/interfaces are housed in a **single file**: `types.go`. **= Convention??**

- Note that this is not an API, rather something that will make a request to it.

# Types
- Client type
    - **composition** from `sync.Mutex` type, all fields and methods related to it promoted to type `Client`.
    - method `NewRequest()`
        - `json.Marshal()` returns the byte-encoded version of a struct, `bytes.NewBuffer()` decodes this to a human-readable fashion, keeping it in bytes.Buffer(`Buffer` is a *`struct`* in `bytes/buffer.go`) form though.
            - notice that it makes all field-names in lowercase.
        - `io.Reader` is an interface with the following signature for method `Read`: `Read(p []byte) (n int, err error)`, and Buffer struct has on line 310 in buffer.go, implemented this method.
        - method `http.NewRequestWithContext`(notice package is **http**)
            - `u, err := urlpkg.Parse(url)` ---> preps the  *`https://api.sandbox.paypal.com/v2/checkout/orders`*.
            - **interface type assertion**(`body.(io.ReadCloser)`): the second variable tells if assertion was successful(true) or not(false).
            - at the end , an http.Request struct gets created.(*API hasn't been hit yet*)
    - method `SendWithAuth(req *http.Request, v interface{})`
        - c.Lock()
            - execute current client object's lock method, to prevent other clients from token acquisition.

# Issue: Make the mutex private in the structure
- refer to `sql.go(in database/sql)` , the DB struct, with all lines where `db.mu.Lock()` is invoked, `mu` is the `sync.Mutex` , which is what is required in this case as well.


# Orders
## Create Order
- 

# Our Contribution - make Mutex private

## Introduction
1. all functions where Mutex methods are invoked:
    1. in `client.go`, `SendWithAuth()` invokes both `Lock()` and `Unlock()`
    2. but this SendWithAuth method is used quite extensively.
        1. so first lets try to create testing code wherever this method is invoked.

## `SendWithAuth()`
1. create a temp `sendwithauth()` and a `createProduct()` inside `client_test.go`.
    1. rename this file to `client_mutex_private_test.go`.
2. use the AssertMutexLocked from [here](https://github.com/trailofbits/go-mutexasserts/blob/master/asserts_debug.go), inside the function body of this temp `sendwithauth()`.

# Our Contribution - Disputes
1. [Paypal Developer API Link- Disputes](https://developer.paypal.com/docs/api/customer-disputes/v1/)
2. for all enum type query parameters(`dispute_state` in List disputes)/response body fields(`dispute_channel` in show dispute details), refer to `types.go`(Possible value for `intent` in CreateOrder)
