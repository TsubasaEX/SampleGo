$(function(){
    $.ajax({
		url: "/reports",
		type: "GET",
        contentType: 'application/json',
        dataType: 'json',
		success: function(result){
			//var txt ='{"files":["./","20191126165004_rpt.csv","20191126165004_stats.csv","Config.yaml","doc.go","go.mod","go.sum","main.go"]}'
			//var obj = JSON.parse(txt);
			$("#demo").text(result.files[1] + ", " + result.files[2]);
//			var i;
//			var len =result.files.length
//			for (i=0;i<len;i++){
//				console.log(result.files[i]);
//			}
			var file;
			for (file of result.files){
				console.log(file);
				$("#reportBody")
					.append($('<tr>')
						.append($('<td>')
							.text('kkkkk')
						)
					)
			}
			//$("#reportBody").html
		
        },
        error: function(result){
			alert("Error: " + result.status + " " + result.statusText);
        }
    });
    $('#reportTable').DataTable();
	$('#statsTable').DataTable();
});