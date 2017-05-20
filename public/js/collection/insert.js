var insert = new Vue({
	delimiters: ['${', '}'],
	el: "#insert",
	data: {
		has_msg: false,
		alert_msg: "", 
		insert_success: false, 
		item: {
			query: "", 
			dbName: dbName, 
			collection: collection
		}, 
		insertUrl: "/server/collection/document/insert"
	},
	methods: {
		execAction: function(){
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
            }

            this.$http.post(this.insertUrl, this.item).then((response) => {
            	this.has_msg = false;
            	this.insert_success = true;
            }).catch(this.requestError)
		},
		requestError: function(response) {
            data = JSON.parse(response.bodyText);
            this.alert_msg = data["errors"][0]["title"];
            this.has_msg = true;
            this.insert_success = false;
        }
	}
})