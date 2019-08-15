var connectionMade = false;
setInterval(function(){
    if(connectionMade){
        connection.send(JSON.stringify({content:"</Alive>"}));
    }
},1000)


var connection = new WebSocket('ws://final160-alexanderpark.codeanyapp.com:2323/commun');
connection.onopen = function(){
    connectionMade=true;
    connection.send("client");
}
connection.onmessage = function(receivedMessage){
    let packet = JSON.parse(receivedMessage.data)
    //packet["text"]
    if(packet["statusOfOther"]=="dead"){
    }else{
        let arrayOfLinks=packet["content"].split(",")

        "<img src='"+arrayOfLinks[0]+"'"
    }


}



// Add floorplan images as the background
$(".bath").css( "background-color", "#020a01");
$(".designBoard>.tab").css("visibility","hidden");
$("#bath").css("visibility","visible");
$("ul>li").click(function() {
    $("ul>li").css( "background-color", "#c1c1c1");
    $(this).css( "background-color", "#020a01");
    console.log(this.className);
    $(".designBoard>.tab").css("visibility","hidden");
    $("#"+ this.className).css("visibility","visible");
    currentDiv = this.className;
});

// Add floorplan images as the background
$(".bath").css( "background-color", "#020a01");
$("#bath").prepend("<img id = 'bath_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Bedroom.svg' width = '800px' height = '800px'/>");
$("#bed").prepend("<img id = 'bed_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Bedroom.svg' width = '800px' height = '800px'/>");
$("#live").prepend("<img id = 'live_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/LivingRoom.svg' width = '800px' height = '800px'/>");
$("#kitchen").prepend("<img id = 'kitchen_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Kitchen.svg' width = '800px' height = '800px'/>");
$(".designBoard>.tab").css("visibility","hidden");
$("#bath").css("visibility","visible");



