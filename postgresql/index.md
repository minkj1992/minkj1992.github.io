# PostgreSQL


Design efficient database structures with `PostgreSQL`
<!--more-->

## How to build a 'Like' System

![](/images/postgresql/like_post.png)

Instagram의 'Like' System을 만들 때, user <-> likes <-> posts 관계를 생각해볼 수 있습니다.

이 경우 like는 아래의 Rule이 필요합니다.

![](/images/postgresql/like_rule.png)

1. 각 유저는 특정 post에 한번만 like를 해야합니다.
2. 자신의 게시물은 like 할 수 없습니다.
3. 유저는 unlike할 수 있어야 합니다.
4. 얼마나 많은 유저들이 post를 like했는지 셀 수 있어야 합니다.
5. 어떤 유저들이 Like했는지 볼 수 있어야 합니다.
6. 어쩌면 post이외에도 like할 수 있어야 합니다. (i.e 댓글)
7. 어쩌면 하트/빈하트가 아니라, 이모티콘 처럼 여러 종류의 Reaction을 관리할 수도 있습니다.

### 1st approach

가장 단순하게 like system을 생각해본다면, posts 테이블안에 likes 필드를 두어 관리하는 방법을 생각해볼 수 있습니다.

![](/images/postgresql/like1.png)

이 경우, 다음과 같은 문제점들이 생깁니다.

1. 어떤 유저가 like했는지 기록하지 않아, 한번만 like 하도록 강제할 수 없습니다.
2. 특정 유저가 unlike할 때, 자신이 한 like를 unlike하는건지 알 수 없습니다.
3. 누가 어떤 post를 like하는지 알 방법이 없습니다.
4. 유저가 탈퇴하면 어떤 포스트의 like 수를 낮춰야 하는지 알 수 없습니다.

### 2nd approach

![](/images/postgresql/like2.png)

다음 방법은, likes 테이블을 따로 두어, user_id와 post_id를 기록하는 방식입니다.
이렇게 할 경우, 아래 4가지 경우를 모두 알 수 있습니다.

1. 특정 post의 # of likes 
2. 특정 post에 누가 like 했는지
3. like 순으로 post들 정렬가능
4. 특정 user가 like한 post들

![](/images/postgresql/like3.png)

이때 한 유저는 한번만 post에 like를 해야하니, unique(user_id, post_id)로 관리합니다.

하지만 여기서 만약 post이외에도 댓글에 대해서 like를 해야하는 요구사항이 생기면 어떻게 관리되어야 할까요?

![](/images/postgresql/like4.png)

### 3rd, Polymorphic association

`post`와 `comment`의 like를 쉽게 관리할 수 있는 첫번째 방법은 다음과 같습니다.
likes 테이블에 post_id 또는 comment_id를 기록하는 like_id와 post/comment에 대한 like_type을 추가하는 방식입니다. 이 방법을 `Polymorphic association`이라고 칭합니다. 

![](/images/postgresql/like5.png)

이 경우에 가장 큰 문제점은 'like_id'가 fk로 관리될 수 없고 일반 int필드로 관리되기 때문에, constraint와 관리에서 큰 문제가 발생할 수 있습니다.

### 4th

`Polymorphic Association`을 피하는 가장 쉬운 방법은 2개의 fk (post, comment)를 관리하던 `like_id`를 나누면 됩니다.

![](/images/postgresql/like6.png)

post_id와 comment_id를 나눈뒤, fk로 관리하면 문제를 해결 할 수 있습니다. 이때 post_id와 comment_id가 동시에 값을 가지거나, 동시에 null인 경우를 처리하기 위해서

`COALESCE((post_id)::BOOLEAN::INTEGER,0) + COALESCE((comment_id)::BOOLEAN::INTEGER,0)`를 사용해 constraint를 걸어줍니다.

![](/images/postgresql/like7.png)

하지만 이 방법 또한 like를 post와 comment이외에도 여러곳에서 사용해야한다면 관리가 복잡해진다는 단점이 존재합니다.


### 5th

이전의 방식의 문제점은 하나의 테이블에서 서로 다른 형태의 like를 관리하려 했기 때문이기 때문에, 이를 해결하기 위해서 like테이블을 각각 나눠서 관리하면 됩니다.

![](/images/postgresql/like8.png)

이 경우 장점은 like가 여러곳에서 사용되더라도, 여러 테이블로 관리하면 되기 때문에 테이블 복잡도를 낮출 수 있습니다.

