<p align="center">
  <img src="/docs/go_robin_logo.svg" width="128" alt="Robin Logo"/>
</p>

# Robin
Decentralized, load balancer in go, utilising a weighted and dynamic Round-Robin algorithm. Robin lives inside a microservice environment as a separate process, balancing outgoing requests to other services. That helps eliminate a single point of failure and achieve resiliency within a crowded system. Robin also follows a self-healing approach creating multiple processes when needed and communicating with other service's load balancers. If Robin fails, other services take over to achieve load balancing, while Robin self-heals, succeeding in decentralization. Robin must have a minimal footprint doing load balancing and load balancing only, in order to keep the service as minimal and as small as possible. It's name is a reference to Robin Hood, stealing from the rich and giving to the poor, kind like it's main purpose as a load balancer.

<p align="center">
  <img src="/docs/usage_diagram.png" alt="Robin diagram"/>
</p>

## Go
Robin is written in Go, because of it's performance and usage in many distributed environments.

### Communicating with other services
Upon deploying Robin creates instances of itself in the cloud, inside the various services it communicates to, as a separate process. It establishes a connection with those instances sending a very small payload, to inform those self-instances that it's still alive, at a variable amount of time. When Robin goes down, another service's Robin instance takes control, providing load balancing, as the robin-daemon tries to bring Robin back to life. Thus, Robin running at the os level provides resiliency to the system.

### Building
To build Robin, from the project directory, simply run:
```
make install
```
Then find and run the compiled binaries at the bin directory of your Go path.
