# wget-server
This is a simple HTTP server for static files scraped by wget.

For example, if you copy a website this way

```shell
wget -k -r -p -l 0 http://example.org
```

you will get a directory called "example.org" with all the files that wget got.

Then you can bring up the server like this:

```
./main -base example.org
```

Browsing to http://localhost:3000/ should then display the copied site.
