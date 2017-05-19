var doc = new Vue ({
	delimiters: ['${', '}'],
	el: "#document",
	data: {
		has_msg: false,
		alert_msg: "",
		show_result: false,
		result: null,
		resultRow: null,
		item: {
			query: "",
			dbName: dbName, 
			collection: collection
		}, 
		queryUrl: "/server/collection/query"
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 4) }
	}, 
	methods: {
		execAction: function() {
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
            	this.has_msg = false;
            	this.show_result = true;
            }).catch(this.requestError)
		}, 
		requestError: function(response) {
            data = JSON.parse(response.bodyText);
            this.alert_msg = data["errors"][0]["title"];
            this.has_msg = true;
            this.show_result = false;
        }
	}
})