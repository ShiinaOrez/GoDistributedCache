## HTTP-REST Cache Service

[![Build Status](https://travis-ci.com/ShiinaOrez/GoDistributedCache.svg?branch=master)](https://travis-ci.com/ShiinaOrez/GoDistributedCache)

------

This part described a single point cache service in **TCP** way. You can deploy this service on your ``localhost`` or any remote linux server. 

**Warning:** this cache service dosen't surpport the persistent data storage. So, if you lose data because you used the service, I will be very sorry that you **did not** read this warning carefully.

------

## Usage:

### Server

```bash
/tcp $: go run server.go
```

### Client

**Get the help information:**

```bash
/tcp/client $: go run client --help
```

**You'll get this:**

```
Usage of /tmp/go-build929116209/b001/exe/client:
  -c string
    	command, must be get/set/del (default "get")
  -h string
    	cache server address. (default "localhost")
  -key string
    	key
  -value string
    	value
```

**Example:**

```bash
/tcp/client $: go run client.go -c set -key cache -value service
```

### Test

```bash
/tcp $: go test
```