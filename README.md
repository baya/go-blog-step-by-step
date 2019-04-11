教程来源于 [https://github.com/EDDYCJY/blog](https://github.com/EDDYCJY/blog) 感谢

## 我自己遇到的一些问题的解决方法

1. ERROR 2059 (HY000): Authentication plugin 'caching_sha2_password' cannot be loaded: dlopen(/usr/local/Cellar/mysql/5.7.16/lib/plugin/caching_sha2_password.so, 2): image not found

通过下面的命令重新启动一个 name 为 mysql5 的 container

```bash
$ docker run --name mysql5 -p 3309:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql --default-authentication-plugin=mysql_native_password
```

或者

```bash
$ docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=rootpwd -d mysql --default-authentication-plugin=mysql_native_password
```

现在可以连上 mysql5 了, 此时里面的数据都是空的，需要重新创建 blog 数据库和相关的表

```bash
$ mysql -h localhost -P 3309 --protocol=tcp -uroot
```

## docker 的一些常用命令

- list stopped containers

```bash
$ docker container ls -a --filter status=exited
```

- remove all stopped containers

```bash
$ docker container prune
```

- list of all Docker containers on your system

```bash
$ docker container ls -aq
```

- stop all running containers

```bash
$ docker container stop $(docker container ls -aq)
```

- remove container by id

```bash
$ docker container rm 701b07250dca
```

- 删除 image

```bash
$ docker rmi -f go-blog-step-by-step
```

- 创建 image

```bash
$ docker build -t go-blog-step-by-step .
```

- 运行 mysql container

```bash
$ docker run --name mysql -p 3307:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -d mysql --default-authentication-plugin=mysql_native_password
```

- 运行 blog container

```bash
$ docker run --link mysql:mysql -p 8000:8848 go-blog-step-by-step
```

- 进入 container bash

```bash
$ docker exec -it your_container_name_or_id bash
```