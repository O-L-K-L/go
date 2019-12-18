## GO와의 첫 만남 (챕터 1)

### GO 언어의 특징

- GO는 코드를 쉽게 공유할 수 있는 프로그래밍 언어이다.
- C나 C++ 같은 언어는 빠른 수행 속도, ruby나 python은 개발기간이 짧다. 이 둘 사이의 균형을 잘 맞추는 언어가 go이다.
- 빠른 컴파일 속도를 자랑한다.
- 내장된 동시성 기능 덕분에 별도의 스레드 라이브러리를 사용하지 않아도 시스템 자원을 효과적으로 활용가능하다.
- 객체 지향 개발 과정에서 발생하는 오버헤드를 줄이고 오로지 코드의 재사용에만 집중할 수 있도록 간결하면서 효과적인 타입 시스템을 제공하고 있다.
- 메모리를 직접관리할 필요가 없도록 가비지 컬렉션도 지원된다.


### 개발 속도
- Java, C 또는 C++ 컴파일러들이 전체 라이브러리의 의존성을 탐색하는 것과 달리 Go 컴파일러는 직접적으로 참조하는 라이브러리의 의존성만을 해석한다.
  
### 동시성
- 고루틴은 스레드와 유사하지만 더 적은 메모리를 소비하며 더 적은 양의 코드로 구현 할 수 있다.

#### 고루틴
- 여러개의 고루틴이 하나의 스레드에서 동작한다.
- 본연의 목적을 달성하기 위한 코드를 실행하는 동안 다른 코드를 동시에 실행하고자 한다면 고루틴을 사용!

#### 채널
- 고루틴 간에 안전한 데이터 전송을 가능하게 하도록 하는 데이터 구조이다.
- 동시성 프로그래밍에 있어 가장 어려운 부분은 고루틴에 의해 의도치 않게 데이터가 변경되는 일을 방지하는 것이다. 이러한 동시에 발생하는 수정요청으로 부터 데이터를 안전하게 보호하기 위한 패턴을 채널에서 적용해 해결하고 있다.

### 타입 시스템
- 합성이라고 불리는 디자인 패턴과 마찬가지로, 기능을 재사용하기 위해 타입을 임베드 한다.
- go는 타입을 모델링 하는 것이 아니라 동작을 모델링할 수 있는 독특한 인터페이스를 구현하고 있다. (덕 타이핑)

### 메모리 관리
- go는 프로그래밍을 지루하고 어려운 작업으로 만드는 여러가지를 걷어내고 프로그램 본연의 목적에 집중할 수 있도록 도움을 준다.