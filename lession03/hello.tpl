<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <p>u1</p>
    {{/* 去除空格 */}}
    <p>Hello {{- .u1.Name -}}</p>
    <p>年龄:{{ .u1.Age }}</p>
    <p>性别:{{ .u1.Gender }}</p>

    {{/* 遇事不决,先写注释 */}}
    <p>m1</p>
    <p>{{.m1.name}}</p>
    <p>{{.m1.age}}</p>
    <p>{{.m1.gender}}</p>
    <hr />
    {{/* 定义变量 */}}
    {{$v1 := 100}}
    {{$age := .m1.name}}

    <hr />
    {{if $v1}}
    {{ $v1 }}
    {{else }}
    啥都没有
    {{ end }}
    <hr />
    {{if lt .m1.age 22}}
    好好上学
    {{ else }}
    好好工作
    {{ end }}
    <hr />
    {{range $idx,$hobby:= .hobby}}
    <p>idx:{{ $idx }} - {{ $hobby }}</p>
    {{else}}
    没啥爱好
    {{ end }}

    <hr />
    <p>m1</p>
    {{ with.m1 }}
    <p>{{.name}}</p>
    <p>{{.age}}</p>
    <p>{{.gender}}</p>
    {{ end }}

    <hr />
    {{index .hobby 2}}
  </body>
</html>
