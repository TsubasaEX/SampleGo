var passPoints =[];
var failPoints =[];

$(document).ready( function () {
	$('#statsTable > tbody > tr').each(function(index, tr) { 
		var name=$(this).find('td:eq(1)').html();
		var pass=$(this).find('td:eq(2)').html();
		var fail=$(this).find('td:eq(3)').html();
		var passObj={
			y:parseInt(pass,10),
			label:name
		}
		var failObj={
			y:parseInt(fail,10),
			label:name
		}
		passPoints.push(passObj);
		failPoints.push(failObj);
//		$(this).find('td').each(function(index, td) {
//			switch(index){
//				case 0:
//					$(this).attr("id","no");
//					break;
//				case 1:
//					$(this).attr("id","name");
//					var name = $(this).find("#name").html();
//					map.set("label",name);
//					console.log(map);
//					passPoints.push(map) 
//					break;
//				case 2:
//					$(this).attr("id","pass");
//					break;
//				case 3:
//					$(this).attr("id","fail");
//					break;
//				case 4:
//					$(this).attr("id","total");
//					break;
//				case 5:
//					$(this).attr("id","passRate");
//					break;
//				case 6:
//					$(this).attr("id","ts");
//					break;
//			}
//		});
	});

    $('#statsTable').DataTable();
} );

//window.onload = function(){
//alert("onload");
//};

$(window).on('load', function () {
	console.log(passPoints);
	console.log(failPoints);
	//Better to construct options first and then pass it as a parameter
	var options = {
		animationEnabled: true,
		theme: "light1", //"light1", "dark1", "dark2"
		title:{
			text: "Test results of Kaos"             
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
			color: "LightSeaGreen",
			toolTipContent: "{label}<br><b>{name}:</b> {y} (#percent%)",
			showInLegend: true, 
			name: "Pass",
			dataPoints: passPoints
		},
		{
			type: "stackedBar100",
			color: "crimson",
			toolTipContent: "<b>{name}:</b> {y} (#percent%)",
			showInLegend: true, 
			name: "Fail",
			dataPoints: failPoints
		}]
	};
	
	$("#chartContainer").CanvasJSChart(options);
});

//Deprecated
//$(window).load(function(){
//	alert("onload");
//});