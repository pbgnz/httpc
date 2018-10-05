# httpc
httpc is a curl - like application but supports HTTP protocol only

## Requirements
1. Go 1.7 or later

## Detailed Usage

### General

``` bash
httpc help

httpc is a curl-like application but supports HTTP protocol only.
Usage:
    httpc [arguments] command
The commands are:
    get     executes a HTTP GET request and prints the response.
    post    executes a HTTP POST request and prints the response.
    help    prints this screen.

Use "httpc help [command]" for more information about a command.
```

### Get Usage

``` bash
httpc help get

usage: httpc [-v] [-h key:value] get URL

Get executes a HTTP GET request for a given URL.

    -v             Prints the detail of the response such as protocol,status,and headers.
    -h key:value   Associates headers to HTTP Request with the format 'key:value'.
```

### Post Usage

``` bash
httpc help post

usage: httpc [-v] [-h key:value] [-d inline-data] [-f file] [-o output] post URL

Post executes a HTTP POST request for a given URL with inline data or from file.

    -v             Prints the detail of the response such as protocol, status, 
and headers.
    -h key:value   Associates headers to HTTP Request with the format 'key:value'.
    -d string      Associates an inline data to the body HTTP POST request.
    -f file        Associates the content of a file to the body HTTP POST request.
    -o output       Associates the content of a file to the body HTTP POST request.

Either [-d] or [-f] can be used but not both.
```

## Examples

### Get with query parameters

httpc get 'http://httpbin.org/get?course=networking&assignment=1'

Output:
``` bash
{
    "args": {
        "assignment": "1",
        "course": "networking"
    },
    "headers": {
        "Host": "httpbin.org",
        "User-Agent": "Concordia-HTTP/1.0"
    },
    "url": "http://httpbin.org/get?course=networking&assignment=1"
}
```

### Get with query parameters

httpc -v get 'http://httpbin.org/get?course=networking&assignment=1'

Output:
``` bash
HTTP/1.1 200 OK
Server: nginx
Date: Fri, 1 Sep 2017 14:52:12 GMT
Content-Type: application/json
Content-Length: 255
Connection: close
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true
{
    "args": {
        "assignment": "1",
        "course": "networking"
    },
    "headers": {
        "Host": "httpbin.org",
        "User-Agent": "Concordia-HTTP/1.0"
    },
    "url": "http://httpbin.org/get?course=networking&assignment=1"
}
```

### Post with inline data

httpc -h Content-Type:application/json --d '{"Assignment": 1}' post http://httpbin.org/post

Output:
``` bash
{
    "args": {},
    "data": "{\"Assignment\": 1}",
    "files": {},
    "form": {},
    "headers": {
        "Content-Length": "17",
        "Content-Type": "application/json",
        "Host": "httpbin.org",
        "User-Agent": "Concordia-HTTP/1.0"
    },
    "json": {
        "Assignment": 1
    },
    "url": "http://httpbin.org/post"
}
```

### Post content of a file and receive the HTTP response in another file

httpc -f input.txt -o output.txt post 'http://httpbin.org/post'

Where input.txt has: 

```txt
{"Assignment": 1}
```

Output:
```bash
{
  "args": {},
  "data": "{\"Assignment\": 1}",
  "files": {},
  "form": {},
  "headers": {
    "Connection": "close",
    "Content-Length": "17",
    "Host": "httpbin.org"
  },
  "json": {
    "Assignment": 1
  },
  "origin": "135.19.219.135",
  "url": "http://httpbin.org/post"
}

Outputing HTTP response to: output.txt
```