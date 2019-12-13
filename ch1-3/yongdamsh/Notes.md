### Packaging 방식
- 하위 디렉토리에 디렉토리명과 동일한 package name을 정의하면 해당 파일들을 합쳐서 사용할 수 있다.
- https://github.com/goinaction/code/tree/master/chapter2/sample 참고
- 

### 문법 특징
- Multiple return
  - 통상적으로 value, error 순으로 반환한다.
- 연쇄 if 조건절
- 맨 앞문자 대소문자 여부로 public/private scope을 구분한다.
- Reference type은 `make`로 생성한다.
  - map
  - slice
  - channel
- 모든 변수는 명시적으로 값을 할당하지 않으면 zero value를 기본 값으로 가진다.
  - string: `''`
  - number: `0`
  - boolean: `false`
  - pointer, map: `nil`
  - struct: 각 필드가 zero value로 설정됨 ex) { num: 0, str: '' }
- `:=`는 선언과 초기화를 동시에 수행
- 인터페이스 정의 규칙
  - 하나의 메서드만을 정의하고 있다면 `-er` 접미사
  - 

### 타입 시스템
타입을 모델링하는 것이 아니라 동작을 모델링할 수 있는 인터페이스

- empty struct: 메모리 0바이트 할당. 타입은 필요하지만 상태 관리할 필요가 없는 경우 유용함.
- value receiver



### 어플리케이션 구동의 필수 요소
- main package
- `main` function
  - 이 함수가 리턴되면 프로그램이 종료된다.
- (optional) `init` function
  - `main`보다 먼저 호출된다.
  - `init` function 사용 예:
    - 어떤 모듈의 메서드나 변수, 타입 등을 직접적으로 사용하진 않지만 초기화가 필요한 경우, 아래와 같이 blank identifier로 선언한다.
      ```go
      import (
        _ "path/to/package"
      )
      ```

### 추가 학습 필요
- WaitGroup, counting semaphore?
- interface{}, reflection (타입의 메타데이터를 읽어 런타임에 코드를 이용해 타입을 조작하는 방법)