#!/bin/bash

version=$1
arch=${2:-amd64}

# build go api program
cd ../api
make clean $arch

# build web app
cd ../web
npm run build

cd ../build

# remove docker image if exists
docker rmi -f baizeai/baize-geekai-api:$version-$arch
# build docker image for geekai-go
docker build -t baizeai/baize-geekai-api:$version-$arch -f dockerfile-api-go ../

# build docker image for geekai-web
docker rmi -f baizeai/baize-geekai-web:$version-$arch
docker build --platform linux/amd64 -t baizeai/baize-geekai-web:$version-$arch -f dockerfile-vue ../

if [ "$3" = "push" ];then
  docker push baizeai/baize-geekai-api:$version-$arch
  docker push baizeai/baize-geekai-web:$version-$arch
fi
