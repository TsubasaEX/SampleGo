$(function(){
    $.ajax({
		url: "/reports",
		type: "GET",
        contentType: 'application/json',
        dataType: 'json',
		success: function(result){
			try {
				var t = $('#reportTable').DataTable();
				var i;
				var len = result.files.length
				
				for (i=0;i<len;i++){
					var file = result.files[i][0];
					var href = "/reports/results?file="+file;
					var icon = "<a href="+href+"><i class=\"fas fa-dragon\" style=\"font-size:16px;color:blue;\"></i></a>";
					t.row.add( [
			            icon,
			            i+1,
			            file,
			            result.files[i][1],
			        ] ).draw( false );
				}
	//			var i;
	//			var len = result.files.length
	//			if(len != 0){
	//				$(".dataTables_empty").remove();
	//			}
	//			for (i=0;i<len;i++){
	//				$("#reportBody")
	//					.append($('<tr>')var file =result.files[i][0]	
	//						.append($('<td>')
	//							.text(result.files[i][0])
	//						)
	//						.append($('<td>')
	//							.text(i+1)
	//						)
	//						.append($('<td>')
	//							.text(result.files[i][0])
	//						).append($('<td>')
	//							.text(result.files[i][1])
	//						)
	//					)
	//			}
			} catch(e){
				console.log(e)
			}
		
        },
        error: function(result){
			alert("Error: " + result.status + " " + result.statusText);
        }
    });
    $.ajax({
		url: "/statistics",
		type: "GET",
        contentType: 'application/json',
        dataType: 'json',
		success: function(result){
			try {
				var t = $('#statsTable').DataTable();
				var i;
				var len = result.files.length
				
				for (i=0;i<len;i++){
					var file = result.files[i][0];
					var href = "/statistics/results?file="+file;
					var icon = "<a href="+href+"><i class=\"fas fa-pastafarianism\" style=\"font-size:16px;color:green;\"></i></a>";
					t.row.add( [
			            icon,
			            i+1,
			            file,
			            result.files[i][1],
			        ] ).draw( false );
				}
	//			var i;
	//			var len = result.files.length
	//			if(len != 0){
	//				$(".dataTables_empty").remove();
	//			}
	//			for (i=0;i<len;i++){
	//				$("#reportBody")
	//					.append($('<tr>')var file =result.files[i][0]	
	//						.append($('<td>')
	//							.text(result.files[i][0])
	//						)
	//						.append($('<td>')
	//							.text(i+1)
	//						)
	//						.append($('<td>')
	//							.text(result.files[i][0])
	//						).append($('<td>')
	//							.text(result.files[i][1])
	//						)
	//					)
	//			}
			} catch(e){
				console.log(e)
			}
		
        },
        error: function(result){
			alert("Error: " + result.status + " " + result.statusText);
        }
    });
    $('#reportTable').DataTable();
	$('#statsTable').DataTable();
});