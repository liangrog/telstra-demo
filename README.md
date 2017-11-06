Telstra Test Demo
====
This repo contains configurationf files and codes providing an example of provisioning a  multi-tier server architecture environment using Vagrant, Virturalbox, Ansible and Docker.

The original requirment is [here](requirement.md)

Prerequisite
----
You will need to have Vagrant and Virturalbox installed in your host machine. You can find the details of installation below:

- [Installing Vagrant](https://www.vagrantup.com/docs/installation/)
- [Installing Virturalbox](https://www.virtualbox.org)

Spin Up
----
Simple run below command on your console:

    $ vagrant up

After all the process is done, you can visit the result on your local browser using address `http://localhost:8333`

Note: I chose to use port `8333` instead of requirement of port `80` to avoid port conflict on the host machine.

Overview
----
Architeture:

                                           Host Machine

                                                     +
                                                     |8333
        +-----------------------------------------------------------------------------+
        |                                            |8080                            |
        |                                            |                                |
        |                                +-----------+------------+                   |
        |     +-----------+              |                        |                   |
        |     |           |       +----> |    WEB Server          |                   |
        |     |           |       |      |   192.168.51.3         |                   |
        |     |           |       |      +-----------+------------+                   |
        |     |           |       |                  |                                |
        |     |   Ansible |  ssh  |                  |                     VMs        |
        |     |   Control +-------+                  | 3306                           |
        |     |   Server  |       |                  |                                |
        |     |           |       |      +-----------+------------+                   |
        |     |           |       |      |                        |                   |
        |     +-----------+       +----> |      Mysql Server      |                   |
        |      192.168.51.2              |    192.168.51.4        |                   |
        |                                +------------------------+                   |
        |                                            |3306                            |
        +-----------------------------------------------------------------------------+
                                                     |33306
                                                     +


There are three VMs in this example:

- Ansible Control which works as the provisioning server using Ansible
- Web server. It has a simple web app runs in a docker container, which listens to incoming request via port `8080` then serve it with result from query to a table in the mysql server
- Mysql server runs mysql and has a demo database contains one table that has a list of movies

Provisioning work flow:

                                  vagrant up
                                      +
                                      |
                                      v
                           Download centos7 vagrant box
                                      +
                                      |
                                      v
                             Create web server vm
                                      +
                                      |
                                      v
                             Create mysql server vm
                                      +
                                      |
                                      v
                     Create Ansible control vm then:
                     1. run provision.sh to setup Ansible runtime
                     2. run Ansible playbook to provision web server and mysql server



                Web Sever provisioning                      Mysql server provisioning
                         +                                             +
                         |                                             |
                         v                                             v
           Install Docker and Docker compose           Install mariadb sever and setup users
                         +                                             +
                         |                                             |
                         v                                             v
        Run docker compose to spin up web app                 Load demo database

