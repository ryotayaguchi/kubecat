#!/bin/bash
set -eu

function f_build(){
    cd ${BASE_DIR}
    docker build -t ${NAME} .
}

function f_undeploy(){
    C=$(docker ps -a -q --filter ancestor=${NAME})
    if [[ `echo ${C}` ]]; then
        echo "cURL /prestop/: $(curl -s localhost:8080/prestop/)"
        docker rm $(docker stop ${C})
    fi
    I=$(docker images -a -q --filter reference=${NAME})
    if [[ `echo ${I}` ]]; then
        docker rmi ${I}
    fi
}

function f_deploy(){
    f_undeploy
    f_build

    EV=(
    "-e ENABLE_LIVENESS=true"
    "-e ENABLE_READINESS=true"
    "-e TIME_UNTIL_LIVENESS_TRUE=5"
    "-e TIME_UNTIL_READINESS_TRUE=10"
    "-e TIME_FOR_PRESTOP=5"
    "-e TIME_UNTIL_AUTO_STOP=60"
    "-e PORT_ADMIN=8080"
    )
    docker run -d -p 8080:8080 --name ${NAME} ${EV[@]} ${NAME}
}

function f_test(){
    echo "cURL /uptime/: $(curl -s localhost:8080/uptime/)"
    echo "cURL /liveness/: $(curl -s localhost:8080/liveness/)"
    echo "cURL /readiness/: $(curl -s localhost:8080/readiness/)"
}

NAME=kubecat
BASE_DIR=$(cd `dirname $0`; pwd)

[[ $# != 1 ]] && exit 1
case ${1} in
    "build")    f_build;;
    "undeploy") f_undeploy;;
    "deploy")   f_deploy;;
    "test")     f_test;;
    *)          ;;
esac
