# GraphQL

## Documentation

- [GraphQL](https://graphql.org/)

## Steps

1. Create tools.go file
2. Run `go run github.com/99designs/gqlgen init` to create the necessary files
3. Change the schema.graphql file. Add your schema.
4. Run `go run github.com/99designs/gqlgen generate` to generate the code
5. Implement the resolvers. Add your dependencies on the resolver.go file and implement the resolvers on the schema.resolvers.go file

### Nested Queries

1. Create a new type on a new file in the model folder
2. Add the new type to the schema.graphql file
3. Remove the reference between the new types
4. Run `go run github.com/99designs/gqlgen generate` to generate the code

## Skills

- [x] Go
- [x] GraphQL
- [x] Active Record Pattern
