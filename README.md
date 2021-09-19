# xerrors

Package `xerrors` provides extended error handling primitives to add a bit more information to errors returning from the function.
```
go get github.com/eugeneradionov/xerrors
```

## Motivation
Every API that handles HTTP requests, you need to work with error handling.
The standard `errors` package does not bring the easy way to know what the error is and with which code to respond.
With `xerrors` package, you can extend the error with missing information, such as HTTP status code to respond, 
and better control the application flow.

## Usage
Use `XError` interface instead of standard `error`

```go
func GetUserByID(id string) (*User, xerrors.XError) {
    user, err := db.GetUserByID(id)
    if err != nil {
        return nil, xhttp.NewInternalServerError(err)
    }
    
return user, nil
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=utf-8")    

    user, xErr := GetUserByID("user_id_1")
    if xErr != nil {
        LogXError(xErr)
        SendXError(w, xErr)
        return
    }
}

func SendExtError(w http.ResponseWriter, xErr xerrors.XError) {
    var statusCode = xErr.GetExtra()["http_code"].(int)

    w.Header().Set("Content-Type", "application/json;charset=utf-8")
    errResp, err := json.Marshal(xErr)
    if err != nil {
        statusCode = http.StatusInternalServerError
    }
    
    w.WriteHeader(statusCode)
    w.Write(errResp)
}

func LogXError(xErr xerrors.XError) {
    log.Printf("[ERR] %s: %s", xErr.Error(), xErr.GetInternalExtra()["error"].(error).Error())
}
```

The same with `xerror`

```go
func GetUserByID(id string) (*User, xerrors.XError) {
    user, err := db.GetUserByID(id)
    if err != nil {
        return nil, xerrors.New("Internal Server Error",
            xerrors.WithExtra(map[string]interface{}{"http_code": http.StatusInternalServerError}),
            xerrors.WithInternalExtra(map[string]interface{}{"error": err}),
        )
    }
    
    return user, nil
}
```

## Caveats

As `XError` requires implementation of standard `error` interface to be compatible with it,
be careful, when trying to assign function result `XError` to the variable with standard `error` type. 
This could cause unpredictable behavior.
