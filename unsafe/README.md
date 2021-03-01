# unsafe.Pointer编译规则

转化规则 类型安全指针 <==> 非类型安全指针 <==> uintptr

编译规则，编译能通过，运行是可能存在问题

# 运行时原则

    1. 保证使用的值的unsafe操作前后时时刻刻都被有效的引用着，无论安全指针还是非安全类型指针。否则可能被垃圾回收器回收掉
    2. 任何指针都不应该引用未知内存块

  

![](https://tva1.sinaimg.cn/large/008eGmZEly1gnt7lxngcyj31ll0u013v.jpg)

