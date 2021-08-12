package _8interface_reflect

// 读和写是软件中很普遍的行为，提起它们会立即想到读写文件、缓存（比如字节或字符串切片）、标准输入输出、标准错误以及网络连接、管道等等，或者读写我们的自定义类型。
// 为了让代码尽可能通用，Go 采取了一致的方式来读写数据。
//
// io 包提供了用于读和写的接口 io.Reader 和 io.Writer：
//
// type Reader interface {
//    Read(p []byte) (n int, err error)
// }
// type Writer interface {
//    Write(p []byte) (n int, err error)
// }
// 只要类型实现了读写接口，提供 Read() 和 Write 方法，就可以从它读取数据，或向它写入数据。
// 一个对象要是可读的，它必须实现 io.Reader 接口，这个接口只有一个签名是 Read(p []byte) (n int, err error) 的方法，它从调用它的对象上读取数据，并把读到的数据放入参数中的字节切片中，然后返回读取的字节数和一个 error 对象，如果没有错误发生返回 nil，如果已经到 达输入的尾端，会返回 io.EOF("EOF")，如果读取的过程中发生了错误，就会返回具体的错误信息。
// 类似地，一个对象要是可写的，它必须实现 io.Writer 接口，这个接口也只有一个签名是 Write(p []byte) (n int, err error) 的方法，它将指定字节切片中的数据写入调用它的对象里，然后返回实际写入的字节数和一个 error 对象（如果没有错误发生就是 nil）。
//
// io 包里的 Readers 和 Writers 都是不带缓冲的，bufio 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时它们尤其有用。在 09read_write 我们会看到很多在实战中使用它们的例子。
//
// 在实际编程中尽可能的使用这些接口，会使程序变得更通用，可以在任何实现了这些接口的类型上使用读写方法。
//
// 例如一个 JPEG 图形解码器，通过一个 Reader 参数，它可以解码来自磁盘、网络连接或以 gzip 压缩的 HTTP 流中的 JPEG 图形数据，或者其他任何实现了 Reader 接口的对象。
