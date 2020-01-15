# CH5 Go의 타입 시스템

Go는 정적 타입 시스템을 사용한다. 정적 타입 시스템의 장점은 컴파일러가 이미 타입을 알고있어 컴파일 타임에서 내재적인 문제점을 걸러낼 수 있다는 점과 메모리 할당을 효율적으로 할 수 있다는 점 등이 있다.

## 사용자정의 타입

1. 구조체(struct)를 이용한 합성

```go
type user struct {
  name string
  email string
  ext int
  priviliged bool
}

// var를 붙이면 제로값으로 초기화됨. 구조체의 모든 값이 제로값으로 초기화된다.
var bill user

// 특정값을 이용한 초기화
lisa := user {
  name: "kevin.koo"
  email: "kevin@myrealtrip.com"
  ext: 123
  priviliged: true
}

// 구조체 정의에 나열된 순서대로 초기화
lisa := user {"kevin.koo", "kevin@naver.com", 123, true}

// 이미 정의된 타입을 이용한 새로운 구조체 선언
type admin {
  person user
  level string
}

// 초기화
fred: admin {
  person: user {
    name: "Lisa",
    email: "List@email.com",
    ext: 123,
    priviliged: true,
  }
  level: "admin",
}
```

2. 기존에 존재하는 타입을 기반으로 새로운 타입을 생섯

```go
// type {새로운 타입} {기반타입: base type}
// 컴파일러는 두 개의 타입이 호환가능해도 혼용을 허용하지 않고 별개의 타입으로 인식한다. 따라서 묵시적 변환이 이루어지지 않는다.
type Duration int64
```

## 메서드

타입에 행위를 정의하는 방법.
func 키워드와 함수 이름 사이에 리시버를 넣어서 정의한다.
=> `리시버`가 정의된 함수를 `메서드`라고 한다.

```go
type user struct {
  name string
  email string
}

// 리시버는 user 타입으로 설정되어 항상 값을 복사해서 바인딩한다.
func (u user) notify() {
	fmt.Printf("사용자에게 메일을 전송합니다: %s<%s>\n", u.name, u.email)
}

// user 참조형으로 리시버가 정의되어 주소값을 복사해서 바인딩함.
func (u *user) changeEmail(email string) {
	u.email = email
}

// 실제 값을 bill 변후에 저장
bill := user{"Bill", "bill@email.com"}
bill.notify()
bill.changeEmail("bill@newDomail.com") // Go 컴파일러가 적절히 호출하기 위해 내부적으로 (&bill).changeEmail()로 변환한다.
bill.notify()

lisa := &user{"Lisa", "lisa@email.com"}
lisa.notify() // 참조 타입의 lisa 변수에서 notify를 호출하기 위해 Go 컴파일러가 내부적으로 (*lisa).notify()로 변환한다.
lisa.changeEmail("lisa@newDomail.com")
lisa.notify()
```

## 타입의 본질

타입의 메서드를 정의할 때는 새로운 값이어야 하는지 기존의 값이 변경되어야 하는지 잘 판단해야 한다. 새로운 값이라면 리시버의 타입이 값이어야 하고 기존 값의 변경이라면 포인터값이어야 한다.

### 1. 내장 타입

숫자, 문자열, 불리언 타입 등이 있고 primitive한 성질이 있고 함수나 메서드에 전달하면 참조값이 아닌 복사본이 전달된다.

### 2. 참조 타입

슬라이스, 맵, 채널, 인터페이스, 함수 타입 등이 있다.
헤더값이 생성되고 이는 기반 데이터 구조에 대한 포인터를 이미 갖고 있다. 따라서 참조 값 자체를 전달하지 않고 헤더값의 복사본을 전달하면 참조 값을 공유할 수 있다.

```go
type IP []byte
// Go는 사용자정의 타입에 대해서만 메서드를 정의할 수 있으므로 기반 타입에 새로운 동작을 정의하기 위해서는 기반 타입을 통해 새로운 타입을 정의해서 추가한다.

// 참조 타입의 값을 공유할 것이 아니므로 값 수신자를 이용해 선언
func (ip IP) MarshalText() ([]byte, error) {
  if len(ip) == 0 {
    return []byte(""), nil
  }
  if len(ip) != IPv4len && len(ip) != IPv6len {
    return nil, errors.New("invalid IP address")
  }

  return []byte(ip.String()), nil
}
```

### 3. 구조체 타입

기본형과 비기본형 성질을 모두 가지는 데이터 값을 표현할 수 있음.

