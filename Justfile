alias b := build
alias t := test

set dotenv-load

default: test

tools:
  go list -f '{{ '{{range .Imports}}{{.}} {{end}}' }}' tools.go | xargs go install
  asdf reshim golang

test *args:
  go test {{args}} ./...

build *args:
  go build {{args}} ./...

run *args:
  go run {{args}} ./cmd/app

sqlc:
  sqlc generate --experimental

get_env_vars:
  bw logout || true
  BW_SESSION="$(bw login --sso)"; export BW_SESSION
  secret="$(bw list items --search address-book-api .env | jq -r '.[0].notes')"; [[ "$secret" != "null" ]] && echo "$secret" > .env

