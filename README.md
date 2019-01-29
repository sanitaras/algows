# algows
algorithms RESTful service

# Deploy on Windows
1. Install Git
2. Install GO
3. Configure GO Work-space and GOPATH env
  3.1 Create folder e.g. `C:\Projects\Go`
  3.2 Set GOPATH Environment Variable pointing to `C:\Projects\Go` (Control Panel -> Environmental Variables)
  3.3 Verify GOPATH from powershell run: `[Environment]::GetEnvironmentVariable("GOPATH","User")`
  3.4 Verify GOPATH from cmd run: `echo %GOPATH%`
4. From powershell console go to `C:\Projects\Go` folder and run: `go get https://github.com/sanitaras/algows`
5. To start algows run: `bin\algows`

# Deploy on Linux
1. Same procedure as Windows, just in Linux environment.

# Usage
1. After start of algows go to browser and open: http://localhost:9000
2. You should see the following output below.

> Welcome to algorithms RESTful service.
> Services available: 
> - Fibonacci Sequence e.g. http://localhost:9000/fib 
> Services not implemented yet:
> - algorithm1 e.g. http://localhost:9000/alg1 
> - algorithm2 e.g. http://localhost:9000/alg2 
> - algorithm3 e.g. http://localhost:9000/alg3 

