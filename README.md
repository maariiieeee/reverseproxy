**How can someone build and run your code?**
1) Edit the config.env file, with your desired configuration for the reverse proxy (proxy adress, backend adress, max concurrent requests and max retries)
2) Run `go get` command in the root directory to get the required dependencies
3) Run `go run reverse_proxy.go` to get the proxy software started
4) Connect to the specified proxy address in `config.env` via a browser or Postman to test and use the reverse proxy


**What resources did you use to build your implementation?**
I started by gaining a high-level understanding of the concept of a reverse proxy. I watched YouTube videos and read articles to learn about the role and functions of a reverse proxy in web architecture.
To build the reverse proxy from scratch with low-level connections, I delved into socket programming. I researched and studied the fundamentals of socket-based TCP communication to establish connections between the reverse proxy and the target server, as well as between the reverse proxy and clients.
Throughout the development process, I encountered various challenges and issues. I utilized online resources like Stack Overflow for troubleshooting and finding solutions to specific problems.


**Explain any design decisions you made, including limitations of the system.**

Design Decisions:
Concurrency Control: I use a semaphore-based approach to limit the number of concurrent requests my reverse proxy can handle. This helps prevent the system from becoming overwhelmed by too many requests at once.
Asynchronous Processing: Each client request is handled in a separate goroutine, allowing for concurrent processing of multiple clients. This asynchronous approach enhances the efficiency of the reverse proxy.
Error Handling: The code includes error handling for connection issues and data copying errors. This is crucial for robustness, ensuring that errors are logged for diagnosis and effective issue resolution.
Dynamic Configuration: Configuration settings are loaded from an external file using godotenv, enabling the modification of parameters like the maximum concurrent requests, server addresses, and the maximum number of retries without changing the source code.
Exponential Backoff: The introduction of exponential backoff when retrying connections to the target server in case of errors enhances the system's resilience and reduces the impact of transient network issues.

Limitations:
Security: The code doesn't address security aspects like encryption (e.g., TLS/SSL), authentication, and authorization. For a production-ready system, securing communication and controlling access to the proxy should be considered.
Logging and Monitoring: While errors and the status of transfer are logged, comprehensive logging and monitoring features could be added to help identify and address issues more effectively in a real-world scenario.


**How would you scale your implementation?**

Implement Load Balancing: I'll introduce load balancing to evenly distribute incoming client requests across multiple backend servers. This ensures that no single server is overloaded, which will improve the performance and redundancy of my application.

Utilize Containers and Enable Auto Scaling: I'll containerize my reverse proxy and the backend application, likely using Docker. This approach offers consistency and portability. To efficiently manage resources, I'll also implement auto-scaling solutions. These solutions will automatically adjust the number of container instances based on the current traffic load. When there's a spike in traffic, new instances will be launched, and during quieter periods, excess instances will be terminated. This dynamic scaling will help me handle varying levels of demand effectively.


**How would you make it more secure?**

Encrypted Connections: Implement secure connections using encryption protocols like TLS/SSL. This ensures that data transmitted between clients, the reverse proxy, and backend servers remains confidential and tamper-proof.
Access Control: Implement access control measures to ensure that only authorized clients can access the reverse proxy. This can involve setting up authentication and authorization mechanisms to verify users' identities and permissions.
Rate Limiting: Implement rate limiting to protect against abuse and Distributed Denial of Service (DDoS) attacks. By setting limits on the number of requests a client can make within a given time frame, I can prevent excessive traffic.




