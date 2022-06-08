GO = go1.18

all:
	${GO} build ./...

get:
	${GO} get ./...

mod:
	${GO} mod tidy

test:
	${GO} test ./...

bench:
	${GO} test ./... -bench=. -benchmem
