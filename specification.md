# 设计思路

参照 `spf13/cobra` 的 `command.go` 进行精简版的实现。从其核心的 `Command struct` 入手，仿照借鉴 `cobra` 的源代码，从只支持单个根命令，逐步拓展至支持多个子命令，并辅以部分测试以保证实现过程中重构的安全性。此外为了兼容 Hugo，另外实现了 Hugo 所需要的一部分字段与方法。

# 单元测试结果
```
=== RUN   TestSimpleCommand
--- PASS: TestSimpleCommand (0.00s)
=== RUN   TestSingleCommandWithArgs
--- PASS: TestSingleCommandWithArgs (0.00s)
=== RUN   TestChildCommand
--- PASS: TestChildCommand (0.00s)
PASS
ok      github.com/robinWongM/cobra     0.003s
```

# 功能测试结果

## Simple MAIN

使用 `examples/main.go` 进行测试，该文件包含了一个简单的子命令，通过读取第一个 arbitrary argument 来替换输出内容。

```bash
codespace ➜ ~/workspace/cobra/examples (main ✗) $ ./main
codespace ➜ ~/workspace/cobra/examples (main ✗) $ ./main hello
Hello, World!
codespace ➜ ~/workspace/cobra/examples (main ✗) $ ./main go
codespace ➜ ~/workspace/cobra/examples (main ✗) $ ./main hello test
Hello, test!
```

## Hugo

Hugo 是一个静态博客生成器，其 CLI 使用 `spf13/cobra` 构建。本次通过替换 Hugo 源代码中对 `spf13/cobra` 的引用为 `robinWongM/cobra`，以测试本包的简单功能可用性。

替换后的源码在 [robinWongM/hugo 的 cobra-replacement 分支](https://github.com/robinWongM/hugo/tree/cobra-replacement)。可使用该源代码编译，并使用 hugo 提供的各种功能，例如构建静态博客，本地运行服务器等。

```bash
rwong@hugo-test:~/hugo$ ./hugo
Start building sites … 
WARN 2020/11/10 00:04:33 found no layout file for "HTML" for kind "home": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.
WARN 2020/11/10 00:04:33 found no layout file for "HTML" for kind "taxonomy": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.
WARN 2020/11/10 00:04:33 found no layout file for "HTML" for kind "taxonomy": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.

                   | EN  
-------------------+-----
  Pages            |  3  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     |  0  
  Processed images |  0  
  Aliases          |  0  
  Sitemaps         |  1  
  Cleaned          |  0  

Total in 31 ms

rwong@hugo-test:~/hugo$ ./hugo list all
path,slug,title,date,expiryDate,publishDate,draft,permalink

rwong@hugo-test:~/hugo$ ./hugo server
Start building sites … 
WARN 2020/11/10 00:05:37 found no layout file for "HTML" for kind "home": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.
WARN 2020/11/10 00:05:37 found no layout file for "HTML" for kind "taxonomy": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.
WARN 2020/11/10 00:05:37 found no layout file for "HTML" for kind "taxonomy": You should create a template file which matches Hugo Layouts Lookup Rules for this combination.

                   | EN  
-------------------+-----
  Pages            |  3  
  Paginator pages  |  0  
  Non-page files   |  0  
  Static files     |  0  
  Processed images |  0  
  Aliases          |  0  
  Sitemaps         |  1  
  Cleaned          |  0  

Built in 4 ms
Watching for config changes in /home/rwong/hugo/go.mod
Environment: "development"
Serving pages from memory
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at //localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop

```