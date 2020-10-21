mv resolver.go resolver.old
go run github.com/99designs/gqlgen # generate graphQL

cp schema.resolvers.go resolver.go.new
git merge-file resolver.go resolver.go resolver.old
rm schema.resolvers.go