```go
// 1. 기본형의 성질(값을 복사하여 사용)을 이용하는 경우
type Time struct {
  sec int64
  nsec int32
  loc *Location
}

func Now() Time {
  sec, nsec := now();
  return Time{sec + unixToInternal, nsec, Local} // 새로 값을 만들어 리턴
}

// 값 수신자로 되어있어 호출자의 Time 값을 복사해 가공하고 다시 리턴한다.
func (t Time) Add(d Duration) Time {
  t.sec += int64(d / 1e9);
  nsec := int32(t.nsec) + int32(d%1e9)

  if nsec >= 1e9 {
    t.sec++
    nsec -= 1e9
  } else if nsec < 0 {
    t.sec--
    nsec += 1e9
  }
  t.nsec = nsec
  return t
}

// 2. 비기본형 성질(원래 값을 변경)하는 경우
// 파일 종료 등의 기능을 수행하기 위해 file은 값을 포함해야 하는데 값 자체가 공유되기에는 위험하므로 숨겨진 타입에 대한 포인터를 포함하도록 정의되어있다.
// file은 *File을 나타냄?
type File struct {
  *file
}

type file struct {
  fd int
  name string
  dirinfo *dirinfo
  nepipe int32
}

func Open(name string) (file *File, err error) {
  return OpenFile(name, O_RDONLY, 0)
}

func (f *File) Chdir() error {
  if f == nil {
    return ErrInvalid
  }
  if e := syscall.Fchdir(f.fd); e != nil {
    return &PathError("chdir", f.name, e)
  }
  return nil
}
```

메서드 정의 시, 값 리시버와 포인터 리시버를 결정하는데 변경 여부가 아니라 값의 본질에 따라 선택되어야 한다. 예외적인 케이스는 인터페이스 값이 다를 때, 값 타입 리시버에서 유연성을 제공해야 할 때 뿐이다. 이런 경우엔 타입의 본질이 비기본형이라도 값 리시버를 사용하기도 한다.

## 인터페이스

인터페이스는 정의된 행동을 외부에 노출할 수 있도록 하는 기능이자 그걸 기반으로 통신할 수 있도록 하는 기능을 한다.

### 인터페이스의 구현 기법

사용자 정의 타입이 인터페이스의 메서드를 구현한다면 인터페이스 타입에 대입하여 사용할 수 있고 실제 실행은 정의된 메서드로 실행되고 이것이 다형성이다.
인터페이스의 값은 `iTable`과 저장된 사용자 정의 객체의 주소값이 저장되어 있는데 `iTable`은 저장된 값(사용자 정의 객체)의 타입과 메서드의 목록을 가지고 있다.

### 메서드 집합

메서드 집합은 인터페이스를 준수하는 규칙을 정의한다.
인터페이스에 대입되는 값이 `포인터`면 메서드 수신자는 `포인터`만이 될 수 있고 `값`이면 `포인터` or `값`이 될 수 있다. 즉, 메서드 수신자의 입장에서는 `포인터`일 때, `포인터`만, `값`일 때는 `포인터` or `값`만이 올 수 있다. 그 이유는 메서드 수신자의 리시버가 `포인터`일 때, `값`이 온다면 주소값을 알 수 없는 경우(ex. 리터럴 상수)가 있기 때문이다.

### 다형성

대입된 인터페이스의 메소드를 호출하면 메소드에서 사용되는 인터페이스 값에 저장된 `iTable`의 메서드 집합의 메서드를 쓰고 실제 값은 객체의 값(ex. user)의 값을 이용하여 다형성을 구현한다.

### 타입 임베딩

기존에 선언된 타입을 새로운 구조체 타입의 내부에 선언하는 방법이다. 내부 타입에 선언된 식별자는 `승격`하여 외부 타입에서 직접 호출이 가능하다. 타입 임베딩을 하려면 타입 이름과 필드 이름을 동일하게 작성하면 된다.
내부 타입이 구현한 인터페이스도 마찬가지로 승격되어 외부 타입에서 인터페이스 메소드를 호출할 수 있다.
cf. 외부 타입이 메서드나 인터페이스를 직접 구현했다면 내부 타입의 메서드나 인터페이스는 승격하지 않는다.

## 외부 노출 식별자와 비노출 식별자

모든 식별자를 패키지로 접근이 가능하게 하면 안 된다. 대문자로 시작하면 외부에 노출되고 소문자로 시작하면 노출되지 않는다.
cf. 어떤 타입이 비노출 타입의 필드를 갖고 있는데 그 필드의 타입이 노출 타입인 경우, 내부 타입의 필드들은 승격된다. 따라서 내부 타입으로의 직접 접근은 허용되지 않지만(비노출) 외부 타입에서 내부 타입의 필드로 접근하는 것은 허용된다(내부 타입의 승격).
