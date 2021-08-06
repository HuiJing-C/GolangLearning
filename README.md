# GolangLearning
go语言自学



纠错笔记：

1：想使用单测的方式来进行语言学习，但是发现以下形式并不能在IDEA运行

```
func TestA(t *testing.T) {
}
```

Answer：Go语言单测只能是以test结尾的才可以使用



2：同一project下不同package之间不可以互相import

Answer：由于上一步全是以test结尾的.go文件，Go语言不导出package下以test结尾的公开方法或者属性



3：Go mod 命令介绍

- `go help mod`查看帮助。
- `go mod init <项目模块名称>`初始化模块，会在项目根目录下生成 `go.mod` 文件。参数`<项目模块名称>`是非必写的，但如果你的项目还没有代码编写，这个参数能快速初始化模块。如果之前使用其它依赖管理工具(比如dep，glide等)，mod会自动接管原来依赖关系。
- `go mod tidy`根据go.mod文件来处理依赖关系。
- `go mod vendor`将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令`go build -mod=vendor`编译。
- `go list -m all`显示依赖关系。`go list -m -json all`显示详细依赖关系。
- `go list -m -versions <path>`显示包有哪些已发布版本
- `go mod download <path@version>`下载依赖。参数`<path@version>`是非必写的，path是包的路径，version是包的版本。
- 其它命令可以通过`go help mod`来查看。



4：Go mod两个文件的介绍

**go.mod** 提供了`module`, `require`、`replace`和`exclude` 四个命令

- `module` 语句指定包的名字（路径）
- `require` 语句指定的依赖项模块
- `replace` 语句可以替换依赖项模块
- `exclude` 语句可以忽略依赖项模块

**go.sum** 的每一行都是一个条目，大致是这样的格式：

```shell
<module> <version>/go.mod <hash>
```

或者

```shell
<module> <version> <hash>
<module> <version>/go.mod <hash>
```

备注：其中module是依赖的路径，version是依赖的版本号。hash是以`h1:`开头的字符串，表示生成checksum的算法是第一版的hash算法（sha256）

- 项目没有打 tag，会生成一个版本号，格式如下：v0.0.0-commit日期-commitID 
- 引用一个项目的特定分支，比如 develop branch，也会生成类似的版本号： v当前版本+1-commit日期-commitID
- 项目有用到 go module，那么就是正常地用 tag 来作为版本号。 如果项目打了 tag，但是没有用到 go module，为了跟用了 go module 的项目相区别，需要加个 `+incompatible` 的标志。比如`:<module>+<version>+incompatible/go.mod+<hash>`
- 对于使用了 v2+ go module 的项目，项目路径会有个版本号的后缀。比如： `<module/v2>+<version> + <hash>`

```shell
// 没有打tag的场景
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
// 打tag的指定branch的场景
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
// 使用了go mod的场景
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
```

- 有go.mod后，如果想升级某一个依赖，可以直接使用go get 依赖地址@版本号/分支，来更新依赖，mod文件会自动更新，比如`go get github.com/aliyun/aliyun-log-go-sdk@v0.1.21 `或者 `go get github.com/aliyun/aliyun-log-go-sdk@master`

