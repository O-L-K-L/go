# Chapter 4. Arrays, slices, and maps

## Array

### 특징
- Value type이다. 따라서 함수에 인자로 넘기면 값이 복사된다. 참조 형태로 전달하려면 포인터를 사용해야 한다. 하지만 Go의 방식에 어울리지 않는다. 필요하다면 slice를 활용하자.
  ```go
  func Sum(arr *[5]int) (sum int) {
    for _, v := range *arr {
      sum += v
    }

    return
  }

  func main() {
    arr := [5]int{1,2,3,4,5}
    
    fmt.Println("total", Sum(&arr))
  }
  ```
- 타입으로 설정된 length만큼 초기값(zero value)이 할당되므로 별도로 초기화할 필요 없다.

## Slice

### 특징
- 참조 타입
  - https://play.golang.org/p/l206Yl9BQhj
  - https://play.golang.org/p/G57qb5Eq7rH

### Capacity
- `make` 생성자의 세번째 파라미터로 지정한다. 지정하지 않으면 length와 같은 값으로 설정된다.
- `cap` 함수로 현재 capacity 값을 조회할 수 있다.
- Capacity를 변경하는 방법
  1. `copy` 함수
    - `func copy(dst, src []T) int`
  2. `append` 함수
    - `append` 함수를 통해 capacity를 확장할 수 있다. 이 때 length에 따라 capacity 증가량이 달라진다.
    - [Capacity grow algorithm](https://github.com/golang/go/blob/master/src/runtime/slice.go#L95-L114)

### 주의사항
Slice는 참조 타입이기 때문에 GC로 정리되기 전까진 참조하는 배열의 메모리가 계속 유지된다. 따라서 파일 전체를 메모리에 유지하는 등의 비효율적인 경우를 피해야 한다.
- [관련 사례와 해결책](https://blog.golang.org/go-slices-usage-and-internals#TOC_6.)

## Map
- 참조 타입
- 초기화 방법
  ```go
  // make
  dict := make(map[string]int)

  // map literal
  dict := map[sting]int{}

  // nil map
  var dict map[string]int
  ```go
  
  ```
- key에는 equality(`==`)가 구현된 타입만 가능하다.
  - slice, function, slice가 포함된 struct 등은 불가
- key 존재 여부 확인
  ```go
  legs := map[string]int{
    "dog": 4,
    "bird": 2,
  }

  key := "dog"
  value, exists := legs[key]

  if exists {
    fmt.Printf("The key '%s' has value %d\n", key, value)
  } else {
    fmt.Printf("The key '%s' does not exist\n", key)
  }
  ```

## Pointer

- HTTP 서버에서 request를 포인터로 받는 이유?
  - [`http.ResponseWriter`](https://golang.org/pkg/net/http/#ResponseWriter)는 인터페이스고, [`http.Request`](https://golang.org/pkg/net/http/#Request)는 struct 타입이기 때문
  - [Stack Overflow Question](https://stackoverflow.com/questions/13255907/in-go-http-handlers-why-is-the-responsewriter-a-value-but-the-request-a-pointer)


## Learn More

### Slice Tricks
- js의 Array method와 같은 기능 구현을 위한 팁: 
https://github.com/golang/go/wiki/SliceTricks

### Slice Usage & Internals
https://blog.golang.org/go-slices-usage-and-internals

- Slice는 array 상의 추상 타입이다.
- Array를 slice로 변환하기
  - https://play.golang.org/p/sWGYBTyRPj5

### Slice의 length와 capacity를 다르게 초기화하는 경우는 언제 필요한가?
```go
slice := make([]int, 3, 5)
```
자동으로 capacity를 늘려 주는 [grow 알고리즘](https://github.com/golang/go/blob/master/src/runtime/slice.go#L76-L191) 연산 비용을 줄이기 위함일까?

### 다양한 데이터베이스 적용
1. 메모리
2. 파일 시스템
  - [json](https://golang.org/pkg/encoding/json/)
    - 관련 [commit](https://github.com/yongdamsh/estimator/commit/60e860c9dc8fac61b8ea1fac1d197dc4c6cb0a07#diff-bb7589ef3d57eaea1d896c9d874d303d)
  - [csv](https://golang.org/pkg/encoding/csv/)
    - 관련 [commit](https://github.com/yongdamsh/estimator/commit/c009a7a4fe7a7a8ec6b111d41ffc91ded36fb87d)
3. [sql](https://golang.org/pkg/database/sql)
  - 단일 사용자 환경의 소규모 앱에 적합한 sqlite 적용
    - 관련 [commit](https://github.com/yongdamsh/estimator/commit/1b97c4d552ca9528736720068131d815f0942ade)
