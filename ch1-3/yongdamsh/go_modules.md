# Learn Go Modules
Go v1.11 부터 도입된 의존성 관리 도구


### Initialize

패키지 root 경로에서 아래 명령어를 실행한다.

```sh
go mod init path/to/package
```

위 명령어를 실행하면,
- [go.mod](wc/go.mod) 파일이 자동 생성되고, 모듈명, go 버전, 의존성이 명시되어 있다.
- 버전 무결성 체크를 위한 [go.sum](wc/go.sum) 파일도 생성된다.


### Upgrade

아래 명령어로 해당 패키지의 tagged version 목록을 조회할 수 있다.

```sh
go list -m -versions path/to/package
```

실행 결과:
```
path/to/package v1.0.0 v1.1.0
```

특정 버전 설치:
```
go get path/to/package@v1.1.0
```

Major 버전은 다른 convention을 따른다.
- import 경로가 분리된다.
- 호환되지 않는 함수는 버전 접미사를 붙인다. ex) `func HelloV2() string`

```go
import (
  "path/to/package"
  pkgV2 "path/to/package/v2"
)

func Hello() string {
  return package.Hello()
}

func NewHello() string {
  return pkgV2.HelloV2()
}
```

버전 분기가 필요 없어지는 시점엔 패키지 alias를 제거해도 된다.
```go
import (
  "path/to/package/v2"
)

func NewHello() string {
  return package.HelloV2()
}
```

아래 명령어로 쓰이지 않는 의존성을 제거할 수 있다.

```sh
go mod tidy
```

### Release

태그를 추가해 릴리즈할 수 있다. 버전은 [semantic version](https://semver.org/) rule을 따른다.

```sh
git add ...
git commit ...

git tag v0.1.0
git push origin v0.1.0
```

### References
- Guide Part 1: https://blog.golang.org/using-go-modules
- Semantic Import Versioning: https://research.swtch.com/vgo-import