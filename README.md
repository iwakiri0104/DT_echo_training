# DT_echo_training
DeepTrack's Echo training sample.

## Install
Preparing Sources
```
cd $WORKDIR
git clone git@github.com:SeiyaNakamura/DT_echo_training.git
```

## Launching Applications
Running with Docker
```
cd $WORKDIR/DT_echo_training
docker-compose up -d
```

## Enter Docker containers
Please check <container_name> by `docker-compose ps` etc.
```
cd $WORKDIR/DT_echo_training
docker container exec -it <container_name> sh
```

## Technologies used
- Language: Go v1.17.7
- Web Framework: Echo v4.1.6
- RDB: MySQL
- Hot reload tool: Air @latest
- Container technology: Docker
- Architecture: Onion Architecture（without DIP）
