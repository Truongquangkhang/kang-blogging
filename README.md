## Backend Mono Repo

BE is a **Go DDD** project that are **easy to develop, maintain, and fun to work with, especially in the long term!**

No application is perfect from the beginning. We will fix issues and achieve clean implementation by refactoring.

### Concept

- Apply DDD: [**Domain Driven Design**](https://en.wikipedia.org/wiki/Domain-driven_design) not ~~Destroy D**k December~~

>Domain-driven design (DDD) is a software design approach focusing on modelling software to match a domain according to input from that domain's experts.

![DDD](https://developer.ibm.com/developer/default/tutorials/cl-domain-driven-design-event-sourcing/images/ddd-figure-1.png "DDD")

- Apply [**Clean Architecture**]():

>The rule states that outer layers (implementation details) can refer to inner layers (abstractions), but not the other way around. The inner layers should instead depend on interfaces.
>
>The Domain knows nothing about other layers whatsoever. It contains pure business logic.
>
>The Application can import domain but knows nothing about outer layers. It has no idea whether it’s being called by an HTTP request, a Pub/Sub handler, or a CLI command.
>Ports can import inner layers. Ports are the entry points to the application, so they often execute application services or commands. However, they can’t directly access Adapters.
>
>Adapters can import inner layers. Usually, they will operate on types found in Application and Domain, for example, retrieving them from the database.

![Clean](https://d33wubrfki0l68.cloudfront.net/010a7625cfa822d0644ade7688b7cf94668b8946/9b1bf/media/introducing-clean-architecture/clean-arch-2.jpg "Clean")

- Apply CQRS: [**Command Query Responsibility Separation**](https://docs.microsoft.com/en-us/azure/architecture/patterns/cqrs)

>CQRS stands for Command and Query Responsibility Segregation, a pattern that separates read and update operations for a data store. Implementing CQRS in your application can maximize its performance, scalability, and security. The flexibility created by migrating to CQRS allows a system to better evolve over time and prevents update commands from causing merge conflicts at the domain level.

![CQRS](https://d33wubrfki0l68.cloudfront.net/53fe3d600a1d0c7396e67a3d5c748c3182615af3/807e2/media/introducing-cqrs/cqrs-architecture.jpg "CQRS")

### Directories

- [api](api/) OpenAPI and gRPC definitions
- [docker](docker/) Dockerfiles
- [internal](internal/) Application code
- [scripts](scripts/) Deployment and development scripts
- [sql](sql/) - Database definition and migration
- [tools](tools/) - Tools

### Running locally

Require Go [1.18+](https://go.dev/doc/install) and [MySQL 8 (InnoDB)](https://dev.mysql.com/downloads/) to run

- Install Go and make sure Go BIN is in PATH
- Install MySQL, can use Docker image: https://hub.docker.com/_/mysql
- Export your envs. Please check .env.sample for refs
- Run `make prepare` and `make run service=[service_name]` to build and run a service at local.
  You can only run `make prepare` 1st time
- Run `make migrate-up service=[service_name]` to migrate DB if need

```
$ make prepare
$ make migrate-up service=voucher_hub
$ make run service=voucher_hub
```