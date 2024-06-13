# Packaging

Starting a new module, initializing a go.mod file in the current directory.

```bash
go mod init github.com/your-username/your-package-name
```

Add missing and remove unused modules.

```bash
go mod tidy
```

- `-e` ignore not found packages

Add a module to the current module and install it.

```bash
go get github.com/your-username/your-package-name
```

Replace a module with a local directory. (Not recommended for production) Use go mod tidy after replacing.

```bash
go mod edit -replace=github.com/your-username/your-package-name=../your-package-name
```

Create a go workspace. A workspace is a directory that contains multiple packages. If you are importing a package that aren't deployed to a remote repository, It'll search for the package in the workspace.

```bash
go work init relative/path/to/your-package-name github.com/your-username/your-package-name
```

Use `go mod tidy -e` in workspaces because it'll ignore not found packages.
