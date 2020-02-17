# Chapter 6. Concurrency

## 고루틴

고루틴은 가벼운 스레드와 같은 것으로 현재 수행 흐름과 별개 흐름을 만든다.
별개의 흐름은 다음과 같이 나뉜다.
1. 병렬성 (Parallelism)
두 사람이 동시에 각각 업무를 보고 있는 경우

2. 동시성 또는 병행성 (Concurrency)
커피를 마시면서 신문을 보고 있는 사람

동시성은 병렬성과 다르지만 동시성이 있어야 병렬성이 생긴다.

### 생성 방법

https://play.golang.org/p/A9smhuLeKHu
```go
func main() {
  go func() {
    fmt.Println("goroutine")
  }()

  fmt.Println("main routine")
}
```

두 로그 중 어느 것이 먼저 실행될지 알 수 없다. 또는 고루틴이 실행되지 않을 수도 있다.

### 고루틴 기다리기

[`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup) 패키지를 통해 동시성 흐름을 제어한다.

https://play.golang.org/p/BbOU4aV0jl_x
```go
func main() {
  var wg sync.WaitGroup

  wg.Add(1)

  go func() {
    defer wg.Done()
    fmt.Println("goroutine")
  }()

  fmt.Println("main routine")

  wg.Wait()
}
```

### 고루틴 간의 통신

고루틴 간의 통신이 없는 경우 공유 메모리를 통해 문제 없이 병렬화할 수 있다.

https://play.golang.org/p/DGq8usJ-b_j
```go
func ParallelMin(a []int, n int) int {
  if len(a) < n {
    Min(a)
  }

  mins := make([]int, n)
  size := (len(a) + n - 1) / n

  var wg sync.WaitGroup

  for i := 0; i < n; i += 1 {
    wg.Add(1)

    go func(int i) {
      defer wg.Done()

      begin, end := i * size, (i + 1) * size

      if end > len(a) {
        end = len(a)
      }

      mins[i] = Min(a[begin:end])
    }(i)

    wg.Wait()

    return Min(mins)
  }
}


func Min(a []int) int {
  if len(a) == 0 {
    return 0
  }

  min := a[0]

  for _, v := a[1:] {
    if v < min {
      min = v
    }
  }

  return min
}
```

하지만 고루틴 간의 통신이 필요한 경우, 채널(channel)을 사용한다.
채널은 맵과 같이 참조 타입이고, `make`로 생성한다.

```go
c1 := make(chan int)
var chan int c2 = c1 // c2와 c1은 동일한 채널
var <-chan int c3 = c1 // Receive only channel
var chan<- int c4 = c1 // Send only channel
```

단방향 채널을 선언하는 이유? 
값을 받고 싶어 넘기거나 주고 싶어 넘기는 것을 명확히 하고자 할 때 실수를 방지하기 위함

채널을 기준으로 보내는 데이터의 수와 받는 데이터의 수가 일치해야 한다. 
그렇지 않으면 고루틴은 동작을 멈추고 다른 고루틴으로 문맥 전환(context switching)을 한다.
https://play.golang.org/p/1PYfOXuP-ge

이를 방지하기 위해 채널을 닫고, `range`로 받은 데이터 수만큼 수행하고, 단방향 채널을 함수로 넘기는 패턴을 자주 사용한다.

https://play.golang.org/p/D9eyq264L_9
```go
func main() {
  c := func() <-chan int { // 함수로 단방향 채널을 반환하기
    c := make(chan int)

    go func() {
      defer close(c) // 채널 닫기
      c <- 1
      c <- 2
      c <- 3
    }()

    return c
  }()

  for num := range c {
    fmt.Println(num)
  }
}
```


### 닫힌 채널
값을 받을 때 두 번째 boolean 값으로 채널이 열려 있는지 알 수 있다.
채널이 닫혀 있으면 zero value와 false를 반환한다.

```go
num, ok := <-c 
```


### 동시성 패턴

1. 파이프라인 패턴
한 단계의 출력이 다음 단계의 입력으로 이어지는 구조

https://play.golang.org/p/O9_pycVmeMM
```go
func PlusOne(in <-chan int) <-chan int {
  out := make(chan int)

  go func() {
    defer close(out)

    for num := range in {
      out <- num + 1
    }
  }()

  return out
}

func ExamplePlusOne() {
  c := make(chan int)

  go func() {
    defer close(c)

    c <- 5
    c <- 3
    c <- 8
  }()

  for num := range PlusOne(PlusOne(c)) {
    fmt.Println(num)
  }
}
```

2. 팬아웃
> 팬아웃 (Fan-out): 논리회로에서 주로 쓰이는 용어. 게이트 하나의 출력이 여러 입력으로 들어가는 경우

채널 하나를 여럿에게 공유하는 방식이다.

```go
func main() {
  c := make(chan int)

  for i := 0; i < 3; i++ {
    go func(i int) {
      for n := range c {
        time.Sleep(1)
        fmt.Println(i, n)
      }
    }(i)
  }

  for i := 0; i < 10; i++ {
    c <- i
  }

  close(c)
}
```


3. 팬인
> 팬인 (Fan-in): 논리회로에서 주로 쓰이는 용어. 하나의 게이트에 여러 개의 입력선이 들어가는 경우

```go
func FanIn(ins ...<-chan int) <-chan int {
  out := make(chan int)
  var wg sync.WaitGroup

  wg.Add(len(ins))

  for _, in := range ins {
    go func(in <-chan int) {
      defer wg.Done()

      for num := range in {
        out <- num
      }
    }(in)
  }

  go func() {
    wg.Wait()
    close(out)
  }()

  return out
}

func main() {
  c := FanIn(c1, c2, c3)
}
```


4. 분산처리

Fan-out으로 파이프라인을 통과시킨 뒤, Fan-in

5. Select

- 모든 case가 계산된다. 
- 각 case는 채널에 입출력하는 형태가 되며 막히지 않고 입출력이 가능한 case가 있으면 해당 case의 코드만 수행된다.
- default가 있으면 모든 case에 해당되지 않을 때 실행되고, 없으면 가능한 case가 발생할 때까지 기다린다.

```go
select {
  case n := <-c1:
    fmt.Println(n, "is from c1")
  case n := <-c2:
    fmt.Println(n, "is from c2")
  case c3 <- f():
    fmt.Println("f() return value is sent to c3")
  default:
    fmt.Println("No channel is ready")
}
```

6. 파이프라인 중단하기

7. context.Context 활용하기

8. 요청과 응답 짝짓기

9. 동적으로 고루틴 이어붙이기


### 고루틴과 채널 사용 시 주의점
![고루틴과 채널 사용 시 주의점](https://user-images.githubusercontent.com/4126644/73361315-11ff2080-42e8-11ea-99b9-a50ee166ff4f.jpg)


### 경쟁 상태

모든 고루틴이 막혀 교착 상태(deadlock)이 된 경우엔 에러가 발생하지만 
경쟁 상태(race condition)인 경우에는 버그를 쉽게 발견하기 어렵다.

아래 코드는 경쟁 상태에 놓일 수 있다. 실제로 `go run -race race.go`를 실행해보면 경쟁 상태 에러가 발생한다.
```go
func main() {
  cnt := int64(10)

  for i := 0; i < 10; i++ {
    go func() {
      cnt--
    }()
  }

  if cnt > 0 {
    time.Sleep(100 * time.Millisecond)
  }

  fmt.Println(cnt)
}
```

cnt 값을 메모리에서 가져와 1을 감소시키고 다시 저장하는 사이에 또다른 고루틴이 아직 저장되지 않은 cnt 값을 조회해 간다면
2가 감소되어야 할 값이 1만 감소될 수 있다. 이 경우 코드가 끝나지 않을 수 있다.

이러한 경쟁 상태 해결을 위해 [`atomic`](https://golang.org/pkg/sync/atomic/) 패키지를 활용할 수 있다.

`cnt--`를 `atomic.AddInt64(&cnt, -1)`로 변경  
`cnt > 0`를 `atomic.LoadInt64(&cnt) > 0`로 변경  

또는 채널을 이용해 atomic 없이도 동시성 문제를 해결할 수 있다.
`go run -race no_race.go`를 실행해 보자.

그 밖에 가장 먼저 실행되어야 하는 초기화에 쓸 수 있는 `sync.Once`와  
외부 자원 접근 시 상호 배타 잠금 기능을 위한 `sync.Mutex`, `sync.RWMutex` 등을 참고하자.


### 문맥 전환 (Context switching)
프로그램이 여러 프로세스 혹은 스레드에서 동작할 때 기존에 하던 작업들을 메모리에 보관해두고 다른 작업을 시작하는 것

Go 컴파일러는 아래의 경우에 문맥 전환을 하는 코드를 생성할 수 있다.
- 파일이나 네트워크 연산처럼 시간이 오래 걸리는 입출력 연산이 있을 때
- 채널에 보내거나 받을 때
- go로 고루틴이 생성될 때
- 가비지 컬렉션 사이클이 지난 뒤




## References
- 디스커버리 Go 언어 Chapter 7: http://www.hanbit.co.kr/store/books/look.php?p_code=B5279497767