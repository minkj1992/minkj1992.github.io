# The cheapest way to build GCP


In this article, I'll describe the cheapest way to deploy server with GCP in a side project
<!--more-->
<br />

## tl;dr
- [x] Deploy API server with `Cloud run`
- [x] Deploy DB server with `Cloud SQL`
- [x] Build CI/CD pipeline with `github action` and `Cloub build`

## Deploy MySQL Database server
> with Cloud SQL in Google cloud platform

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_1.png"/>
    <figcaption>
        <h5>Cloud SQL meta</h5>
    </figcaption>
</figure>

At first fill in your database meta data.

1. Set instance id
2. Set databse password
3. MySQL version
4. Region (taiwan)
5. Set `single zone`

The `myply` mainly serves to korean users, but I selected `Taiwan(asia-east1)` because of pricing issue. [detail price description](https://minkj1992.github.io/gcp_pricing/)


<figure>
    <img width="600" height="450"src="/images/gcp/gcp_2.png"/>
    <figcaption>
        <h5>Cloud SQL Status</h5>
    </figcaption>
</figure>

- Select `db-f1-micro` 

I set a machine type to `Shared core` with `1vCPU, 0.614`. This spec is called `db-f1-micro`. It costs $7.665 per month, which is the cheapest spec in `cloud sql`. If I calculated it in `₩`, the currency of the south korea.

> $7.665 * 1300 = ₩9,964 per month

{{< admonition warning "db-f1-micro" >}}
Based on [Official cloud sql docs](https://cloud.google.com/sql/docs/mysql/instance-settings)

_The db-f1-micro and db-g1-small machine types aren't included in the Cloud SQL SLA. These machine types are configured to use a shared-core CPU, and are designed to provide low-cost test and development instances only. Don't use them for production instances._

*Note: The db-f1-micro and db-g1-small machine types are not included in the Cloud SQL SLA. These machine types are designed to provide low-cost test and development instances only. Do not use them for production instances.* 

{{< /admonition  >}}

- Select HDD storage
  - 10GB is the lowest storage size.

> $0.09 per GB/month * 10GB(min) = ₩1,170 per month



<figure>
    <img width="600" height="450"src="/images/gcp/gcp_3.png"/>
    <figcaption>
        <h5>Cloud SQL misc.1</h5>
    </figcaption>
</figure>

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_4.png"/>
    <figcaption>
        <h5>Cloud SQL misc.2</h5>
    </figcaption>
</figure>

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_5.png"/>
    <figcaption>
        <h5>Cloud SQL misc.3</h5>
    </figcaption>
</figure>

**So total `GCP cloud SQL` database server will cost `₩ 11,134`($8.56) per month.**


## Buil CI/CD pipeline
> with github action and `Cloud Build`

Before deploy `cloud run`, you should set `cloud build` to apply `continuous integration and continuous deployment(CICD)`.

<br />

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_6.png"/>
    <figcaption>
        <h5>Link with github repository</h5>
    </figcaption>
</figure>

- Enable Cloud Build API (almost free)
- Enable Container analysis API (free)


<figure>
    <img width="600" height="450"src="/images/gcp/gcp_7.png"/>
    <figcaption>
        <h5>Install GCP Build plugin to your repository</h5>
    </figcaption>
</figure>

- Set your repository to `cloud build`

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_8.png"/>
    <figcaption>
        <h5></h5>
    </figcaption>
</figure>

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_9.png"/>
    <figcaption>
        <h5>Select branch to be triggered and Dockerfile locaiton</h5>
    </figcaption>
</figure>

- [x] Select branch to be triggered
- [x] Enter Dockerfile loaction
  - If you have multi phase (e.g. `local`, `sandbox`, `beta` ,`prod`) it would be useful to set dockerfile name as `dockerfile.prod`.

```dockerfile
FROM golang:1.18-alpine AS builder

LABEL maintainer="leoo.j <minkj1992@gmail.com> (https://minkj1992.github.io)"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver ./application/cmd/main.go

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
```

## Deploy API server
> with Cloud run in Google cloud platform

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_10.png"/>
    <figcaption>
        <h5></h5>
    </figcaption>
</figure>

- Select `Continuously deploy new revisions from a source repository.`
- Select region as `Taiwan(asia-east1)`
- Select `Cpu is only allocated during request processing`
- Set number of instance (Autoscaling) (0 ~ 4)

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_11.png"/>
    <figcaption>
        <h5>Set Container status</h5>
    </figcaption>
</figure>

- [x] Set application server's port number
- [x] Set memory `128Mib`(lowest)
- [x] Set `Number of vCPUs` less than 1.
- [x] Set Execution environment to `First generation` (slower than 2nd generation)


<figure>
    <img width="600" height="450"src="/images/gcp/gcp_12.png"/>
    <figcaption>
        <h5>Connect Cloud SQL instance</h5>
    </figcaption>
</figure>

- Click `Connections` tab > `Cloud SQL connections` > `+Add Connection` button

<figure>
    <img width="600" height="450"src="/images/gcp/gcp_13.png"/>
    <figcaption>
        <h5>Connect Cloud SQL instance 2</h5>
    </figcaption>
</figure>


## Conclusion
Through the steps so far, I have covered `the cheapest way to deploy a gcp service for a side project` topic.
Now if you have followed all the steps so far, go to your github repository and merge it into the main branch.
Then the service will deployed according to the `Dockerfile` you set.
