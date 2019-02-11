# algows
algorithms RESTful service

### Deploy on Windows
1. Install Git
2. Install GO
3. Configure GO Work-space and GOPATH env
  - Create folder e.g. `C:\Projects\Go`
  - Set GOPATH Environment Variable pointing to `C:\Projects\Go` (Control Panel -> Environmental Variables)
  - Verify GOPATH from powershell run: `[Environment]::GetEnvironmentVariable("GOPATH","User")`
  - Verify GOPATH from cmd run: `echo %GOPATH%`
4. From powershell console go to `C:\Projects\Go` folder and run: `go get github.com/sanitaras/algows`
5. To start algows run: `bin\algows`

### Deploy on Linux
1. Same procedure as Windows, just in Linux environment.

### Usage
1. After start of algows go to browser and open: http://localhost:9000
2. You should see the following output below.

<pre>
Welcome to algorithms RESTful service.
Services available: 
    -Fibonacci Sequence e.g. http://localhost:9000/fib
Services not implemented yet:
    -algorithm1 e.g. http://localhost:9000/alg1
    -algorithm2 e.g. http://localhost:9000/alg2
    -algorithm3 e.g. http://localhost:9000/alg3
</pre>

### Testing with CURL
<pre>
PS C:\curl> .\curl.exe -si -X GET http://localhost:9000
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Tue, 29 Jan 2019 16:15:08 GMT
Content-Length: 288

Welcome to algorithms RESTful service.
 Services available:
 - Fibonacci Sequence e.g. http://localhost:9000/fib
 Services not implemented yet:
 - algorithm1 e.g. http://localhost:9000/alg1
 - algorithm2 e.g. http://localhost:9000/alg2
 - algorithm3 e.g. http://localhost:9000/alg3

PS C:\curl> .\curl.exe -si -X POST http://localhost:9000
HTTP/1.1 405 Method Not Allowed
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Tue, 29 Jan 2019 16:15:25 GMT
Content-Length: 49

Invalid request method. Only HTTP GET supported.

PS C:\curl> .\curl.exe -si -X GET http://localhost:9000/fib/
HTTP/1.1 200 OK
Date: Tue, 29 Jan 2019 16:15:54 GMT
Content-Length: 91
Content-Type: text/plain; charset=utf-8

Welcome to Fibonacci Sequence RESTful service.
 - Usage e.g. http://localhost:9000/fib/5

PS C:\curl> .\curl.exe -si -X GET http://localhost:9000/fib/5
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 29 Jan 2019 16:16:03 GMT
Content-Length: 11

[0,1,1,2,3]

PS C:\curl> .\curl.exe -si -X GET http://localhost:9000/fib/-5
HTTP/1.1 200 OK
Date: Tue, 29 Jan 2019 16:16:11 GMT
Content-Length: 22
Content-Type: text/plain; charset=utf-8

negative integer:  -5

PS C:\curl> .\curl.exe -si -X GET http://localhost:9000/fib/a
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Tue, 29 Jan 2019 16:16:14 GMT
Content-Length: 15

not an integer

PS C:\curl> .\curl.exe -si -X GET http://localhost:9000/fib/55.00
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Tue, 29 Jan 2019 16:20:37 GMT
Content-Length: 15

not an integer

</pre>

## Docker: deploy algows as a service

Algows is available at Docker HUB.

URL: https://hub.docker.com/r/sanitaras/algows/tags

### Environment

<pre>
ekar@docker1:~$ hostnamectl

   Static hostname: docker1
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 36d92979bc934051bbb70cdac0f5845b
           Boot ID: e8123f8ee8174a299d8847d1f637c3dc
    Virtualization: bhyve
  Operating System: Ubuntu 18.10
            Kernel: Linux 4.18.0-15-generic
      Architecture: x86-64

ekar@docker1:~$ docker version
Client:
 Version:           18.09.1
 API version:       1.39
 Go version:        go1.10.6
 Git commit:        4c52b90
 Built:             Wed Jan  9 19:35:36 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.1
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.6
  Git commit:       4c52b90
  Built:            Wed Jan  9 19:02:44 2019
  OS/Arch:          linux/amd64
  Experimental:     false
</pre>

### Creating docker-compose.yml file

<pre>
ekar@docker1:~$ cat docker-compose.yml
version: "3"
services:
  web:
    image: sanitaras/algows:v1.0.0
    deploy:
      replicas: 2
      resources:
        limits:
          cpus: "0.1"
          memory: 50M
      restart_policy:
        condition: on-failure
    ports:
      - "80:9000"
    networks:
      - webnet
networks:
  webnet:
</pre>

### Validating docker-compose.yml config

