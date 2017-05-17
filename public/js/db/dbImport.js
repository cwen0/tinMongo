(function() {
    var dom = {
        jsFile : $('#js-file'),
        jsFileList : $('#js-file-list'),
        jsonFile : $('#json-file'),
        jsonFileList : $('#json-file-list'), 
        js_msg_alert: $("#js_msg_alert"),
        js_msg_success: $("#js_msg_success"),
        json_msg_alert: $("#json_msg_alert"), 
        json_msg_success: $("#json_msg_success"), 
        alert: $(".am-alert"), 
        js_import_btn: $("#js_import_btn"), 
        json_import_btn: $("#json_import_btn"), 
        dbName: $("#dbName")
    }

    var dbImport = {
        init : function() {
            this.eventFn();
        },
        eventFn : function () {
            dom.jsFile.bind('change',function() {
                var fileNames = '';
                $.each(this.files, function() {
                  fileNames += '<span class="am-badge">' + this.name + '</span> ';
                });
                dom.jsFileList.html(fileNames);
            });

            dom.jsonFile.bind('change',function() {
                var fileNames = '';
                $.each(this.files, function() {
                  fileNames += '<span class="am-badge">' + this.name + '</span> ';
                });
                dom.jsonFileList.html(fileNames);
            });

            dom.js_import_btn.bind("click", function() {
                var files = $('#js-file').prop('files');
                if(files.length == 0) {
                    dom.js_msg_alert.find("p").text("No files are selected!!");
                    dom.js_msg_alert.show();
                    dom.js_msg_success.hide();
                    return 
                }
                var formData = new FormData();
                formData.append("file", files[0]);
                formData.append("fileType", "js");
                formData.append("dbName", dom.dbName.val());
                $.ajax({
                    url: "/server/db/dbImport",
                    type: "post", 
                    data: formData,
                    processData : false,
                    contentType : false,
                    success: function(data) {
                        dom.js_msg_success.find("p").text("Import success!!");
                        dom.js_msg_alert.hide();
                        dom.js_msg_success.show();
                    }, 
                    error: function(response) {
                        data = JSON.parse(response.responseText);
                        dom.js_msg_alert.find("p").text(data["errors"][0]["title"]);
                        dom.js_msg_alert.show();
                        dom.js_msg_success.hide();
                    }
                })
            });

            dom.json_import_btn.bind("click", function() {
                var files = $("#json-file").prop("files");
                if(files.length == 0) {
                    dom.json_msg_alert.find("p").text("No files are selected!!");
                    dom.json_msg_alert.show();
                    dom.json_msg_success.hide();
                    return 
                }
                coll_name = $("#import_collection").val();
                if(coll_name == "") {
                    dom.json_msg_alert.find("p").text("Collection is required!!");
                    dom.json_msg_alert.show();
                    dom.json_msg_success.hide();
                    return 
                }
                var formData = new FormData();
                formData.append("file", files[0]); 
                formData.append("fileType", "json");
                formData.append("coll_name", coll_name);
                $.ajax({
                    url: "/server/db/dbImport", 
                    type: "post", 
                    data: formData, 
                    processData : false,
                    contentType : false,
                    success: function(data) {
                        dom.json_msg_success.find("p").text("Import success!!");
                        dom.json_msg_alert.hide();
                        dom.json_msg_success.show();
                    }, 
                    error: function(response) {
                        data = JSON.parse(response.responseText);
                        dom.json_msg_alert.find("p").text(data["errors"][0]["title"]);
                        dom.json_msg_alert.show();
                        dom.json_msg_success.hide();
                    }
                })
            });

            dom.alert.find("button").bind("click",function() {
                $(this).parent().hide();
            }); 
        }
    }

    dbImport.init();
})()