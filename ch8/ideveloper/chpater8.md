## 표준 라이브러리

- Go의 표준 라이브러리는 언어자체를 향상시키고 확장할 수 있는 핵심 패키지의 집합

- go 커뮤니티의 개발자들은 다른 언어 개발자들에 비해, 표준 라이브러리 의존도가 높은 편
  - 표준 라이브러리의 디자인이 훌륭하며 전통적인 라이브러리에 비해 훨씬 많은 기능을 제공하기 떄문

### 문서화와 소스 코드

- https://golang.org/pkg/

  - 각 패키지의 godoc 문서 제공

- 인터랙티브한 문서 => Sourcegraph

  - https://sourcegraph.com/github.com/sourcegraph/go-sourcegraph

- GOROOT/pkg 안에 이미 go를 설치한 운영체제에서 동작하는 플랫폼별로 폴더에 나뉘어 컴파일된 아카이브 파일이 존재한다.
  - 따라서, go의 빌드과정이 빠르게 됨

## 로깅

- go의 표준 라이브러리는 약간의 설정만으로 편리하게 활용할 수 있는 log 패키지를 제공한다. 사용자가 정의한 로거를 작성해 필요한 로깅 기능을 물론 구현할 수도 있다.
- UNIX 아키텍트는 stderr라는 장치를 추가했고, 이장치는 로그를 위한 기본 출력 장치로 사용되기 위해 만들어졌다. 또 이 장치를 이용해 개발자들은 프로그램의 출력과 로그를 분리할 수 있었다.
- 그러나 프로그램이 로그만 기록하고 있다면, 로그 정보는 `stdout 장치`에, 오류와 경고는 `stderr`에

![image](https://img.velog.io/post-images/jakeseo_me/ecf11ca0-6d70-11e9-8ea3-211446efebf3/stdinstdoutstderr.png?w=1024)

https://velog.io/@jakeseo_me/%EC%9C%A0%EB%8B%89%EC%8A%A4%EC%9D%98-stdin-stdout-stderr-%EA%B7%B8%EB%A6%AC%EA%B3%A0-pipes%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90 참고

여기서 말하는 stdin, stdout, stderr은 stream이라고 불리는 것들입니다.

- stdin은 받은 입력 값을 프로그램에 나타내주는 stream입니다. (예를 들면, 비밀번호를 입력하기 위한 프롬프트 같은 것들이 있습니다.)
- stdout은 모든 출력값들이 가는 곳입니다. C로 프로그래밍을 할 때, printf를 생각해보시거나 Java로 프로그래밍을 할 때는 System.out.println 파이썬으로 프로그래밍할 때는 print를 생각하면 됩니다.
- stderr은 또 다른 출력 채널입니다. 주로 디버깅 정보를 출력하거나 에러를 출력하는데에 쓰입니다.

### log 패키지

- 로깅의 목적 : 프로그램이 현재 어떤 동작을 수행하고 있는지, 그리고 언제 어떤일이 발생했는지를 추적하는 것
- flag들에 대한 설명: https://golang.org/pkg/log/#pkg-constants

```go

// This sample program demonstrates how to use the base log package.
package main

import (
	"log"
)

func init() {
  log.SetPrefix("TRACE: ")
  //  비트 OR을 이용해 출력될 항목을 조정
  // 이 플래그들이 출력되는 순서는 조정할수는 없다.
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger.
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	log.Panicln("panic message")
}
```

#### Ldate 상수, iota

```go
Ldate = 1 << iota
```

- `iota 키워드`는 상수 블록을 선언할 때 사용하는 특별한 키워드
  - 이 키워드는 코드 블록의 끝에 도달하거나 대입 구문이 발견될 때까지 동일한 표현식을 매 상수마다 중복해서 적용할 것을 컴파일러에 지시
  - iota의 기본 값을 0으로 하되, 상수를 정의할 때마다 1씩 증가시키는 것

ex) log flag

```go
const {
  Ldate = 1 << iota // 1 << 0 = 00000001 = 1
  Ltime // 1 << 1 00000010 = 2
  Lmicroseconds // 1 << 2 00000100 = 4
}
```

- 위 예제에서는 << 연산자가 왼쪽 값의 비트를 1비트씩 옮기게 되고, 각 상수가 독자적인 비트위치를 갖게 됨.

```go
func init(){
  ...
  log.setFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
```

- 같은 맥락으로, 각각의 독자적인 flag값들이 비트 OR 연산을 통해 독자적인 flag가 만들어짐 => 우리가 표현하고자 하는 모든 로그 옵션을 표현하는 하나의 값 생성 가능
  - 위에서는 00001101 , 13이라는 값으로 표현

```go
func main() {
	// 표준 로거에 메시지를 출력
	log.Println("message")

	// Println() 함수를 실행한 후 os.Exit(1)을 추가로 호출하여 프로그램 종료
	log.Fatalln("fatal message")

	// Panicln 함수는 Println() 함수를 호출한 후 panic() 함수를 추가로 호출하여 패닉을 발생시켜 이 상황을 복구하지 못하면, 프로그램 및 스택 추적 종료
	log.Panicln("panic message")
}
```

