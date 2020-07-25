pathDir=${1}
fileName=${2}

echo 'Running into Docker container...'

echo "Directory and filename: ${pathDir}"

docker run --rm -v ${pathDir}:/usr/src/app/${fileName} -w /usr/src/app golang:1.14 go run ${fileName}