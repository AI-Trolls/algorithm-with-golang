# algorithm-with-golang
golang으로 알고리즘 연습하기
스크립트 언어만 하다보니 불안
컴파일 언어 하나쯤은 제대로 해보고 싶은데 c 쪽은 너무 식상

## golang
- static typed(type 미리 선언), strong typed(implicit 형변환 x) --> 모두 컴파일에러로 이어짐
- compile language --> 실행파일 바로 떨어짐
- garbage collection --> 컬렉터 내장되어 메모리관리 자동, 생산성 높일 수
- concurrency --> go 키워드를 통한 여러개 함수 실행(goroutine) 및 채널을 통한 통신 가능 vs. 쓰레드(커널의 도움 받아 생성 부담스러움)
- multicore env
- module, package --> 자체 모듈화 제공, 다른 사람 소스 그대로 가져와 바로 사용 가능(go get, go install), 의존성 관리 쉬움
- fast compile speed
- docker도 이거로 만들어짐 --> 웹브라우저, 서버, 디비 등 큰 규모의 복잡한 애플리케이션 개발에 적합 (메모리 관리 철저해야 하는 것에는 부적합)

