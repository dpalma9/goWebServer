#HOW TO TEST/DEV
You can dev and test this app by using the container ready for that porpuse.

```bash
$ docker build -t gowebserver_dev -f Dockerfile.dev .
$ docker run -d --mount type=bind,src="$(pwd)",dst=/web --name dev gowebserver_dev
```
