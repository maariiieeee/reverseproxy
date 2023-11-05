package main

import (
    "net"
    "log"
    "io"
    "sync"
    "github.com/joho/godotenv"
    "fmt"
    "os"
    "strconv"
    "time"
)

func main() {
    if err := godotenv.Load("config.env"); err != nil {
        fmt.Println(err)
    }

    maxConcurrentRequests, err := strconv.Atoi(os.Getenv("maxConcurrentRequests"))

    semaphore := make(chan struct{}, maxConcurrentRequests)

    listener, err := net.Listen("tcp", os.Getenv("PROXY_ADDRESS"))
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        if len(semaphore) == maxConcurrentRequests {
            log.Println("Request limit reached. Waiting for slots to become available...")
        }

        clientConn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        semaphore <- struct{}{}

        go func() {
            handleClient(clientConn)
            <-semaphore
        }()
    }
}



func handleClient(clientConn net.Conn) {
    maxRetries, err := strconv.Atoi(os.Getenv("MAX_RETRIES"))
    if err != nil {
        log.Println("Error reading maxRetries:", err)
        return
    }

    for retry := 1; retry <= maxRetries; retry++ {
        if retry > 1 {
            sleepDuration := time.Duration(2<<uint(retry-2)) * time.Second
            time.Sleep(sleepDuration)
            log.Printf("Retry %d after waiting for %s\n", retry, sleepDuration)
        }

        targetServer, err := net.Dial("tcp", os.Getenv("BACKEND_SERVER")) 
        if err != nil {
            log.Printf("Error connecting to the target server (retry %d): %v\n", retry, err)
            if retry >= maxRetries {
                log.Println("Max retries reached. Giving up.")
                
                response := "HTTP/1.1 503 Service Unavailable\r\n"
                response += "Content-Type: text/plain\r\n"
                response += "\r\n"
                response += "The service is currently unavailable. Please try again later."
                
                _, err := clientConn.Write([]byte(response))
                if err != nil {
                    log.Println("Error writing error response to the client:", err)
                }
            
                clientConn.Close()
                return
            }
        } else {
            done := make(chan struct{})
            var wg sync.WaitGroup

            copyData := func(src, dest net.Conn) {
                _, err := io.Copy(dest, src)
                if err != nil {
                    log.Println("Error copying data:", err)
                } else {
                    log.Println("Transfer successful")
                }
                wg.Done()
            }

            wg.Add(2)
            go copyData(clientConn, targetServer)
            go copyData(targetServer, clientConn)

            go func() {
                wg.Wait()
                close(done)
            }()

            <-done
            clientConn.Close()
            targetServer.Close()
            return
        }
    }

    
}
