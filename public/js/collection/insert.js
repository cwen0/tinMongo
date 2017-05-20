var insert = new Vue({
	delimiters: ['${', '}'],
	el: "#document",
	data: {
		has_msg: false,
		alert_msg: ""
	},
	filters: {
		json: (value) => { return JSON.stringify(value, null, 4) }
	},
	methods: {

	}
})