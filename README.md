# Introduction

This repository is useful to check that the behavior of service check monitors can be closely emulated with metrics monitors. This includes both check and cluster alerts.

# How to use?

### Start the DD agent

Export a DD api key as an env variable:

```bash
export DD_API_KEY="<YOUR_API_KEY>"
```

### Start the DD agent as a container:

```bash
docker-compose up -d
```

### Create the different monitors by importing their JSON definition in DD (see an example below):

![alt text](./import_definition.gif)

> Make sure to create these monitors at the same time so that evaluations also occur at the same time

### Start submitting data to DD by running:

```bash
go run submit.go
```

Wait a couple of minutes and compare the different monitor detail pages to confirm that, over time, the statuses of the different groups are behaving the same way (including the NO DATA status).

# How to tear down?

Run the following commands to tear down the agent:

```bash
docker-compose stop
docker-compose rm
```

