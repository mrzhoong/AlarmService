#!/bin/sh
set -e

#install requirements.sh
echo "install depends..."
rpm -Uvh --nodeps --force /opt/yum/*.rpm
echo "depends done..."

#yum -y install redis
#yum -y install mariadb mariadb-server
#yum -y install mariadb*

#��װredis
systemctl start redis.service
#systemctl enable redis.service
#systemctl is-enabled redis.service

#��װmysql
systemctl start mariadb
#systemctl enable mariadb.service
#systemctl is-enabled mariadb.service
mysql -uroot < /opt/golang/src/AlarmService/conf/alarm_user.sql
systemctl stop mariadb
sleep 5
systemctl start mariadb

#�رշ���ǽ
systemctl stop firewalld
systemctl disable firewalld