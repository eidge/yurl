**Notice:** *WIP, it does not work yet.*

# Yurl

**Yurl** lets you define requests in a simple and readable sintaxe, extract repeated
data and save these to use later or share with fellow developers.

## Usage

Yurl consists of **yml** configuration files and the **yurl CLI**.

### YML Files

A yurl file can represent one or more HTTP Requests and respects the following
format:

```yml
# ~/(...)/users_endpoint.yml

base_url: https://example.com
headers: # headers to include for every request in this file
  AUTHENTICATION_TOKEN: 123456

users: # yurl users_endpoint.yml users
  path: /users
  method: GET
  query_str:
    age: 30
    created_on: 30/02/1992

create_user: # yurl users_endpoint.yml create_user
  path: /users
  method: POST
  body_format: json # Can be one of: form_encoded|json|raw . Defaults to JSON
  body:
    first_name: John
    last_name:  Doe
    age:        34
```

### Yurl CLI

The Yurl CLI lets you make requests based on the yml files you've previously
defined.

The most basic usage is ```yurl YML_FILE REQUEST_NAME``` which allows you to
make a single request defined in the ```YML_FILE```.


## WHYs

### Why not Curl?

Curl is a great tool but ends up having a complicated sintaxe and it has no
builtin support to replay or save requests to perform later.

### Why YAML?

YAML sintax is simple enough that everyone can edit or create a file even if
it is the first time using it.

It also maps well to most common data types used for API communication.

### Why Go?

I was curious about Go.

It's fast and easily supports cross-compiling and static linking, both being
important when you want to release a simple self-contained program to run on the command
line.

## Developing

Make sure you have ```go``` installed and your ```$GO_PATH``` defined.

To get the code, run:

```bash
go get github.com/eidge/yurl
```

The standard go tools are used, so to run/test/build/install, you should only
need to run:

```bash
go run|test|build|install
```

## Contributing

1. Create an issue for the feature/bug you're implementing
2. Fork this repository
3. Create a failing test if you're solving a bug
4. Implement away! (All features must be tested)
6. Make sure all tests pass
7. Create a PR against master

## License

The MIT License (MIT)

Copyright (c) 2015 Hugo Ribeira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
