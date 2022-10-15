# OverlayFS 文件系统练习
```bash
cjx@cjx-0004:~$ mkdir overlayFS_test
cjx@cjx-0004:~$ cd overlayFS_test/
cjx@cjx-0004:~/overlayFS_test$ mkdir upper lower merged work
cjx@cjx-0004:~/overlayFS_test$ echo "from lower" > lower/in_lower.txt
cjx@cjx-0004:~/overlayFS_test$ echo "from upper" > upper/in_upper.txt
cjx@cjx-0004:~/overlayFS_test$ echo "from lower" > lower/in_both.txt
cjx@cjx-0004:~/overlayFS_test$ echo "from upper" > upper/in_both.txt
cjx@cjx-0004:~/overlayFS_test$ cat lower/in_lower.txt 
from lower
cjx@cjx-0004:~/overlayFS_test$ cat lower/in_both.txt 
from lower
cjx@cjx-0004:~/overlayFS_test$ cat upper/in_upper.txt 
from upper
cjx@cjx-0004:~/overlayFS_test$ cat upper/in_both.txt 
from upper
cjx@cjx-0004:~/overlayFS_test$ sudo mount -t overlay overlay -o lowerdir=`pwd`/lower,upperdir=`pwd`/upper,workdir=`pwd`/work `pwd`/merged
cjx@cjx-0004:~/overlayFS_test$ ll merged/
total 20
drwxrwxr-x 1 cjx cjx 4096 10月 15 16:51 ./
drwxrwxr-x 6 cjx cjx 4096 10月 15 16:51 ../
-rw-rw-r-- 1 cjx cjx   11 10月 15 16:51 in_both.txt
-rw-rw-r-- 1 cjx cjx   11 10月 15 16:51 in_lower.txt
-rw-rw-r-- 1 cjx cjx   11 10月 15 16:51 in_upper.txt
cjx@cjx-0004:~/overlayFS_test$ cat merged/in_both.txt 
from upper
cjx@cjx-0004:~/overlayFS_test$ cat merged/in_lower.txt 
from lower
cjx@cjx-0004:~/overlayFS_test$ cat merged/in_upper.txt 
from upper

# work 目录没有权限，并且内部为空
cjx@cjx-0004:~/overlayFS_test$ ll work/
total 12
drwxrwxr-x 3 cjx  cjx  4096 10月 15 16:51 ./
drwxrwxr-x 6 cjx  cjx  4096 10月 15 16:51 ../
d--------- 2 root root 4096 10月 15 16:51 work/
cjx@cjx-0004:~/overlayFS_test$ ll work/work/
ls: cannot open directory 'work/work/': Permission denied
cjx@cjx-0004:~/overlayFS_test$ sudo ls work/work/

```
```
cjx@cjx-0004:~/overlayFS_test$ mount | grep '/home/cjx/overlayFS_test'
overlay on /home/cjx/overlayFS_test/merged type overlay (rw,relatime,lowerdir=/home/cjx/overlayFS_test/lower,upperdir=/home/cjx/overlayFS_test/upper,workdir=/home/cjx/overlayFS_test/work)
```
## Delete merged file
```shell
cjx@cjx-0004:~/overlayFS_test$ rm -rf merged/
rm: cannot remove 'merged/': Device or resource busy

cjx@cjx-0004:~/overlayFS_test$ sudo umount /home/cjx/overlayFS_test/merged
cjx@cjx-0004:~/overlayFS_test$ rm -rf merged/
```

# docker 文件系统
```
cjx@cjx-0004:~/yamls$ docker inspect k8s_redis_redis-5c9986b94b-4ghx7_default_c91587b1-6c38-4d1f-bdb0-41bce72a5df4_0
[
    {
        "Id": "c7e29a977e2aedd702927eb751690f52ddc75c16fa2da5da4936cc9a0cb6a3ec",
        "Created": "2022-10-15T08:58:23.227973646Z",
        "Path": "docker-entrypoint.sh",
        "Args": [
            "redis-server"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 40197,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2022-10-15T08:58:23.36684249Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:f8528f17261c4a2c94ef702ff483ba7e4b998aa734cba60fa689ca5ecc14705f",
        "ResolvConfPath": "/var/lib/docker/containers/dbd93c5418749ece2e581deebae6a1460d1c35d1dcd16a9451b1fb1b04d32f5b/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/dbd93c5418749ece2e581deebae6a1460d1c35d1dcd16a9451b1fb1b04d32f5b/hostname",
        "HostsPath": "/var/lib/kubelet/pods/c91587b1-6c38-4d1f-bdb0-41bce72a5df4/etc-hosts",
        "LogPath": "/var/lib/docker/containers/c7e29a977e2aedd702927eb751690f52ddc75c16fa2da5da4936cc9a0cb6a3ec/c7e29a977e2aedd702927eb751690f52ddc75c16fa2da5da4936cc9a0cb6a3ec-json.log",
        "Name": "/k8s_redis_redis-5c9986b94b-4ghx7_default_c91587b1-6c38-4d1f-bdb0-41bce72a5df4_0",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
......
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/449fa6581a87c254462fbc964143eec337b9ea24e6e0fef7c56e14da818cddbc-init/diff:/var/lib/docker/overlay2/e99428626115b3d5627b11887db9d50a77447e09abb6ca7f88c1d4c3292cb678/diff:/var/lib/docker/overlay2/d00bb66979c0e13b5ecf75e59053eb2dac96cd95bd8cb7a62666dad9fc37ffa3/diff:/var/lib/docker/overlay2/d44528f62b5370fbfea16daacbe5113f8a904fee956c2930d4c1bf5a02bd3575/diff:/var/lib/docker/overlay2/c7e66a7126a5b274953254d2b45c8ec818110390b0c12d0526394be8d8697f3d/diff:/var/lib/docker/overlay2/95447c303743a9f6ecab7068f6e0912f859a83ca22145a7fb169a17204d3fae5/diff:/var/lib/docker/overlay2/b6fe7e056518a6c95ce2ddf19f74dd01b897acc160ab547a54343ad315a01874/diff",
                "MergedDir": "/var/lib/docker/overlay2/449fa6581a87c254462fbc964143eec337b9ea24e6e0fef7c56e14da818cddbc/merged",
                "UpperDir": "/var/lib/docker/overlay2/449fa6581a87c254462fbc964143eec337b9ea24e6e0fef7c56e14da818cddbc/diff",
                "WorkDir": "/var/lib/docker/overlay2/449fa6581a87c254462fbc964143eec337b9ea24e6e0fef7c56e14da818cddbc/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [
            {
                "Type": "bind",
                "Source": "/var/lib/kubelet/pods/c91587b1-6c38-4d1f-bdb0-41bce72a5df4/volumes/kubernetes.io~projected/kube-api-access-kwjh8",
                "Destination": "/var/run/secrets/kubernetes.io/serviceaccount",
                "Mode": "ro",
                "RW": false,
                "Propagation": "rprivate"
            },
            {
                "Type": "bind",
......
```
