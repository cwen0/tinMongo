{{define "Tpl/collectionNavigationTwo"}}
<div>
  <ul class="am-nav am-nav-tabs">
    <li  {{if equal .active "home" }}class="am-active" {{end}}><a href="/server/collection/home/{{ .DBName }}/{{.Collection}}">查询</a></li>
    <li {{if equal .active "insert" }}class="am-active" {{end}}><a href="/server/collection/document/insert/{{.DBName}}/{{.Collection}}">插入</a></li>
    <li {{if equal .active "indexs" }}class="am-active" {{end}}><a href="/server/collection/document/insert/{{.DBName}}/{{.Collection}}">主键</a></li>
    <li {{if equal .active "explain" }}class="am-active" {{end}}><a href="/server/collection/document/explain/{{.DBName}}/{{.Collection}}">查询分析</a></li>
  </ul>
</div>
{{end}}