또한 트래픽 차원에서도 만약 post-like에 대한 접근과 comment에 대한 like 접근이 8:2라면, 이전 방식에서는 하나의 테이블에서 부하를 모두 감당했지만, 나눠진 테이블에서는 분리해서 관리하기 때문에 
관리가 더 편합니다.

또한 Reaction 타입의 like를 Post에 도입한다 하더라도, col한개만 추가하면 되기 때문에 간편합니다.

다만 정규화를 할 수록 join에 대한 요구가 더 많아질 수 있기 때문에, 이에 대해서는 trade-off가 존재합니다.

만약 post와 comment 외에 like가 필요없다면 기획적으로 사용할 예정이라면, 이전 방법도 좋은 방법입니다.

## How to build a 'Mention' System
> `@user_id``

다음은 `mention`에 대해서 db를 디자인 해봅니다.

![](/images/postgresql/mention1.png)

가장 먼저 post안의 `@renan_ozturk`나, `@stepenwikes`와 같이 멘션들을 넣을 수 있습니다. 


![](/images/postgresql/mention3.png)

또한 post를 생성 시, `Tag People`를 하여, 사용자를 태그할 수 있으며, Image의 특정 x,y 좌표에 mention을 할수도 있습니다. 이외에도 post에 대한 좌표를 저장해야 합니다.

우선 여기에선 photo에 대한 tag와 post안에 들어있는 caption_tag에 대한 schema를 디자인 해보자면
다음과 같은 2가지 방식이 존재할 수 있습니다.

![](/images/postgresql/mention2.png)

1. 첫번째 방식은 photo mention인 경우에는 x,y를 저장하고, 그렇지 않으면 x,y를 Null로 관리하는 방식이고
2. 두번째 방식은 2개의 테이블로 나눠서 관리하는 방식입니다.

## How to build a 'HashTag' System

![](/images/postgresql/hash1.png)

다음은 인스타그램 hashtag 시스템입니다. 여러 곳에서 hashtag가 사용되고 있으며 이에 대해서 기록하면 스키마는 다음과 같습니다.

![](/images/postgresql/hash2.png)

하지만 실제로 hashtag는 post에 대한 검색만을 제공해주고 있기 때문에 rdb에는 post에 대한 hashtag만 저장해두면 될 것 같습니다. 

![](/images/postgresql/hash3.png)

다른 곳에서 사용되는 hashtag들은 rdb가 아닌 분석을 위해, preprocess 단계를 거쳐 데이터웨어하우스에 저장하는 방식으로 저장될 수는 있어 보입니다.

결국 추가적으로 관리해야 하는 table은 hashtag <-> post에 대한 테이블입니다.

![](/images/postgresql/hash4.png)

그럼 이 hashtag를 어떻게 효율적으로 저장할 수 있을까요?

### hash tag table

![](/images/postgresql/hash5.png)

해시태그는 그 성격상 중복이 많을 수 있기 때문에, 위의 테이블과 같이 저장하게 되면 매우 비효율적으로 테이블이 관리되게 됩니다.

![](/images/postgresql/hash6.png)

또한 hashtags: posts = n:m 관계인 것을 고려하면, 중간에 매핑 테이블로 저장해서 관리하면 됩니다.
이떄 hashtags 테이블의 title을 unique로 저장하면 중복관리에도 좋습니다.

## How to build a 'Follow' System

![](/images/postgresql/follower1.png)

인스타그램에서는 총 follwer 숫자, following 숫자 그리고 누구를 follow 또는 following되는지를 저장하고 있습니다.

이는 아래와 같은 스키마로 저장합니다.

![](/images/postgresql/follower2.png)



## Final Schema

최종적으로 schema는 아래의 diagram을 따릅니다.


<center>

![](/images/postgresql/postgresql_diagrams.png)

</center>


아래는 table statement입니다. 이때 null, default는 아래 기준으로 정의했습니다.

1. value가 주어지든 아니든 상관없다 -> `비워둠`
2. 100% 유저 또는 engineer가 값을 주어야 한다 -> `NOT NULL`
3. 언제나 값이 테이블에 존재하길 원하지만, 생성될 때 optional로 처리되어도 된다 -> `NOT NULL + DEFAULT`


```sql
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	username VARCHAR(30) NOT NULL UNIQUE,
	bio VARCHAR(400),
	avatar VARCHAR(200),
	phone VARCHAR(25),
	email VARCHAR(40),
	password VARCHAR(50),
	status VARCHAR(15),
	CHECK(COALESCE(phone, email) IS NOT NULL)
);

