//Global JS function for greeting
function greet() {
    //Call Go Greet function
    window.go.main.App.Greet().then(result => {
        //Display result from Go
            layer.msg("机器ID已复制到剪切板,可以直接粘贴使用");
    }).catch(err => {
        console.log(err);
    }).finally(() => {
        console.log("finished!")
    });
}