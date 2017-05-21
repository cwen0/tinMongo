var insert = new Vue({
	delimiters: ['${', '}'],
	el: "#explain",
	data: {
		has_msg: false,
		alert_msg: "", 
		insert_success: false, 
		show_result: false,
		result: "",
		item: {
			query: "", 
			dbName: dbName, 
			collection: collection
		}, 
		explainUrl: "/server/collection/document/explain"
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

            this.$http.post(this.explainUrl, this.item).then((response) => {
            	//this.insert_success = true;
            	data = JSON.parse(response.bodyText);
            	this.result = data["datas"][0]["context"];
            	this.has_msg = false;
            	this.show_result = true;
            }).catch(this.requestError)
		},
		requestError: function(response) {
            data = JSON.parse(response.bodyText);
            this.alert_msg = data["errors"][0]["title"];
            this.has_msg = true;
            //this.insert_success = false;
            this.show_result = false;
        }
	}
})