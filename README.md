# lolproxy

"some fuckshit that changes your response codes"

## running
```
Usage of ./lolproxy:
  -host string
        host to bind to (default "localhost:6969")
  -only404
        only change status code to 404
  -upstream string
        upstream host to proxy to (required)
```
example: `./lolproxy --upstream https://example.com`

## docker
the docker image binds to `0.0.0.0:6969` by default, so you'll need to expose the port (`--port 127.0.0.1:6969:6969`).
you'll also need to set the upstream as the command

example: `docker run --rm -it -p 127.0.0.1:6969:6969 docker.io/galenguyer/lolproxy --upstream https://example.com`

### docker-compose
a compose file is also provided for your inconvenience. it can be ran with `docker-compose up --build`

you can also command out the `build` directive and uncomment the `image` directive to use the latest image from docker hub