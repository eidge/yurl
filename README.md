# Yurl

**A simple tool to make and save API HTTP requests.**

## WHYs

### Why not Curl?

Curl is a great tool, but ends up having a complicated sintaxe for beginners and
it doens't really solve the problem of saving requests for using later or
actually reuse repeated parts such as the authentication headers.

### Why Go?

I was curious about Go.

It's fast and easily supports cross-compiling and static linking, both being
important when you want to release a simple self-contained program to run on the command
line.

## Usage

Yurl consists of **yml** configuration files and the **yurl CLI**.

### YML Files

A yurl file can represent one or more HTTP Requests and respects the following
format:

```yml
# ./users_endpoint.yml
base_url: https://example.com
headers: # headers to include for every request in this file
  AUTHENTICATION_TOKEN: 123456

users: # yurl users_endpoint users
  path: /users
  method: GET
  query_str:
    age: 30
    created_on: 30/02/1992

create_user: # yurl users_endpoint create_user
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

The most basic usage is ```yurl yml_file request_name``` which allows you to
make a single request defined in the ```yml_file```.

## Developing

## Contributing
