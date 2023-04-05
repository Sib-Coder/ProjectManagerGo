$(".alert").hide()

$("#send").click(function (e) {
    $('#inputLogin').removeClass("is-invalid");
    $('#inputName').removeClass("is-invalid");
    $('#inputLName').removeClass("is-invalid");
    $('#inputPassword').removeClass("is-invalid");
    $('#reInputPassword').removeClass("is-invalid");

    $('#inputLogin').addClass("is-valid");
    $('#inputName').addClass("is-valid");
    $('#inputLName').addClass("is-valid");
    $('#inputPassword').addClass("is-valid");
    $('#reInputPassword').addClass("is-valid");

    let login = $("#inputLogin").val();
    let name = $("#inputName").val();
    let last_name = $("#inputLName").val();
    let pass = $("#inputPassword").val();
    let rep_pass = $("#reInputPassword").val();
    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(login))
    {
        is_valid = false
        $('#inputLogin').removeClass("is-valid");
        $("#inputLogin").addClass("is-invalid")
        // show error
    }
    if (re.test(name))
    {
        is_valid = false
        $('#inputName').removeClass("is-valid");
        $("#inputName").addClass("is-invalid")
        // show error
    } 
    if (re.test(last_name))
    {
        is_valid = false
        $('#inputLName').removeClass("is-valid");

        $("#inputLName").addClass("is-invalid")
        // show error
    } 
    if (re.test(pass))
    {
        is_valid = false
        $('#inputPassword').removeClass("is-valid");

        $("#inputPassword").addClass("is-invalid")
        // show error
    } 
    if (pass != rep_pass)
    {
        is_valid = false
        $('#inputPassword').removeClass("is-valid");
        $('#reInputPassword').removeClass("is-valid");
        $("#inputPassword").addClass("is-invalid")
        $("#reInputPassword").addClass("is-invalid")
        // show error
    } 
    let send_data = JSON.stringify({
        login: login,
        password: pass,
        lname: last_name,
        fname: name
    })

    if(!is_valid){
        return
    }

    $(".alert").hide()

    $.ajax({
        type: "POST",
        url: "../api/new_login",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        statusCode:{
            200:function(response) {
                if (response && response["error"]){
                    $('#inputLogin').addClass("is-invalid");
                    $(".alert").show()
                } else {
                    let toast = new bootstrap.Toast(document.getElementById('done'))
                    toast.show()
                }
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
});