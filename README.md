# GO Redis

## Instal Redis On Windows

1. Install WSL
2. Open WSL and Run This Command

### Install Redis

```bash
sudo add-apt-repository ppa:redislabs/redis
```

```bash
sudo apt-get update
```

```bash
sudo apt-get install redis
```

### Start Redis

```bash
sudo service redis-server start
```

```bash
redis-cli
```

### Stop Redis

```bash
sudo service redis-server stop
```
