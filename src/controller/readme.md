##### 路由配置规则

直接通过 function 绑定 controller，function name 映射到 路由的地址

function named rule：

```
func(*Controller) GetLoginBy(id int64)
绑定的controller HTTP-method+routePath
map to - GET:/user/login/{param:long}
```



通过控制器方法的输入参数访问动态路径参数，不需要绑定。当您使用Iris的默认语法来解析来自控制器的处理程序时，您需要使用By字来为方法添加后缀，大写是一个新的子路径。例： 如这种形式 mvc.New(app.Party("/user")).Handle(new(user.Controller)) 则:

- func(*Controller) Get() - GET:/user
- func(*Controller) Post() - POST:/user
- func(*Controller) GetLogin() - GET:/user/login
- func(*Controller) PostLogin() - POST:/user/login
- func(*Controller) GetProfileFollowers() - GET:/user/profile/followers
- func(*Controller) PostProfileFollowers() - POST:/user/profile/followers
- func(*Controller) GetBy(id int64) - GET:/user/{param:long}
- func(*Controller) PostBy(id int64) - POST:/user/{param:long}

Go 编程规范不用下划线分割单词
