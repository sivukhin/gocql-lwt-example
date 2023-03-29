## Steps

- Run local Scylla node:
```
docker run -p 9042:9042 --name scylla-node -d scylladb/scylla:latest
```
- Create simple table:
```
CREATE KEYSPACE ks WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'} AND durable_writes = true;

CREATE TABLE ks.tbl (id text PRIMARY KEY);
```
- Build docker image
```
docker build . -t gocql-lwt-example
```
- Run docker image
```
docker --network host run gocql-lwt-example
```
- Binary in image will print
```
...
2023/03/29 07:28:28 err: <nil>
2023/03/29 07:28:28 isLwt: true
...
```
