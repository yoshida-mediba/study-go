bash-3.2$ pwd
/Users/yoshida/github/study-go/1
bash-3.2$ ls
hello.go
bash-3.2$ go run hello.go
Hello, world!
bash-3.2$ ls
hello.go
bash-3.2$ cd ../
bash-3.2$ ls
1
bash-3.2$ mv 1 1-hello
bash-3.2$ ls
1-hello
bash-3.2$ git status
On branch master

Initial commit

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        1-hello/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 1-hello/
bash-3.2$ git status
On branch master

Initial commit

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)

        new file:   1-hello/hello.go

bash-3.2$ git commit -m 'hello' .
[master (root-commit) a49fc5b] hello
 1 file changed, 9 insertions(+)
 create mode 100644 1-hello/hello.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (2/2), done.
Writing objects: 100% (4/4), 330 bytes | 0 bytes/s, done.
Total 4 (delta 0), reused 0 (delta 0)
To github.com:yoshida-mediba/study-go.git
 * [new branch]      master -> master
bash-3.2$ rm ~/test.go
bash-3.2$ ls
1-hello
bash-3.2$ pwd
/Users/yoshida/github/study-go
bash-3.2$ go run 2-package/
go run: no go files listed
bash-3.2$ go run 2-package/main.go
Grass
Banana
Carrot
bash-3.2$ cd 2-package/
bash-3.2$ go build
bash-3.2$ ls
2-package	animals		main.go
bash-3.2$ ./2-package
Grass
Banana
Carrot
bash-3.2$ ls
2-package	animals		main.go
bash-3.2$ ls -al
total 3160
drwxr-xr-x  5 yoshida  staff      170 12 19 14:30 .
drwxr-xr-x  5 yoshida  staff      170 12 19 14:27 ..
-rwxr-xr-x  1 yoshida  staff  1611712 12 19 14:30 2-package
drwxr-xr-x  5 yoshida  staff      170 12 19 14:28 animals
-rw-r--r--  1 yoshida  staff      169 12 19 14:29 main.go
bash-3.2$ ls -alh
total 3160
drwxr-xr-x  5 yoshida  staff   170B 12 19 14:30 .
drwxr-xr-x  5 yoshida  staff   170B 12 19 14:27 ..
-rwxr-xr-x  1 yoshida  staff   1.5M 12 19 14:30 2-package
drwxr-xr-x  5 yoshida  staff   170B 12 19 14:28 animals
-rw-r--r--  1 yoshida  staff   169B 12 19 14:29 main.go
bash-3.2$ cd ..
bash-3.2$ rm 2-package/2-package
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        2-package/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 2-package/
bash-3.2$ git commit -m '2-package' .
[master 961bc65] 2-package
 4 files changed, 27 insertions(+)
 create mode 100644 2-package/animals/elephant.go
 create mode 100644 2-package/animals/monkey.go
 create mode 100644 2-package/animals/rabbit.go
 create mode 100644 2-package/main.go
bash-3.2$ git push
Counting objects: 8, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (8/8), done.
Writing objects: 100% (8/8), 790 bytes | 0 bytes/s, done.
Total 8 (delta 0), reused 0 (delta 0)
To github.com:yoshida-mediba/study-go.git
   a49fc5b..961bc65  master -> master
