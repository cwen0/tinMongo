var login = new Vue({
	delimiters: ['${', '}'],
	el: "#login-auth",
	data: {
		msg: '',
		has_msg: false,
		open: false,
		selected: "0",
		item: {
			hostname: "localhost",
			port: 27017
		},
		loginUrl: "/login"
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 2) }
	}, 
	watch: {
		selected: function(val) {
			 this.open = val == "1" ? true : false;
		}
	},
	methods: {
		loginAction: function() {
			if(this.selected == 1 ) {
				if(this.item.username == null) {
					this.has_msg = true; 
					this.msg = "Username is required! "			
					return 
				}
				if(this.item.password == null ) {
					this.has_msg = true; 
					this.msg = "Password is required! "	
					return 
				}
			}
            this.$http.post(this.loginUrl, this.item)
                .then((response) => {
                
            }).catch(this.requestError)
		},
		requestError: function(response) {
				this.msg = response.json()
		}
	}
});
