// let router = new VueRouter();
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
			port: 27017, 
			database: "admin",
			isAuth: 0 ,
		},
		loginUrl: "/login"
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 4) }
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
					this.msg = "Username is required! "	;		
					return ;
				}
				if(this.item.password == null) {
					this.has_msg = true; 
					this.msg = "Password is required! "	;
					return ;
				}
			}
			if(this.selected == "1") {
				this.item.isAuth = 1;
			} else {
				this.item.isAuth = 0;
			}
            this.$http.post(this.loginUrl, this.item)
                .then((response) => {
                //router.redirect("/server/home")
                location.href = "/server/home";
            }).catch(this.requestError)
		},
		requestError: function(response) {
			//console.log(response);
			data = JSON.parse(response.bodyText);
			this.msg = data["errors"][0]["title"];
			this.has_msg = true;
		}
	}
});