bash-3.2$ mkdir 3-test
bash-3.2$ rmdir 3-test/
bash-3.2$ cp -r 2-package 3-test
bash-3.2$ cd 3-test/
bash-3.2$ ls
animals	main.go
bash-3.2$ ls
animals	main.go
bash-3.2$ ls
animals	main.go
bash-3.2$ pwd
/Users/yoshida/github/study-go/3-test
bash-3.2$ go test ./animals/
ok  	_/Users/yoshida/github/study-go/3-test/animals	0.007s
bash-3.2$ go test -v ./animals/
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
PASS
ok  	_/Users/yoshida/github/study-go/3-test/animals	0.007s
bash-3.2$ ls
animals	main.go
bash-3.2$ pwd
/Users/yoshida/github/study-go/3-test
bash-3.2$ test go
bash-3.2$ go test
# _/Users/yoshida/github/study-go/3-test
./main_test.go:9: main() used as value
FAIL	_/Users/yoshida/github/study-go/3-test [build failed]
bash-3.2$ go test
# _/Users/yoshida/github/study-go/3-test
./main_test.go:9: undefined: AppName
FAIL	_/Users/yoshida/github/study-go/3-test [build failed]
bash-3.2$ ls
animals		main.go		main_test.go
bash-3.2$ ls
animals		main.go		main_test.go
bash-3.2$ mv ../2-package/app.go .
bash-3.2$ mv main_test.go app_test.go
bash-3.2$ go test
PASS
ok  	_/Users/yoshida/github/study-go/3-test	0.008s
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        ./

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ cd ../
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        3-test/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 3-test/
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   3-test/animals/animals_test.go
        new file:   3-test/animals/elephant.go
        new file:   3-test/animals/monkey.go
        new file:   3-test/animals/rabbit.go
        new file:   3-test/app.go
        new file:   3-test/app_test.go
        new file:   3-test/main.go

bash-3.2$ git commit -m 'test' .
[master 5075fd3] test
 7 files changed, 60 insertions(+)
 create mode 100644 3-test/animals/animals_test.go
 create mode 100644 3-test/animals/elephant.go
 create mode 100644 3-test/animals/monkey.go
 create mode 100644 3-test/animals/rabbit.go
 create mode 100644 3-test/app.go
 create mode 100644 3-test/app_test.go
 create mode 100644 3-test/main.go
bash-3.2$ git push
Counting objects: 7, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (7/7), done.
Writing objects: 100% (7/7), 945 bytes | 0 bytes/s, done.
Total 7 (delta 0), reused 0 (delta 0)
To github.com:yoshida-mediba/study-go.git
   961bc65..5075fd3  master -> master
bash-3.2$ pwd
/Users/yoshida/github/study-go
bash-3.2$ ls
1-hello		2-package	3-test		4-print
bash-3.2$ cd 4-print/
bash-3.2$ go run
go run: no go files listed
bash-3.2$ go run main.go
# command-line-arguments
./main.go:4: imported and not used: "fmt"
bash-3.2$ go run main.go
stdout
stderr
bash-3.2$ go run main.go 2>/dev/null
stdout
bash-3.2$ go run main.go 1>/dev/null
stderr
bash-3.2$ go run main.go
stdout
stderr
bash-3.2$ go run main.go
stdout
stderr
2016-12-19%!(EXTRA int=15)bash-3.2$
bash-3.2$
bash-3.2$ go run main.go
stdout
stderr
2016-12-19
%!(EXTRA int=15)bash-3.2$ go run main.go
stdout
stderr
2016-12-19
%!(EXTRA int=15)bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19
%!(EXTRA int=15)bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19%!(EXTRA int=15)
bash-3.2$
bash-3.2$
bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19%!(EXTRA int=15)
数値=5, 文字列=Golang, 配列=[1 2 3]bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19%!(EXTRA int=15)
数値=5, 文字列=Golang, 配列=[1 2 3]
bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19%!(EXTRA int=15)
数値=5, 文字列=Golang, 配列=[1 2 3]
数値=5, 文字列="Golang", 配列=[3]int{1, 2, 3}
bash-3.2$ go run main.go
stdout
stderr
2016-12-%!d(MISSING)
2016-12-19%!(EXTRA int=15)
数値=5, 文字列=Golang, 配列=[1 2 3]
数値=5, 文字列="Golang", 配列=[3]int{1, 2, 3}
数値=int, 文字列=string, 配列=[3]int
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        ./

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ cd ..
bash-3.2$ git add 4-print/
bash-3.2$ git commit -m '4-print'
[master f7c7c33] 4-print
 1 file changed, 19 insertions(+)
 create mode 100644 4-print/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 506 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   5075fd3..f7c7c33  master -> master
