# 라이브러리 개발

## 1. 패키지 구성

#### 1.1 코드 작성

github를 기본 경로 패키지의 이름은 fibonacci다.

```
# folder 생성 및 이동
$ mkdir -p $GOPATH/src/github.com/socialpercon/fibonacci
$ cd $GOPATH/src/github.com/socialpercon/fibonacci

# fibonacci.go 코드 생성
$ vi fibonacci.go
package fibonacci

func Fibonacci(n int) chan int {
        c := make(chan int)
        go func() {
                a, b := 0, 1
                for i := 0; i < n; i++ {
                        a, b = b, a+b
                        c <- a
                }
                close(c)
        }()
        return c
}
```

#### 1.2 코드 빌드


```
# build를 해봤다.
$ go build
$ echo $?
0
```



#### 1.3 패키지 생성

빌드는 성공적으로 끝낸 것 같은데, 아무런 파일도 만들어지지 않는다. 이런 패키지의 경우에는 반드시 install 옵션으로 컴파일해야 한다. install 로 컴파일 하면, 작업공간의 "pkg" 디렉토리에 object file이 만들어진다.

```
$ go install
$ cd $GOPATH/pkg/linux_amd64/github.com/socialpercon
$ ls
fibonacci.a
```

## 2. custom 패키지 사용

#### 2.1 import

```
$ mkdir -p $GOPATH/src/github.com/socialpercon/hello
$ cd $GOPATH/src/github.com/socialpercon/hello
# fibonacci 패키지를 사용하도록 hello.go 코드를 수정했다.
$ vi hello_fibonacci.go
package main

import (
        "fmt"
        "github.com/socialpercon/fibonacci"
)

func main() {
        fmt.Println("Hello World. Fibonacci(10)=")
        for x := range fibonacci.Fibonacci(10) {
                fmt.Println(x)
        }
}
```

#### 2.2 compile

install로 컴파일 하면 $GOLANG/bin 밑에 실행파일이 만들어진다.



```
$ cd $GOPATH/src/github.com/socialpercon/hello
$ go install
$ cd $GOPATH/bin
$ ./hello
```
#### 2.3 생성된 파일들의 경로

bin 디렉토리 밑에 install한 hello 실행파일이 보인다. pkg 밑에 fibonacci.a 파일이 만들어진 걸 확인할 수 있다. pkg디렉토리 구성을 보면 linux_amd64가 있는데, 이 부분은 크로스 컴파일시 사용한다. 디렉토리 이름은 시스템의 운영체제와 아키텍처를 반영한다.

```
$ tree
.
├── bin
│   └── hello 
├── pkg                                    
│   └── linux_amd64
│          └── socialpercon
│              └── fibonacci.a 
└── src 
       └── socialpercon
           ├── fibonacci
           │   └── fibonacci.go
           └── hello
               └── hello_fibonacci.go
```

#### 2.4 실행

```
Hello World. Fibonacci(10)=
1
1
2
3
5
8
13
21
34
55
```
