#/bin/zsh

echo '{"message": "Hello %s"}' | go run cli.go greet --name=Test