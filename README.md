# captian is a buid script for docker builds

## usage
```
 captain -h
Usage of captain:
  -build
        bulid Flag
  -help
        help Flag
  -test
        test Flag
  -version
        version Flag
```

example
```
captain -version
0.0.1-f4a4713-alpha.1
```

## properties
set your version parameters. define your build commands.

Version string: %versionPrefix-%versionSuffix-%versionGitHash

```
versionPrefix=0.0.1
versionSuffix=alpha.1
versionGitHash=git rev-parse --short HEAD

cmd.build=docker build -f Dockerfile --tag demoimage:#VERSIONSTRING# .
cmd.test=docker build -target unittest -f Dockerfile -t demoimage:#VERSIONSTRING#-test .
```



## run
```
go run captain.go
```

## test
```
go test captain*.go
```

## install
```
go install github.com/trieder83/captain
```

## build with docker image 
```
 sh build
```


