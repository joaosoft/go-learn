> Reference
https://go.dev/doc/tutorial/workspaces

# get dependencies
go get "golang.org/x/example/stringutil"

# initialize a module
go mod init github.com/joaosoft/go-learn/81_multi_modules/workspace/hello_world

# update the module
go mod tidy

# create a vendor folder and add the dependencies there
go mod vendor

# initialize the workspace
go work init ./hello_world

# run the module
go run hello_world/main.go

# Next, we’ll add a local copy of the golang.org/x/example module to the workspace. 
# We’ll then add a new function to the stringutil package that we can use instead of Reverse.
git clone https://go.googlesource.com/example

# use the local module instead of the remote! 
# we have changed the local reverse function to also use a ToUpper function
go work use ./example

# run the module and see what have changed
go run hello_world/main.go

# go.work can be used instead of adding replace directives to work across multiple modules.