# flomo-cli

![Go](https://github.com/geekdada/flomo-cli/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/geekdada/flomo-cli/branch/master/graph/badge.svg?token=FJ3Y2ZB8YS)](https://codecov.io/gh/geekdada/flomo-cli)

A CLI tool for [flomo](https://flomoapp.com/register2/?Mzk3).

## 安装

在 [Releases](https://github.com/geekdada/flomo-cli/releases) 页面下载对应的二进制文件。

## 使用

### 添加一条新的墨

```bash
$ flomo new --api <YOUR API> "一条新的墨"
```

### 添加一条带标签的墨

```bash
$ flomo new --api <YOUR API> --tag "随手记" "一条新的墨"
```

**效果**

![CleanShot 2020-12-24 at 20.27.55@2x.png](https://i.loli.net/2020/12/24/g3v7c6fwOKyauRT.png)

### 使用环境变量来指定 API

```bash
$ export FLOMO_API=<YOUR API>
$ flomo new --tag "随手记" "一条新的墨"
```

### 将文本文件添加到浮墨

```bash
$ cat memo.txt | flomo new --tag "Quote"
```

## LICENCE

[MIT](./LICENSE)
