# (WIP) Go Channel and Scheduler


go channel and scheduler
<!--more-->

# 1. Go channels
> Understanding Channels

## Make channel

1. G1, G2

- channel has mutex
- copy elements


## If full?

P: context for scheduling


1. P는 context를 들고 있으며, channel buffer가 full 되면 full된 버퍼로 task를 send하는 goroutine은 pausing되며 (receiver goroutine이 pause되는게 아니다.) 
2. go runtime scheduler로 `gopark`이 호출하여, pause를 진행한다. 
3. sender goroutine은 waiting state가 되며, 해당 고루틴의 os thread (M)과의 association을 지운다.
4. Q. wait되는 goroutine은 어디에 저장되는가? (https://www.linkedin.com/pulse/golang-how-does-goroutine-parks-chung-yi-kao/) gobuf는 g struct안의 sched에서 사용된다.
5. P가 들고있던 run queue에서 runnable goroutine을 pop한다.


# 2. Go scheduler
> Go scheduler: Implementing language with lightweight concurrency



## Refs

- [원본 유튜브 영상](https://www.youtube.com/watch?v=-K11rY57K7k)
- [한글: go scheduler](https://changhoi.kim/posts/go/go-scheduler/)
- [Schduling In Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
- [Illustrated Tales of Go Runtime Scheduler.](https://medium.com/@ankur_anand/illustrated-tales-of-go-runtime-scheduler-74809ef6d19b)
- [Go Concurrency Series: Deep Dive into Go Scheduler(I)](https://www.linkedin.com/pulse/go-concurrency-series-deep-dive-scheduleri-pratik-pandey-mhx4e/)
- c.f [Understanding the python GIL](https://www.dabeaz.com/GIL/)
    - https://dabeaz.blogspot.com/2010/02/revisiting-thread-priorities-and-new.html
