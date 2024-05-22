# Connect Issues
https://stackoverflow.com/questions/36732875/cant-connect-to-public-ip-for-ec2-instance  

If my Mac's IP changes, I have to update the firewall rule for the instance, or set it to 0.0.0.0/0

**To see what app is listening on a por**  
sudo netstat -tulpn | grep LISTEN  

**[ec2-user@ip-172-31-20-12 ~]$ netstat -tulpn | grep LISTEN**  
(Not all processes could be identified, non-owned process info
will not be shown, you would have to be root to see it all.)
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      -                   
tcp6       0      0 :::22                   :::*                    LISTEN      -                   
tcp6       0      0 :::8080                 :::*                    LISTEN      31553/./linux_binar **

**Deborahs-iMac:vtserver debjo$ nc -v 44.234.131.118 8080**  
Connection to 44.234.131.118 port 8080 [tcp/http-alt] succeeded!

### On my Mac in Chrome  
**http://44.234.131.118:8080/**

**Output**  
*Web server is running!  
To test: http://localhost:8080/upload/vtserver_test*


**http://44.234.131.118:8080/upload/vtserver_test**  

**Output**  
*Filename  
    ./vtserver_test  
Sha256  
    c449a93fb4598c689a895784a4ffe20cad0b8e2e0af6c204090cbf51e927330c    
Bytes  
    7f454c460101010000000000000000000200030001000000c0 ...*


### Resources
https://www.cyberciti.biz/faq/unix-linux-check-if-port-is-in-use-command/
https://stackoverflow.com/questions/36732875/cant-connect-to-public-ip-for-ec2-instance

# Build
**Build LInux on Mac**  
*env GOOS=linux GOARCH=386  go build*

# Upload to AWS
**Deborahs-iMac:~ debjo$ scp -i ".ssh/vtserver-key-ED25519.pem" ~/vtserver  ec2-user@ec2-44-234-131-118.us-west-2.compute.amazonaws.com:linux_binary/vtserver**  
vtserver  

**Deborahs-iMac:~ debjo$ scp -i "vtserver-key-ED25519.pem" ~/vtserver  ec2-user@ec2-44-234-131-118.us-west-2.compute.amazonaws.com:linux_binary/vtserver**  
vtserver



# SSH to EC2 in AWS
**Deborahs-iMac:~ debjo$ ssh -i ".ssh/vtserver-key-ED25519.pem" ec2-user@ec2-34-223-223-68.us-west-2.compute.amazonaws.com**  
   ,     #_
   ~\_  ####_        Amazon Linux 2023
  ~~  \_#####\
  ~~     \###|
  ~~       \#/ ___   https://aws.amazon.com/linux/amazon-linux-2023
   ~~       V~' '->
    ~~~         /
     ~~._.   _/
       _/ _/
      _/m/'

**[ec2-user@ip-172-31-20-12 ~]$ printenv**    
SHELL=/bin/bash  
HISTCONTROL=ignoredups  
SYSTEMD_COLORS=false  
HISTSIZE=1000  
HOSTNAME=ip-172-31-20-12.us-west-2.compute.internal  
PWD=/home/ec2-user  
LOGNAME=ec2-user  
XDG_SESSION_TYPE=tty  
MOTD_SHOWN=pam  
HOME=/home/ec2-user  
LANG=C.UTF-8  

SSH_CONNECTION=50.35.90.164 60755 172.31.20.12 22    
XDG_SESSION_CLASS=user  
SELINUX_ROLE_REQUESTED=  
TERM=xterm-256color  
LESSOPEN=||/usr/bin/lesspipe.sh %s  
USER=ec2-user  
SELINUX_USE_CURRENT_RANGE=  
SHLVL=1  
XDG_SESSION_ID=1  
XDG_RUNTIME_DIR=/run/user/1000  
S_COLORS=auto  
SSH_CLIENT=50.35.90.164 60755 22  
which_declare=declare -f  
PATH=/home/ec2-user/.local/bin:/home/ec2-user/bin:/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin  
SELINUX_LEVEL_REQUESTED=  
DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus  
MAIL=/var/spool/mail/ec2-user  
SSH_TTY=/dev/pts/0  
BASH_FUNC_which%%=() {  ( alias;    
eval ${which_declare} ) | /usr/bin/which --tty-only --read-alias --read-functions --show-tilde --show-dot "$@"
}  
_=/usr/bin/printenv  


**[ec2-user@ip-172-31-20-12 ~]$ ls -laf  /**   
.  ..  boot  dev  etc  local  proc  run  sys  tmp  var  usr  bin  sbin  lib  lib64  home  media  mnt  opt  root  srv



## Runcloud?
https://runcloud.io/blog/aws