- panic

https://www.joinc.co.kr/w/GoLang/example/panic

### 사용자정의 로거

- 사용자정의 로거를 구현하려면 Logger 타입의 값을 직접 생성해야 함.
- log.New
  - 다른 타입의 로거를 생성하기 위해 호출
  - 로그 데이터가 기록될곳, prefix, flag 값들을 매개변수로 받음
  - https://golang.org/src/log/log.go?s=2897:2953#L52

```go
// This sample program demonstrates how to create customized loggers.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

func init() {
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}

```

`ioutil.Discard`

- https://golang.org/pkg/io/ioutil/#pkg-variables
- Discard 변수를 이용하면, 해당 수준의 로그가 필요하지 않을 때 그 로그 수준을 비활성화 할수있다.

`io.MultiWriter`

- io.Writer 인터페이스를 구현하는 값이라면 몇 개라도 전달할수 있는 함수
- io.MultiWriter(file, os.Stderr)
  - 위 예에선 Error 로거를 이용해 로그를 출력하면 그 메시지는 파일 및 stderr 장치에 모두 출력

시간되면 다른 라이브러리 로깅 살펴보기..? github cli..?

## 인코딩/디코딩

### JSON 데이터 디코딩

- 파일이나 웹 응답을 통해, JSON 데이터를 소비하는 경우라면 NewDecoder 함수와 Decode 메소드를 반드시 사용해야 함.

https://github.com/goinaction/code/blob/master/chapter8/listing24/listing24.go

`태그`

- 아래 구조체들의 선언을 살펴보면, 각 필드 선언의 마지막에 문자열이 덧붙여져있는데 이 문자열들을 `태그`라고함.
- json문서와 구조체 타입 간의 필드 매핑을 위한 메타데이터를 제공
- 태그가 지정되지 않으면 디코딩 및 인코딩이 처리될 때 대소문자 구분 없이 필드의 이름에 해당하는 항목이 JSON 문서에 존재하는지 찾음, 없으면 제로값으로 초기화

```go
type (
	// gResult maps to the result document received from the search.
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// gResponse contains the top level document.
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)
```

`Decode 메소드`

```go
var gr *gResponse
err = json.NewDecoder(resp.Body).Decode(&gr)
```

`JSON문서에 문자열 값이 포함되어 있는 경우`

- 문자열을 바이트 슬라이스로 변환한 후 json 패키지의 Unmarshal 함수 사용
- https://github.com/goinaction/code/blob/master/chapter8/listing27/listing27.go
- 위 링크 결과 값 {Gopher programmer {415.333.3333 415.555.5555}}

`JSON문서를 맵으로 디코딩하거나 언마샬링`

- https://github.com/goinaction/code/blob/master/chapter8/listing29/listing29.go
- https://play.golang.org/p/Vpat_HQZ_19

### JSON 데이터 인코딩

- json 패키지의 MarshalIndent 함수, Marshal함수
  - 들여쓰기가 잘 된 JSON 문자열을 만들고 싶으면 MarshalIndent
- Go의 맵이나 구조체 타입의 값으로부터 JSON 문서를 도출할 때 매우 편리한 함수
- https://play.golang.org/p/HSOZ95Xvnb7

## 입력과 출력

### Writer 인터페이스와 Reader 인터페이스

- http://www.codingnuri.com/golang-book/13.html 13.2 입출력 부분
- bytes, fmt, os 패키지에서 Buffer와 문자열 결합 이용하여 stdout 장치에 문자열 출력
- bytes.Buffer와 os.File 타입이 모두 io.Writer 인터페이스를 구현하고 있음

`fmt`

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

`buffer`

```go

// io.Writer 인터페이스
// type Writer interface {
//  Write(p []byte) (n int, err error)
// }
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	m, ok := b.tryGrowByReslice(len(p))
	if !ok {
		m = b.grow(len(p))
	}
	return copy(b.buf[m:], p), nil
}

func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.lastRead = opInvalid
	if nBytes := b.Len(); nBytes > 0 {
		m, e := w.Write(b.buf[b.off:])
		if m > nBytes {
			panic("bytes.Buffer.WriteTo: invalid Write count")
		}
		b.off += m
		n = int64(m)
		if e != nil {
			return n, e
		}
		// all bytes should have been written, by definition of
		// Write method in io.Writer
		if m != nBytes {
			return n, io.ErrShortWrite
		}
	}
	// Buffer is now empty; reset.
	b.Reset()
	return n, nil
}

```

`os`

```go
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout") => file 타입에 대한 포인터 리턴
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)


func (f *File) Write(b []byte) (n int, err error) {
	if err := f.checkValid("write"); err != nil {
		return 0, err
	}
	n, e := f.write(b)
	if n < 0 {
		n = 0
	}
	if n != len(b) {
		err = io.ErrShortWrite
	}

	epipecheck(f, e)

	if e != nil {
		err = f.wrapErr("write", e)
	}

	return n, err
}
```

- https://play.golang.org/p/SK9fMHRbaKl
