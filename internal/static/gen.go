// Package static contains the static HTTP resources served at /static/.
package static

//go:generate gobin -m -run github.com/mjibson/esc -o=static.esc.go -pkg=static -ignore=\.go$ .
