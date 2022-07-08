# git
![](https://www.runoob.com/wp-content/uploads/2015/02/git-command.jpg)

+ workspace：工作区
+ staging area：暂存区/缓存区
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

![](https://img-blog.csdnimg.cn/20191201114346620.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2t1c2VkZXhpbmdmdQ==,size_16,color_FFFFFF,t_70)


index:暂存区
HEAD:最近一次commit
working copy:工作区


git merge 合并分支


## 合并冲突的解决



## git diff 

比较文件的不同，即暂存区和工作区的差异


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