<pre>
ekar@docker1:~$ docker-compose config
WARNING: Some services (web) use the 'deploy' key, which will be ignored. Compose does not support 'deploy' configuration - use `docker stack deploy` to deploy to a swarm.
networks:
  webnet: {}
services:
  web:
    deploy:
      replicas: 2
      resources:
        limits:
          cpus: '0.1'
          memory: 50M
      restart_policy:
        condition: on-failure
    image: sanitaras/algows:v1.0.0
    networks:
      webnet: null
    ports:
    - 80:9000/tcp
version: '3.0'
</pre>

### Run scalable service with load-balance (DNS round-robin)

<pre>
ekar@docker1:~$ docker swarm init
Swarm initialized: current node (6fex3tbwsnaahghnsncq7vwrj) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-3ah9gyszt0xhw96k5m2nedyc6q6etom6ssigawao4s6on9ocbb-68xa47zdxfcnhsjvrczjy5z56 192.168.1.15:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.


ekar@docker1:~$ docker stack deploy -c docker-compose.yml algows-service
Creating network algows-service_net
Creating service algows-service_web
</pre>


### Verify deployment

<pre>
ekar@docker1:~$ docker service ls
ID                  NAME                 MODE                REPLICAS            IMAGE                     PORTS
amddb9twsuno        algows-service_web   replicated          2/2                 sanitaras/algows:v1.0.0   *:80->9000/tcp

ekar@docker1:~$ docker ps
CONTAINER ID        IMAGE                     COMMAND             CREATED             STATUS              PORTS               NAMES
0e74e9c0a268        sanitaras/algows:v1.0.0   "bin/algows"        13 seconds ago      Up 11 seconds       9000/tcp            algows-service_web.2.uyletehv5110c3n6cs57w28kj
cb5243cdafe0        sanitaras/algows:v1.0.0   "bin/algows"        13 seconds ago      Up 11 seconds       9000/tcp            algows-service_web.1.4702uo6h5w4onch5ay55p3608


ekar@docker1:~$ docker network ls
NETWORK ID          NAME                    DRIVER              SCOPE
bwzu55kcwnqc        algows-service_webnet   overlay             swarm
18fbddefcf29        bridge                  bridge              local
30394fb61b44        docker_gwbridge         bridge              local
e8cabada2a8f        host                    host                local
jlh0ff9sbp2a        ingress                 overlay             swarm
29f5d361048c        none                    null                local

ekar@docker1:~$ docker network inspect algows-service_webnet
[
    {
        "Name": "algows-service_webnet",
        "Id": "bwzu55kcwnqc6ejsjm8uqxxwh",
        "Created": "2019-02-11T21:56:36.547055564Z",
        "Scope": "swarm",
        "Driver": "overlay",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "10.0.0.0/24",
                    "Gateway": "10.0.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "0e74e9c0a268ac2e6591184766e5113ed9f30cedfbd603bcd47f5c4bdc04294b": {
                "Name": "algows-service_web.2.uyletehv5110c3n6cs57w28kj",
                "EndpointID": "61cd5bc8dad2c3776ca7bd0490196be2f981de2db84baff2ffe117fc7b871855",
                "MacAddress": "02:42:0a:00:00:04",
                "IPv4Address": "10.0.0.4/24",
                "IPv6Address": ""
            },
            "cb5243cdafe08e1fbdd93a9cd7e6b57bbf578b237630c839339f56a3f999306a": {
                "Name": "algows-service_web.1.4702uo6h5w4onch5ay55p3608",
                "EndpointID": "1b3c2501e6bd6bb0b57e567d805755fff3e7b2c0c84baa3aa794d3a79a0a43fa",
                "MacAddress": "02:42:0a:00:00:03",
                "IPv4Address": "10.0.0.3/24",
                "IPv6Address": ""
            },
            "lb-algows-service_webnet": {
                "Name": "algows-service_webnet-endpoint",
                "EndpointID": "1f747dc6c567d9d0714644e0f12628514186cbe87fc5a4bf21b9835083788d9c",
                "MacAddress": "02:42:0a:00:00:05",
                "IPv4Address": "10.0.0.5/24",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.driver.overlay.vxlanid_list": "4097"
        },
        "Labels": {
            "com.docker.stack.namespace": "algows-service"
        },
        "Peers": [
            {
                "Name": "99023cf9e0bf",
                "IP": "192.168.1.15"
            }
        ]
    }
]

</pre>


### Do a test for round-robin

<pre>
ekar@docker1:~$ for i in {1..5}; do curl -sS http://192.168.1.15 || break; done

Hostname: 0e74e9c0a268

Welcome to algorithms RESTful service.
 Services available:
 - Fibonacci Sequence e.g. /fib
 Services not implemented yet:
 - algorithm1 e.g. /alg1
 - algorithm2 e.g. /alg2
 - algorithm3 e.g. /alg3

Hostname: cb5243cdafe0

