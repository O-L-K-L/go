## make 함수

- make 함수는 슬라이스 타입 오브젝트를 확보하고 초기화 한다.

```go
// Slice
// make(type, length[, capacity])
s1 := make([]int, 10)
s2 := make([]int, 10, 100)

// Map
// make(type)
m := make(map[string]string)
```

- 참고: go에서 new와 make 생성 방식 차이
  - https://yojkim.me/posts/golang-differences-between-new-and-make/

## map

- 맵(map)은 순서가 없는 키-값(key-value) 쌍의 집합이다. 맵은 연관 배열 또는 해시 테이블(hash table), 딕셔너리(dictionary)로도 알려져 있으며, 연관 키를 통해 값을 찾는 데 사용된다. 다음은 Go에서 맵을 사용하는 예제다.

```go
var x map[string]int

var x map[string]int
x["key"] = 10
fmt.Println(x)
```

- http://www.codingnuri.com/golang-book/6.html

## sync

- go루틴 사이의 동기화 지원

## go 에서의 식별자

- 패키지 외부로 노출 되는 것
  - 대문자로 시작
- 노출이 되지 않는 것
  - 소문자로 시작

## go에서의 변수 초기화

- go 에서는 모든 변수 제로값으로 초기화
  - 숫자에서는 0
  - 문자열에서는 빈문자열
  - Boolean은 false
  - 포인터는 nil
  - 참조 타입의 경우는 nil (ex: Map)

## :=

- 변수를 선언하는 또 다른 방식으로 `Short Assignment Statement ( := )` 를 사용할 수 있다. 즉, var i = 1 을 쓰는 대신 i := 1 이라고 var 를 생략하고 사용할 수 있다. 하지만 이러한 표현은 함수(func) 내에서만 사용할 수 있으며, 함수 밖에서는 var를 사용해야 한다. Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.

- http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98

## go의 채널

- 맵이나 슬라이스와 마찬가지로 참조 타입이지만 다른 타입들과 달리 채널은 고루틴 사이의 데이터 통신에 사용될 특정 타입의 값들을 위한 `큐`를 구현

## WaitGroup

- 실행하게 될 모든 고루틴 추척
  - 특정 고루틴이 작업을 완료했는지 추적할 수 있는 편리한 기능을 제공하기 때문에 가능
- 카운팅 세마포어여서 고루틴의 실행이 종료될 때마다 전체개수를 하나씩 줄여나간다.
  - `카운팅 세마포어` : 세마포어는 shared data의 개수를 말하는데, 카운팅 세마포어는 공유자원의 개수를 나타내는 단순한 변수이다.
  - https://jhnyang.tistory.com/101 (카운팅 세마포어 글)
- 제공함수
  - Add (대기 그룹에 고루틴 개수 추가)
  - Done (고루틴이 끝났다는 것을 알려줄때 사용)
  - Wait (모든 고루틴이 끝날 때까지 기다림)
  - http://pyrasis.com/book/GoForTheReallyImpatient/Unit35/06 참고

## 고루틴

- 프로그램 내의 다른 함수와는 독립적으로 실행되는 함수
- 고루틴을 실행하고 동시적 실행을 위한 스케줄링을 시도할때는 go 키워드를 사용하면 된다.

## defer함수

- 함수가 리턴된 직후 실행될 작업을 예약하기 위한 키워드

## Decode 메서드

- 어떤 타입이든 받아들일 수 있도록 설계

## interface{}

- go에서 특별하게 취급하는 타입
- reflect 패키지를 이용한 reflection 지원이 가능한 타입
  - 리플렉션은 실행 시점(Runtime, 런타임)에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능입니다.
  - http://pyrasis.com/book/GoForTheReallyImpatient/Unit36

## 인터페이스

- 구조체나 다른 명명된 타입들이 어떤 조건을 만족하기 위해 구현해야 하는 동작을 정의하는 타입이 인터페이스 타입
- 인터페이스의 동작은 인터페이스 타입 내부에 선언된 메서드에 의해 정의된다.
- 인터페이스를 정의할때는 Go의 이름 규칙을 준수하는 편이 좋다. 인터페이스가 하나의 메서드만을 선언하고 있다면 인터페이스의 이름은 `er` 접미사로 끝나야 한다. ex) Matcher

## 패키지 import 하는것 앞에 \_(빈 식별자)

-     _ "./matchers"
- 이렇게 함으로써 컴파일러가 오류 없이 패키지를 가져오고 init 함수를 호출할 수 있다.
