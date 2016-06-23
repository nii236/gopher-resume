# Project Name

This is an early stage project attempting to port JSON resume over to Go

## Installation

```
$ go get github.com/nii236/gopher-resume
```

## Usage


CLI tool to generate CVs from the command line.

```
$ go run main.go process -name ExampleGuy
```

Web server with a REST API that generates CVs.

```
$ go run main.go serve
$ curl localhost:8080/?name=john%20guy

Name: john guy
Label:
Picture:
Email:
Phone:
Website:
Summary:
Location: {    }
Profiles: []
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
