# Fundamentals of Database Engineering


Learn ACID, Indexing, Partitioning, Sharding, Concurrency control, Replication, DB Engines, Best Practices and More!

<!--more-->

# 2. ACID

## 2.1 What is Transaction?

- A Collection of a queries
- One unit of work

### Transaction Lifespan
- BEGIN
- COMMIT
- (opt) Rollback

## 2.2. Atomicity

- 트랜잭션안의 모든 쿼리들을 succeed 되어야한다.
- 트랜잭션안의 한 쿼리가 실패하면, 이전의 모든 성공적인 쿼리들은 rollback 되어야 한다.
- 트랜잭션이 commit 되기 전에, db가 crash (went down)되면, 모든 트랜잭션 안의 successful queries는 rollback 되어야 한다.
- 꺼진 db가 다시 켜지게 되면, should clean this up 해야 한다. 

## 2.3. Isolation
> 실행되고 있는 트랜잭션이, 다른 실행되고 있는 트랜잭션이 변경 시킨 change를 볼 수있을까?

1. Read Phenomena (읽기 이상 현상)
2. Isolation Levels (격리 수준)
3. Impl of isolation

### 2.3.1. Read Phenomena
> dnpl

- Dirty reads
- Non-repeatable reads
- Phantom reads
- Lost updates


#### 2.3.1.1. Dirty Reads
> middle update
![](/images/db1.png)

동시에 진행되고 있는 다른 트랜잭션(아직 커밋하지 않은 상태)에서 변경한 데이터를 현재 진행 중인 트랜잭션에서 읽어 들이는 것을 뜻합니다.

#### 2.3.1.2. Non-repeatable reads
> middle update commit

![](/images/db2.png)

하나의 트랜잭션 중 읽어 들였던 특정 row의 값을 같은 트랜잭션 내에서 다시 읽어 들이는데 중간에 변경사항이 생겨 (실제로 COMMIT이 된 변경사항) 결괏값이 다르게 나오는 현상

#### 2.3.1.3. Phantom reads
> insert

![](/images/db3.png)

Phantom read란, 트랜잭션 시작 시점 데이터를 읽었을 때 존재하지 않았던 데이터가 다시 같은 조건으로 데이터를 읽어 들였을 때 존재해 (유령처럼) INCONSISTENT 한 결괏값을 반환하는 현상

#### 2.3.1.4. Lost updates
> duplicate update

![](/images/db4.png)

Lost Update란, 한 트랜잭션에서 데이터를 변경한 뒤 아직 커밋을 하지 않은 상태에서 읽어 들일 때, 다른 트랜잭션으로 인해 내가 작성한 변경사항이 덮어씌워지는 현상을 뜻합니다.

### 2.3.2. Isolation Level

1. Read uncommitted, No isolation
2. Read committed, tx only sees commited changed by other tx
3. Repeatable Read, make sure query read unchanged while tx running
    1. 트랜잭션안에 있는 여러 쿼리들이 같은 row를 참조할 때, 항상 같은 값을 주도록 한다.
    2. lock 또는 snapshot으로 구현시킬 수 있다.
    3. PostgreSQL은 (4) Snapshot으로 구현하여 phantom read 방지
    4. InnoDB는 MVCC (multi version concurrency control), 즉 undo 백업을 관리함.
4. Snapshot, snapshot으로 version 관리
    1. phantom read 없음
5. Serializable, no concurrency


![](/images/db5.png)

### 2.3.3. Implementation of Isolation

- Pessimistic, lock (row, table, page level)
- Optimistic, no lock (track and 충돌 발생시 뒤의 요청은 에러)
    - version, hashcode, timestamp등의 col 추가해서 충돌을 예방
- Serializable은 주로 Optimistic concurrency control로 구현된다. Pessimistic하게 구현하기 위해서는 SELECT FOR UPDATE

## 2.4. Consistency

#### 2.4.1. Consistency in Data
> 개발자 테이블 설계 관련 문제

![](/images/db6.png)

#### 2.4.2. Consistency in Read
> tx 커밋된 변화를, 새로운 tx가 곧바로 참조할 수 있을까?

Sharding된 db1,db2에서 db1에 update x to a를 했고, db2에서 select x를 할 경우에 
Inconsistency가 발생한다.

scale horizontally한 또는 caching을 도입한, 
RDB 그리고 NoSQL 모두 이에 대해서 suffer하고 있다.

- 이를 해결하기 위해서 Eventual consistency 전략이 도입되었다.
- or 무시 가능한 경우(예를들어 좋아요 수), inconsistent를 허용하는 방법.


## 2.5. Durability