bash-3.2$ ls
1-hello		2-package	3-test		4-print
bash-3.2$ cd 5-variable/
bash-3.2$ go run
go run: no go files listed
bash-3.2$ go run main.go
# command-line-arguments
./main.go:11: unknown escape sequence: %
./main.go:11: unknown escape sequence: #
./main.go:11: unknown escape sequence: %
bash-3.2$ go run main.go
0=int
bash-3.2$ go run main.go
0 [int]
bash-3.2$ go run main.go
# command-line-arguments
./main.go:16: syntax error: unexpected {, expecting name
./main.go:17: syntax error: unexpected }
./main.go:17: cannot declare name
bash-3.2$ go run main.go
0 [int]
0 [int]
0 [int]
0 [int]
bash-3.2$ go run main.go
# command-line-arguments
./main.go:20: syntax error: unexpected {, expecting name
./main.go:21: cannot declare name
./main.go:21: syntax error: unexpected comma at end of statement
./main.go:22: syntax error: unexpected string at end of statement
./main.go:24: syntax error: non-declaration statement outside function body
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
-----
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
-----
1 [int]
10 [int]
100 [int]
-----
bash-3.2$ go run main.go
# command-line-arguments
./main.go:49: n redeclared in this block
        previous declaration at ./main.go:8
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
-----
1 [int]
10 [int]
100 [int]
-----
1 [int]
"true" [string]
-----
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
-----
1 [int]
10 [int]
100 [int]
-----
1 [int]
"true" [string]
-----
bash-3.2$ go run main.go
0 [int]
-----
0 [int]
0 [int]
0 [int]
-----
0 [int]
0 [int]
"" [string]
-----
true [bool]
1 [int]
3.14 [float64]
"abc" [string]
-----
1 [int]
10 [int]
100 [int]
-----
1 [int]
"true" [string]
-----
101 [int]
-----
bash-3.2$ cd ..
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        5-variable/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 5-variable/
bash-3.2$ git commit -m 'variable'
[master 21cef8a] variable
 1 file changed, 62 insertions(+)
 create mode 100644 5-variable/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 647 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   f7c7c33..21cef8a  master -> master
bash-3.2$ ls
1-hello		2-package	3-test		4-print		5-variable
bash-3.2$ mkdir 6-overflow
bash-3.2$ cd 6-overflow/
bash-3.2$ ls
bash-3.2$ pwd
/Users/yoshida/github/study-go/6-overflow
bash-3.2$ go run main.go
uint32 max value= 4294967295
bash-3.2$ pwd
/Users/yoshida/github/study-go/6-overflow
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
0xa [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
overflow
0xa [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
overflow
0xa [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
overflow
0xa [uint32]
0x9 [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
overflow
0xa [uint32]
0x9 [uint32]
0x8 [uint32]
bash-3.2$ go run main.go
uint32 max value= 4294967295
0x0 [uint32]
0xa [uint32]
overflow
0xa [uint32]
-----
0x9 [uint32]
0x8 [uint32]
bash-3.2$ pwd
/Users/yoshida/github/study-go/6-overflow
bash-3.2$ cd ..
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        6-overflow/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 6-overflow/
bash-3.2$ git commit -m '6-overflow'
[master c0e869d] 6-overflow
 1 file changed, 38 insertions(+)
 create mode 100644 6-overflow/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 573 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   21cef8a..c0e869d  master -> master
bash-3.2$ ls
1-hello		2-package	3-test		4-print		5-variable	6-overflow
bash-3.2$ ls
1-hello		2-package	3-test		4-print		5-variable	6-overflow
bash-3.2$
bash-3.2$
bash-3.2$ mkdir 7-array
bash-3.2$ cd 7-array/
bash-3.2$ ls
bash-3.2$ go run main.go
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
bash-3.2$ go run main.go
# command-line-arguments
./main.go:13: invalid array index 6 (out of bounds for 5-element array)
bash-3.2$ go run main.go
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{0, 0, 0, 0, 0} [[5]int]
-----
bash-3.2$ go run main.go
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{0, 0, 0, 0, 0} [[5]int]
-----
[5]int{1, 2, 3, 4, 5} [[5]int]
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
bash-3.2$ go run main.go
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{0, 0, 0, 0, 0} [[5]int]
-----
[5]int{1, 2, 3, 4, 5} [[5]int]
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{1, 11, 3, 4, 5} [[5]int]
[5]int{1, 2, 22, 4, 5} [[5]int]
-----
bash-3.2$ go run main.go
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{0, 0, 0, 0, 0} [[5]int]
-----
[5]int{1, 2, 3, 4, 5} [[5]int]
[5]int{1, 2, 3, 4, 5} [[5]int]
-----
[5]int{11, 2, 3, 4, 5} [[5]int]
[5]int{1, 22, 3, 4, 5} [[5]int]
-----
bash-3.2$ pwd
/Users/yoshida/github/study-go/7-array
bash-3.2$ cd ../
bash-3.2$ git add 7-array/
bash-3.2$ git commit -m 'array'
[master f614aa1] array
 1 file changed, 26 insertions(+)
 create mode 100644 7-array/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 465 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   c0e869d..f614aa1  master -> master
bash-3.2$ ls
1-hello		2-package	3-test		4-print		5-variable	6-overflow	7-array
bash-3.2$ mkdir 8-function
bash-3.2$ cd 8-function/
bash-3.2$ ls
bash-3.2$ go run main.go
Hello
bash-3.2$ go run main.go
Hello
voidFunc()
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bash-3.2$ go run main.go
# command-line-arguments
./main.go:10: voidFunc() used as value
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
%T true
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
bash-3.2$ go run main.go
# command-line-arguments
./main.go:10: syntax error: unexpected :=
bash-3.2$ go run main.go
# command-line-arguments
./main.go:10: voidFunc() used as value
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
bash-3.2$ go run main.go
# command-line-arguments
./main.go:12: multiple-value multiFunc() in single-value context
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
int
bash-3.2$ go run main.go
# command-line-arguments
./main.go:18: err declared and not used
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
int
int
string
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
int
int
string
20
bash-3.2$ go run main.go
# command-line-arguments
./main.go:28: cannot use func literal (type func(int)) as type func() in return argument
./main.go:29: too many arguments to return
./main.go:33: too many arguments in call to f2()
./main.go:33: f2()(100) used as value
bash-3.2$ go run main.go
# command-line-arguments
./main.go:29: too many arguments to return
./main.go:33: f2()(100) used as value
bash-3.2$ go run main.go
# command-line-arguments
./main.go:29: too many arguments to return
bash-3.2$ go run main.go
# command-line-arguments
./main.go:29: too many arguments to return
./main.go:33: f2()() used as value
bash-3.2$ go run main.go
# command-line-arguments
./main.go:28: cannot use func literal (type func() int) as type func() in return argument
./main.go:33: f2()() used as value
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
int
int
string
20
2
bash-3.2$ go run main.go
# command-line-arguments
./main.go:33: not enough arguments in call to f2()
bash-3.2$ go run main.go
Hello
voidFunc()
stdFunc()
bool
int, int
int
int
string
20
200
bash-3.2$ pwd
/Users/yoshida/github/study-go/8-function
bash-3.2$ cd ..
bash-3.2$ git add 8-function/
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   8-function/main.go

bash-3.2$ git commit -m '8-function'
[master fb7b948] 8-function
 1 file changed, 52 insertions(+)
 create mode 100644 8-function/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 610 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   f614aa1..fb7b948  master -> master
bash-3.2$ ls
1-hello		3-test		5-variable	7-array
2-package	4-print		6-overflow	8-function
bash-3.2$ mkdir 9-scope
bash-3.2$ cd 9-scope/
bash-3.2$ go run main.go
main.go:5:2: cannot find package "conf" in any of:
        /usr/local/Cellar/go/1.7.4_1/libexec/src/conf (from $GOROOT)
        /Users/yoshida/golang/src/conf (from $GOPATH)
bash-3.2$ ls
conf.go	main.go
bash-3.2$ mkdir conf
bash-3.2$ mv conf.go conf
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ find .
.
./conf
./conf/conf.go
./conf.go
./main.go
bash-3.2$ rm conf.go
bash-3.2$ go run main.go
# _/Users/yoshida/github/study-go/9-scope/conf
conf/conf.go:5: extra expression in const declaration
conf/conf.go:5: syntax error: unexpected =, expecting semicolon, newline, or )
bash-3.2$ go run main.go
# command-line-arguments
./main.go:9: cannot refer to unexported name conf.testFunc
./main.go:9: undefined: conf.testFunc
bash-3.2$ go run main.go
# command-line-arguments
./main.go:5: imported and not used: "_/Users/yoshida/github/study-go/9-scope/conf"
./main.go:9: undefined: testFunc
bash-3.2$ go run main.go
# command-line-arguments
./main.go:9: cannot refer to unexported name conf.testFunc
./main.go:9: undefined: conf.testFunc
bash-3.2$ go run main.go
10
bash-3.2$ go run main.go
10
100
bash-3.2$ go run main.go
# command-line-arguments
./main.go:10: cannot refer to unexported name conf.y
bash-3.2$ go run main.go
10 10
100
bash-3.2$ go run main.go
10 10
100
bash-3.2$ go run main.go
10 10
100
3.141592653589793
bash-3.2$ pwd
/Users/yoshida/github/study-go/9-scope
bash-3.2$ cd ..
bash-3.2$ git add 9-scope/
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   9-scope/conf/conf.go
        new file:   9-scope/main.go

bash-3.2$ git commit -m 'scope'
[master a89473d] scope
 2 files changed, 28 insertions(+)
 create mode 100644 9-scope/conf/conf.go
 create mode 100644 9-scope/main.go
bash-3.2$ git push
Counting objects: 6, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (5/5), done.
Writing objects: 100% (6/6), 589 bytes | 0 bytes/s, done.
Total 6 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   fb7b948..a89473d  master -> master
bash-3.2$
bash-3.2$
bash-3.2$ ls
1-hello		3-test		5-variable	7-array		9-scope
2-package	4-print		6-overflow	8-function
bash-3.2$ ls -al
total 0
drwxr-xr-x  12 yoshida  staff  408 12 19 16:50 .
drwxr-xr-x  17 yoshida  staff  578 12 19 14:24 ..
drwxr-xr-x  13 yoshida  staff  442 12 19 17:01 .git
drwxr-xr-x   3 yoshida  staff  102 12 19 14:24 1-hello
drwxr-xr-x   4 yoshida  staff  136 12 19 14:46 2-package
drwxr-xr-x   6 yoshida  staff  204 12 19 14:46 3-test
drwxr-xr-x   3 yoshida  staff  102 12 19 15:00 4-print
drwxr-xr-x   3 yoshida  staff  102 12 19 15:16 5-variable
drwxr-xr-x   3 yoshida  staff  102 12 19 15:46 6-overflow
drwxr-xr-x   3 yoshida  staff  102 12 19 16:16 7-array
drwxr-xr-x   3 yoshida  staff  102 12 19 16:27 8-function
drwxr-xr-x   4 yoshida  staff  136 12 19 16:56 9-scope
bash-3.2$ pwd
/Users/yoshida/github/study-go
bash-3.2$
bash-3.2$
bash-3.2$ pwd
/Users/yoshida/github/study-go
bash-3.2$ ls
1-hello		3-test		5-variable	7-array		9-scope
2-package	4-print		6-overflow	8-function
bash-3.2$ ls -al
total 0
drwxr-xr-x  12 yoshida  staff  408 12 19 16:50 .
drwxr-xr-x  17 yoshida  staff  578 12 19 14:24 ..
drwxr-xr-x  13 yoshida  staff  442 12 19 17:01 .git
drwxr-xr-x   3 yoshida  staff  102 12 19 14:24 1-hello
drwxr-xr-x   4 yoshida  staff  136 12 19 14:46 2-package
drwxr-xr-x   6 yoshida  staff  204 12 19 14:46 3-test
drwxr-xr-x   3 yoshida  staff  102 12 19 15:00 4-print
drwxr-xr-x   3 yoshida  staff  102 12 19 15:16 5-variable
drwxr-xr-x   3 yoshida  staff  102 12 19 15:46 6-overflow
drwxr-xr-x   3 yoshida  staff  102 12 19 16:16 7-array
drwxr-xr-x   3 yoshida  staff  102 12 19 16:27 8-function
drwxr-xr-x   4 yoshida  staff  136 12 19 16:56 9-scope
bash-3.2$ mkdir 10-control_flow
bash-3.2$ cd 10-control_flow/
bash-3.2$ ls
bash-3.2$ go run main.go
# command-line-arguments
./main.go:11: syntax error: unexpected ++, expecting comma or )
bash-3.2$ go run main.go
1
2
3
4
5
6
7
8
9
10
11
bash-3.2$ go run main.go
1
2
3
4
5
6
7
8
9
10
bash-3.2$ go run main.go
1
2
3
4
5
0
1
2
3
4
bash-3.2$ go run main.go
0
1
2
3
4
5
0
1
2
3
4
bash-3.2$ go run main.go
0
1
2
3
4
5
6
0
1
2
3
4
bash-3.2$ go run main.go
0
1
2
3
4
0
1
2
3
4
bash-3.2$ go run main.go
# command-line-arguments
./main.go:23: undefined: err
./main.go:24: undefined: err
bash-3.2$ go run main.go
# command-line-arguments
./main.go:23: invalid operation: err != nil (mismatched types string and nil)
bash-3.2$ go run main.go
0
1
2
3
4
0
1
2
3
4
ERROR!!
bash-3.2$ go run main.go
0
1
2
3
4
0
1
2
3
4
ERROR!!
bash-3.2$ go run main.go
0
1
2
3
4
0
1
2
3
4
0
1
2
3
4
ERROR!!
bash-3.2$ go run main.go
i 0
i 1
i 2
i 3
i 4
j 0
j 1
j 2
j 3
j 4
p 0
p 1
p 2
p 3
p 4
ERROR!!
bash-3.2$ go run main.go
i= 0
i= 1
i= 2
i= 3
i= 4
j= 0
j= 1
j= 2
j= 3
j= 4
p= 0
p= 1
p= 2
p= 3
p= 4
ERROR!!
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
bash-3.2$ go run main.go
# command-line-arguments
./main.go:33: undefined: fruits
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
10!
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
10!
n == 10!
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
bash-3.2$ go run main.go
# command-line-arguments
./main.go:54: xi declared and not used
./main.go:55: xf declared and not used
./main.go:56: xs declared and not used
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
%x, %x, %x false true false
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
%!x(bool=false), %!x(bool=true), %!x(bool=false)
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
x is float: %v
 3.14
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
x is float: 3.14
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
x is float: 3.14
float64
bash-3.2$ go run main.go
i = 0
i = 1
i = 2
i = 3
i = 4
j = 0
j = 1
j = 2
j = 3
j = 4
p = 0
p = 1
p = 2
p = 3
p = 4
ERROR!!
fruits[0]=Apple
fruits[1]=Banana
fruits[2]=Cherry
n == 10!
n == 10!
false, true, false
x is float: 3.14
float64 3.14
bash-3.2$ pwd
/Users/yoshida/github/study-go/10-control_flow
bash-3.2$ cd ..
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        10-control_flow/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 10-control_flow/
bash-3.2$ git commit -m 'control_flow'
[master 07e7a47] control_flow
 1 file changed, 80 insertions(+)
 create mode 100644 10-control_flow/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 855 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   a89473d..07e7a47  master -> master
bash-3.2$ mkdir 11-panic_recover
bash-3.2$ cd 11-panic_recover/
bash-3.2$ go run main.go
# command-line-arguments
./main.go:4: imported and not used: "fmt"
bash-3.2$ go run main.go
panic: ERROR!!

goroutine 1 [running]:
panic(0x56020, 0xc42000a150)
        /usr/local/Cellar/go/1.7.4_1/libexec/src/runtime/panic.go:500 +0x1a1
main.doPanic()
        /Users/yoshida/github/study-go/11-panic_recover/main.go:12 +0x6d
main.main()
        /Users/yoshida/github/study-go/11-panic_recover/main.go:8 +0x14
exit status 2
bash-3.2$ go run main.go
# command-line-arguments
./main.go:10: undefined: fmt in fmt.Println
bash-3.2$ go run main.go
ERROR!!
bash-3.2$ cd ..
bash-3.2$ git add 11-panic_recover/
bash-3.2$ git commit -m '11-panic_recover'
[master ae4e376] 11-panic_recover
 1 file changed, 18 insertions(+)
 create mode 100644 11-panic_recover/main.go
bash-3.2$ git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)
nothing to commit, working tree clean
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 456 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   07e7a47..ae4e376  master -> master
bash-3.2$ mkdir 12-go_init
bash-3.2$ cd 12-go_init/
bash-3.2$ pwd
/Users/yoshida/github/study-go/12-go_init
bash-3.2$ go run main.go
bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!bash-3.2$
bash-3.2$
bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!bash-3.2$ go run main.go
# command-line-arguments
./main.go:11: cannot use time.Millisecond (type time.Duration) as type int in assignment
./main.go:16: cannot use sleep (type int) as type time.Duration in argument to time.Sleep
bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
131
bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
AB
bash-3.2$ go run main.go
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
SUB!!
ABC
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        ./

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ cd ..
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        12-go_init/

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git add 12-go_init/
bash-3.2$ git commit -m 'go_init'
[master b441b24] go_init
 1 file changed, 27 insertions(+)
 create mode 100644 12-go_init/main.go
bash-3.2$ git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 488 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local objects.
To github.com:yoshida-mediba/study-go.git
   ae4e376..b441b24  master -> master
bash-3.2$ ls
1-hello			12-go_init		4-print			7-array
10-control_flow		2-package		5-variable		8-function
11-panic_recover	3-test			6-overflow		9-scope
bash-3.2$ mkdir 13-slice
bash-3.2$ cd 13-slice/
bash-3.2$ ls
bash-3.2$ pwd
/Users/yoshida/github/study-go/13-slice
bash-3.2$ go run main.go
8
bash-3.2$ go run main.go
8 10
bash-3.2$ go run main.go
8 10
[1 2 3 4 5 6 7 8]bash-3.2$ go run main.go
8 10
[6 7 8]bash-3.2$ go run main.go
8 10
[6 7 8]
bash-3.2$ go run main.go
8 10
[6 7 8]
[1 2 3 4 5]
[1 2 3 4 5 6 7 8]
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
ld!
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
orld!
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
World
bash-3.2$ go run main.go
# command-line-arguments
./main.go:19: missing '
./main.go:19: syntax error: unexpected DOYA, expecting comma or )
bash-3.2$ go run main.go
# command-line-arguments
./main.go:19: first argument to append must be slice; have string
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
World
[1 2 3 4 5 6 7 8 9 10]
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
World
[1 2 3 4 5 6 7 8 9 10]
3 := [1 2 3 4 5 1 2 3]
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
World
[1 2 3 4 5 6 7 8 9 10]
3 := [6 7 8 4 5 6 7 8]
bash-3.2$ git branch
* master
bash-3.2$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        ./

nothing added to commit but untracked files present (use "git add" to track)
bash-3.2$ git push
Everything up-to-date
bash-3.2$ git push
Everything up-to-date
bash-3.2$ go run main.go
8 10
[1 2 3 4 5]
[6 7 8]
[1 2 3 4 5 6 7 8]
World
[1 2 3 4 5 6 7 8 9 10]
3 := [6 7 8 4 5 6 7 8]
15
bash-3.2$
