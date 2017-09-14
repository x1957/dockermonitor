set -e

cd `dirname $0`
tag=docker-exporter
docker build -t $tag .
docker run -it $tag
