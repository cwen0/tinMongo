{{define "Tpl/sidebar"}}
<div class="am-cf admin-main" id = "adminMain">
  <!-- sidebar start -->
  <div class="admin-sidebar am-offcanvas" id="admin-offcanvas">
    <div class="am-offcanvas-bar admin-offcanvas-bar">
    <div class="me-sidebar">
      <ul class="am-tree" id="firstTree">
        <li class="am-tree-branch am-hide" data-template="treebranch">
          <div class="am-tree-branch-header">
            <button class="am-tree-branch-name">
              <span class="am-tree-icon am-tree-icon-folder"></span>
              <span class="am-tree-label"></span>
            </button>
          </div>
          <ul class="am-tree-branch-children"></ul>
          <div class="am-tree-loader"><span class="am-icon-spin am-icon-spinner"></span></div>
        </li>
        <li class="am-tree-item am-hide" data-template="treeitem">
          <button class="am-tree-item-name">
            <span class="am-tree-icon am-tree-icon-item"></span>
            <span class="am-tree-label"></span>
          </button>
        </li>
      </ul>
    </div>
      <div class="am-panel am-panel-default admin-sidebar-panel">
        <div class="am-panel-bd">
          <p><span class="am-icon-bookmark"></span> 公告</p>
          <p>时光静好，与君语；细水流年，与君同。—— CWen</p>
        </div>
      </div>
      <div class="am-panel am-panel-default admin-sidebar-panel">
        <div class="am-panel-bd">
          <p><span class="am-icon-tag"></span> wiki</p>
          <p>Welcome to the TinMongo wiki!</p>
        </div>
      </div>
    </div>
  </div>

  <!-- sidebar end -->

<script type="text/javascript" src="/public/js/jquery-3.2.1.js"></script>
<script type="text/javascript" src="/public/js/amazeui.js"></script>
<script type="text/javascript" src="/public/js/amazeui.tree.min.js"></script>
<script type="text/javascript" src="/public/js/app.js"></script>
<script>

  var dbNames = {{.DBNames}};
  var data = []; 
  for (var i = 0; i < dbNames.length; i++) {
      data.push({
        title: dbNames[i],
        type:"folder",
        products: []
      });
  }
  //console.log({{.ServerCmdLineOpts}})
  $('#firstTree').tree({
    dataSource: function(options, callback) {
      // 模拟异步加载
      setTimeout(function() {
        callback({data: options.products || data});
      }, 200);
    },
    folderIcon:'database',
    itemIcon:'table',
    multiSelect: false,
    cacheItems: true,
    folderSelect: false
  });
</script>
{{end}}
