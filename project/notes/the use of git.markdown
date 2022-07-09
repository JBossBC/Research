# git


## git 对象

git 对象一共有三种:数据对象、树对象以及提交对象

### 数据对象(blob)

数据对象的产生是在使用git add 命令将文件或者目录加入到暂存区时产生的，git会把一个文件中要存的数据和一个头部信息一起做一个hash(sha-1)散列运算，将得到的散列值作为这个文件得到路径。

根据git cat-file -p hash来看，其实这个散列值路径文件中保存的只是源文件的一种压缩形式


## 树对象(tree)

git中的树对象能解决文件名保存的问题，也允许我们将多个文件组织到一起，**git以一种类似于unix系统的方式存储内容**，但做了许多简化。所有内容均已树对象和数据对象的形式进行存储，其中树对象对应了unix中的目录项，数据对象大致对应了文件内容


git ls-files --stage中前面的数字代表Git为每一个文件设置了文件模式100644，表明这两个文件都是普通文件， 其他两种种模式选择包括：100755，表示一个可执行文件；120000，表示一个符号链接。上述三种模式即是Git数据对象的所有合法模式（当然对于目录项和子模块还有其他的一些模式);模式为040000(表示该文件名为一个目录)


git会将一个目录存储为一个树对象进行保存，这个树对象中会有每个子目录的树对象条目和每个文件的数据对象条目，这样就可以根据一种树结构保存整个目录

## 提交对象(commit)

提交对象中包含一个树对象条目，代表者当前项目快照，从这树对象开始我们就能找到所有提交的数据对象，从而形成git中的一个版本

其他之外还有作者一些信息，留空一行，最后是提交注释

## 总结

git会对所有文件内容进行压缩，即使仓库存储了非常多内容，.git文件也不会很大，如果我们真正的修改了一个文件，这个文件的散列值会被修改，然后将这个文件压缩存储在objects中，这样我们需要创建一个相应的树结构来对原来的提交进行修改，这其实并不是一个困难的过程，我们只要为每一次修改都创建一个顶层的树对象来表示这个提交快照。git可能会对比前一个提交的顶层树对象，然后将没有修改的树对象或数据对象直接复制到新创建的这个顶层树对象中，将改变的树对象像这样递归的进行。也就是说决定你仓库大小的并完全是每个文件的大小，而是你修改提交的次数，修改的次数越多，产生的树对象和数据对象也就越多。

**每一次commit的顶层树结构都包含了不管有没有修改的所有子树对象和数据对象，如果没有修改，就引用上一次提交的相关树对象hash**

