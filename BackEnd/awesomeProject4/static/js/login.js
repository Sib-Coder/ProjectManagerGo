$(".alert").hide()

$("#send").click(function (e) {

    $('#inputLogin').removeClass("is-invalid");
    $('#inputPassword').removeClass("is-invalid");


    $('#inputLogin').addClass("is-valid");
    $('#inputPassword').addClass("is-valid");

    let name = $("#inputLogin").val();
    let pass = $("#inputPassword").val();

    $(".alert").hide()

    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(name)) {
        is_valid = false
        $('#inputLogin').removeClass("is-valid");
        $("#inputLogin").addClass("is-invalid")
        // show error
    }

    if (re.test(pass)) {
        is_valid = false
        $('#inputPassword').removeClass("is-valid");
        $("#inputPassword").addClass("is-invalid")
        // show error
    }

    if(!is_valid){
        return
    }


    let send_data = JSON.stringify({
        login: name,
        pass: pass
    })
    $.ajax({
        type: "POST",
        url: "../api/login",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            if (response && response["error"]){
                $('#inputLogin').addClass("is-invalid");
                $('#inputPassword').addClass("is-invalid");
                $(".alert").show()

            }
            else {
                location.reload()
            }
        }
    });
    
});