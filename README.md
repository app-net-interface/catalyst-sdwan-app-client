# catalyst-sdwan-app-client (Catalyst SDWAN App client for application owner (DevOps/CloudOps)

This package is a set of low level catalyst sdwan internal functionalities exposed as high level golang functions. 
It is primarily meant for developers looking to interact with Catalyst SDWAN programatically for various use cases. 

In the context of AWI, this library is used by catalyst-sdwan-grpc repository for SDWAN interations. 

## Available functions

CRUD operation available on following functions for Cisco Catalyst SDWAN.

```go
  
    ACL() ACL
    Connection() Connection
    Device() Device
    Feature() Feature
    VPN() VPN
    VPC() VPC
    Status() Status
    Site() Site
    Policy() Policy
    URLFiltering() URLFiltering
    URLDenylist() URLDenylist
    URLAllowlist() URLAllowlist
    -------------------
    Login(ctx context.Context, username, password string) error
    GetToken() string
    SetToken(token string)
```

## Interface Definition

https://github.com/app-net-interface/catalyst-sdwan-app-client/blob/main/vmanage/vmanage.go

## Contributing

Thank you for interest in contributing! Please refer to our
[contributing guide](CONTRIBUTING.md).

## License

catalyst-sdwan-app-client is released under the Apache 2.0 license. See
[LICENSE](./LICENSE).

catalyst-sdwan-app-client is also made possible thanks to
[third party open source projects](NOTICE).
