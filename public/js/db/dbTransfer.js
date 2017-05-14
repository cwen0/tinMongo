(function() {
    var dom = {
        dbTransferAll : $('#dbTransferAll'),
        dbCollections : $("form div:first input[name='collection']"), 
        transfer_btn: $("form").find("button"), 
        msg_alert: $("#msg_alert"), 
        msg_success: $("#msg_success")
    }

    var dbTransfer = {
        init : function() {
            this.eventFn();
        },

        eventFn : function() {
            dom.dbTransferAll.bind('click',function() {
                
                if (this.checked == true) {
                    dom.dbCollections.prop("checked",true);
                } else {
                    dom.dbCollections.prop("checked", false); 
                }
                
            });
            dom.transfer_btn.bind("click", function() {
                socket = $("#socket").val();
                host = $("#host").val();
                port = $("#port").val();
                isAuth = $("#isAuth").is(":checked");
                username = $("#username").val();
                password = $("#password").val();
                dbName = $("#dbName").val();
                isCopyIndex = $("#isCopyIndex").is(":checked");
                var collections = [] 
                $("input[name=collection]:checked").each(function() {
                    collections.push(this.value);
                }); 
                // console.log(collections);
                // return;
                if(collections.length <1 ) {
                    dom.msg_alert.find("p").text("Collection is required!!");
                    dom.msg_alert.show();
                    dom.msg_success.hide();
                    return 
                }
                if(isAuth) {
                    if(username == "") {
                        dom.msg_alert.find("p").text("Username is required!!");
                        dom.msg_alert.show();
                        dom.msg_success.hide();
                        return 
                    }
                    if(password == "") {
                        dom.msg_alert.find("p").text("Password is required!!");
                        dom.msg_alert.show();
                        dom.msg_success.hide();
                        return
                    }
                }
                $.ajax({
                    url: "/server/db/dbTransfer", 
                    type:"post", 
                    data: {
                        socket: socket, 
                        host: host, 
                        port: port, 
                        isAuth: isAuth, 
                        username: username, 
                        password: password, 
                        collections: collections,
                        dbName: dbName,
                        isCopyIndex: isCopyIndex,

                    }, 
                    success: function(data) {
                        dom.msg_success.find("p").text("Copy success!!");
                        dom.msg_success.show();
                        dom.msg_alert.hide();
                    }, 
                    error: function(response) {
                        data = JSON.parse(response.responseText);
                        dom.msg_alert.find("p").text(data["errors"][0]["title"]);
                        dom.msg_alert.show();
                        dom.msg_success.hide();
                    } 
                })
            }); 
            dom.msg_alert.find("button").bind("click", function() {
                dom.msg_alert.hide();
            });
            dom.msg_success.find("button").bind("click", function() {
                dom.msg_success.hide();
            })
        }
    }
    dbTransfer.init();
})();