$("#update").click(function (e) {

    $('#title').removeClass("is-invalid");
    $('#color').removeClass("is-invalid");
    $('#text_color').removeClass("is-invalid");


    $('#title').addClass("is-valid");
    $('#color').addClass("is-valid");
    $('#text_color').addClass("is-valid");


    let title_name = $("#title").val();
    let description_val = $("#description").val();
    let color_val = $("#color").val();
    let text_color_val = $("#text_color").val();

    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(title_name)) {
        is_valid = false
        $('#title').removeClass("is-valid");
        $("#title").addClass("is-invalid")
        // show error
    }

    if (re.test(color_val)) {
        is_valid = false
        $('#color').removeClass("is-valid");
        $("#color").addClass("is-invalid")
        // show error
    }

    if (re.test(text_color)) {
        is_valid = false
        $('#text_color').removeClass("is-valid");
        $("#text_color").addClass("is-invalid")
        // show error
    }

    if(!is_valid){
        return
    }




    let send_data = JSON.stringify({
        name: title_name,
        description: description_val,
        colour: color_val,
        textcolor: text_color_val
    })
    let page_url = window.location.href.split('/')
    let project_id = page_url[page_url.length - 2]
    $.ajax({
        type: "POST",
        url: "../../../api/project/"+project_id,
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done_update'))
                toast.show()
            }
        }
    })
        .fail(function() {
        let toast = new bootstrap.Toast(document.getElementById('fail_update'))
        toast.show()
    });
    
});

$("#delete").click(function (e) {

    let page_url = window.location.href.split('/')
    let project_id = page_url[page_url.length - 2]
    $.ajax({
        type: "POST",
        url: "../../../api/delete_project/"+project_id,
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done_del'))
                toast.show()
                location.reload()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
});