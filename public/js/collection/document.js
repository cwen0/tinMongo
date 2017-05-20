Vue.component('row-result', {
    template: `
    <div class="am-panel am-panel-default admin-sidebar-panel">
        <div class="am-panel-bd am-g">
    		<div class="am-u-sm-10">
    			<pre>{{ title }}</pre>
    		</div> 
    		<div class="am-u-sm-2" >
    			<a v-on:click="$emit('edit')" >
                    <i class="am-icon-pencil am-icon-fw"></i>
                </a>
                <a v-on:click="$emit('remove')" >
                    <i class="am-icon-trash am-icon-fw"></i>
                </a>
    		</div>
    	</div>
    </div>
  	`,
  	props: ['title']
})
var doc = new Vue ({
	delimiters: ['${', '}'],
	el: "#document",
	data: {
		has_msg: false,
		alert_msg: "",
		show_result: false,
		result: null,
		resultRow: null,
		del_row:null,
		rows: [],
		item: {
			query: "",
			dbName: dbName, 
			collection: collection
		}, 
		edit_row: null,
		edit_index: null,
		edit_alert_msg:"", 
		edit_has_msg: false,
		edit_send: {
			id: null, 
			edited: null,
			dbName: dbName, 
			collection: collection
		}, 
		// del_alert_msg: "",
		// del_has_msg: false,
		queryUrl: "/server/collection/document/query", 
		editUrl: "/server/collection/document/update"
		//deleteUrl: "/server/collection/document/delete"
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 4) }
	}, 
	methods: {
		execAction: function() {
			this.rows = [];
			if(this.item.query == "") {
				this.show_result = false; 
				this.has_msg = true; 
				this.alert_msg = "Query is required!!"; 
				return;
			}

			if(this.item.dbName == "") {
				this.show_result = false;
			    this.has_msg = true;
                this.alert_msg = "Something is wrong!!";
                return;
			}

			if(this.item.collection == "") {
				this.show_result = false;
			    this.has_msg = true;
                this.alert_msg = "Something is wrong!!";
                return;
			}

			this.$http.post(this.queryUrl, this.item).then((response) => {
            	data = JSON.parse(response.bodyText);
            	this.result = data["datas"][0]["context"];
            	this.resultRow = this.result.length;
            	for(var i = 0; i < this.resultRow; i++) {
            		this.rows.push(this.result[i]);
            	}
            	this.has_msg = false;
            	this.show_result = true;
            }).catch(this.requestError)
		}, 
		requestError: function(response) {
            data = JSON.parse(response.bodyText);
            this.alert_msg = data["errors"][0]["title"];
            this.has_msg = true;
            this.show_result = false;
        }, 
        collectionRemove: function(index, row) {
        	this.del_row = row;
        	jQuery("#del-confirm").modal({
        		relatedTarget: this,
            	closeOnConfirm: false,
            	closeOnCancel: true, 
            	onConfirm: function(e) {
            		jQuery.ajax({
            			url: "/server/collection/document/delete", 
            			dataType: "json", 
            			type: "post", 
            			data: {
            				rowID: row._id, 
            				dbName: dbName, 
            				collection: collection
            			}, 
            			success: function(data) {
            				jQuery("#del-confirm").modal('close');
            				// location.reload();
            				jQuery("#exec_btn").trigger("click");  
            			}, 
            			error: function(response) {
            				data = JSON.parse(response.bodyText);
                      		// this.del_alert_msg = data["errors"][0]["title"]; 
                      		// this.del_has_msg = true;
                      		// return;
                      		alert(data["errors"][0]["title"]);
            			}
            		})
            	}
        	})
    		// this.rows.splice(index, 1);
      //   	this.resultRow -= 1;
        }, 

        collectionEdit: function(index, row) {
        	this.edit_row = row;
        	this.edit_index = index;
        	jQuery("#edit-confirm").modal({
        		relatedTarget: this, 
        		closeOnConfirm: false, 
        		closeOnCancel: true
        	})
        }, 

        editSave: function() {
        	edited = jQuery("#edit_row").val();
        	editedT = JSON.parse(edited);
        	delete editedT._id;
        	
        	this.edit_send.edited = JSON.stringify(editedT);

        	this.edit_send.id = this.edit_row._id;
        	if(edited == "") {
        		this.edit_alert_msg = "Please, fill out form corrently!!";
        		this.edit_has_msg = true;
        	}


			this.$http.post(this.editUrl, this.edit_send).then((response) => {
            	jQuery("#edit-confirm").modal('close');
            	//this.rows[this.edit_index] = JSON.parse(edited);
            	//console.log(this.rows);
            	jQuery("#exec_btn").trigger("click");  
            }).catch(function(response) {
            	console.log(response);
            	data = JSON.parse(response.bodyText);
            	this.edit_alert_msg = data["errors"][0]["title"];
            	this.edit_has_msg = true;
            })
        }

	}
})