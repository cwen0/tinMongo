(function() {
    var dom = {
        usersListA : $('#users-list-a'),
        usersList : $('#users-list'),
        addUserA : $('#add-user-a'),
        addUser : $('#add-user'),
        del_alert_msg: $("#del_alert_msg"),
        delete_user_btn: $(".am-icon-trash"), 
        delete_user_confirm: $("#delete_user_confirm"),
        delete_user_input: $("#delete_user_input"), 
        create_user_btn: $("#create_user_a"),
        create_user_confirm: $("#create_user_confirm"),
        create_alert_msg: $("#create_alert_msg"), 
        alert: $(".am-alert")
    }

    var true_delete_user = ""; 
    var can_del = false;
    var dbUsers = {
        init : function() {
            this.eventFn();
        },
        eventFn : function() {
            dom.usersListA.bind('click', function(e){
                e.preventDefault();  
                dom.usersList.show();
                dom.addUser.hide();
            }),

            dom.addUserA.bind('click', function(e) {
                e.preventDefault();
                dom.usersList.hide();
                dom.addUser.show();
            }), 

            dom.delete_user_btn.bind("click", function() {
                dom.del_alert_msg.hide();
                $("#delete_user").modal({
                    open: true,
                    relatedTarget: this,
                    closeOnConfirm: false,
                });
                dom.delete_user_input.val("");
                $("#delete_user").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
                true_delete_user = $(this).next().val();
            }), 

            dom.delete_user_input.on("change input propertychange blur", function() {
                modal_val = dom.delete_user_input.val();
                if(modal_val == true_delete_user) {
                    $("#delete_user").find(".data-am-modal-confirm").css("background-color", "");
                    can_del = true;
                } else {
                    can_del = false;
                }
            }), 

            dom.delete_user_confirm.bind("click", function() {
                if(can_del == false) {
                    return 
                }
                if(true_delete_user.length < 1) {
                    dom.del_alert_msg.find("p").text("Database Name is required!!");
                    dom.del_alert_msg.show();
                    return 
                }
                $.ajax({
                    url: "/server/db/dbUsers/" + $("#dbName").val() + "/" + true_delete_user + "/delete", 
                    type: "post",
                    success: function(data) {
                        $("#delete_user").modal("close");
                        location.reload();
                    }, 
                    error: function(response) {
                        data = JSON.parse(response.bodyText);
                        dom.del_alert_msg.find("p").text(data["errors"][0]["title"]);
                        dom.del_alert_msg.show();
                    }
                })
            }), 

            dom.create_user_btn.bind("click", function() {
                dom.create_alert_msg.hide();
                $("#create_user").modal({
                    open: true,
                    relatedTarget: this,
                    closeOnConfirm: false,
                    onConfirm: function(e) {
                        userName = $("#new_user").val();
                        if(userName == "") {
                            dom.create_alert_msg.find("p").text("User name is required!!");
                            dom.create_alert_msg.show();
                            return;
                        }
                        password1 = $("#password").val();  
                        if(password1 == "") {
                            dom.create_alert_msg.find("p").text("Password is required!!");
                            dom.create_alert_msg.show();
                            return;
                        }
                        password2 = $("#password2").val(); 
                        if(password2 != password1) {
                            dom.create_alert_msg.find("p").text("Confirm Password is wrong!!");
                            dom.create_alert_msg.show();
                            return;
                        }
                        isReadonly = $("input[name='isReadonly']:checked").val();
                        $.ajax({
                            url: "/server/db/newDBUser", 
                            data: {
                                dbName: $("#dbName").val(),
                                user: userName, 
                                password: password1, 
                                isReadonly: isReadonly == 1 ? true : false,
                            }, 
                            dataType: "json", 
                            type: "post", 
                            success: function(data) {
                                $("#create_user").modal('close');
                                location.reload();
                            }, 
                            error: function(response) {
                                data = JSON.parse(response.bodyText);
                                dom.create_alert_msg.find("p").text(data["errors"][0]["title"]);
                                dom.create_alert_msg.show();
                            }
                        })
                    }
                });
            }), 

            dom.alert.find("button").bind("click", function() {
                $(this).parent().hide();
            })
        }
    }
    dbUsers.init();
})()