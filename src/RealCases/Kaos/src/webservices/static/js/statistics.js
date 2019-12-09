$(document).ready( function () {
	$('#statsTable > tbody > tr').each(function(index, tr) { 
		$(this).find('td').each(function(index, td) {
			console.log(index);
			console.log(td);
		});
	});
	
      $('#statsTable').DataTable();
} );

//window.onload = function(){
//alert("onload");
//};

$(window).on('load', function () {
	 
//	$('#statsTable > tbody > tr').each(function(index, tr) { 
////		alert(this);
//		var name = $(this).find('#name').html();
//  		console.log(index);
//// 		console.log(name);
//		console.log(tr);
//	});
	
	//Better to construct options first and then pass it as a parameter
	var options = {
		animationEnabled: true,
		theme: "light2", //"light1", "dark1", "dark2"
		title:{
			text: "Division of Products Sold in 2nd Quarter"             
		},
		axisX:{
			interval: 1
		},
		axisY:{
			interval: 10,
			suffix: "%"
		},
		toolTip:{
			shared: true
		},
		data:[{
			type: "stackedBar100",
			toolTipContent: "{label}<br><b>{name}:</b> {y} (#percent%)",
			showInLegend: true, 
			name: "April",
			dataPoints: [
				{ y: 550, label: "Water Filter" },
				{ y: 450, label: "Modern Chair" },
				{ y: 70, label: "VOIP Phone" },
				{ y: 200, label: "Microwave" },
				{ y: 70, label: "Water Filter" },
				{ y: 324, label: "Expresso Machine" },
				{ y: 300, label: "Lobby Chair" },
		{ y: 500, label: "Lobby Chair1" },
		{ y: 500, label: "Lobby Chair2" },
		{ y: 500, label: "Lobby Chair3" },
		{ y: 500, label: "Lobby Chair4" },
		{ y: 500, label: "Lobby Chair5" },
		{ y: 500, label: "Lobby Chair6" },
		{ y: 500, label: "Lobby Chair7" }
			]
		},
		{
			type: "stackedBar100",
			toolTipContent: "<b>{name}:</b> {y} (#percent%)",
			showInLegend: true, 
			name: "May",
			dataPoints: [
				{ y: 450, label: "Water Filter" },
				{ y: 550, label: "Modern Chair" },
				{ y: 270, label: "VOIP Phone" },
				{ y: 400, label: "Microwave" },
				{ y: 270, label: "Water Filter" },
				{ y: 524, label: "Expresso Machine" },
				{ y: 500, label: "Lobby Chair" },
		{ y: 500, label: "Lobby Chair1" },
		{ y: 500, label: "Lobby Chair2" },
		{ y: 500, label: "Lobby Chair3" },
		{ y: 500, label: "Lobby Chair4" },
		{ y: 500, label: "Lobby Chair5" },
		{ y: 500, label: "Lobby Chair6" },
		{ y: 500, label: "Lobby Chair7" }
			]
		}, 
		{
			type: "stackedBar100",
			toolTipContent: "<b>{name}:</b> {y} (#percent%)",
			showInLegend: true, 
			name: "June",
			dataPoints: [
				{ y: 350, label: "Water Filter" },
				{ y: 660, label: "Modern Chair" },
				{ y: 265, label: "VOIP Phone" },
				{ y: 271, label: "Microwave" },
				{ y: 125, label: "Water Filter" },
				{ y: 360, label: "Expresso Machine" },
				{ y: 340, label: "Lobby Chair" },
		{ y: 500, label: "Lobby Chair1" },
		{ y: 500, label: "Lobby Chair2" },
		{ y: 500, label: "Lobby Chair3" },
		{ y: 500, label: "Lobby Chair4" },
		{ y: 500, label: "Lobby Chair5" },
		{ y: 500, label: "Lobby Chair6" },
		{ y: 500, label: "Lobby Chair7" }
			]
		}]
	};
	
	$("#chartContainer").CanvasJSChart(options);
});

//Deprecated
//$(window).load(function(){
//	alert("onload");
//});