module chap_1

go 1.14

replace github.com/matryer/goblueprints/chapter1/trace => ./trace

require (
	github.com/gorilla/websocket v1.4.2
	github.com/matryer/goblueprints/chapter1/trace v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d // indirect
)
