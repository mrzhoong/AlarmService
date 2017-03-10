#!/bin/sh
set -e

#install requirements.sh
echo "install depends..."
rpm -Uvh --nodeps --force /opt/yum/*.rpm
echo "depends done..."

#yum -y install redis
#yum -y install mariadb mariadb-server
#yum -y install mariadb*

#安装redis
systemctl start redis.service
#systemctl enable redis.service
#systemctl is-enabled redis.service

#安装mysql
systemctl start mariadb
#systemctl enable mariadb.service
#systemctl is-enabled mariadb.service
mysql -uroot < /opt/golang/src/AlarmService/conf/alarm_user.sql
systemctl stop mariadb
sleep 5
systemctl start mariadb

#关闭防火墙
systemctl stop firewalld
systemctl disable firewalld