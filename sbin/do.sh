#!/bin/bash

BASEDIR="$(dirname "$(perl -e 'use Cwd qw/abs_path/; print abs_path($ARGV[0]);' "$0")")"

cd "${BASEDIR}"/..

#
PROJECT_NAME=${PROJECT_NAME:-golang-learn}
# Docker
DOCKER_REGISTRY=
# Dockerfiles
DOCKER_BASE_DOCKERFILE=./sbin/docker/base-alpine.dockerfile
DOCKER_FINAL_DOCKERFILE=./sbin/docker/final.dockerfile
# Docker compose (here we split between development environment and a deployment environment)
if [[ -d ./vendor/ && -f docker-compose.dev.yml ]]; then
	DOCKER_COMPOSE_CMD="docker-compose -f docker-compose.dev.yml"
else
	DOCKER_COMPOSE_CMD="docker-compose"
fi

set -e

#===============================================================================
function __create_build_properties {
#===============================================================================
	local GIT_HASH=$(git rev-parse --short HEAD)
	local VERSION=${VERSION:-build-$GIT_HASH}
	local VERSION=$(echo $VERSION | sed -E s/[^\._a-zA-Z0-9]+/-/g | sed -E s/^-+\|-+$//g | tr A-Z a-z)
	local DOCKER_IMAGE=${PROJECT_NAME}:$VERSION

	# The build properties is essencially a list of ENV vars for feeding subsequent
	# CD (Jenkins) jobs
	echo '::::::: create build properties'
	cat <<EOF > build.properties
GIT_HASH=${GIT_HASH}
PROJECT_NAME=${PROJECT_NAME}
DOCKER_IMAGE=${DOCKER_IMAGE}
VERSION=${VERSION}
EOF
}

#===============================================================================
function __load_build_properties {
#===============================================================================
	if [ -f build.properties ]; then
		source build.properties
	fi

	# make sure certain env vars are defined
	for VARNAME in GIT_HASH PROJECT_NAME VERSION DOCKER_IMAGE; do
		if [[ -z ${!VARNAME} ]]; then
			echo "::eee:: missing build property: $VARNAME"
			exit 127
		fi
	done
}

#===============================================================================
function __remove_container_by_name {
#===============================================================================
	local CONTAINER_NAME=$1

	local container_ids=$(docker container ls -a -q -f name=${CONTAINER_NAME})
	if [ ! -z "${container_ids}" ]; then
		docker container rm ${container_ids}
	fi
}

#===============================================================================
function __build {
#===============================================================================
	__create_build_properties
	__load_build_properties

	echo "::::::: building the docker image used for building the project"
	__build_base_image

	# run the image so it builds the source code and places the executable
	echo "::::::: building the project"
	__remove_container_by_name ${PROJECT_NAME}_${VERSION}_build

	_HAVE_BUILD=0
	for BIN_FOLDER in ./bin/*; do
		if [[ -d "$BIN_FOLDER" && -f "$BIN_FOLDER/main.go" ]]; then
			BIN_NAME=$(basename $BIN_FOLDER)
			echo "::::::: building service ${BIN_NAME}"
			docker run \
				--name ${PROJECT_NAME}_${VERSION}_build \
				-e CGO_ENABLED=0 \
				-e GOOS=linux \
				${DOCKER_REGISTRY}/${PROJECT_NAME}:base \
				bash ./sbin/build.sh /bin/${BIN_NAME} ${BIN_FOLDER}
				# build -a -installsuffix cgo -o /bin/${BIN_NAME} ${BIN_FOLDER}
			echo "::::::: copying the program to dist folder"
			if [ ! -d "./dist" ]; then
				mkdir "./dist"
			fi
			docker cp ${PROJECT_NAME}_${VERSION}_build:/bin/${BIN_NAME} dist/${BIN_NAME}
			echo "::::::: remove the ${BIN_NAME} build container"
			__remove_container_by_name ${PROJECT_NAME}_${VERSION}_build
			_HAVE_BUILD=1
		fi
	done

	if [ "$_HAVE_BUILD" == "0" ]; then
		echo "Nothing to build"
		exit 1
	fi

	echo "::::::: creating the final docker image"
	docker build \
		-f ${DOCKER_FINAL_DOCKERFILE} \
		-t ${DOCKER_REGISTRY}/${DOCKER_IMAGE} .

	echo "::::::: tagging the docker image as the latest"
	docker tag ${DOCKER_REGISTRY}/${DOCKER_IMAGE} ${DOCKER_REGISTRY}/${PROJECT_NAME}:latest
}


#===============================================================================
function __build_base_image {
#===============================================================================
	docker build -f ${DOCKER_BASE_DOCKERFILE} -t ${DOCKER_REGISTRY}/${PROJECT_NAME}:base .
}

#===============================================================================
function __start_integration_environment {
#===============================================================================
	echo "TODO"
}

#===============================================================================
function __stop_integration_environment {
#===============================================================================
	${DOCKER_COMPOSE_CMD} down --rmi local -v
}

#===============================================================================
function __run_unit_tests {
#===============================================================================
	go test $(glide nv)
}

#===============================================================================
function __run_docker_compose {
#===============================================================================
	${DOCKER_COMPOSE_CMD} $@
}

#===============================================================================
function __run_coverage {
#===============================================================================
	mkdir -p ./sbin/coverage

	set +e

	echo "mode: count" > ./sbin/coverage.out

	export env=local-tests

	local paths_to_packages
	if [ -z "$1" ]; then
		paths_to_packages=$(glide nv)
	else
		for _pkg in $@; do
			paths_to_packages="$paths_to_packages ./$_pkg/..."
		done
	fi

	for pkg in $(go list $paths_to_packages); do
		_pkg=$(echo $pkg | sed 's/\//_/g')
		if ! go test -coverprofile ./sbin/coverage/$_pkg.tmp $pkg; then
			echo "Failed to run tests!"
		fi
		if [ -f ./sbin/coverage/$_pkg.tmp ]; then
			tail -n +2 ./sbin/coverage/$_pkg.tmp > ./sbin/coverage/$_pkg
			rm ./sbin/coverage/$_pkg.tmp
		fi
	done

	for _file in ./sbin/coverage/*; do
		cat $_file >> ./sbin/coverage.out
	done

	go tool cover -html=./sbin/coverage.out -o ./sbin/coverage.html

	set -e
}

op=$1; shift

case "$op" in
	'build')
		__build
	;;

	'test')
		__run_unit_tests
	;;

	'clean')
		__run_docker_compose down --rmi all -v
	;;

	'distclean')
		rm ./dist/*
	;;

	'coverage')
		__run_coverage $@
	;;

	'compose')
		${DOCKER_COMPOSE_CMD} $@
	;;
esac
