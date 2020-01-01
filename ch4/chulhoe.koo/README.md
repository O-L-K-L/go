### CH4 배열, 슬라이스, 맵

#### 배열

`배열의 특징: 정해진 자료형과 길이가 있다면 메모리에 순차적으로 저장되어 효율적인 데이터 구조`

1. 배열의 선언

```go
// 배열 선언
var array = [5]int

// 배열 리터럴로 선언
array := [5]int{10, 20, 30, 40, 50}

// ...을 이용해 자동으로 배열 길이 지정
array := [...]int{10, 20, 30, 40, 50}

// 일부만 초기화
array := [5]int{0: 10, 3: 20}
```

2. 배열의 활용

```go
// 포인터 연산자를 활용
array := [5]*int{0: new(int), 2: new(int)}
*array[0] = 10
*array[2] = 30

// 배열은 값으로 인식되어 같은 타입일 때, 대입 연산이 가능하다. 값, 길이 등의 정보가 모두 대입된다.
var array1 [5]string
array2 := [5]string{'a', 'b', 'c', 'd', 'e'}

array1 = array2
```

3. 다차원 배열

```go
var array [3][2]int
array2 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
array3 := [3][2]int{0: {1, 2}, 2: {5, 6}}
```

#### 슬라이스

`동적 배열의 개념으로 내부적으로 배열을 가리키는 포인터 주소, 길이, 최대 용량으로 구성되어 있다.`

1. 생성 및 초기화

```go
// make를 이용한 생성
slice := make([]string, 5) // 길이와 최대 용량이 5
slice := make([]string, 3, 5) // 길이가 3, 최대 용량이 5

// 리터럴을 이용한 생성
slice := []string{"red", "blue", "yellow", "green", "purple"} // 동적 배열이므로 []안에 숫자가 안 들어감

// nil 슬라이스: 포인터 주소가 아직 nil임
var slice []int

// 빈 슬라이스: 내부 배열을 가지고 있지 않은 슬라이스
slice := make([]int, 0)
slice := []int{}
```

2. 슬라이스의 활용

```go
// 배열과 동일하게 index로 대입 및 읽기 가능

// 잘라내기: 두 슬라이스는 내부 배열을 공유한다. 따라서 하나마 수정되면 다른 슬라이스 값도 변경된다.
slice := []int{10, 20, 30, 40, 50}
newSlice := slice[1:3] // 20, 30

// append: 용량이 남아있으면 기존 내부 배열에 추가하고 용량이 꽉 차 있으면 새로운 내부 배열을 생성한다.
slice := []int{10, 20, 30, 40, 50}
newSlice := slice[1:3]
newSlice = append(newSlice, 60) // 10, 20, 30, 60, 50

// append 세번째 인자: 용량을 제한하여 슬라이스 조작을 안전하게 하도록 한다.
source := []int{1, 2, 3, 4, 5}
slice := source[2:3:4] // 인덱스 2부터, 길이 1, 최대 용량 2
slice := source[2:3:3] // 두번째 인자와 같은 값으로 지정하면 내부 배열을 복사하므로 데이터 가공이 안전해짐
fmt.Printf("%v\n", append(s1, s2...)) // s1과 s2는 두 슬라이스로 합칠 수 있음

// range: value는 '복사본'을 생성한다.
slice := []int{10, 20, 30, 40, 50}

for index, value := range slice {
  fmt.Printf("인덱스: %d 값: %d\n", index, value)
}

// for 루프
for index := 2; index < len(slice); index++ {
  fmt.Printf("인덱스: %d 값: %d\n", index, slice[index])
}

// 다차원 슬라이스
slice := [][]int{{10}, {100, 200}}
slice[0] = append(slice[0], 20)
```

- 슬라이스는 포인터, 길이, 용량 필드만이 있으므로 함수 등을 통해 전달할 때, 24바이트만이 사용되므로 가볍다.

#### 맵

`key와 value로 구성되고 정렬이 없는 데이터 구조`

1. 생성 및 초기화

```go
// make를 이용
dict := make(map[string]int) // key: string, value: int

// 리터럴을 이용
dict := map[string][int]{"a": 23, "b": 24}

// key에는 슬라이스, 함수, 슬라이스를 가진 구조체 등이 될 수 없다.
```

2. 맵의 활용

```go
// 대입
colors := map[string]string{}
colors["Red"] = "#da1337"

// nil 맵: 맵에 새로운 키/값을 추가할 수 없음
var colors map[string]string

// 키/값이 존재하는지 확인. exists
value, exists := colors["Blue"]

// 키가 존재하지 않아도 항상 값이 리턴된다. 값 타입의 제로값이 리턴됨

// range, map에 사용될 경우, 키와 값이 리턴됨
colors := map[string]string{
  "Red": "#123sfa",
  "Yellow": "#ds2141",
}

for key, value := range colors {
  fmt.Printf("키: %s 값: %s\n", key, value)
}

// 키 삭제
delete(colors, "Red")

// 맵은 함수에 전달될 때, 복사본이 전달되지 않으므로 함수 내에서 변경이 있으면 원 맵에도 영향이 미친다.
```
