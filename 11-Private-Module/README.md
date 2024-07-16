# Private Module

## Private Repository

Configure GOPRIVATE to use private modules.

```sh
export GOPRIVATE=github.com/yourusername
```

Configura your credentials on ~/.netrc

```txt
machine github.com
login yourusername
password yourtoken
```

## Save Dependencies

Add a vendor directory to your project to store dependencies.

```sh
go mod vendor
```
