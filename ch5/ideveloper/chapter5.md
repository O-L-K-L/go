## Go의 타입 시스템

- GO는 정적타입 프로그래밍언어
  - 이를 통해 잠재적인 메모리 문제와 버그를 줄일 수 있음은 물론 컴파일러가 더 나은 코드를 생성할 수 있는 기회를 가질 수 있게 됨
- 컴파일러는 값의 타입을 통해 두가지 정보를 얻을 수 있음
  - 첫번째: 값에 할당해야 하는 메모리의 크기
  - 두번째: 할당된 메모리를 통해 표현할 수 있는 값의 종류

#### 사용자정의 타입

- Go는 사용자가 직접 타입을 정의하는것을 허용함
- Go의 사용자정의 타입을 선언하는 방법

  - struct 키워드를 이용하여 합성 (가장 일반적)

  ```go
  type user struct {
    name  string
    email string
    ext int
    privileged bool
  }
  ```

  - 기존의 타입을 새로운 타입의 명세로 활용하는 방법

  ```go
  type Duration int64
  ```

**추가 정보**

- https://golang.org/pkg/time/

```go
type Duration int64
type Month int
```

- type Duration

A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately `290 years`.

#### 메서드

`설명`

- 사용자가 정의한 타입에 행위를 정의하기 위한 방법
- 실제로는 func 키워드와 함수 이름 사이에 추가 매개변수를 정의한 함수

`자세한 설명`

- func키워드와 함수 이름 사이의 매개변수는 **수신자(receiver)**라고 부르며, 함수를 특정 타입에 바인딩 하는 역할이고, 이렇게 수신자가 정의된 함수를 **메서드**라고 함

```go
func (u *user) changeEmail(email string){
  u.email = email
}
```

##### 수신자

```go
func (u user) notify() {

}

ideveloper := &user{"ideveloper","ideveloper@email.com"}
ideveloper.notify()

// 실제로 go가 수행하는 작업
(* ideveloper).notify()

```

`값 수신자`

- 호출 시점에 항상 그 값의 복사본을 대상으로 실행
- go 컴파일러는 포인터 값을 역참조하여 값 수신자에 정의된 메서드를 호출하도록 도움

```go
func (u *user) changeEmail(email string){
  u.email = email
}

ideveloper := user{"ideveloper","ideveloper@email.com"}
ideveloper.changeEmail("ideveloper@naver.com")

// 실제로 go가 수행하는 작업
(&ideveloper).changeEmail("ideveloper@naver.com")
```

`포인터 수신자`

- 호출 시점에 그 값의 실제값을 전달 받는다.
- go 컴파일러는 값을 참조하여 메서드 호출에 적합한 수신자 타입으로 변환해준다.

**추가정보**

- 값과 포인터 수신자 중 어느것을 사용해야 하는지 결정하는 것은 혼란스러울 수 있는데, 표준 라이브러리는 몇가지 기본 가이드라인을 제시하고 있다.
- https://golang.org/doc/faq#methods_on_values_or_pointers
  - 첫째: (method need to modify the receiver) receiver를 통해 받는 것을 변경시켜야 하는지 여부
  - 둘째: 효율성을 고려해야 할때 (receiver로 받는게 큰지, instance의 struct가 방대한지 등등)
  - 셋째: consistency 일관성 (예: 메소드 집합)
    - some of the methods of the type must have pointer receivers, the rest should too

#### 타입의 본질

##### 내장 타입

- 언어 차원에서 지원되는 타입
  - 숫자
  - 문자열
  - Boolean
- 이런 타입들은 primitive한 성질을 가지고 있어, 이 값들을 함수나 메서드에 전달하면 이 값들의 복사본이 전달된다.

##### 참조 타입

- 종류

  - 슬라이스
  - 맵
  - 채널
  - 인터페이스
  - 함수 타입

- 이런 타입의 변수를 선언하면 **헤더** 값이라고 불리는 값이 생성된다.
  - 헤더 값들을 기반 데이터 구조에 대한 포인터를 가지고 있다.

##### 구조체 타입

- 기본형 (primitive)
- 비기본형 (non-primitive)

```go
type File struct {
   *file
}

type file struct {
  fd int
  name string
  dirinfo *dirinfo
  nepipe int32
}
```

#### 인터페이스

- 다형성이란 타입을 작성할 때 `다양한 동작을 수행할 수 있는 코드 작성`을 가능하게 해주는 기법이다.

  - 125~127 이해가 잘안됨 (io 관련 다시 학습)

##### 인터페이스의 구현 기법

- 행위를 선언하기 위한 타입
- 이러한 행위는 인터페이스 타입이 직접 구현하지 않고, 사용자 정의 타입이 메서드 형태를 구현해야 한다.
  - 사용자 정의 타입을 종종 구현타입이라고 부르기도 함
- 인터페이스 값의 메서드에 대한 호출은 본질적으로 다형성을 갖는다.

`사용자정의 타입값에 저장된 인터페이스 타입의 모습`

![image](https://user-images.githubusercontent.com/26598542/72215274-708c7680-3554-11ea-9caf-af780af1939d.png)

![image](https://user-images.githubusercontent.com/26598542/72215275-71250d00-3554-11ea-92d5-8125d84170f7.png)

##### 메서드 집합

- 메서드 집합은 인터페이스를 준수하는 것과 관련된 규칙들을 정의한다.
- 메서드 집합은 주어진 타입 값이나 포인터와 관련된 메서드 집합을 정의한다.

메서드 집합규칙

| 값  |  메서드 수신자   |
| :-- | :--------------: |
| T   |      (t T)       |
| \*T | (t T) 와 (t \*T) |

| 메서드 수신자 |    값    |
| :------------ | :------: |
| (t T)         | T 와 \*T |
| (t \*T))      |   \*T    |

##### 다형성

- 매개변수에 전달된 구현 타입이 notify 메서드를 구현하고 있다면, 구현 타입의 종류와 관계없이 그 메서드를 호출 가능
- https://play.golang.org/p/mWYSQB4D2P9

```go

type notifier interface {
	notify()
}

func main() {
	// Create a user value and pass it to sendNotification.
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create an admin value and pass it to sendNotification.
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
```

#### 타입 임베딩

- go에서는 타입을 확장하거나 그 동작을 변경하는 것이 가능
  - 코드의 재사용은 물론 새로운 수요에 따라 기존 타입의 동작을 변경할 때 유용함
- 타입 임베딩은 기존 선언된 타입을 새로운 구조체 타입의 내부에 선언하는 것이다.
  - 이렇게 포함된 타입은 새로운 외부 타입의 내부 타입으로 활용

```go
type admin struct{
  user
  level string
}
```

- 내부 타입은 자신의 정체성을 잃어버리지 않으며 언제든지 직접 접근 가능

  - https://github.com/goinaction/code/blob/master/chapter5/listing60/listing60.go

#### 외부 노출 식별자와 비노출 식별자

- 식별자 노출에 대한 규칙은 좋은 API 디자인과 관련해 아주 중요한 부분
  - 외부에 노출되는 것 : 대문자
  - 외부에 노출되지 않는 것 : 소문자
