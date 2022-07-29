# Golang builtin


## make

make内置函数分配和初始化仅限类型为object，slice，map或者chan。与new一样，第一个参数是type，而不是value。与new不同的是，make的返回类型和他的参数类型一样，而不是指向参数的指针。

## new

new内置函数分配内存。他的第一个参数是类型而不是value，返回值是一个指针指向一个新分配的该类型的零值