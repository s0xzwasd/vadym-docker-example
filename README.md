# go-docker-environment 

![Type of License](https://img.shields.io/github/license/s0xzwasd/go-docker-environment) ![Go version is 1.16](https://img.shields.io/github/go-mod/go-version/s0xzwasd/go-docker-environment) ![Size of the repository](https://img.shields.io/github/repo-size/s0xzwasd/go-docker-environment)

Baisc example of usage Docker with Go and GoLand. Pay attention to requirements and steps sections before clone the repository.

## Getting started

### Requirements

- [GoLand 2020.1.1 version](https://www.jetbrains.com/go/download) or newer.
- [Docker Desktop 17.06](https://www.docker.com/products/docker-desktop) or newer.

### Steps

1. Navigate to `Settings / Preferences | Build, Execution, Deployment | Docker`, add a new entry and select the local Docker instance. ![Example of adding a new entry](https://i.imgur.com/ps7stWr.png)
2. Go to Dockerfile, click on the gutter icon over `FROM golang:1.16.3` and select Run 'Docker' option. ![Dockerfile with gutter icons in GoLand](https://i.imgur.com/MnLfhSQ.png)
3. Open http://localhost:8000/ in the browser. ![Successful run in the docker container](https://i.imgur.com/D8wmiNf.png)

## Debugging

- Select the Docker debugging configuration and start the container. ![How to run the container](https://i.imgur.com/qB8e8Xl.png)
- Set breakpoints. ![Example of setting breakpoints](https://i.imgur.com/c5lULVI.png)
- Choose the Dockerization debugger configuration and press Debug button. ![How to ren the debug configuration](https://i.imgur.com/f8UYT9X.png)

## Roadmap

- [x] Add a basic example of usage
- [x] Pre-define run configurations
- [x] Add helpful README
- [x] Configure debugger options and describe the way to do it properly
- [ ] Add docker-compose and how to debug it
- [ ] Manage run targets and save it to the project

## Troubleshooting

TBD

## License

This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
