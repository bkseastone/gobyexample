# sync

## sync.Once
全局只会执行一次

## sync.Pool
有些对象会经常用到,为了减少gc的压力 将经常用到的对象放到池中,
要用的时候去取,不用了就放回去.每次gc时会把对象池清空
比如 String builder 和 bytes.buffer 会经常重用