Commit된 tx의 변화는 durable, non-volatile storage에 persisted 되어야한다.

Durability techniques

1. WAL - Write ahead log segments
2. Async snapshot, memory -> background snapshot
3. AOF


### WAL

연산에 사용되는 모든 데이터를 Disk에 저장하는 건 비싸다. 
WAL (Compressed version of the change)은 변화 segments를 compress해서 version 관리를 한다.

### OS Cache

![](/images/db7.png)

Os에서 write request는 대부분 os cache로 들어간다. 예를들어 file 저장을 요구했고, os가 succeed 되었다고 하지만, 알고보니 os cache memory에 기록되었고, os crush 상황에서 이 데이터는 유실될 수 있다는 뜻입니다. (os는 cache에 모아두었다가 batch로 처리해서 i/o 사용을 줄이는 전략을 취하기 때문에)

이를 해결하는 방법은 `fsync` command를 사용해서, write되는 데이터가 항상 go to disk 하도록 할 수 있습니다. 하지만 trade-off는 expensive하고 commit을 느려지게 할 수 있다는 점입니다.

## 2.6. Practical Exercise ACID

```sh
docker run --name pgacid -d -e POSTGRES_PASSWORD=postgres postgres:13
docker exec -it pgacid psql -U postgres
```

```postgres
create table products (
    pid serial primary key,
    name text,
    price float,
    inventory integer
);

create table sales (
    saleid serial primary key,
    pid integer,
    price float,
    quantity integer
);

insert into products(name, price, inventory) values('phone', 999.99, 100);
```

### 2.6.1 Atomicity

```
postgres=# begin transaction;
BEGIN
postgres=*# select * from products;
 pid | name  | price  | inventory
-----+-------+--------+-----------
   1 | phone | 999.99 |       100
(1 row)

postgres=*# update products set inventory = inventory - 10;
UPDATE 1
postgres=*# select * from products;
 pid | name  | price  | inventory
-----+-------+--------+-----------
   1 | phone | 999.99 |        90
(1 row)

postgres=*# exit;
> docker exec -it pgacid psql -U postgres
psql (13.12 (Debian 13.12-1.pgdg120+1))
Type "help" for help.

postgres=# select * from products;
 pid | name  | price  | inventory
-----+-------+--------+-----------
   1 | phone | 999.99 |       100
(1 row)
```
### 2.6.2 Isolation
> mvcc

```postgres
postgres=*# begin transaction isolation level repeatable read;
```

### 2.6.3. Durability

![](/images/db8.png)

TV를 insert하고 tx commit 시점에 docker stop 하더라도, durability
로 TV 정보가 남아있다.

![](/images/db9.png)

### 2.6.4 phantom read

```postgres
postgres=*# begin transaction isolation level serializable;
```

다른 tx에서 commit 하더라도 phantom read가 일어나지 않는다.

```postgres
postgres=*# begin transaction isolation level repeatable read;
```

하지만 PostgreSQL serializable보다 더 낮은 레벨인 repeatable read level에서도 tx안에서의 query에 대한 snapshot을 관리하기 때문에, phantom read를 방지 가능하다.

### 2.6.5. Serializable vs Non-Repeatable

- reapeatable read level인 경우

![](/images/db10.png)

```py
a
a
b
b
```

테이블에서

a->b (where not b), b->a (where not a)하는 2개의 tx(reapeatable read level)에서 동시에 commit을 할 경우, 

```py
b
b
a
a
```
가 발생한다.

만약 이를 원하지 않는 경우에는, Serializable level을 사용하면 된다.

- Serializable level인 경우

![](/images/db11.png)

```
1. start tx1, tx2


(tx1) select * from test;
1 a
2 a
3 b
4 b

(tx1) update  test set t = 'a' where t = 'b';

1 b
2 b
3 b
4 b

(tx1) commit


(tx2) update  test set t = 'a' where t = 'b';

ERROR: could not serialize access due to read/write dependencies among transactions


(tx2) rollback 하고 다시 commit하면 성공
```

backend 개발자는 에러 처리에 대해서 어떻게 동작할지를 서비스에 적용해 둬야한다.

# 3. Understanding Database Internals

## 3.1. How tables and Indexes are stored on disk
> 인덱스와 테이블이 disk에 저장되는 방식 그리고 query 어떻게 되는지

Storage Concepts

- Table
- Row_id
- Page
- IO
- Heap data structure
- Index data structure b-tree
- Example of a query

### Row id

- 내부 시스템에서 관리되는 row 아이디
- innodb는 pk와 같지만, postgres는 row_id(tuple_id)가 따로 관리된다.

