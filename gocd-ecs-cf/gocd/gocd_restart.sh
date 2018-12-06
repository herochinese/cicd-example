#!/bin/bash
# chmod 664 /etc/systemd/system/gocd-server.service
# systemctl daemon-reload
# systemctl enable gocd-server.service
# systemctl start gocd-server.service

source /home/ec2-user/.bash_profile

thing=$1

kill_gocd()
{
  p=$@
  if [ -z $p ]
  then
    echo "Nothing is running."
  else
    echo "kill -9 $p"
    kill -9 $p
  fi
}

if [ $thing = "server" ]
then
  echo "Work for server MoDe."
  ps -ef|grep cruise.server|grep -v grep|grep java
  pid=`ps -ef|grep cruise.server|grep -v grep|grep java|awk '{print $2}'|tr "\\n" " "`
  kill_gocd $pid
  cd /home/ec2-user/go-server-18.10.0
  ls -F
  ./server.sh
  echo $?
fi


if [ $thing = "client" ]
then
  echo "Work for client MoDe."
  ps -ef|grep gocd.agent |grep -v grep|grep java
  pid=`ps -ef|grep gocd.agent |grep -v grep|grep java|awk '{print $2}'|tr "\\n" " "`
  kill_gocd $pid
  cd /home/ec2-user/go-agent-18.10.0
  rm -rf .agent-bootstrapper.running
  ls -Fa
  ./agent.sh
fi
