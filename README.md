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

# Dev with VSCode

.devcontainer.json is a file that VSCode will read if you install the Dev Containers extensions.
This file is used to pre-configure the VSCode with all the plugins you need on your dev container.
