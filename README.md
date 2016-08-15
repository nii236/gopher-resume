[![Stories in Ready](https://badge.waffle.io/nii236/gopher-resume.png?label=ready&title=Ready)](https://waffle.io/nii236/gopher-resume)
[![Go Report Card](https://goreportcard.com/badge/github.com/nii236/gopher-resume)](https://goreportcard.com/report/github.com/nii236/gopher-resume)
[![Build Status](https://travis-ci.org/nii236/gopher-resume.svg?branch=master)](https://travis-ci.org/nii236/gopher-resume)
# Gopher Resume

This is an early stage project attempting to port JSON resume over to Go.

The scope has since change and now the server will host a copy of a user's resume. Recruiters will be given a expirable token with which they can login and view the candidate's private details.

There will be a CLI proxy from the client side, containing admin, member and guest levels of information.

At some point there will be the ability for other people to upload their own CVs and generate their own tokens for their own recruiters.

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
