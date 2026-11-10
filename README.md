# Go binding

`github.com/EuroZig/git-native-issuez-go/issuez` wraps the C API
(`issuez/include/issuez.h`) with cgo. Everything the library returns is
copied into plain Go values and the C memory is freed at once; only
`*issuez.Repo` holds C state. Calls take a `context.Context`, errors are
`*issuez.Error` (`errors.Is(err, issuez.ErrVersionMismatch)`), writes are
`issuez.Edit` values.

```go
repo, err := issuez.Open("/srv/git/ada/app.git", issuez.WithCommitter("Host", "noreply@example.com"))
w, err := repo.Update(ctx, id, new(issuez.Edit).SetState("closed"), actor, issuez.Reload(), issuez.ExpectTip(tip))
page, err := repo.Query(ctx, issuez.Query{Labels: []string{"bug"}, Order: issuez.SortUpdated})
```

See [the embedding guide](https://github.com/EuroZig/git-native-issuez/blob/v0.5.0/docs/embedding.md) for exact-tip reads,
text-version preconditions, external image storage, staged imports, and
how to handle a committed write whose reload fails.

## Archives

This module has `libissuez.a` for macOS (arm64, amd64) and Linux (amd64, arm64
and arm, which is ARMv6 and runs on ARMv7 too, with glibc or musl), so `go get`
is all it takes. Builds need cgo and a C compiler, and Linux builds for musl
(Alpine) need `-tags musl`.

```sh
go get github.com/EuroZig/git-native-issuez-go/issuez@v0.5.0
```

Cross-compiling needs a C compiler for the target; `zig cc` works:
`CC="zig cc -target aarch64-linux-musl" GOOS=linux GOARCH=arm64 go build -tags musl`.

Each [release](https://github.com/EuroZig/git-native-issuez/releases) writes this repository, so changes and
issues go to [git-native-issuez](https://github.com/EuroZig/git-native-issuez), where the binding is
`bindings/go`.

## Go versions

The `go` line is the oldest supported Go release (1.26 now), raised only
when a newer feature is needed or a release leaves support.
