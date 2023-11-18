# JS garbage collection (2018 v8)


[Concurrent marking in V8](https://v8.dev/blog/concurrent-marking)의 내용을 참조해서 v8의 garbage collection의 원리를 정리합니다. 이 GC는 Chrome 64 및 Node.js v10 부터 적용 되어 있습니다.
<!--more-->

## 기존 GC in JS (v8)

- **자바스크립트는 reachability(도달 가능성) 개념으로 메모리 관리를 수행합니다.** 즉 도달 가능한 값은 메모리에서 삭제되지 않습니다.

- mark-and-sweep, 루트들로 부터 시작해서 reach하는 영역들 마크하고, 마크되지 않은 영역 sweep
- generational collection, 오래된 객체와 새로운 객체를 나눠서 관리하는 방법

이전 JS의 GC는 주로 메인 스레드에서 동작했습니다. 당연히 marking 알고리즘의 마킹이 진행되는 동안, application은 일시 중지 된 경우에만 동작하기 떄문에 stop-the-world가 발생하게 됩니다.


![](/images/v8gc/v8_gc0.svg)

마킹 상태는 3가지 상태로 구분되며, 최초 root에서 출발합니다. (grey로 색칠)

1. white(00): 초기 상태 
2. grey(10): gc collector가 발견하여, marking worklist로 push한 상태
3. black(11): worklist에서 pop하여 Object의 모든 필드를 visit한 상태

![](/images/v8gc/v8_gc00.svg)

더이상 grey object가 없게되면 마킹은 중단되고, 남아있는 white node들은 unreachable로 간주되어 제거됩니다.

### incremental collection
> 가비지 컬렉션의 검사해야 하는 heap을 여러 부분으로 분리한 다음, 각 부분들을 별도로 수행하는 방법.

![](/images/v8gc/v8_gc1.svg)


Stop-the-World 시간을 줄이기 위해서 v8은 2011년 incremental collection을 도입했습니다. **이를 통해 GC는 더 작은 청크로 분할하고 애플리케이션이 청크 사이에서 실행될 수 있도록 합니다.**


![](/images/v8gc/v8_gc2.svg)

단, `Incremental marking`은 공짜로 이뤄지지 않습니다. application은 object graph (heap에서 root 부터 존재하는 instance들의 graph)가 변경될 때마다, GC에 notify해줘야 합니다. v8은 이 notification을 Dijkstra-style의 `write-barrier`를 통해 구현했습니다. 


아래 코드는 object.field = value와 같은 할당이 일어나게 될 경우, object의 색깔을 grey로 변경시키고, worklist에 다시 push하는 코드를 나타냅니다.
```js
// Called after `object.field = value`.
write_barrier(object, field_offset, value) {
  if (color(object) == black && color(value) == white) {
    set_color(value, grey);
    marking_worklist.push(value);
  }
}
```

`write-barrier`는 black object는 모든 필드가 검사된 상태이기 때문에, 흰색 entity를 가리키지 않는다는 사실에 근거해 있습니다.


### idle-time colleciton
> 가비지 컬렉터는 실행에 주는 영향을 최소화하기 위해, cpu가 idle 상태일 때만 GC를 실행하는 방법

v8의 incremental collection으로 쪼개진 chunk단위로 gc가 이뤄지는 것은 cpu idle time에 gc가 스케쥴링 되도록 하는 idle-time collection 기법과 매우 효과적으로 작동합니다.

## 2. 새로운 v8 마킹 방법

2018년 v8에서는 새로운 접근방식인 [Concurrent marking in V8](https://v8.dev/blog/concurrent-marking)을 발표했습니다. (Chrome 64 및 Node.js v10에 적용)

이 접근방식은 크게 2가지로 진행됩니다.

1. Parallel Marking
2. Concurrent Marking

![](/images/v8gc/v8_gc3.svg)

- `Parallel Marking`은 main thread와 모든 worker thread를 중단 시킨 뒤, parallel하게 mark작업을 진행하는 방식입니다.

![](/images/v8gc/v8_gc4.svg)

- `Concurrent Marking`은 마킹을 주로 worker thread에 위임하고 marking이 되더라도, main thread에서는 application을 지속적으로 실행하는 방식입니다.



@TODO CONTINUE

- https://v8.dev/blog/concurrent-marking#parallel-marking 
