# Memory 子系统练习
在 cgroup cpu 子系统目录中创建目录结构
```
cd /sys/fs/cgroup/memory
mkdir memorydemo
cd memorydemo

ls -la
总用量 0
drwxr-xr-x 2 root root 0 10月  7 21:01 .
dr-xr-xr-x 7 root root 0 10月  7 07:21 ..
-rw-r--r-- 1 root root 0 10月  7 21:02 cgroup.clone_children
--w--w--w- 1 root root 0 10月  7 21:02 cgroup.event_control
-rw-r--r-- 1 root root 0 10月  7 21:02 cgroup.procs
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.failcnt
--w------- 1 root root 0 10月  7 21:02 memory.force_empty
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.failcnt
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.limit_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.max_usage_in_bytes
-r--r--r-- 1 root root 0 10月  7 21:02 memory.kmem.slabinfo
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.tcp.failcnt
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.tcp.limit_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.kmem.tcp.max_usage_in_bytes
-r--r--r-- 1 root root 0 10月  7 21:02 memory.kmem.tcp.usage_in_bytes
-r--r--r-- 1 root root 0 10月  7 21:02 memory.kmem.usage_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.limit_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.max_usage_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.move_charge_at_immigrate
-r--r--r-- 1 root root 0 10月  7 21:02 memory.numa_stat
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.oom_control
---------- 1 root root 0 10月  7 21:02 memory.pressure_level
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.soft_limit_in_bytes
-r--r--r-- 1 root root 0 10月  7 21:02 memory.stat
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.swappiness
-r--r--r-- 1 root root 0 10月  7 21:02 memory.usage_in_bytes
-rw-r--r-- 1 root root 0 10月  7 21:02 memory.use_hierarchy
-rw-r--r-- 1 root root 0 10月  7 21:02 notify_on_release
-rw-r--r-- 1 root root 0 10月  7 21:02 tasks
```
## 运行 malloc（在linux机器make build）
查看内存使用情况
```
watch 'ps -aux|grep malloc|grep -v grep‘
```
### 通过 cgroup 限制 memory
把进程添加到cgroup进程配置组
```
echo ps -ef|grep malloc |grep -v grep|awk '{print $2}' > cgroup.procs
```
设置 memory.limit_in_bytes
```
echo 104960000 > memory.limit_in_bytes
```
等待进程被 oom kill
