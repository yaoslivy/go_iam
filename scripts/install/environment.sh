#!/usr/bin/env bash


# IAM project source code root directory
IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

# Generate file storage directory
LOCAL_OUTPUT_ROOT="${IAM_ROOT}/${OUT_DIR:-_output}"

readonly PASSWORD=${PASSWORD:-'123'}

# Linux username and password
readonly LINUX_USERNAME=${LINUX_USERNAME:-going}
readonly LINUX_PASSWORD=${LINUX_PASSWORD:-${PASSWORD}}

# Install directory
readonly INSTALL_DIR=${INSTALL_DIR:-/tmp/installation}
mkdir -p ${INSTALL_DIR}
readonly ENV_FILE=${IAM_ROOT}/scripts/install/environment.sh

# MariaDB configuration information
readonly MARIADB_ADMIN_USERNAME=${MARIADB_ADMIN_USERNAME:-root} 
readonly MARIADB_ADMIN_PASSWORD=${MARIADB_ADMIN_PASSWORD:-${PASSWORD}} 
readonly MARIADB_HOST=${MARIADB_HOST:-127.0.0.1:3306} 
readonly MARIADB_DATABASE=${MARIADB_DATABASE:-iam} # MariaDB iam database name
readonly MARIADB_USERNAME=${MARIADB_USERNAME:-iam} # iam database username
readonly MARIADB_PASSWORD=${MARIADB_PASSWORD:-${PASSWORD}} 

# Redis configuration information
readonly REDIS_HOST=${REDIS_HOST:-127.0.0.1} # Redis address 
readonly REDIS_PORT=${REDIS_PORT:-6379} # Redis port
readonly REDIS_USERNAME=${REDIS_USERNAME:-''}   
readonly REDIS_PASSWORD=${REDIS_PASSWORD:-${PASSWORD}}

# MongoDB configuration information
readonly MONGO_ADMIN_USERNAME=${MONGO_ADMIN_USERNAME:-root} 
readonly MONGO_ADMIN_PASSWORD=${MONGO_ADMIN_PASSWORD:-${PASSWORD}} 
readonly MONGO_HOST=${MONGO_HOST:-127.0.0.1} # MongoDB address
readonly MONGO_PORT=${MONGO_PORT:-27017} # MongoDB port
readonly MONGO_USERNAME=${MONGO_USERNAME:-iam} 
readonly MONGO_PASSWORD=${MONGO_PASSWORD:-${PASSWORD}} 

# iam configuration information
readonly IAM_DATA_DIR=${IAM_DATA_DIR:-$HOME/iam/data/iam} # iam data catalog of each component
readonly IAM_INSTALL_DIR=${IAM_INSTALL_DIR:-$HOME/iam/opt/iam} # iam installation file storage directory
readonly IAM_CONFIG_DIR=${IAM_CONFIG_DIR:-$HOME/iam/etc/iam} # iam config files
readonly IAM_LOG_DIR=${IAM_LOG_DIR:-$HOME/iam/var/log/iam} # iam log files
readonly CA_FILE=${CA_FILE:-${IAM_CONFIG_DIR}/cert/ca.pem} # CA

# iam-apiserver configuration information
readonly IAM_APISERVER_HOST=${IAM_APISERVER_HOST:-127.0.0.1} # iam-apiserver deployment machine ip address
readonly IAM_APISERVER_GRPC_BIND_ADDRESS=${IAM_APISERVER_GRPC_BIND_ADDRESS:-0.0.0.0}
readonly IAM_APISERVER_GRPC_BIND_PORT=${IAM_APISERVER_GRPC_BIND_PORT:-8081}
readonly IAM_APISERVER_INSECURE_BIND_ADDRESS=${IAM_APISERVER_INSECURE_BIND_ADDRESS:-127.0.0.1}
readonly IAM_APISERVER_INSECURE_BIND_PORT=${IAM_APISERVER_INSECURE_BIND_PORT:-8080}
readonly IAM_APISERVER_SECURE_BIND_ADDRESS=${IAM_APISERVER_SECURE_BIND_ADDRESS:-0.0.0.0}
readonly IAM_APISERVER_SECURE_BIND_PORT=${IAM_APISERVER_SECURE_BIND_PORT:-8443}
readonly IAM_APISERVER_SECURE_TLS_CERT_KEY_CERT_FILE=${IAM_APISERVER_SECURE_TLS_CERT_KEY_CERT_FILE:-${IAM_CONFIG_DIR}/cert/iam-apiserver.pem}
readonly IAM_APISERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE=${IAM_APISERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE:-${IAM_CONFIG_DIR}/cert/iam-apiserver-key.pem}

