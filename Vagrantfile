
Vagrant.configure(2) do |config|

    config.vm.box = "centos/7"
    # Sync folder
    config.vm.synced_folder ".", "/home/vagrant/sync"

    # Web server
    config.vm.define "web" do |web|
        web.vm.hostname = "web"
        web.vm.network "private_network", ip: "192.168.51.3"
        web.vm.network "forwarded_port", guest: 8080, host: 8333
        web.vm.provision :shell, inline: 'cat /home/vagrant/sync/ssh_keys/id_rsa.pub >> /home/vagrant/.ssh/authorized_keys'
    end

    # Database server
    config.vm.define "db" do |db|
        db.vm.hostname = "db"
        db.vm.network "private_network", ip: "192.168.51.4"
        db.vm.network "forwarded_port", guest: 3306, host: 33306
        db.vm.provision :shell, inline: 'cat /home/vagrant/sync/ssh_keys/id_rsa.pub >> /home/vagrant/.ssh/authorized_keys'
    end

    # Ansible command server
    config.vm.define "ansible", primary: true do |ansible|
        ansible.vm.hostname = "ansible"
        ansible.vm.network "private_network", ip: "192.168.51.2"        
        config.vm.provision "shell", path: "./provision.sh", privileged: true
    end

end
