# CH6 동시성

Go에서 동시성(Concurrency)은 함수들을 다른 함수들과 독립적으로 실행할 수 있는 기능이고 고루틴을 통해서 구현되었다. Go 런타임 스케쥴러는 고루틴 관리하는데 운영체제를 기반으로 동작한다. 동시성에서 동기화는 `CSP(Communicating Sequential Processes)`를 기반으로 설계되었고 이는 고루틴 사이에 `채널`을 통해 메시지를 교환한다.

## 동시성과 병렬성

Go 스케쥴러의 논리적 프로세스는 운영체제의 스레드에 1대 1로 바인딩되고 이 한 스레드에서 여러개의 고루틴이 실행된다.

- 실행을 중단해야하는 시스템 콜을 수행하는 경우: 스레드와 고루틴이 분리되고 논리 프로세서는 새로운 스레드를 할당받고 여기서 고루틴을 실행한다. 그리고 시스템 콜이 리턴되면 다시 실행된다.
- 네트워크 I/O가 발생한 경우: 고루틴이 논리 프로세서에서 분리되고 네트워크 풀러(poller)로 이동되고 I/O를 처리할 준비가 되면 다시 논리 프로세서로 돌아와 처리된다.

1. 병렬성: 여러 코드가 각각의 다른 물리적 프로세서에서 동시에 실행되어 물리적 프로세서 사이에 메모리 공유 등이 불가능한 방법이다.
2. 동시성: 논리적 프로세서 안에서 한 번에 여러 작업을 관리하는 것을 말한다.

## 고루틴

하나의 논리 프로세서 안에서 여러개의 고루틴은 스레드를 번걸아가며 점유하며 실행된다. 완료되지 않은 고루틴은 실행 큐로 되돌아간다.
[고루틴 기본 코드](https://play.golang.org/p/_ZXrLutLKcz)

cf. play.golang에서 병렬적으로 실행되지 않는데 물리적 프로세서가 하나로만 되어있어서 그런걸까?

```go
import "runtime"

// 코어마다 하나의 논리 프로세서를 지정
runtime.GOMAXPROCS(runtime.NumCPU())
```

## 경쟁 상태(Race condition)

공유 자원에 동시에 접근하여 데이터가 변경되는 현상. 이를 방지하기 위해 한 시점에는 한 고루틴만이 접근해야 한다.
[경쟁 상태](https://play.golang.org/p/2y9P0EP-HRX)

## 공유 자원 잠금

동기화를 위한 도구들

### 원자성 함수들

```go
for count := 0; count < 2; count++ {
	atomic.AddInt64(&counter, 1) // counter 변수에 한 고루틴만이 접근 가능하도록 함

	runtime.Gosched()
}

// cf.
atomic.StoreInt64(&variable, 1) // 경쟁 상태에 놓이지 않고 variable 변수에 할당
atomic.LoadInt64(&variable) // variable 값을 로드
```

[Atomic 라이브러리를 이용한 예제](https://play.golang.org/p/-uNf2akWru8)

### 뮤텍스(Mutual exclusion)

공유 자원을 사용하는 코드 주변에 임계 지역을 설정해 그 지역 안에서 하나의 고루틴만이 접근 가능하도록 한다(Java의 synchronized 같은 개념).

```go
mutex sync.Mutex

mutex.Lock()
//... 임계 구역에는 한번에 하나의 고루틴만이 접근할 수 있음
mutex.Unlock()
```

## 채널

고루틴 사이에 데이터를 교환할 때, 동기화를 보장해주는 도구.

```go
unbuffered := make(chan int) // chan 키워드와 타입(내장 타입, 사용자정의 타입, 구조체, 참조 타입 등 모두 가능)

buffered := make(chan string, 10)

buffered <- "Gopher" // 버퍼 채널을 통해 "Gopher" 스트링을 전달
value := <- buffered // 다른 고루틴에서 버퍼 채널을 통해 "Gopher" 스트링을 전달받음
```

### 버퍼가 없는 채널

값을 전달받기 전에 얼마나 보유할지 모르므로 두 고루틴이 같은 시점에 채널을 사용할 준비가 되어있어야 함.
[테니스 게임 예제](https://play.golang.org/p/nYP_LiCO3S4)
[계주 달리기 예제](https://play.golang.org/p/kFTlS4R8OKm)

### 버퍼가 있는 채널

채널에 보관할 수 있는 값의 개수를 지정할 수 있고 이로 인해 동시에 고루틴이 실행될 필요가 없다.
[작업자 예제](https://play.golang.org/p/rZpsVTdfW3T)
