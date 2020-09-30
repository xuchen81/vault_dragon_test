deploymenthost: local
user: vagrant
server_name: local.vaultdragon.com

db_name: "{{project}}"
db_password: "password"
db_hostname: "localhost"
db_port: "3306"

env:
  VD_ENV: "local"
