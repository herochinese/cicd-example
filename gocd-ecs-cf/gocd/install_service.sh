#!/bin/bash

sudo cp /home/ec2-user/inst/gocd-server.service /etc/systemd/system/.
sudo chmod 664 /etc/systemd/system/gocd-server.service
sudo systemctl daemon-reload
sudo systemctl enable gocd-server.service
sudo systemctl start gocd-server.service

sudo cp /home/ec2-user/inst/gocd-agent.service /etc/systemd/system/.
sudo chmod 664 /etc/systemd/system/gocd-agent.service
sudo systemctl daemon-reload
sudo systemctl enable gocd-agent.service
sudo systemctl start gocd-agent.service
