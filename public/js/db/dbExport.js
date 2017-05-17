(function() {
    var dom = {
        dbExportAll : $('#dbExportAll'),
        dbCollections : $("form div:first input[name='collection']"), 
        export_btn: $("#export_btn"), 
        msg_alert: $("#msg_alert"), 
        msg_success: $("msg_success"),
        alert: $(".am-alert")
    }

    var dbTransfer = {
        init : function() {
            this.eventFn();
        },

        eventFn : function() {
            dom.dbExportAll.bind('click',function() {
                
                if (this.checked == true) {
                    dom.dbCollections.prop("checked",true);
                } else {
                    dom.dbCollections.prop("checked", false); 
                }
                
            });
            dom.export_btn.bind("click", function() {
                dbName = $("#dbName").val();
                isDownload = $("#isDownload").is(":checked");
                isGzip = $("#isGzip").is(":checked");
                var collections = [] 
                $("input[name=collection]:checked").each(function() {
                    collections.push(this.value);
                })
                if(collections.length < 1) {
                    dom.msg_alert.find("p").text("Collection is required!!");
                    dom.msg_alert.show();
                    dom.msg_success.hide();
                    return 
                }
                //console.log(collections);
                // console.log(dbName);
                colls =  collections.join(",");
                $.ajax({
                    url: "/server/db/dbExport",
                    type: "post", 
                    data: {
                        isDownload: isDownload, 
                        isGzip: isGzip, 
                        colls: colls,
                        dbName: dbName,
                    }, 
                    success: function(data) {
                        //console.log(data);
                        //data = JSON.parse(data.responseText);
                        text = data["datas"][0]["context"];
                        //window.location = "data:text/plain;charset=UTF-8,"+text;
                        if(data["datas"][0]["type"] == "gzip") {
                            document.location = 'data:Application/x-gzip,' + encodeURIComponent(text);
                        } else {
                            document.location = 'data:Application/octet-stream,' + encodeURIComponent(text);
                        }
                        
                    }, 
                    error: function(response) {
                        data = JSON.parse(response.responseText);
                        dom.msg_alert.find("p").text(data["errors"][0]["title"]);
                        dom.msg_alert.show();
                        dom.msg_success.hide();
                    }
                }) 
            });

            dom.alert.find("button").bind("click",function() {
                $(this).parent().hide();
            }); 

        }
    }
    dbTransfer.init();
})();