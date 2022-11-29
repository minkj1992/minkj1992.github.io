# Gevent



Let's deep dive into gevent based python programming
<!--more-->
<br />

## tl;dr

```python
from gevent import (
    event,
    joinall,
    queue,
    sleep,
    threadpool,
    Timeout,
    wait,
    kill,
    get_hub,
    signal,
)
```

## `gevent`

*gevent is a coroutine-based Python networking library that uses greenlet to provide a high-level synchronous API on top of the libev or libuv event loop.*

- **deterministic**
- 


## Intro

### `gevent` vs `eventlet`
gevent is inspired by eventlet has some diff

1. gevent is built on top of `libevent` event loop
  - epoll on Linux
  - Signal handling is integrated with the event loop.
  - Other libevent-based libraries can integrate with your app through single event loop.
  - DNS requests are resolved asynchronously rather than via a threadpool of blocking calls.
  - WSGI server is based on the libevent’s built-in HTTP server
2. gevent’s interface follows the conventions set by the standard library


### backgrounds
{{< admonition tip "greenlet">}}
**Greenlet은 C 확장 모듈 형태로 제공되는 경량 코루틴 입니다**. Greenlet들은 메인 프로그램을 실행하는 OS process 안에서 모두 실행되지만 상호작용하며 스케줄링됩니다. 운영체제에 의해 스케줄링되는 process들과 POSIX 쓰레드들을 사용하여 실제로 병렬로 실행되는 multiprocessing 나 threading을 이용한 병렬처리들과는 달리 **한 번에 오직 하나의 greenlet만이 실행됩니다.**

greenlet은 deterministic 합니다. 같은 greenlet 세팅과 같은 입력이 주어졌을때, 언제나 같은 결과를 출력합니다.

```python
import time

def echo(i):
    time.sleep(0.001)
    return i

# Non Deterministic Process Pool

from multiprocessing.pool import Pool

p = Pool(10)
run1 = [a for a in p.imap_unordered(echo, range(10))]
run2 = [a for a in p.imap_unordered(echo, range(10))]
run3 = [a for a in p.imap_unordered(echo, range(10))]
run4 = [a for a in p.imap_unordered(echo, range(10))]

print(run1 == run2 == run3 == run4)

# Deterministic Gevent Pool

from gevent.pool import Pool

p = Pool(10)
run1 = [a for a in p.imap_unordered(echo, range(10))]
run2 = [a for a in p.imap_unordered(echo, range(10))]
run3 = [a for a in p.imap_unordered(echo, range(10))]
run4 = [a for a in p.imap_unordered(echo, range(10))]

print(run1 == run2 == run3 == run4)

False
True

```

{{< /admonition >}}

{{< admonition tip "libev vs libuv vs libevent" >}}

The three network libraries Libevent , libev , and libuv are all asynchronous event libraries implemented in C language. ( Asynchronousevent library ).

- `Libevent`: the most famous, most widely used and long-standing cross-platform event library;

- `libev`: Compared with libevent , the design is more concise and the performance is better, but the support for Windows is not good enough;

- `libuv`: A cross-platform event library is needed in the process of developing node . They preferred libev , but they also need to support Windows , so they repackaged a set, implemented with libev under Linux , and IOCP under Windows
{{< /admonition >}}


