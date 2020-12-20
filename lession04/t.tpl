<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <h1>测试嵌套template语法</h1>
    <hr />
    {{template "ul.tpl"}}
    <hr />
    {{template "ol.tpl"}}
  </body>
</html>

{{/* 通过define定义一个模版 */}}
{{define "ol.tpl"}}
<ol>
  <li>吃饭</li>
  <li>睡觉</li>
  <li>打游戏</li>
</ol>
{{ end }}