### Page
> "고정된 크기의 block으로, disk에서 데이터와 인덱스를 저장하고 관리하는 기본 단위.

- innoDB(16KB)
- postgres(8KB)

- db는 single row를 읽는 것이 아닌, 한번의 IO에서 page or more을 읽는다. 
- 1001 rows가 존재하고, 각 page가 3rows를 가진다면, 333 ~ pages를 가지게 된다.
- 한편 Postgresql에서는 Heap안에 여러 page들로 데이터를 disk에 저장한다.


### IO
> input/output request to the disk

- IO는 expensive operation이기 때문에, 개발자들은 io를 최대한 최소화 하려고 노력
- 한번의 IO는 1page 이상을 fetch한다.
- 몇몇의 IO는 os에 따라서, disk가 아닌 os system cache에 접근한다. (i.e. postgresql)

### Heap

페이지들을 담고 있는 heap data structure.
Heap 안에서 효율적으로 데이터를 search하기 위해 index 필요.


### Index

- Heap과 분리된 data structure로 heap에 대한 pointer들을 가지고 있다.
- Index는 Heap의 어떤 page에서 정보를 가져와야할지 알려준다.
- 인덱스도 page로 관리되고, index entries들을 가져오기 위해 IO가 소요된다.
- Index가 작으면 작을 수로 memory에 fit되고 search가 빨라진다.

여기서 마지막 문장이 모호하게 들려, 좀 더 자세히 풀어보겠습니다.
index의 크기가 작으면 작을 수록 memory에 fit하다는 뜻은, Index가 작을 수록 RAM의 Virtual Memory에서의 page 하나 또는 적은수의 페이지에 포함되게 되며, 이를 통해 페이징 작업이 줄어들게 됩니다.

또한 페이지 fault시에도 적은 데이터로 인해 디스크 I/O작업을 최소화 시켜 페이지 폴트 처리 시간을 줄일 수 있습니다.

**사실 index도 page로 저장되기 때문에 disk에 저장되지만, buffer pool(innodb), shared Buffer Cache(postgre)같이 RAM 메모리 내의 buffer pool에 데이터를 cache**할 수 있기 때문에 적용되는 내용이라 생각된다.

## 3.2. Row vs Column oriented Databases

![](/images/db12.png)

### Row Databases
- Optimal for read / writes
- OLTP
- Compression isn't efficient
- Aggregation isn't efficient
    - aggr에 필요하지 않은 필드들도 mem에 올라와서 더 많은 page들을 뒤져서 aggr함
- Efficient queries w/multi-columns
    - 칼럼형은 row데이터를 얻기 위해서는, column 갯수 만큼 io 해서 찾아야 하기 때문에.
### Columnar Databases
- Writes are slower
- OLAP (online analytical Processing)
    - Google Bigquery, parquet
- Compress greatly
    - 칼럼형의 값들은 같은 타입의 값들이기 때문에 row형은 여러 타입들의 데이터가 저장되어있음.
    - 또한 col 값마다 중복된 값이 더 많을 수 있으므로 압축효율이 높음
- Amazing for aggregation (i.e. SUM() ...)
- Inefficient queries w/multi-columns

## 3.3. Primary key vs Secondary Key
> - pk는 clustering과 관련 되어 있습니다. 
> - [Mysql 인덱스 - 클러스티드 인덱스와 논클러스티드 인덱스 개념편 ](https://sihyung92.oopy.io/database/mysql-index) 참조


기본적으로 table은 ordering하지 않고 있지만 oracle의 IOT(Index organized Table), Postgres의 Clustered Index 따위가 있고, InnoDB는 Primary Key 클러스터링 인덱싱이 default입니다.

<center>

![](/images/db_index.png)

![](/images/db_index2.png)

</center>

InnoDB 기준 Primary key는 자동으로 clustered index가 되기 때문에 PK = clustered index로 생각하겠습니다. B Tree 구조에서 table이 저장될 때, PK(Clustered index)로 정렬되어 저장되어 tree형식으로 page들이 저장됩니다. 

![](/images/clustered_index.png)

Secondary Index(non clustered index)는 Index Page를 따로 생성하여 관리되며, `CREATE INDEX`로 생성된 index에 대한 B-tree를 생성하여 관리합니다.

![](/images/non_clustered_index.png)

즉 non clustered index는 index table이 따로 관리되며, B-Tree로 관리되기 때문에, Read 효율을 얻는 대신, CUD를 trade off로 가집니다.

## 3.4. Databases Pages - Deep dive
- 데이터 베이스 페이지란?
- 어떻게 Disk에 write 되는가?
- 어떻게 Disk에 store 되는가?
- Postgres의 page layout은 어떻게 되는가?

