# CPU 子系统练习
在 cgroup cpu 子系统目录中创建目录结构
```
cd /sys/fs/cgroup/cpu
sudo mkdir cpudemo
cd cpudemo

➜  cpudemo ll
总用量 0
drwxr-xr-x 2 root root 0 10月  7 20:37 .
dr-xr-xr-x 7 root root 0 10月  7 07:21 ..
-rw-r--r-- 1 root root 0 10月  7 20:37 cgroup.clone_children
-rw-r--r-- 1 root root 0 10月  7 20:37 cgroup.procs
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.stat
-rw-r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_all
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_percpu
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_percpu_sys
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_percpu_user
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_sys
-r--r--r-- 1 root root 0 10月  7 20:37 cpuacct.usage_user
-rw-r--r-- 1 root root 0 10月  7 20:37 cpu.cfs_period_us
-rw-r--r-- 1 root root 0 10月  7 20:37 cpu.cfs_quota_us
-rw-r--r-- 1 root root 0 10月  7 20:37 cpu.shares
-r--r--r-- 1 root root 0 10月  7 20:37 cpu.stat
-rw-r--r-- 1 root root 0 10月  7 20:37 cpu.uclamp.max
-rw-r--r-- 1 root root 0 10月  7 20:37 cpu.uclamp.min
-rw-r--r-- 1 root root 0 10月  7 20:37 notify_on_release
-rw-r--r-- 1 root root 0 10月  7 20:37 tasks

➜  cpudemo cat cpu.shares                 
1024

限制使用多少 CPU 时间片，-1 表示不限制
➜  cpudemo cat cpu.cfs_quota_us          
-1

一个 CPU 时间片数
➜  cpudemo cat cpu.cfs_period_us                
100000

监控的进程 ID，新建出来的为空
➜  cpudemo cat cgroup.procs
```
## 运行 busyloop
```
➜  busyloop git:(main) ✗ go build -o busyloop main.go
➜  busyloop git:(main) ✗ ll 
总用量 1.2M
-rwxrwxr-x 1 cjx cjx 1.2M 10月  7 20:42 busyloop
-rw-rw-r-- 1 cjx cjx   73 10月  7 20:34 main.go
➜  busyloop git:(main) ✗ ./busyloop
```
执行 top 查看 CPU 使用情况，CPU 占用 200%
```
top - 20:43:20 up 13:21,  1 user,  load average: 1.62, 0.87, 0.91
任务: 531 total,   3 running, 527 sleeping,   0 stopped,   1 zombie
%Cpu(s): 11.1 us,  0.3 sy,  0.0 ni, 87.2 id,  0.0 wa,  0.0 hi,  1.4 si,  0.0 st
MiB Mem : 128749.5 total, 117783.0 free,   5873.8 used,   5092.6 buff/cache
MiB Swap:      0.0 total,      0.0 free,      0.0 used. 121378.7 avail Mem 

 进程号 USER      PR  NI    VIRT    RES    SHR    %CPU  %MEM     TIME+ COMMAND      
  52435 cjx       20   0  702364   1160    648 R 200.0   0.0   0:44.07 busyloop  
......
```
```
# 通过 cgroup 限制 cpu
cd /sys/fs/cgroup/cpu/cpudemo

# 把进程添加到 cgroup 进程配置组
# echo ps -ef|grep busyloop|grep -v grep|awk '{print $2}' > cgroup.procs
➜  cpudemo ps -ef|grep busyloop|grep -v grep|awk '{print $2}' 
52435
➜  cpudemo cat cgroup.procs
52435

# 设置 cpuquota
echo 10000 > cpu.cfs_quota_us
cat cpu.cfs_quota_us
```
执行 top 查看 CPU 使用情况，CPU 占用变为10%(cpu.cfs_quota_us/cpu.cfs_period_us -> 10000/100000)
```
top - 20:49:11 up 13:27,  1 user,  load average: 1.28, 1.58, 1.26
任务: 531 total,   3 running, 527 sleeping,   0 stopped,   1 zombie
%Cpu(s):  1.7 us,  0.5 sy,  0.0 ni, 97.7 id,  0.0 wa,  0.0 hi,  0.1 si,  0.0 st
MiB Mem : 128749.5 total, 117765.6 free,   5888.9 used,   5095.0 buff/cache
MiB Swap:      0.0 total,      0.0 free,      0.0 used. 121364.8 avail Mem 

 进程号 USER      PR  NI    VIRT    RES    SHR    %CPU  %MEM     TIME+ COMMAND      
   4507 cjx       20   0 1125.5g 481848 208608 S  11.0   0.4  46:28.85 chrome       
  52435 cjx       20   0  702364   1160    648 R  10.3   0.0  11:28.43 busyloop
```
# 删除
```
sudo apt install cgroup-tools
sudo cgdelete cpu:cpudemo
```