CREATE TABLE posts (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	url VARCHAR(200) NOT NULL,
	caption VARCHAR(240),
	lat REAL CHECK(lat IS NULL OR (lat >= -90 AND lat <= 90)), 
	lng REAL CHECK(lng IS NULL OR (lng >= -180 AND lng <= 180)),
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE comments (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	contents VARCHAR(240) NOT NULL,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE
);

CREATE TABLE likes (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
	comment_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
	CHECK(
		COALESCE((post_id)::BOOLEAN::INTEGER, 0)
		+
		COALESCE((comment_id)::BOOLEAN::INTEGER, 0)
		= 1
	),
	UNIQUE(user_id, post_id, comment_id)
);

CREATE TABLE photo_tags (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
	x INTEGER NOT NULL,
	y INTEGER NOT NULL,
	UNIQUE(user_id, post_id)
);

CREATE TABLE caption_tags (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
	UNIQUE(user_id, post_id)
);

CREATE TABLE hashtags (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	title VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE hashtags_posts (
	id SERIAL PRIMARY KEY,
	hashtag_id INTEGER NOT NULL REFERENCES hashtags(id) ON DELETE CASCADE,
	post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
	UNIQUE(hashtag_id, post_id)
);

CREATE TABLE followers (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	leader_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	follower_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	UNIQUE(leader_id, follower_id)
);
```


## The internals of PostgreSQL

Postgresql에서 데이터들은 어떻게 저장될까요? 한번 알아보도록 하겠습니다.

![](/images/postgresql/internal1.png)

먼저 pgAdmin에서 data_directory를 요청하면, postgresql이 데이터를 저장하고 있는 dir을 알려줍니다.

![](/images/postgresql/internal2.png)

해당 디렉토리에 들어가면 여러 폴더들이 존재하는데, 그 중에서 ./base 폴더에 들어가보면

![](/images/postgresql/internal4.png)

oid로 구분된 디렉토리에 데이터들이 저장되어있는 것을 확인할 수 있습니다.

![](/images/postgresql/internal5.png)

이를 pgAdmin에서 dataname에 대한 oid를 찾아보면, instagram은 22442로 저장되어있는 것을 확인할 수 있습니다.

![](/images/postgresql/internal6.png)

./base/22442안에 들어가게 되면 여러 파일들로 disk에 데이터가 저장되어있는 것을 확인할 수 있습니다.

![](/images/postgresql/internal7.png)

이때 pg_class로 instagram과 관련된 table의 oid를 얻어내면, 아래와 같이 저장된 것을 확인가능합니다.

- users 테이블은 22445 
- posts 테이블은 22459
- comments 테이블은 22476 

그럼 postgreSQL는 어떤 구조로 이렇게 disk에 데이터를 저장하고 있을까요?


### Heaps, Blocks and Tuples

![](/images/postgresql/internal8.png)

PostgreSQL는 다음 3가지로 데이터를 disk에 저장합니다.

1. Heap (File)
2. Pages (Block)
3. Tuple (Item)

하나의 Heap 단위 즉 file단위로 테이블의 data(rows)를 저장해둡니다. 이때 저장되는 row(tuple)은 page단위로 group화 되어있으며, i/o의 단위가 됩니다. 즉 하나의 row를 찾더라도 최소한 하나의 page를 i/o 해야 한다는 뜻입니다.

![](/images/postgresql/internal9.png)

앞서 user 테이블을 담고있는 22445 파일을 예시로 보자면, 22445라는 heap file로 디렉토리에 저장되어있으며, user row들은 Page 단위로 구분되어있습니다.

![](/images/postgresql/internal10.png)

하나의 page는 8kb로 구성되어있습니다. 

![](/images/postgresql/page_layout1.png)

PostgreSQL 공식 홈페이지의 [Page Layout](https://www.postgresql.org/docs/current/storage-page-layout.html)를 살펴보면 8kb인 페이지의 레이아웃은 크게 4가지로 구성되어있습니다.

![](/images/postgresql/page_layout2.png)

1. 헤더 데이터
    1. 24bytes long
    2. `pd_lower`	LocationIndex	2 bytes	Offset to start of free space
    3. `pd_upper`	LocationIndex	2 bytes	Offset to end of free space
2. 아이템 id(페이지내의 tuple의 위치를 가리키는 포인터)들
    1. 4byte per item
3. Free Space
4. Items

헤더 데이터 안에는 free space의 시작 지점에 대한 포인터와 free space가 끝나는 지점에 대한 포인터값이 존재합니다. 

페이지 레이아웃을 추상적으로 도식화 해보자면 아래와 같습니다.

![](/images/postgresql/internal11.png)

물론 물리적으로는 2진수(실제로는 16진수)로 저장되어 있기 때문에 page(block)은 실제로는 다음과 같은 형태가 됩니다.

![](/images/postgresql/internal12.png)

실제 22445 heap 파일의 첫번째 page를 vs code에서 꺼내보면 다음과 같이 보입니다. (with extension `hex editor` by ms)

![](/images/postgresql/internal13.png)

이 값을 postgresql의 구조에 따라 decode해보자면 아래와 같습니다.

| | |
|:---:|:---:|
| ![](/images/postgresql/internal14.png) | ![](/images/postgresql/internal15.png) |

## Index

### Full Scan
pk 또는 uniq 키 또는 secondary index를 적용시키지 않은 경우 postgresql은 full scan을 합니다.

![](/images/postgresql/full_scan1.png)

Disk에 있는 heap 파일을 찾아서, 메모리에 해당 heap에 존재하는 모든 tuple들을 memory에 올린뒤, 하나하나 찾습니다.

![](/images/postgresql/full_scan2.png)


### Index Scan

성능을 높이기 위해서는 index를 사용하면 됩니다. `B+ Tree` 형태로 관리되는 index는 index의 크기에 따라 level이 다르긴 하지만, root <-> leaf node 사이즈의 index를 기준으로 tree를 그려보면 아래와 같습니다.

![](/images/postgresql/index1.png)

Root 또는 branch node를 통해서, 검색 조건을 O(logn) 성능으로 타고 내려가서, leaf node에서 검색 결과에 맞는 CTID(=ROWID) 즉 (page number, row id)를 찾습니다. 이를 기반으로 index가 처리된 table의 heap파일에서 page 위치에 있는 row_id를 찾습니다.

다만 index 또한 trade-off가 존재합니다.

![](/images/postgresql/index2.png)

예를 들어 880kb 사이즈의 user table에 대해서 username에 대한 index를 생성하면 184kb의 index 파일을 추가로 저장해야합니다. RDB는 클라우드상에서 기본적으로 비싼 값을 내고 관리되어야 하니, 대규모 시스템에서는 상당한 규모의 index 관리를 위해 돈을 지불해야합니다.


| | |
|:---:|:---:|
| ![](/images/postgresql/index3.png) | ![](/images/postgresql/index4.png) |


index의 단점을 정리하면 총 3가지입니다.

1. 추가 disk 요구 (큰 disk 사이즈)
2. B+-Tree를 유지 관리해야 하기 때문에 insert / update / delete에 대한 추가 오버헤드 필요
3. Postgresql planner에 의해서 index를 실제로는 사용안할 수도 있다. (성능과 돈을 들였는데, index를 사용하지 않기 때문에 최악의 경우라고 할 수 있음)


### Default Index

Secondary Index를 생성하지 않더라도, 기본적으로 index를 생성해주는 index가 존재합니다. (FYI, postgresql의 모든 index는 secondary index입니다. innodb와 달리 pk가 클러스터링 index가 아닙니다.)

![](/images/postgresql/index5.png)

기본적으로 PK와 uniq키는 index를 생성해서 관리됩니다. 이들은 pgAdmin의 index 섹션에 보이지 않기 때문에, query로 찾아보면 아래와 같습니다.

![](/images/postgresql/index6.png)

### Index in detail

인덱스를 좀 더 상세히 살펴보자면, index는 table과 마찬가지로 Heap으로 disk에 저장되며, heap안에는 여러 pages들로 구분됩니다. 

| | |
|:---:|:---:|
| ![](/images/postgresql/index7.png) | ![](/images/postgresql/index8.png) |

Heap안에 page들을 memory에 올려 O(logn) scan을 실시합니다.

이제 그럼 pgAdmin으로 실제 index가 어떻게 생겼는지 확인해보록 하겠습니다. 먼저 index를 상세히 보기위해서 extension을 하나 설치해줍니다.

```sql
CREATE EXTENSION pageinspect;
```

그 뒤, users 테이블에 존재하는 username에 대한 index 형태를 확인해줍니다. `users_username_idx` 인덱스에 대해서 3번째 page를 확인해보면 아래와 같습니다.

![](/images/postgresql/index10.png)

이때 가장 첫번째 row를 확인하면 이 page가 root(or branch) page인지, leaf page인지 알 수 있습니다. 만약 첫번째 row의 data가 비워져있다면, 이는 root 또는 branch page이며, data가 존재한다면 이는 leaf page입니다. 

참고로 B+Tree에서 leaf node끼리는 연결되어있는데, 이를 표현한 것이 leaf page의 첫번째 row입니다. 해당 row의 data는 다음 leaf node의 index col 데이터를 의미하며, ctid를 찾아서 가게되면 다음 leaf page로 이동할 수 있습니다.

![](/images/postgresql/28-1.svg)

위의 다이어그램을 봐서 알듯이, page 3는 root page입니다. 이제 page 1번을 확인해보자면

![](/images/postgresql/index11.png)

첫번쨰 row의 data가 채워져있는 것을 확인할 수 있습니다. 2번째 row부터는 leaf page에 들어있는 element들입니다. 여기에서는 크게 (ctid, 특정 username) 정도가 들어있다 생각하면 됩니다.

![](/images/postgresql/index12.png)

(33,43)을 타고 들어가보면 실제 Users 테이블의 ctid와 동일하다는 것을 알 수 있습니다.

마지막으로 relkind = 'i' (index타입) 찾아보면 22713 이라는 heap 파일에 index가 저장된 것을 확인가능하며, 

![](/images/postgresql/index13.png)

이 heap파일은 이전 table 파일과 마찬가지로 page단위로 섹션이 구분된 것을 알 수 있습니다.

![](/images/postgresql/index14.png)

**User 테이블은 그 숫자가 적어, index의 B+tree가 높이 2로 존재했지만, 더 큰 likes 테이블의 경우 아래와 같이 더 깊이 있는 tree로 구성됩니다.**


![](/images/postgresql/28-2.svg)

## Planner

- `EXPLAIN`, Query Plan + Show
- `EXPLAIN ANALYZE`, Query Plan + Run + Show

쿼리 플랜에서 해석할 때, 밑에서 부터 위로 올라오는 순으로 쿼리가 실행되는 순서라고 생각하면 됩니다. 
또한 `->`이 들어 있는 문장 또는 가장 위에 존재하는 문장은 `plan node`로 해석되며, query plan들은 plan node의 순서로 진행됩니다.

```sql
SELECT username, contents
FROM users
JOIN comments ON comments.user_id = users.id
WHERE username = 'Alyson14';
```

![](/images/postgresql/plan3.png)

위의 다이어 그램은 실제 query plan을 표현한 그림으로, 아래에서 위로 올라오면서 paln node들이 진행된다고 생각하면 됩니다. 

![](/images/postgresql/plan1.png)

1. WHERE = 'Alyson14'인 `users_username_idx` 인덱스를 scan하여 users를 가져옵니다.
2. 9kB 버킷을 메모리에 올려 Hash를 생성합니다.
3. Comments 테이블을 Full Scan(=Seq Scan)하여 60410(guess)개의 row를 memory에 올리고
4. Hash join 하여 comments.user_id = users.id인 row들을 뽑아냅니다.

가장 상위의 plan node를 분석하면 아래와 같습니다.

![](/images/postgresql/plan2.png)

1. return할 데이터를 어떤 방식으로 plan했는지
2. plan과 execute에 대한 소요시간 추측 
3. 생성될 row 수를 추측
4. row의 평균 byte를 추측

이처럼 postgresql은 실제로 planning할 때, 시간 및 row 수를 guess할 수 있는데, 이에 필요한 메타데이터들은 pg_stats에서 확인가능합니다.

![](/images/postgresql/plan4.png)



## View


### Common Table Expression (CTE)

Common Table Expression(CTE)란 set of a query로 쿼리 가독성을 위해서 긴 쿼리문을 tmp하게 변수로 관리하는 것을 뜻합니다.

Recursive common table Expression은 CTE에서 recursive하게 참조해서 복잡한 쿼리문을 간단하게 표현하는 방식입니다. 특히 graph관계나 tree 관계에서 유용한데, 예를들어서 user와 follow 테이블이 존재할 때, 연결 관계가 4 depth까지인 모든 user들을 찾고 싶을 때와 같은 상황에서 유용하게 관리될 수 있습니다.

예를 들어 **tag가 많이된 username, 태그된 수를 태그순으로 정렬**해서 보여주려 할 때, 사진에 들어있는 tag와 post를 작성시 caption에서 태그한 사용자들을 모아서 user와 join한 뒤 count하는 방식으로 아래와 같은 query를 만들 수 있습니다. 

```sql
SELECT username, COUNT(*)
FROM users
JOIN (
	SELECT user_id FROM photo_tags
	UNION ALL
	SELECT user_id FROM caption_tags
) AS tags ON tags.user_id = users.id
GROUP BY username
ORDER BY COUNT(*) DESC;
```

이를 CTE를 활용하면, 좀 더 가시성 있게 쿼리를 만들 수 있습니다.

```sql
WITH tags AS (
	SELECT user_id FROM photo_tags
	UNION ALL
	SELECT user_id FROM caption_tags
)

SELECT username, COUNT(*)
FROM users
JOIN tags ON tags.user_id = users.id
GROUP BY username
ORDER BY COUNT(*) DESC;
```


### View

**하지만 CTE는 오직 선언된 다음의 SQL 쿼리에 국한되어 사용됩니다.** 즉 이후의 다른 쿼리에서 해당 tags CTE를 사용하고 싶어도 사용할 수 없습니다. 이를 위해서 view를 사용합니다.

![](/images/postgresql/view1.png)


메모리에 임시로 존재하는 CTE와 달리, `VIEW`는 쿼리 텍스트를 disk에 저장해서 관리합니다. 

![](/images/postgresql/view2.png)

즉 위의 경우 tag가 많이 된 user에 대한 쿼리를 매번 작성하기 보다, view를 사용해서 간편하게 쿼리를 실행할 수 있습니다. (view의 결과값은 disk에 저장되는게 아니라, 쿼리문 그 자체가 저장됨)

![](/images/postgresql/view3.png)

tag에 대한 view를 생성하려면 아래와 같습니다.

```sql
CREATE VIEW tags AS (
	SELECT id, created_at, user_id, post_id, 'photo_tag' AS type FROM photo_tags
	UNION ALL
	SELECT id, created_at, user_id, post_id, 'caption_tag' AS type FROM caption_tags
);
```

![](/images/postgresql/view4.png)

### Materialized Views

하지만 `view`의 경우, 매번 refer될 떄마다 execute해야 한다는 단점이 있습니다. 이를 해결하기 위해서 `Materialized Views`라는 개념이 도입됩니다.

![](/images/postgresql/view5.png)

View와 달리 `Materialized View`는 만들어지는 시점에 snapshot으로 쿼리된 데이터를 disk에 저장합니다. 다만 만들어지는 시점의 snapshot을 만들기 때문에, table이 변화될 때마다 refresh를 해주어야 합니다.

![](/images/postgresql/view6.png)

만약 매주, post와 comment에 대한 like 수를 테이블로 표현하고 싶은 경우, 쿼리문의 경우 아래와 같이 표현할 수 있습니다.

```sql
SELECT
	date_trunc('week', COALESCE(posts.created_at, comments.created_at)) AS week,
	COUNT(posts.id) AS num_post_likes,
	COUNT(comments.id) AS num_comment_likes
FROM likes
LEFT JOIN posts ON posts.id = likes.post_id
LEFT JOIN comments ON comments.id = likes.comment_id
GROUP BY week
ORDER BY week
```

![](/images/postgresql/view7.png)

이때 이를 Materialized View로 생성하면 아래와 같습니다.

```sql
CREATE MATERIALIZED VIEW weekly_likes AS (
	[[쿼리문]]
) WITH DATA;
```

```sql
CREATE MATERIALIZED VIEW weekly_likes AS (
	SELECT
		date_trunc('week', COALESCE(posts.created_at, comments.created_at)) AS week,
		COUNT(posts.id) AS num_post_likes,
		COUNT(comments.id) AS num_comment_likes
	FROM likes
	LEFT JOIN posts ON posts.id = likes.post_id
	LEFT JOIN comments ON comments.id = likes.comment_id
	GROUP BY week
	ORDER BY week	
) WITH DATA;
```

