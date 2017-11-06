Requirement
====

Demonstrate the automation of the spin-up and installation of a multi-machine n-tier architecture. It is up to your interpretation which layers you would like to introduce, provided they can all be automatically started from a single command and serve some basic data.

Requirements:

Using automation we want to spin up an environment which will allow us to connect to a web server on port 80 and serve a bit of simple HTML content from a data storage source. You will be required to write a small application in the language/framework of your choice to connect to the database, query it, and return the result to the user.

Criteria:

Developed within a git repository with frequent commits Vagrant machine to spin up/down a multi-machine environment. We require you to make a selection of technical choices you're most comfortable with, whether it be Linux vs. Windows, Java vs. Ruby, Puppet vs Ansible etc. Choose the platforms, frameworks and languages you're most familiar with.

Process:

Please establish a git repository from the beginning and demonstrate frequent commits. We would like to be able to spin up the environment using a single command of `vagrant up`. We should only have to then connect to the instances IP address with a browser to see the resulting data returned. We would then like you to demonstrate the stack and talk us through the processes and choices you made via documentations. Please upload your results and artefacts in a repository and share the link to review them.

Bonus points for including Docker and Ansible
