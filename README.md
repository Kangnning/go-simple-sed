# go-simple-sed
使用go实现的简易文件处理sdk，目前初步版本准备可以根据上送的pattern来插入需要的内容

## quick start
```Go
    s := sed.New()
    conf := sed.Config{
        FileName:  "./QueryRoute.go",
        Opt:       sed.InsertBefore,
        Pattern:   "test.*",
        DesString: "This is inerst before test\n next line",
    }
    s.Run(conf)
```