# Building Go Applications

## Dockerfile

Normally, we have 2 Dockerfiles

### Dockerfile.dev

Used for development

It only contains the base image, working directory, and the command to run the application. It is used to build the image and run the application locally.

Often,  I use the command `tail -f /dev/null` to keep the container running.

### Dockerfile.prod

Used for production

It must have every step to build the application, including the dependencies, and the final image must be as small as possible.

For generate a build with go, we can use a multi-stage build. The first stage is to build the application, and the second stage is to run the application.

- Build the application a go image
  - set the OS and the architecture with `GOOS` and `GOARCH`
    - GOOS is the operating system, linux, windows, etc.
    - GOARCH is the architecture, amd64, arm, etc.
  - remove DWARF information with `-ldflags="-w -s"`.
    - ldflags are flags that are passed to the linker. The `-w` flag tells the linker to remove the DWARF information, and the `-s` flag tells the linker to remove the symbol table.
    - DWARF is the debugging information that is added to the binary. It is useful for debugging, but it is not necessary for production.
  - Disable CGO
    - CGO is the Go compiler that allows Go code to call C code. It is not necessary for this application, so we can disable it with the `CGO_ENABLED=0` environment variable.
- Run the application in a smaller image
  - use the `scratch` image as the base image.
    - It is an empty image, and it is the smallest image possible and it can run binaries compiled with Go.
  - copy the binary from the first stage to the second stage and run it.

## Kubernetes Files

### Service

The service is used to expose the application to the outside world. It is a load balancer that forwards the traffic to the pods.

### Deployment

The deployment is used to create and manage the pods. It is responsible for creating the pods, updating the pods, and deleting the pods.
