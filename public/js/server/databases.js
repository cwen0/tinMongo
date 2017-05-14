(function(){ 
     var dom = {
       alert_msg: $("#alert_msg"), 
       del_alert_msg: $("#del_alert_msg"),
       delete_database_btn: $(".am-icon-trash"),
       delete_database_input: $("#delete_database").find(".am-modal-prompt-input"),
       delete_database_confirm: $("#delete_database_confirm")
     }

     dom.alert_msg.hide();

     dom.alert_msg.find("button").on('click', function() {
        dom.alert_msg.hide();
     });
     var true_delete_database = "";
     var can_del = false;
     var del = {

        init: function() {
          this.eventFn();
        }, 
        eventFn: function() {
          dom.delete_database_btn.on("click", function() {
              dom.del_alert_msg.hide();
              $("#delete_database").modal({
                  open: true,
                  relatedTarget: this,
                  closeOnConfirm:false,
              });
              dom.delete_database_input.val("");
              $("#delete_database").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
              true_delete_database = $(this).next().val();
          })

          dom.delete_database_input.on("change input propertychange blur", function() {
             modal_val = dom.delete_database_input.val();
             if(modal_val == true_delete_database) {
               $("#delete_database").find(".data-am-modal-confirm").css("background-color", "");
                can_del = true;
             }else {
                can_del = false;
             }
          })

          dom.delete_database_confirm.on("click", function(){
            if(can_del == false) {
              return
            }
            if(true_delete_database.length < 1) {
                dom.del_alert_msg.find("p").text("Database Name is required!!");
                dom.del_alert_msg.show();
                return 
            }
            $.ajax({
              url : "/server/database" + "/" + true_delete_database + "/delete",
              type : "post",
              success : function(data) {
                $('#delete_database').modal('close');
                location.reload();
              },
              error : function(response) {
                data = JSON.parse(response.bodyText);
                dom.del_alert_msg.find("p").text(data["errors"][0]["title"]);
                dom.del_alert_msg.show();
              }
            });
         })


        }
     }

     del.init();

    
     $('#add_database_btn').on('click', function() {
          var $promt = $('#add_database').modal({
            relatedTarget: this,
            closeOnConfirm: false,
            closeOnCancel: true,
            onConfirm: function(e) {
                databaseName = $("#new_database_name").val();
                collectionName = $("#new_collection_name").val();
                if(databaseName.length  < 1) {
                    dom.alert_msg.find("p").text("Database Name is required!!");
                    dom.alert_msg.show();
                    return 
                }
                if(collectionName.length < 1 ) {
                    dom.alert_msg.find("p").text("Collection Name is required!!");
                    dom.alert_msg.show();
                    return 
                }
                $.ajax({
                    url : "/server/newDatabase",
                    data : {
                        databaseName: databaseName, 
                        collectionName: collectionName
                    },
                    dataType : "json",
                    type : "post",
                    success : function(data) {
                      $('#add_database').modal('close');
                      location.reload();
                    },
                    error : function(response) {
                      data = JSON.parse(response.bodyText);
                      dom.alert_msg.find("p").text(data["errors"][0]["title"]);
                      dom.alert_msg.show();
                    }
                  });
              }
            });
        });
  })();