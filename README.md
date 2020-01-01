# Wait

Wait is a simple package extending go's official sync package with a context aware ```Wait``` method, thus making cancellable via timeout.

## Get started

```bash
go get github.com/shirikiorg/wait
```

## Usage

```go
  var wg wait.Group
  ctx, cancel := context.WithTimeout(context.Background, 1*time.Second)
  defer cancel()

  wg.Add(1)
  go func() {
    defer wg.Done()
    // do something time consuming here
  }()

  wg.WaitWithContext(ctx)
```

## Contributing

Contributions are welcome :)


