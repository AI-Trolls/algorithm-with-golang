# algorithm-with-golang
golang으로 알고리즘 연습하기  
스크립트 언어만 하다보니 불안  
컴파일 언어 하나쯤은 제대로 해보고 싶은데 c 쪽은 너무 식상  

## golang
- static typed(type 미리 선언), strong typed(implicit 형변환 x) 
  - 모두 컴파일에러로 이어짐
- compile language
  - 실행파일 바로 떨어짐
- garbage collection
  - 컬렉터 내장되어 메모리관리 자동, 생산성 높일 수
- concurrency
  - go 키워드를 통한 여러개 함수 실행(goroutine) 및 채널을 통한 통신 가능 
  - vs. 쓰레드(커널의 도움 받아 생성 부담스러움)
- multicore env
- module, package
  - 자체 모듈화 제공, 다른 사람 소스 그대로 가져와 바로 사용 가능(go get, go install), 의존성 관리 쉬움
- fast compile speed
- docker도 이거로 만들어짐
  - 웹브라우저, 서버, 디비 등 큰 규모의 복잡한 애플리케이션 개발에 적합 (메모리 관리 철저해야 하는 것에는 부적합)


## 설치
```
docker run -it --name go golang:latest /bin/bash 
```

## gopath ?
- 기준 디렉터리, workspace 위치를 정의 하는 것
  - workspace는 bin, pkg, src 디렉터리로 구성됨
    - [workspace 개념](https://github.com/golang-kr/golang-doc/wiki/Go-%EC%BD%94%EB%93%9C%EB%A5%BC-%EC%9E%91%EC%84%B1%ED%95%98%EB%8A%94-%EB%B0%A9%EB%B2%95#workspace)
```
  export GOPATH=위치
```
- 이후에 go env 명령어를 통해 GOPATH 확인 가능
- 소스를 받아오거나 패키지를 컴파일할 때 해당 경로 기준으로 진행됨
- 진행하는 프로젝트 위치로 매번 수정해주기


## 기본적인 실행법 및 코드 구조 [영어원본](https://golang.org/doc/code.html) 참조하세요!

## 기본 문법은 [여기](http://pyrasis.com/book/GoForTheReallyImpatient/Unit01)가 최고입니다.

## vim에서 개발 환경 구축
- https://www.joinc.co.kr/w/man/12/golang/Start


## 명령어 정리
```
go build <source> # 소스가 위치하는 디렉터리에 컴파일
go install <source> # 관련 패키지 모두 받아와 컴파일 후 bin 디렉터리에 실행 파일 생성
go get <주소> # 외부 소스 fetch + build + compile
gofmt -w 소스.go # 자동 정렬!!!
```

## 예시 소스들은 [여기](https://github.com/AI-Trolls/algorithm-with-golang/tree/master/basic/src/github.com/songjein)
