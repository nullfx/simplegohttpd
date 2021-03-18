# Another Simple go httpd Web Server
A simple go based http daemon that reads serves file content. 
This is more of a learning exercise than anything useful.

Loads static content from local directory and writes it out based on uri path. If no file matches URI path, returns a 404 page.

Includes [SB Admin](https://github.com/startbootstrap/startbootstrap-sb-admin) bootstrap template as content that it loads and serves.

## Build 
```
go build httpd.go
```

### Run
browse http://localhost:82/index.html or http://localhost/foo