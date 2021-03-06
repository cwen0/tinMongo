{{define "Tpl/dbNavigationOne"}}
<div class="me-db-navigation">
    <a href="/server/home"><span class="am-icon-globe">&nbsp;数据库</span></a>&nbsp;
    <span class="am-icon-angle-double-right"></span>&nbsp;
    <a href="/server/db/home/{{.DBName}}"><span class="am-icon-database" id="nav_dbName">&nbsp;{{ .DBName }}</span></a>&nbsp;
    <span class="am-icon-angle-double-right"></span>&nbsp;
    
	{{ if .Collection }} 
		<a href="/server/collection/home/{{.DBName}}/{{.Collection}}"><span class="am-icon-table">&nbsp;{{ .Collection}}</span></a>
		<span class="am-icon-angle-double-right"></span>&nbsp;
	{{ end }}
    <span>{{ .navigation}}</span>
  </div>
{{end}}
