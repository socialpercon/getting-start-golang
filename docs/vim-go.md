# vim-go 환경 만들기


## 1. golang 설치

#### 1.1 다운로드 및 압축풀기

```
$ cd ${USR_LOCAL}
$ wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz
$ tar -zxvf go1.4.linux-amd64.tar.gz 
```

#### 1.2 환경변수 추가

```
$ mkdir -p ${WORKSPACE}/gopath
$ vi ~/.bashrc
$ export PATH=${USR_LOCAL}/go/bin:$PATH
$ export GOROOT=${USR_LOCAL}/go
$ export GOPATH=${WORKSPACE}/gopath
$ source ~/.bashrc
```

## 2. vim 패키지(플러그인) 관지자 설치

#### 2.1 pathogen 설치

```
# 다음 디렉토리가 없다면 생성
$ mkdir -p ~/.vim/autoload
$ mkdir -p ~/.vim/bundle

# vim-pathogen 을 ~/.vim/autoload 에 다운 받는다
$ cd ~/.vim/autoload
$ git clone https://github.com/tpope/vim-pathogen.git

# autoload 에 복사
$ cp -v ./vim-pathogen/autoload/pathogen.vim ./

```

#### 2.2 pathogen 환경 설정

```
# .vimrc 편집
$ vi ~/.vimrc
execute pathogen#infect()
filetype plugin indent on
```

#### 2.3 vundle 설치 (vim 패키지(플러그인) 관지자)

```
# ruby 설치
$ sudo apt-get install ruby

# vundle 을 ~/.vim/bundle 에 다운 받는다
$ cd ~/.vim/bundle
$ git clone https://github.com/gmarik/Vundle.vim.git
```

#### 2.4 vundle 환경 설정

```
# .vimrc 설정
$ vi ~/.vimrc
" Vundle 사용하기 위한 설정
" https://github.com/gmarik/Vundle.vim 에서 발췌
set nocompatible              " be iMproved, required
filetype off                  " required
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'gmarik/Vundle.vim'
call vundle#end()            " required
filetype plugin indent on    " required
```


## 3. vim-go 설치


#### 3.1 vim-go 
```
# vim-go 을 ~/.vim/bundle 에 다운 받는다
$ cd ~/.vimrc/bundle
$ git clone https://github.com/fatih/vim-go.git

# .vimrc 설정
$ vi ~/.vimrc
Plugin 'fatih/vim-go'

# 플러그인 설치
$ vim
:PluginInstall
```

#### 3.2 mercurial 설치
```
$ sudo apt-get install mercurial
# vim 실행 후 :GoInstallBinaries 를 실행하면 $GOPATH/bin 에 필요한 파일들이 설치된다.
$ vim
:GoInstallBinaries
```

## 4. vim 자동완성 기능

```
# YCM 설치전 python-dev, cmake 설치
$ sudo apt-get install python-dev
$ sudo apt-get install cmake

# YCM 을 ~/.vim/bundle 에 다운 받는다
$ cd ~/.vim/bundle
$ git clone https://github.com/Valloric/YouCompleteMe.git

# 설치
$ cd  YouCompleteMe

# 다음 명령이 필요하다면 수행
$ git submodule update --init --recursive
$ ./install.sh
```


## 5. Tagbar (태그 창) 설치

```
# Tagbar 설치전 ctags 설치
$ sudo apt-get install ctags

# Tagbar 을 ~/.vim/bundle 에 다운 받는다
$ cd ~/.vim/bundle
$ git clone https://github.com/majutsushi/tagbar.git

# vim 에서 :TagbarToggle 을 사용하면 오른쪽에 태그 창이 보인다.
# Ctrl+F12 단축키 설정
$ vi ~/.vimrc
nmap <C-F12> :TagbarToggle<CR>
```

## 6. NertdTree (파일 브라우저) 설치

```
$ cd ~/.vim/bundle
$ git clone https://github.com/scrooloose/nerdtree.git

# vim 에서 :NERDTreeToggle 을 사용하면 오른쪽에 태그 창이 보인다.
# F11 단축키 설정
$ vi ~/.vimrc
nmap <C-F11> :NERDTreeToggle<CR>
```



## 7. vim-go 명령들

#### 7.1 vim-go

```
$ vim
:GoRun (go 실행)
:GoBuild (go 빌드)
:GoDoc (go 문서)
:GoDef (go 변수 정의)
:GoFmt(go 형식 맞춤)
:GoImports (go 패키지 자동 임포트)
```

#### 7.2 vim-go 단축키 설정

```
$ ~/.vimrc
nmap <f5> :GoRun<cr>
nmap <f7> :GoBuild<cr>
nmap <c-i> :GoFmt<cr>
nmap <c-p> :GoImports<cr>
```