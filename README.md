## A static file server made by Go for development

I named it after lite-server made in nodejs, but this is not as fancy as lite-server in a sense that it doesn't have web-socket logic to watch your changes
and cause the page to reload. It is stand alone; single binary; when you don't want a million packages and nodejs run time to serve your static files.


- run help for options
```
$ gite-server --help  
```


Cross arch build options GOOS= windows | darwin | linux:
```
$ GOOS=darwin GOARCH=amd64 go build .
```