Welcome to algorithms RESTful service.
 Services available:
 - Fibonacci Sequence e.g. /fib
 Services not implemented yet:
 - algorithm1 e.g. /alg1
 - algorithm2 e.g. /alg2
 - algorithm3 e.g. /alg3

Hostname: 0e74e9c0a268

Welcome to algorithms RESTful service.
 Services available:
 - Fibonacci Sequence e.g. /fib
 Services not implemented yet:
 - algorithm1 e.g. /alg1
 - algorithm2 e.g. /alg2
 - algorithm3 e.g. /alg3
</pre>

Note: `hostname` is equal to `container ID`

### Do benchmark with 2 nodes

<pre>
ekar@docker1:~$ ab -n 10000 -c 20 http://192.168.1.15/fib/15

This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.1.15 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        192.168.1.15
Server Port:            80

Document Path:          /fib/15
Document Length:        42 bytes

Concurrency Level:      20
Time taken for tests:   15.429 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1650000 bytes
HTML transferred:       420000 bytes
Requests per second:    648.12 [#/sec] (mean)
Time per request:       30.858 [ms] (mean)
Time per request:       1.543 [ms] (mean, across all concurrent requests)
Transfer rate:          104.43 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.1      0       6
Processing:     0   30  37.4      7     281
Waiting:        0   27  36.1      6     209
Total:          0   31  37.4      9     282

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     14
  75%     79
  80%     82
  90%     87
  95%     91
  98%     96
  99%    101
 100%    282 (longest request)

</pre>


### Add more nodes

Modify `replicas: <number>` parameter.

<pre>
ekar@docker1:~$ cat docker-compose.yml
version: "3"
services:
  web:
    image: sanitaras/algows:v1.0.0
    deploy:
      replicas: 5
      resources:
        limits:
          cpus: "0.1"
          memory: 50M
      restart_policy:
        condition: on-failure
    ports:
      - "80:9000"
    networks:
      - webnet
networks:
  webnet:
</pre>

<pre>
ekar@docker1:~$ docker stack deploy -c docker-compose.yml algows-service
Updating service algows-service_web (id: amddb9twsunozfq21id248dpv)

ekar@docker1:~$ docker ps
CONTAINER ID        IMAGE                     COMMAND             CREATED             STATUS              PORTS               NAMES
7d838b5176a8        sanitaras/algows:v1.0.0   "bin/algows"        37 seconds ago      Up 34 seconds       9000/tcp            algows-service_web.4.wxp9ljkodra6powshuhiivwb7
d2774cb48e98        sanitaras/algows:v1.0.0   "bin/algows"        37 seconds ago      Up 34 seconds       9000/tcp            algows-service_web.5.6pqqinx79a5rh9z1ymng5r07n
471119fd62e8        sanitaras/algows:v1.0.0   "bin/algows"        37 seconds ago      Up 35 seconds       9000/tcp            algows-service_web.3.f9g93g2dmhcuwkhh7sphr4pml
0e74e9c0a268        sanitaras/algows:v1.0.0   "bin/algows"        9 minutes ago       Up 9 minutes        9000/tcp            algows-service_web.2.uyletehv5110c3n6cs57w28kj
cb5243cdafe0        sanitaras/algows:v1.0.0   "bin/algows"        9 minutes ago       Up 9 minutes        9000/tcp            algows-service_web.1.4702uo6h5w4onch5ay55p3608
</pre>

### Do benchmark again with 5 nodes

<pre>

ekar@docker1:~$ ab -n 10000 -c 20 http://192.168.1.15/fib/15

This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.1.15 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        192.168.1.15
Server Port:            80

Document Path:          /fib/15
Document Length:        42 bytes

Concurrency Level:      20
Time taken for tests:   8.652 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1650000 bytes
HTML transferred:       420000 bytes
Requests per second:    1155.78 [#/sec] (mean)
Time per request:       17.304 [ms] (mean)
Time per request:       0.865 [ms] (mean, across all concurrent requests)
Transfer rate:          186.23 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0      13
Processing:     0   17  21.4      8     200
Waiting:        0   15  19.5      6     200
Total:          0   17  21.4      8     201

Percentage of the requests served within a certain time (ms)
  50%      8
  66%     14
  75%     21
  80%     29
  90%     53
  95%     63
  98%     75
  99%     84
 100%    201 (longest request)
</pre>

### Stop service and swarm (clean up)

<pre>
ekar@docker1:~$ docker stack rm algows-service
Removing service algows-service_web
Removing network algows-service_webnet

ekar@docker1:~$ docker swarm leave --force
Node left the swarm.
</pre>

### Summary
Very easy to scale the service up or down and we see that while with 2 nodes we achieve 648.12 requests per second, with 5 nodes we get 1155.78.
