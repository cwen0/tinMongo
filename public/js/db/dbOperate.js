(function() {
	var dom = {
		del_database_btn: $("#del_database_btn"), 
		del_database_input: $("#delete_database_input"),
		del_database_confirm: $("#delete_database_confirm"),
		del_alert_msg: $("#del_alert_msg"),
		// clear_database_btn: $("#clear_database_btn"), 
		// clear_database_input: $("#clear_database_input"),
		// clear_database_confirm: $("#clear_database_confirm"),
		// clear_alert_msg: $("#clear_alert_msg"),
		alert: $(".am-alert"), 
		del_collection_btn: $(".am-icon-trash"),
		del_collection_input: $("#delete_collection_input"),
		del_collection_confirm: $("#delete_collection_confirm"),
		del_coll_alert_msg: $("#del_coll_alert_msg")
	}
	var true_delete_database = "";
	var true_delete_collection = "";
	var can_del = false;
	var can_del_coll = false;
	// var true_clear_database= ""; 
	// var can_clear = false;
	var drop_db = {
		init: function() {
			this.eventFn();
		}, 
		eventFn: function() {
			dom.del_database_btn.bind("click", function() {
				dom.del_alert_msg.hide();
				$("#delete_database").modal({
                	open: true, 
                	relatedTarget: this, 
                	closeOnConfirm: false,
                });
                dom.del_database_input.val("");
                $("#delete_database").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
                true_delete_database = $("#dbName").val();
			}), 

			dom.del_database_input.on("change input propertychange blur", function() {
                modal_val = dom.del_database_input.val();
                if(modal_val == true_delete_database) {
                    $("#delete_database").find(".data-am-modal-confirm").css("background-color", "");
                    can_del = true;
                } else {
                    can_del = false;
                }
            }), 

			dom.del_database_confirm.bind("click", function() {
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
            }), 

			dom.alert.find("button").bind("click", function() {
                $(this).parent().hide();
            })
		}
	}

	var del_collection = {
		init: function() {
			this.eventFn();
		}, 
		eventFn: function() {
			dom.del_collection_btn.bind("click", function() {
				dom.del_coll_alert_msg.hide();
				$("#delete_collection").modal({
                	open: true, 
                	relatedTarget: this, 
                	closeOnConfirm: false,
                });
                dom.del_collection_input.val("");
                $("#delete_collection").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
                true_delete_collection = $(this).next().val();
			}), 

			dom.del_collection_input.on("change input propertychange blur", function() {
                modal_val = dom.del_collection_input.val();
                if(modal_val == true_delete_collection) {
                    $("#delete_collection").find(".data-am-modal-confirm").css("background-color", "");
                    can_del_coll = true;
                } else {
                    can_del_coll = false;
                }
            }), 

			dom.del_collection_confirm.bind("click", function() {
                if(can_del_coll == false) {
		              return
		         }
		        if(true_delete_collection.length < 1) {
		            dom.del_coll_alert_msg.find("p").text("Database Name is required!!");
		            dom.del_coll_alert_msg.show();
		            return 
		        }
		        $.ajax({
		            url : "/server/db/collection/" + $("#dbName").val() + "/" + true_delete_collection + "/delete",
		            type : "post",
		            success : function(data) {
		            $('#delete_collection').modal('close');
		               location.reload();
		            },
		            error : function(response) {
		                data = JSON.parse(response.bodyText);
		                dom.del_coll_alert_msg.find("p").text(data["errors"][0]["title"]);
		                dom.del_coll_alert_msg.show();
		         	}
		         });
            })
		}
	}

	// var clear_db = {
	// 	init: function() {
	// 		this.eventFn();
	// 	}, 
	// 	eventFn: function() {
	// 		dom.clear_database_btn.bind("click", function() {
	// 			dom.clear_alert_msg.hide();
	// 			$("#clear_database").modal({
 //                	open: true, 
 //                	relatedTarget: this, 
 //                	closeOnConfirm: false,
 //                });
 //                dom.clear_database_input.val("");
 //                $("#clear_database").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
 //                true_clear_database = $("#dbName").val();
	// 		}), 

	// 		dom.clear_database_input.on("change input propertychange blur", function() {
 //                modal_val = dom.clear_database_input.val();
 //                if(modal_val == true_clear_database) {
 //                    $("#clear_database").find(".data-am-modal-confirm").css("background-color", "");
 //                    can_clear = true;
 //                } else {
 //                    can_clear = false;
 //                }
 //            }), 

	// 		dom.clear_database_confirm.bind("click", function() {
 //                if(can_clear == false) {
	// 	              return
	// 	         }
	// 	        if(true_clear_database.length < 1) {
	// 	            dom.clear_alert_msg.find("p").text("Database Name is required!!");
	// 	            dom.clear_alert_msg.show();
	// 	            return 
	// 	        }
	// 	        $.ajax({
	// 	            url : "/server/db/dbOperate/" + true_clear_database + "/clear",
	// 	            type : "post",
	// 	            success : function(data) {
	// 	            $('#clear_database').modal('close');
	// 	               location.reload();
	// 	            },
	// 	            error : function(response) {
	// 	                data = JSON.parse(response.bodyText);
	// 	                dom.clear_alert_msg.find("p").text(data["errors"][0]["title"]);
	// 	                dom.clear_alert_msg.show();
	// 	         	}
	// 	         });
 //            })
	// 	}
	// }
	drop_db.init();
	//clear_db.init();
	del_collection.init();
})()