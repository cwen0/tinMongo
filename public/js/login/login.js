var login = new Vue({
	
	el: "#login-auth",
	data: {
		open: false, 
		selected: "0"
	}, 
	watch: {
		selected: function(val) {
			 this.open = val == "1" ? true : false;
		}
	}
}); 