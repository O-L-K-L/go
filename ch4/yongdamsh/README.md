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
- 

## Slice

### 특징
- 참조 타입
  - https://play.golang.org/p/l206Yl9BQhj
  - https://play.golang.org/p/G57qb5Eq7rH
- 

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
