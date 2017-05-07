{{define "Tpl/navigation"}}
<div>
  <ul class="am-nav am-nav-tabs">
    <li {{if equal .active "home"}}class="am-active" {{end}}><a href="/server/home">服务器</a></li>
    <li {{if equal .active "status"}}class="am-active" {{end}}><a href="/server/status">状态</a></li>
    <li {{if equal .active "databases"}}class="am-active" {{end}}><a href="/server/databases">数据库</a></li>
    <li {{if equal .active "processList"}}class="am-active" {{end}}><a href="/server/processList">进程</a></li>
    <li {{if equal .active "command"}}class="am-active" {{end}}><a href="/server/command">命令</a></li>
    <li {{if equal .active "execute"}}class="am-active" {{end}}><a href="/server/execute">代码执行</a></li>
    <li {{if equal .active "replication"}}class="am-active" {{end}}><a href="/server/replication">主/从</a></li>
    <!-- <li class="am-dropdown" data-am-dropdown>
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
