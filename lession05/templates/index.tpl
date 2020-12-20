{{/*继承新模版*/}}
{{template "base.tpl"}}
{{/* 重新定义快模版 */}}
{{define "content"}}
<h1>这是index页面</h1>
<p>Hello {{ . }}</p>
{{ end }}
