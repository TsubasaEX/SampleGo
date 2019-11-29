$(function(){
    $.ajax({
		url: "/reports",
		type: "GET",
		success: function(result){
			$("#demo").innerHTML=result;
        },
        error: function(result){
			alert("Error: " + result.status + " " + result.statusText);
        }
    });
});