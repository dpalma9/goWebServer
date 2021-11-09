# Overview

# Run

Build:

```bash
$ docker build -t gowebserver -f Dockerfile .
```
And run:

```bash
$ docker run -d --rm -p 8080:8080 --name goserver gowebserver
```
