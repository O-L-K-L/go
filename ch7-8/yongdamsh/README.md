# Chapter 8. 실무 패턴 (Discovery Go)

예제 코드: https://github.com/jaeyeom/gogo#8%EC%9E%A5


## 오버로딩

같은 이름의 함수 및 메서드를 여러 개 둘 수 있는 기능 (Go에서는 지원되지 않음)

오버로딩을 대체할 수 있는 방법
- 자료형에 따라 다른 함수명 붙이기 
  - ex) volumeCube, volumeCylinder, volumeCuboid
- 동일한 자료형의 자료 개수에 따른 오버로딩은 가변 인자를 사용할 수 있다.
- 다양한 인자를 오버로딩하는 경우 구조체를 넘기는 것이 더 나을 수 있다.
  ```go
  type Option struct {
    Idx int
    Lang Language
    ExcludeEmpty bool
  }

  func GetElement(opt Option) *Element {
    // ...
  }

  func main() {
    GetElement(Option{ Idx: 3 })
  }
  ```
- 인터페이스를 활용하는게 나을 수도 있다.
  ```go
  type Stringer interface {
    String() string
  }

  type Int int
  type Double float64

  func (i int) String() string { ... }
  func (d Double) String() string { ... }

  func ToString(s Stringer) string {
    return s.String()
  }
  ```
- 연산자 오버로딩
새로운 문제를 풀기보단 편의성을 위한 기능. 인터페이스를 이용해 해결할 수 있다.
ex) `sort.Interface`의 `Less`는 `<`를 오버로딩하기 위한 것


## 템플릿 및 제네릭 프로그래밍

제네릭은 알고리즘을 표현하면서 자료형을 배제할 수 있는 프로그래밍 패러다임이다. 
Go는 제네릭을 지원하지 않는데 어떻게 같은 문제를 풀 수 있는지 알아본다.

1. 유닛 테스트
2. 컨테이너 알고리즘
3. 자료형 메타 데이터
4. go generate


## 객체지향

### 1.다형성 

메서드가 호출되었을 때 어떤 자료형이냐에 따라 다른 구현을 할 수 있게 하는 것
인터페이스로 쉽게 구현 가능하다.

```go
type Shape interface {
  Area() float32
}

type Square struct {
  Size float32
}

func (s Square) Area() float32 {
  return s.Size * s.Size
}

type Rectangle struct {
  Width, Height float32
}

func (r Rectangle) Area() float32 {
  return r.Width * r.Height
}

func TotalArea(shapes []Shape) float32 {
  var total float32

  for _, shape := range Shape {
    total += shape.Area()
  }

  return total
}

func ExampleTotalArea() {
  fmt.Println(TotalArea([]Shape{
    Square{3},
    Rectangle{4, 5},
  }))
}
```

### 2. 인터페이스

자바 등 다른 언어와의 차이점? 인터페이스 내의 메서드들을 구현하기만 하면 그 인터페이스를 구현하는 것이 된다.
위의 다형성 예제 확장을 위해 외부 라이브러리의 Triangle을 가져온다고 생각해 보자.

```go
type Triangle struct {
  Width, Height float32
}

func (t Triangle) Area() float32 {
  return 0.5 * t.Width * t.Height
}
```

여타 객체 지향 언어라면 implements 키워드를 통해 새로운 상속 클래스를 만들거나 Triangle 코드를 수정해야 할 수도 있다.
Go에서는 이 코드를 가져오는 것만으로 Shape 인터페이스를 구현한 것이 된다. 이 점이 프로그램의 유연한 확장 및 재사용을 가능하게 한다.

### 3. 상속

상속이 풀려는 문제

#### 메서드 추가

Rectangle이 외부 라이브러리에 있다고 가정한다면, 구조체를 내장해 필드 참조나 메서드 호출을 위한 불필요한 코드를 줄일 수 있다.

```go
type RectangleCircum struct {
  Rectangle
}

func (r RectangleCircum) Circum() float32 {
  return 2 * (r.Width + r.Height)
}
```

#### 오버라이딩, 서브타입

기존의 구현을 다른 구현으로 대체하고자 하는 경우에도 구조체 내장으로 해결이 가능하다.

```go
type WrongRectangle struct {
  Rectangle
}

func (r WrongRectangle) Area() float32 {
  return r.Rectangle.Area() * 2
}

func ExampleTotalArea() {
  fmt.Println(TotalArea([]Shape{
    Square{3},
    Rectangle{4, 5},
    Wrongrectangle{4, 5},
  }))
}
```



### 4. 캡슐화

객체 안에 있는 정보를 바깥에 숨기고자 하는 것
대소문자 구분으로 public/private 지원

내가 만든 다른 패키지에서는 접근이 가능하게 하고 싶지만 남은 접근하지 못하게 하고 싶은 경우:
패키지 경로에 internal을 넣으면 internal이 있는 경로에 있는 패키지를 포함한 범위에서만 참조가 가능하다.

Internal Package Proposal: https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU

Go에서는 getter를 생략하고, setter는 선언한다.
ex) `Len()`, `SetLen(x)`

`NewX()`와 같은 형태의 생성자를 통해 내부 정보를 제한적으로 제공할 수 있다.


## 디자인 패턴

### 1. 반복자 패턴

