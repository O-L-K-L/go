# Chapter 5. Go's Type System

## Custom Type

### Declaration
- 내장 타입을 기반으로 원하는 타입을 정의할 수 있다.
- 대표적인 예로 다양한 타입을 합성(composition)하는 struct가 있다.
  ```go
  type user struct {
    name string
    age int
  }
  ```
- Primitive 타입도 사용자 정의 타입으로 선언할 수 있다. 이 때 타입 변환은 암묵적으로 이뤄지지 않는다.
  ```go
  type Duration int64

  func main() {
    var d Duration

    d = int64(10) // cannot use int64(10) (type int64) as type Duration in assignment
  }
  ```

### Method
- 사용자 정의 타입에 한해 행위(behavior)를 정의할 수 있다.
- 메서드는 수신자(receiver)가 정의된 함수를 말한다. 메서드를 통해 타입의 상태와 행위 간의 응집도를 높일 수 있다.
  ```go
  type user struct {
    name string
  }

  func (u user) introduce() {
    fmt.Printf("My name is %s\n", u.name)
  }
  ```
- 수신자는 value receiver와 pointer receiver로 구분된다. 
  - Value receiver는 해당 타입의 값이 복사되기 때문에 조회 용도로 사용된다. 또는 편집된 복사본을 반환하는 immutable 함수 구현으로 사용할 수 있다.
  - Pointer receiver 해당 타입의 주소에 접근하기 때문에 값을 변경하는 용도로 사용된다.
  - 메서드를 정의할 때 위의 용도에 적합한 방식의 receiver를 사용한다.
  ```go
  type user struct {
    name string
  }

  // Value receiver
  func (u user) introduce() {
    fmt.Printf("My name is %s\n", u.name)
  }

  // Pointer receiver
  func (u *user) changeName(newName string) {
    u.name = newName
    fmt.Printf("My name changed to %s\n", u.name)
  }
  ```
  - 근데 '5.3 타입의 본질' section 후반부에서는 `메서드가 수신된 값을 변경하는지는 전혀 관련이 없다. 값의 본질에 따라 결정해야 한다.` 라고 한다. 
  - 그럼 '값의 본질'이란 뭘 뜻하는가? 책에서는 [time](https://golang.org/pkg/time/), [os](https://golang.org/pkg/os/) 표준 라이브러리의 일부분을 예를 들어 비교하고 있다.


## Interface