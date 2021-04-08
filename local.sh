#!/bin/bash
set -eu

function f_build(){
    if [[ `ls -A ${BIN_DIR}` ]]; then
        rm -r ${BIN_DIR}/*
    fi
    cd ${SRC_DIR}
    go build -o ${BIN_DIR}
}

function f_undeploy(){
    if [[ `pgrep -x ${NAME}` ]]; then
        echo "cURL /prestop/: $(curl -s localhost:8080/prestop/)"
        pkill -x ${NAME}
    fi
}

function f_deploy(){
    f_undeploy
    f_build

    export ENABLE_LIVENESS=true
    export ENABLE_READINESS=true
    export TIME_UNTIL_LIVENESS_TRUE=5
    export TIME_UNTIL_READINESS_TRUE=10
    export TIME_FOR_PRESTOP=5
    export TIME_UNTIL_AUTO_STOP=60
    export PORT_ADMIN=8080
    ${BIN_DIR}/${NAME} &> /tmp/${NAME}.log &
}

function f_test(){
    echo "cURL /uptime/: $(curl -s localhost:8080/uptime/)"
    echo "cURL /liveness/: $(curl -s localhost:8080/liveness/)"
    echo "cURL /readiness/: $(curl -s localhost:8080/readiness/)"
}


NAME=kubecat
BASE_DIR=$(cd `dirname $0`; pwd)
SRC_DIR=${BASE_DIR}/src/${NAME}
BIN_DIR=${BASE_DIR}/bin

[[ $# != 1 ]] && exit 1
case ${1} in
    "build")    f_build;;
    "undeploy") f_undeploy;;
    "deploy")   f_deploy;;
    "test")     f_test;;
    *)          ;;
esac