![](https://www.runoob.com/wp-content/uploads/2015/02/git-command.jpg)

+ workspace：工作区
+ staging area：暂存区/缓存区(一个简单的索引文件，指的是.git/index文件,索引文件里面包含的是文件的目录树，**像一个虚拟的工作区**，在这个虚拟工作区的目录树中，记录了文件名、文件的最后修改时间、文件长度、以及最重要的sha-1值，文件的内容并没有存在其中)
+ local repository：版本库或本地仓库
+ remote repository：远程仓库


## 工作区、版本库中的暂存区和版本库之间的关系
![](https://www.runoob.com/wp-content/uploads/2015/02/1352126739_7909.jpg)

index:暂存区,标记为"master"的是master分支所代表的目录树
"HEAD"实际是指向master分支的一个"游标"。
Objects标识的区域为Git的对象库，实际位于".git/objects"目录下，里面包含了创建的各种对象及内容

+ 当对工作区修改(或者新增)的文件执行git add命令的时候，暂存区的目录树被更新，同时工作区修改(或新增)的文件内容被写入到对象库中的一个新对象中，而该对象的ID被记录在暂存区的文件索引中

+ 当执行提交操作(git commit)时，暂存区的目录树写到版本库(对象库)中国，master分支会做出相应的更新。即master指向的目录树就是提交时暂存区的目录树。

+ 当执行git reset HEAD命令时，暂存区的目录树会被重写，被master分支指向的目录树所替换，但是工作区不受影响。
+ 当执行 `git rm --cached<file>`命令时，会直接从暂存区删除文件，工作区不受影响
+ 当执行	`git checkout .`或者`git checkout -- <file>`命令时，会用暂存区全部或指定的文件替换工作区的文件。这个操作很危险，会清除工作区中未添加到暂存区中的改动。(我们使用checkout来转换分支的时候，分支上面的最新目录树就是暂存区的目录树)
+ 当执行git checkout HEAD .或者git checkout HEAD <file>命令时，会用HEAD指向的master分支中的全部或者部分文件替换暂存区和以及工作区中的文件。这个命令也是极具危险性的，因为不但会清楚工作区中未提交的改动，也会清楚暂存区中未提交的改动。

## git 标签

当你达到一个重要的阶段，并希望永远记住那个特别的提交快照，可以使用git tag 给它打上标签。-a选项意为"创建一个带注解的标签"，不用-a选项也可以执行，但不会记录这个标签是啥时候打的，是谁打的，会不会让你添加这个标签的注解。

当执行git log -decorate 的时候，我们就可以看到我们的标签了
## Key word

VSC(version system control)

## what is version control

版本控制是能记录一个文件或一系列文件在过去一段时间发生的变化，以便我们之后能够回退到之前特定的版本的系统。


+ 集中式版本控制的缺点：
    
    单点故障会导致系统不能正常运作，如果磁盘

**git add -p  能够决定将文件的哪部分内容提交到index区域**


**git reset**

>git-reset - Reset current HEAD to the specified state


--soft

    Does not touch the index file or the working tree at all (but resets the head to <commit>, just like all modes do). This leaves all your changed files "Changes to be committed", as git status would put it.
--mixed

    Resets the index but not the working tree (i.e., the changed files are preserved but not marked for commit) and reports what has not been updated. This is the default action.

    If -N is specified, removed paths are marked as intent-to-add (see git-add(1)).
--hard

    Resets the index and working tree. Any changes to tracked files in the working tree since <commit> are discarded. Any untracked files or directories in the way of writing any tracked files are simply deleted.


上面三种参数对应三种级别，确立了 head回退的时候对 work repository，index area，local repository的影响。




![](https://img-blog.csdnimg.cn/20191201114346620.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t1c2VkZXhpbmdmdQ==,size_16,color_FFFFFF,t_70)


index:暂存区
HEAD:最近一次commit
working copy:工作区


git merge 合并分支


## 合并冲突的解决



## git diff 
> git-diff - Show changes between commits, commit and working tree, etc

> git diff [<options>] [<commit>] [--] [<path>…​]
> 
> git diff [<options>] --cached [--merge-base] [<commit>] [--] [<path>…​]
> 
> git diff [<options>] [--merge-base] <commit[<commit>…​] <commit[--] [<path>…​]
> 
> git diff [<options>] <commit>…​<commit[--] [<path>…​]
> 
> git diff [<options>] <blob<blob>
> 
> git diff [<options>] --no-index [--] <path<path>


git diff HEAD 查看工作区与本地仓库之间的状态

## git log 

展示提交日志 

## git show(和 git cat-file -p的作用优点相似)
> 
> git-show - Show various types of objects

     描述:

          + 显示一个或多个对象(blob、树、tag和提交)
          + 对于提交，他显示日志消息和文本差异。他还以git-diff-tree-cc生成的特殊格式显示合并提交
          + 标示标记，它显示标记消息和引用对象
          + 对于树，它显示名称(相当于git ls tree，仅带-name)
          + 对于普通blob，它显示普通内容
        
      该命令采用适用于git diff tree命令的选项来控制提交引入的更改的显示方式。



## git cat-file -p hash


git-cat-file - Provide content or type and size information for repository objects

查看objects对象数据

-p

    Pretty-print the contents of <object> based on its type.



## 查看某次提交的内容

git show hash

## 查看本地某个 tag 的详细信息：

git show \<tagName>

## 查看工作区文件
`git ls-files `


## 查看index文件 

`git ls-files --stage`

## 查看index区文件的内容

`git cat-file  hash`

> <typecan be one of: blob, tree, commit, tag
> 
>     -t                    show object type
>     
>     -s                    show object size
>     
>     -e                    exit with zero when there's no error
>     
>     -p                    pretty-print object's content
>     
>     --textconv            for blob objects, run textconv on object's content
>     
>     --filters             for blob objects, run filters on object's content
>     
>     --path <blob        use a specific path for --textconv/--filters
>     
>     --allow-unknown-type  allow -s and -t to work with broken/corrupt objects
>     
>     --buffer              buffer --batch output
>     
>     --batch[=<format>]    show info and content of objects fed from the standard input
>     
>     --batch-check[=<format>]
>                           show info about objects fed from the standard input
>                           
>     --follow-symlinks     follow in-tree symlinks (used with --batch or --batch-check)
>     
>     --batch-all-objects   show all objects with --batch or --batch-check
>     
>     --unordered           do not order --batch-all-objects output

## git add
git add 的作用就是创建一个blob文件来记录最新的修改代码，并且在index file里面添加一个到blob的链接

## git tag

> 创建一个tag,基于本地分支的commit，与分支的推送是两回事

git tag <tagName>

> 推送tag到远程分支

git push origin --tags


## 删除远程分支

删除远程分支 git push origin --delete branchName
