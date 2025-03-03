# kvstore
Basic KV store in GO

## Dependencies
- [Buf](https://buf.build/docs/cli/installation/#__tabbed_1_1) - Used tool for go code generations for protobufs relating to `.proto` files for gRPC calls. 

## Testing gRPC
To directly test against `gRPC` API call's we can use a tool called [grpcui](https://github.com/fullstorydev/grpcui?tab=readme-ov-file#installation).
1. Use `grpcui` to interact with any internal or public gRPC calls:
    * This will require generating the `protoset` which can be done via buf
    ```bash
    buf build ./proto-public -o api.bin --as-file-descriptor-set
    ```
1. We can now use the generated `protoset` file called `api.bin` to send requests at our gRPC URL
    ```bash
    # Define your $GRPC_ADDR var
    grpcui -plaintext -protoset api.bin "$GPRC_ADDR"
    ```