# algows
algorithms RESTful service

# Deploy on Windows
1. Install Git
2. Install GO
3. Configure GO Work-space and GOPATH env
  - Create folder e.g. `C:\Projects\Go`
  - Set GOPATH Environment Variable pointing to `C:\Projects\Go` (Control Panel -> Environmental Variables)
  - Verify GOPATH from powershell run: `[Environment]::GetEnvironmentVariable("GOPATH","User")`
  - Verify GOPATH from cmd run: `echo %GOPATH%`
4. From powershell console go to `C:\Projects\Go` folder and run: `go get https://github.com/sanitaras/algows`
5. To start algows run: `bin\algows`

# Deploy on Linux
1. Same procedure as Windows, just in Linux environment.

# Usage
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

# Testing with CURL
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
</pre>
