# CreditStore
Nothing interesting, Just learning :D

### Create/Init Module
```shell
go mod init github.com/gunjanpatel/creditStore
go mod tidy
```

### Refer Local package (while working on dev)
Just add following line into go.mod to the practice module
```go
require github.com/gunjanpatel/creditStore v0.0.0

replace github.com/gunjanpatel/creditStore => ../creditStore
```

## Publish Package
```shell
git tag v0.1.0
git push origin v0.1.0
```

## Use published package
go get github.com/gunjanpatel/creditStore