package main

import (

  "encoding/json"
  "net"
  "net/http"
  "fmt"
  "log"
  "time"
  "strings"
  "strconv"
  
)


func Log(handler http.Handler) http.Handler {

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

      ip, _, _ := net.SplitHostPort(r.RemoteAddr)

      t := time.Now()

      logtime := t.Format("2006-01-02 15:04:05\n")

      log.Printf("%s %s %s %s %s %s", logtime, r.Header.Get("User-Agent"), ip, r.Method, r.Proto, r.URL)

      handler.ServeHTTP(w, r)
  })

}



func getFibs(n int) []int {

    var s []int

    v1:=0
    v2:=1
    next:=0

    for i:=1;i<=n;i++ {
        
        if(i==1){

            s = append(s, v1)
            continue
        }

        if(i==2){

            s = append(s, v2)
            continue
        }
        
        next = v1 + v2
        v1=v2
        v2=next

        s = append(s, next)
    }

    return s

  }



func getHandler(w http.ResponseWriter, r *http.Request) {


		if r.Method == "GET" {

		    fibnum := strings.TrimPrefix(r.URL.Path, "/fib/")

		    switch r.URL.Path {

          case "/fib/":

                fmt.Fprintln(w, "Welcome to Fibonacci Sequence RESTful service.\n - Usage e.g. http://localhost:9000/fib/5 \n")

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")


    			case "/fib/"+fibnum:
								
      					fmt.Printf("got - %s. \n", fibnum)

                if i, err := strconv.Atoi(fibnum) 

                    err == nil {
          
                    fmt.Printf("%s is an integer.\n", i)

                     if i < 0 {
  
                          fmt.Printf("%s is a negative integer.", i)

                          fmt.Fprintln(w, "negative integer: ", i)

                      } else if i > 0 {
  
                          fmt.Printf(" %s is a positive integer.", i)

                          c := getFibs(i)

                          js, err := json.Marshal(c)

                          if err != nil {

                              http.Error(w, err.Error(), http.StatusInternalServerError)
    
                              return
                          }

                              w.Header().Set("Content-Type", "application/json; charset=utf-8")
                              w.Write(js)
 

                      } else {

                              w.Header().Set("Content-Type", "text/plain; charset=utf-8")

                              fmt.Printf("%s is Zero.", i)
                              fmt.Fprintln(w, "value is Zero: ", i)
                      
                      }

                        } else {


                              w.Header().Set("Content-Type", "text/plain; charset=utf-8")

                              fmt.Printf("%s is not an integer.\n", i)

                              fmt.Fprintln(w, "not an integer")

                        }


          case "/alg1/":

                fmt.Fprintln(w, "Welcome to algorithm1 RESTful service.\n Not implemented yet! \n")

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")

                
          case "/alg2/":

                fmt.Fprintln(w, "Welcome to algorithm2 RESTful service.\n Not implemented yet! \n")

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")

                
          case "/alg3/":

                fmt.Fprintln(w, "Welcome to algorithm3 RESTful service.\n Not implemented yet! \n")

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")


          case "/":

                w.Header().Set("Content-Type", "text/plain; charset=utf-8")

                fmt.Fprintln(w, "Welcome to algorithms RESTful service.\n Services available: \n - Fibonacci Sequence e.g. http://localhost:9000/fib \n Services not implemented yet:\n - algorithm1 e.g. http://localhost:9000/alg1 \n - algorithm2 e.g. http://localhost:9000/alg2 \n - algorithm3 e.g. http://localhost:9000/alg3 \n")
		}


		} else {
        
            http.Error(w, "Invalid request method. Only HTTP GET supported.", 405) 
    }

}

