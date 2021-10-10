# log

## 原则

1. 记录日志的目的是为了出现问题时能快速找到根因。
2. 日志应该有一套易读的结构。

## 需求

1. 应用内并不管log输出格式等，只管调用方法将信息输出，log的格式化由`log`包来做。
2. log提供等级`Debug`，`Info`，`Warning`，`Error`。
3. log对象可以指定前缀，并且一次指定，包内生效。

## 设计

1. 看了`logrus`的源码后，发现`logrus`完全符合我的需求，并且依赖简单，所以决定使用`logrus`进行日志输出。
2. 为了不强依赖`logrus`，在`log`包内提供`Logger`接口供应用调用。`logger`结构体包装`logrus.Entry`，并实现`Logger`接口。

## 问题

### `logger`设置相关的问题

在我实现`logger`结构体时，主要遇到了以下两个问题：

1. **`Level`的处理**：为了方便，我打算直接在新建(`NewLogger()`)方法中传参`string`类型的`level`，然后通过`switch`将这个`level`传给`logrus`，但这样做还是有问题的，比如当遇到不支持的`Level`时如何处理？常规来说应该返回一个错误，但这种`NewXXX`的新建方法我认为不应该返回一个错误。
2. **输出格式的处理**：通过一种什么姿势，把输出格式加载到`logrus`中？

这两个问题的本质其实是：**当一个包内同时包含对外暴露的接口，和不对外暴露的实现类时，我如何初始化这个实现类，如何将配置加载到这个实现类中。**

那么，开源项目中有哪些这样的例子呢？

1. `go-kratos/kratos`中，`log`包的`stdLogger`结构体：

   `stdLogger`使用标准库`log`实现其定义的`Logger`接口。通过`NewStdLogger()`方法来初始化。标准库的`log.New()`方法需要三个参数，而在`NewStdLogger()`中只要求其中一个`writer`，另外两个以默认值。

   ```go
   // NewStdLogger new a logger with writer.
   func NewStdLogger(w io.Writer) Logger {
   	return &stdLogger{
   		log: log.New(w, "", 0),
           ......
   }
   ```

   而关于`Level`，则是在每次调用`Log`方法时将`Level`传入：

   

1. `Level`的处理。为了方便，我打算直接在新建(`NewLogger`)方法中传参`string`类型的`level`，然后通过`switch`将这个`level`传给`logrus`，但这样做还是有问题的：

   1. 通过`switch`处理`level`时，当遇到不支持的`Level`时如何处理，即`default`怎么写？

      答：应该返回`error`，但由于逻辑是在`NewLogger()`方法中，这种新建的方法我认为不应该返回一个错误，所以应该把等级设置的逻辑从`NewLogger()`中取出。

      并且，可以不用自己`switch`。用`logrus`提供的`ParseLevel()`或者`Level.UnmarshalText()`方法更好。这个方法个估计就是给我这种懒汉准备的，`logrus`包内并没有用这个方法。

      ```go
      // 取自logrus.go
      // ParseLevel takes a string level and returns the Logrus log level constant.
      func ParseLevel(lvl string) (Level, error) {
      	switch strings.ToLower(lvl) {
      	case "panic":
      		return PanicLevel, nil
      	......
      	}
      
      	var l Level
      	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
      }
      
      // UnmarshalText implements encoding.TextUnmarshaler.
      func (level *Level) UnmarshalText(text []byte) error {
      	l, err := ParseLevel(string(text))
      	if err != nil {
      		return err
      	}
      
      	*level = l
      
      	return nil
      }
      ```

      参考：

      正经的`log`包的设计：

      1. `logrus`包内的处理方式：不处理。`logrus`在赋值`level`时，是通过`atomic.StoreUint32()`方法，直接将`Logger.Level`的指针指向入参`level`。反正`Logger.Level`是个`uint32`型，输出时只比大小，管你传什么数字进来。并且，一般来说也只会传一个`logrus.xxxLevel`进来，不会有这个问题。

### WithFields方法

## 阅读列表

- [ ] `logrus`的设计
- [ ] `glog`的设计

