var indexs = new Vue({
	delimiters: ['${', '}'],
	el: "#indexs", 
	data: {
		true_del_name: null,
		del_name: null,
		can_del: false,
		has_del_msg: false, 
		del_alert_msg: null,
		del_item: {
			dbName: dbName, 
			collection: collection,
			index: null
		}, 
		deleteURL: "/server/collection/document/indexs/delete"
	}, 
	methods: {
		deleteAction: function(name) {
			this.true_del_name = name;
			this.can_del = false;
			this.del_name = "";
			jQuery("#delete_index").modal("open");
			jQuery("#delete_index").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
		}, 
		doDeleteAction: function() {
			if(this.can_del == false) {
				return 
			} 
			this.del_item.index = this.true_del_name;
			this.$http.post(this.deleteURL, this.del_item).then((response) => {
            	//this.insert_success = true;
            	jQuery("#delete_index").modal("close");
            	location.reload();
            }).catch(function(response) {
            	data = JSON.parse(response.bodyText);
            	this.del_alert_msg = data["errors"][0]["title"];
            	this.has_del_msg = true;
            })
		}, 
		addAction: function() {
			
		}
	}, 
	watch: {
		"del_name": function(val) {
			if(val == this.true_del_name) {
				jQuery("#delete_index").find(".data-am-modal-confirm").css("background-color", "");
				this.can_del = true;
			}else {
				jQuery("#delete_index").find(".data-am-modal-confirm").css('background-color', '#f2f2f2');
				this.can_del = false;
			}
		}
	}
})