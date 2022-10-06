# 运行容器
## 制作、运行镜像
参照: 99homework/01httpserver/01/02cjx/README.md
### 接口测试
#### 进入容器测试
```
docker exec -it httpserver sh

/opt/modules/httpserver # curl -v localhost:8080/healthz -H "ttt:yyy"
sh: curl: not found
```
#### 宿主机测试
```
curl -v localhost:9090/healthz -H "ttt:yyy"
*   Trying 127.0.0.1:9090...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /healthz HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.68.0
> Accept: */*
> ttt:yyy
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Accept: */*
< Content-Type: application/json; charset=utf-8
< Ttt: yyy
< User-Agent: curl/7.68.0
< Version: v0.0.1
< Date: Thu, 06 Oct 2022 10:07:17 GMT
< Content-Length: 18
< 
* Connection #0 to host localhost left intact
{"message":"pong"}% 
```
### 查看容器信息
```
docker ps | grep httpserver

docker inspect httpserver|grep -i pid
            "Pid": 36309,
            "PidMode": "",
            "PidsLimit": null,
```
### 查看容器 ip
```
docker exec -it httpserver sh

/opt/modules/httpserver # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
23: eth0@if24: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

## 练习
查看当前系统的 namespace：`lsns –t <type>`
```
lsns -t net
        NS TYPE NPROCS   PID USER    NETNSID NSFS COMMAND
4026531992 net     124  9158 cjx  unassigned      /lib/systemd/systemd --user
4026533051 net       2 19842 cjx  unassigned      /usr/share/code/code --type=zygo
4026533633 net      25  9679 cjx  unassigned      /opt/google/chrome/chrome --type
4026533725 net       1  9680 cjx  unassigned      /opt/google/chrome/nacl_helper
```
查看某进程的 namespace：`ls -la /proc/<pid>/ns/`
```
sudo ls -la /proc/36309/ns/
总用量 0
dr-x--x--x 2 root root 0 10月  6 17:55 .
dr-xr-xr-x 9 root root 0 10月  6 17:55 ..
lrwxrwxrwx 1 root root 0 10月  6 18:03 cgroup -> 'cgroup:[4026531835]'
lrwxrwxrwx 1 root root 0 10月  6 17:56 ipc -> 'ipc:[4026533336]'
lrwxrwxrwx 1 root root 0 10月  6 17:56 mnt -> 'mnt:[4026533334]'
lrwxrwxrwx 1 root root 0 10月  6 17:55 net -> 'net:[4026533339]'
lrwxrwxrwx 1 root root 0 10月  6 17:56 pid -> 'pid:[4026533337]'
lrwxrwxrwx 1 root root 0 10月  6 18:03 pid_for_children -> 'pid:[4026533337]'
lrwxrwxrwx 1 root root 0 10月  6 18:03 user -> 'user:[4026531837]'
lrwxrwxrwx 1 root root 0 10月  6 17:56 uts -> 'uts:[4026533335]'
```
进入某 namespace 运行命令：`nsenter -t <pid> -n ip addr`

- -n: 指定网络 ns
```
sudo nsenter -t 36309 -n ip a 
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
23: eth0@if24: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
### 测试接口(容器内没有安装 curl，只能通过这个方法测试，否则需要把接口端口开放出来)
```
sudo nsenter -t 36309 -n curl -v localhost:8080/healthz -H "ttt:yyy"      
*   Trying 127.0.0.1:8080...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /healthz HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.68.0
> Accept: */*
> ttt:yyy
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Accept: */*
< Content-Type: application/json; charset=utf-8
< Ttt: yyy
< User-Agent: curl/7.68.0
< Version: v0.0.1
< Date: Thu, 06 Oct 2022 10:05:41 GMT
< Content-Length: 18
< 
* Connection #0 to host localhost left intact
{"message":"pong"}%                             
```

## 练习2
在新 network namespace 执行 sleep 指令：`sudo unshare -fn sleep 60`

- -fn: 放入独立的网络 ns

查看进程信息`sudo ps -ef|grep 'sleep 60'`
```
sudo ps -ef|grep 'sleep 60'
root       42079   26372  0 18:19 pts/1    00:00:00 sudo unshare -fn sleep 60
root       42080   42079  0 18:19 pts/1    00:00:00 unshare -fn sleep 60
root       42081   42080  0 18:19 pts/1    00:00:00 sleep 60
cjx        42170   32973  0 18:19 pts/3    00:00:00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox sleep 60
```
进入改进程所在 Namespace 查看网络配置，与主机不一致`sudo nsenter -t 42080 -n ip a`
```
sudo nsenter -t 42080 -n ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
```
查看网络 Namespace`lsns -t net`
```
lsns -t net 
        NS TYPE NPROCS   PID USER    NETNSID NSFS COMMAND
4026531992 net     126  9158 cjx  unassigned      /lib/systemd/systemd --user
4026533051 net       2 19842 cjx  unassigned      /usr/share/code/code --type=zygo
4026533633 net      25  9679 cjx  unassigned      /opt/google/chrome/chrome --type
4026533725 net       1  9680 cjx  unassigned      /opt/google/chrome/nacl_helper
```
