{{define "Tpl/dbNavigationTwo"}}
<div>
  <ul class="am-nav am-nav-tabs">
    <li  {{if equal .active "home" }}class="am-active" {{end}}><a href="/server/db/home/{{ .DBName }}">统计</a></li>
    <li {{if equal .active "newCollection" }}class="am-active" {{end}}><a href="/server/db/newCollection/{{.DBName}}">创建集合</a></li>
    <li><a href="/server/command">命令</a></li>
    <li><a href="/server/execute">执行代码</a></li>
    <li {{if equal .active "dbTransfer" }}class="am-active" {{end}}><a href="/server/db/dbTransfer/{{.DBName}}">克隆</a></li>
    <li {{if equal .active "dbExport" }}class="am-active" {{end}}><a href="/server/db/dbExport/{{.DBName}}">导出</a></li>
    <li {{if equal .active "dbImport" }}class="am-active" {{end}} ><a href="/server/db/dbImport/{{.DBName}}">导入</a></li>
    <li {{if equal .active "dbUsers" }}class="am-active" {{end}} ><a href="/server/db/dbUsers/{{.DBName}}">用户</a></li>
    <li {{if equal .active "dbOperate" }}class="am-active" {{end}}  ><a href="/server/db/dbOperate/{{.DBName}}">操作</a></li>
<!--     <li class="am-dropdown" data-am-dropdown>
      <a class="am-dropdown-toggle" data-am-dropdown-toggle href="javascript:;">
        菜单 <span class="am-icon-caret-down"></span>
      </a>
      <ul class="am-dropdown-content">
        ...
      </ul>
    </li> -->
  </ul>
</div>
{{end}}