## key features
> [refs](http://leekchan.com/gevent-tutorial-ko/)

### Synchronous & Asynchronous Execution on gevent

동시성 처리의 핵심 개념은 큰 단위의 task를 한번에 동기로 처리하는 대신, 작은 단위의 subtask들로 쪼개서 동시에 비동기로 실행시키는 것입니다. 두 subtask간의 스위칭을 컨텍스트 스위칭이라고 합니다.
**gevent에서는 컨텍스트 스위칭을 yielding을 이용해서 합니다.**

gevent의 진짜 힘은 상호작용 하며 스케쥴링 될 수 있는 네트워크와 IO bound 함수들을 작성할때 발휘됩니다. **gevent는 네트워크 라이브러리들이 암시적으로 greenlet 컨텍스트들이 가능한 시점에 암시적으로 yield 하도록 보장합니다.**

```python
import time
import gevent
from gevent import select

start = time.time()
tic = lambda: 'at %1.1f seconds' % (time.time() - start)


def gr1():
    # Busy waits for a second, but we don't want to stick around...
    print('Started Polling: %s' % tic())
    select.select([], [], [], 2)
    print('Ended Polling: %s' % tic())


def gr2():
    # Busy waits for a second, but we don't want to stick around...
    print('Started Polling: %s' % tic())
    select.select([], [], [], 2)
    print('Ended Polling: %s' % tic())


def gr3():
    print("Hey lets do some stuff while the greenlets poll, %s" % tic())
    gevent.sleep(1)
    print('gr3 fin')


gevent.joinall([
    gevent.spawn(gr1),
    gevent.spawn(gr2),
    gevent.spawn(gr3),
])



Started Polling: at 0.0 seconds
Started Polling: at 0.0 seconds
Hey lets do some stuff while the greenlets poll, at 0.0 seconds
Ended Polling: at 2.0 seconds
Ended Polling: at 2.0 seconds
```

### Determinism
greenlet은 deterministic 합니다. 같은 greenlet 세팅과 같은 입력이 주어졌을때, 언제나 같은 결과를 출력합니다.

gevent가 일반적으로 deterministic 하다고 해도, 소켓과 파일과 같은 외부 서비스와 연동할 때 non-deterministic한 입력들이 들어올 수 있습니다. 그러므로 green 쓰레드가 "deterministic concurrency" 형태라고 해도, POSIX 쓰레드들과 프로세스들을 다룰 때 만나는 문제들을 경험할 수 있습니다.

동시성을 다룰 때 만날 수 있는 문제로 race condition이 있습니다. 간단히 요약하자면, race condition은 두 개의 동시에 실행되는 쓰레드나 프로세스들이 같은 공유 자원을 수정하려고 할 때 발생합니다. 이때 해당 공유자원의 결과 값은 실행 순서에 따라 달라지게 됩니다. 이런 결과는 non-deterministic한 프로그램 동작을 야기하기 때문에 발생시키지 않기 위해 노력해야 합니다.
**Best practice는 공유자원을 사용하지 않도록 하는 것입니다.**

### Spawning Greenlets

gevent는 greenlet 초기화를 위한 몇 가지 wrapper들을 제공합니다.

```python
import gevent
from gevent import Greenlet

def foo(message, n):
    gevent.sleep(n)
    print(message)


c1 = Greenlet.spawn(foo, "Hello", 1)
c2 = gevent.spawn(foo, "Hello2", 2)
c3 = gevent.spawn(lambda x: (x+1), 5)

coroutines = [c1,c2, c3]

gevent.joinall(coroutines)


6
Hello
Hello2
```

Greenlet 클래스를 상속하고 _run 함수를 override 하는 방법도 있습니다.

```python
import gevent
from gevent import Greenlet

class MyGreenlet(Greenlet):

    def __init__(self, message, n):
        Greenlet.__init__(self)
        self.message = message
        self.n = n

    def _run(self):
        print(self.message)
        gevent.sleep(self.n)

g = MyGreenlet("Hi there!", 3)
g.start()
g.join()

Hi there!
```

### Greenlet State

다른 코드 예시들처럼, greenlet도 다양한 경우에 실패할 수 있습니다. greenlet은 예외를 발생시키는것이 실패하거나, 정지에 실패할 수도 있고, 시스템 자원을 과도하게 사용할 수도 있습니다.

**greenlet의 내부 상태는 대체로 time-dependent합니다. greenlet에는 쓰레드의 상태를 모니터링 할 수 있는 다양한 flag들이 있습니다.**

- started: bool
  - Greenlet이 실행되었는지 여부를 나타냅니다
- ready(): bool
  - Greenlet이 정지되었는지 여부를 나타냅니다
- successful(): bool
  - Greenlet이 예외를 발생시키지 않고 정지되었는지 여부를 나타냅니다.
- value: Any
  - Greenlet에 의해서 리턴된 값입니다.
- exception: Exception
  - Greenlet안에서 발생한 예외입니다.

```python

import gevent

def win():
    return 'You win!'

def fail():
    raise Exception('You fail at failing.')

winner = gevent.spawn(win)
loser = gevent.spawn(fail)
print(winner.started) # True
print(loser.started)  # True

try:
    gevent.joinall([winner, loser])
except Exception as e:
    print('This will never be reached')
print(winner.value) # 'You win!'
print(loser.value)  # None

print(winner.ready()) # True
print(loser.ready())  # True

print(winner.successful()) # True
print(loser.successful())  # False

# It is possible though to raise the exception again outside
# raise loser.exception
# or with
# loser.get()
print(loser.exception) # You fail at failing.
```

### Program Shutdown

메인 프로그램이 SIGQUIT 시그널을 받은 시점에 yield를 실패한 Greenlet은 예상보다 오래 실행이 정지되어 있을 수 있습니다. 이런 프로세스는 "좀비 프로세스"라고 불리고, 파이썬 인터프리터 외부에서 kill되어야 합니다.

일반적인 패턴은 메인 프로그램에서 SIGQUIT 시그널에 대기하고 있다가 프로그램이 종료되기 전에 
gevent.shutdown 호출하는 것입니다.

```python
import gevent
import signal

def run_forever():
    gevent.sleep(1000)

if __name__ == '__main__':
    # gevent.signal(signal.SIGQUIT, gevent.shutdown)
    gevent.signal(signal.SIGQUIT, gevent.kill)
    thread = gevent.spawn(run_forever)
    thread.join()
```

아래는 realworld example입니다.

```python
def main():
    parser = ArgumentParser()
    parser.add_argument("--start",
                        action="something start",
                        help="Launch the processes")
    args = parser.parse_args()

    threading.currentThread().setName('something')
    launcher = Launcher()
    launcher.setup_procmanager()
    if args.start:
        launcher.start_proc_manager()
    try:
        gevent.signal.signal(gevent.signal.SIGTERM, sig_handler)
        gevent.signal.signal(gevent.signal.SIGINT, sig_handler)
        gevent.wait()
    except KeyboardInterrupt:
        print('Shutdown requested. Exiting...')
        launcher.shutdown()


if __name__ == '__main__':
    main()
```

### Timeouts
```python
import gevent
from gevent import Timeout

seconds = 10

timeout = Timeout(seconds)
timeout.start()

def wait():
    gevent.sleep(10)

try:
    gevent.spawn(wait).join()
except Timeout:
    print('Could not complete')
```

- `with statement`

```python
import gevent
from gevent import Timeout

time_to_wait = 5 # seconds

class TooLong(Exception):
    pass

with Timeout(time_to_wait, TooLong):
    gevent.sleep(10)
```

### Events

Event는 Greenlet 간의 비동기 통신에 사용됩니다.

```python
import gevent
from gevent.event import Event
'''
Illustrates the use of events
'''

evt = Event()


def setter():
    '''After 3 seconds, wake all threads waiting on the value of evt'''
    print('A: Hey wait for me, I have to do something')
    gevent.sleep(3)
    print("Ok, I'm done")
    evt.set()


def waiter():
    '''After 3 seconds the get call will unblock'''
    print("I'll wait for you")
    evt.wait()  # blocking
    print("It's about time")


def main():
    gevent.joinall([
        gevent.spawn(setter),
        gevent.spawn(waiter),
        gevent.spawn(waiter),
        gevent.spawn(waiter),
        gevent.spawn(waiter),
        gevent.spawn(waiter)
    ])


if __name__ == '__main__':
    main()


A: Hey wait for me, I have to do something
I'll wait for you
I'll wait for you
I'll wait for you
I'll wait for you
I'll wait for you
Ok, I'm done
It's about time
It's about time
It's about time
It's about time
It's about time
```

Event 객체의 확장은 wakeup call과 함께 값을 전송할 수 있는 `AsyncResult`입니다. AsyncResult는 임의의 시간에 할당될 미래 값에 대한 레퍼런스를 갖고 있기 때문에, **때때로 `future`나 `deferred`로 불리기도 합니다.**

```python
import gevent
from gevent.event import AsyncResult

a = AsyncResult()


def setter():
    a.set('leoo is cool')
    gevent.sleep(0.5)
    a.set('minwook is cool')


def waiter():
    print(a.get())
    gevent.sleep(1)
    print(a.get())


gevent.joinall([
    gevent.spawn(setter),
    gevent.spawn(waiter),
])

leoo is cool
minwook is cool
```

### Queues

Queue는 일반적인 put 과 get 연산을 지원하지만 Greenlet 사이에서 안전하게 조작되는 것이 보장되는 순서를 가진 데이터들의 집합입니다.

```python
import gevent
from gevent.queue import Queue

tasks = Queue()

def worker(n):
    while not tasks.empty():
        task = tasks.get()
        print('Worker %s got task %s' % (n, task))
        gevent.sleep(0)

    print('Quitting time!')

def boss():
    for i in xrange(1,25):
        tasks.put_nowait(i)

gevent.spawn(boss).join()

gevent.joinall([
    gevent.spawn(worker, 'steve'),
    gevent.spawn(worker, 'john'),
    gevent.spawn(worker, 'nancy'),
])


Worker steve got task 1
Worker john got task 2
Worker nancy got task 3
Worker steve got task 4
Worker john got task 5
Worker nancy got task 6
Worker steve got task 7
Worker john got task 8
Worker nancy got task 9
Worker steve got task 10
Worker john got task 11
Worker nancy got task 12
Worker steve got task 13
Worker john got task 14
Worker nancy got task 15
Worker steve got task 16
Worker john got task 17
Worker nancy got task 18
Worker steve got task 19
Worker john got task 20
Worker nancy got task 21
Worker steve got task 22
Worker john got task 23
Worker nancy got task 24
Quitting time!
Quitting time!
Quitting time!
```

**즉 Queue는 put이나 get 연산 시 block 됩니다.**
만약 *non-blocking 연산이 필요할 때는 block이 되지 않는 put_nowait과 get_nowait을 사용할 수 있습니다. 대신 연산이 불가능할 때는 `gevent.queue.Empty` 나 `gevent.queue.Full` 예외를 발생시킵니다.*

아래 코드는 상사가 3명의 작업자(steve, john, nancy)에게 동시에 일을 시키는데 Queue가 3개 이상의 요소를 담지 않도록 제한하는 예시입니다. 이 제한은 put연산이 Queue에 남은 공간이 있을 때 까지 block 되어야 함을 의미합니다. 반대로 get 연산은 Queue에 요소가 없으면 block 되는데, 일정 시간이 지날 때 까지 요소가 들어오지 않으면 gevent.queue.Empty 예외를 발생시키면서 종료될 수 있도록 타임아웃 파라미터를 설정할 수 있습니다.

```python

import gevent
from gevent.queue import Queue, Empty

tasks = Queue(maxsize=3)

def worker(name):
    try:
        while True:
            task = tasks.get(timeout=1) # decrements queue size by 1
            print('Worker %s got task %s' % (name, task))
            gevent.sleep(0)
    except Empty:
        print('Quitting time!')

def boss():
    """
    Boss will wait to hand out work until a individual worker is
    free since the maxsize of the task queue is 3.
    """

    for i in xrange(1,10):
        tasks.put(i)
    print('Assigned all work in iteration 1')

    for i in xrange(10,20):
        tasks.put(i)
    print('Assigned all work in iteration 2')

gevent.joinall([
    gevent.spawn(boss),
    gevent.spawn(worker, 'steve'),
    gevent.spawn(worker, 'john'),
    gevent.spawn(worker, 'bob'),
])


Worker steve got task 1
Worker john got task 2
Worker bob got task 3
Worker steve got task 4
Worker john got task 5
Worker bob got task 6
Assigned all work in iteration 1
Worker steve got task 7
Worker john got task 8
Worker bob got task 9
Worker steve got task 10
Worker john got task 11
Worker bob got task 12
Worker steve got task 13
Worker john got task 14
Worker bob got task 15
Worker steve got task 16
Worker john got task 17
Worker bob got task 18
Assigned all work in iteration 2
Worker steve got task 19
Quitting time!
Quitting time!
Quitting time!
```

