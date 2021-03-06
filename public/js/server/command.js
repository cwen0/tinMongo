var command = new Vue({
	delimiters: ['${', '}'],
	el: "#command",
	data: {
		has_msg: false,
		alert_msg: "",
		show_result: false,
		result: null,
		item: {
			command: "",
			dbName: "admin",
			format: "json"
		},
        execCommandUrl: "/server/command"
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 4) }
	}, 
	methods: {
		execAction: function() {
			if(this.item.command == "")  {
				this.show_result = false;
				this.has_msg = true;
				this.alert_msg = "Command is required!!";
				return;
			}
			if(this.item.dbName == "") {
				this.show_result = false;
			    this.has_msg = true;
                this.alert_msg = "Database is required!!";
                return;
			}

            this.$http.post(this.execCommandUrl, this.item).then((response) => {
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
            this.show_result = false;
        }
	}
})
