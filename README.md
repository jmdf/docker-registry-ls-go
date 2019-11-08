# Docker registry list

Tool to list information from Docker registries

## Instructions

Print registry catalog
```
docker-registry-ls -s https://exampleServer:5000
```

Print tags from a specific repository
```
docker-registry-ls -s https://exampleServer:5000 -r group/image
```

Note: `-s https://exampleServer:5000` can be replaced by setting `DOCKER_REGISTRY` environment variable