# iam-authz-server configuration information
readonly IAM_AUTHZ_SERVER_HOST=${IAM_AUTHZ_SERVER_HOST:-127.0.0.1} # iam-authz-server deployment machine ip address
readonly IAM_AUTHZ_SERVER_INSECURE_BIND_ADDRESS=${IAM_AUTHZ_SERVER_INSECURE_BIND_ADDRESS:-127.0.0.1}
readonly IAM_AUTHZ_SERVER_INSECURE_BIND_PORT=${IAM_AUTHZ_SERVER_INSECURE_BIND_PORT:-9090}
readonly IAM_AUTHZ_SERVER_SECURE_BIND_ADDRESS=${IAM_AUTHZ_SERVER_SECURE_BIND_ADDRESS:-0.0.0.0}
readonly IAM_AUTHZ_SERVER_SECURE_BIND_PORT=${IAM_AUTHZ_SERVER_SECURE_BIND_PORT:-9443}
readonly IAM_AUTHZ_SERVER_SECURE_TLS_CERT_KEY_CERT_FILE=${IAM_AUTHZ_SERVER_SECURE_TLS_CERT_KEY_CERT_FILE:-${IAM_CONFIG_DIR}/cert/iam-authz-server.pem}
readonly IAM_AUTHZ_SERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE=${IAM_AUTHZ_SERVER_SECURE_TLS_CERT_KEY_PRIVATE_KEY_FILE:-${IAM_CONFIG_DIR}/cert/iam-authz-server-key.pem}
readonly IAM_AUTHZ_SERVER_CLIENT_CA_FILE=${IAM_AUTHZ_SERVER_CLIENT_CA_FILE:-${CA_FILE}}
readonly IAM_AUTHZ_SERVER_RPCSERVER=${IAM_AUTHZ_SERVER_RPCSERVER:-${IAM_APISERVER_HOST}:${IAM_APISERVER_GRPC_BIND_PORT}}

# iam-pump configuration information
readonly IAM_PUMP_HOST=${IAM_PUMP_HOST:-127.0.0.1} # iam-pump deployment machine ip address
readonly IAM_PUMP_COLLECTION_NAME=${IAM_PUMP_COLLECTION_NAME:-iam_analytics}
readonly IAM_PUMP_MONGO_URL=${IAM_PUMP_MONGO_URL:-mongodb://${MONGO_USERNAME}:${MONGO_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/${IAM_PUMP_COLLECTION_NAME}?authSource=${IAM_PUMP_COLLECTION_NAME}}

# iam-watcher configuration information
readonly IAM_WATCHER_HOST=${IAM_WATCHER_HOST:-127.0.0.1} # iam-watcher deployment machine ip address

# iamctl configuration information
readonly CONFIG_USER_USERNAME=${CONFIG_USER_USERNAME:-admin}
readonly CONFIG_USER_PASSWORD=${CONFIG_USER_PASSWORD:-Admin@2021}
readonly CONFIG_USER_CLIENT_CERTIFICATE=${CONFIG_USER_CLIENT_CERTIFICATE:-${HOME}/iam/cert/admin.pem}
readonly CONFIG_USER_CLIENT_KEY=${CONFIG_USER_CLIENT_KEY:-${HOME}/iam/cert/admin-key.pem}
readonly CONFIG_SERVER_ADDRESS=${CONFIG_SERVER_ADDRESS:-${IAM_APISERVER_HOST}:${IAM_APISERVER_SECURE_BIND_PORT}}
readonly CONFIG_SERVER_CERTIFICATE_AUTHORITY=${CONFIG_SERVER_CERTIFICATE_AUTHORITY:-${CA_FILE}}

