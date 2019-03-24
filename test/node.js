var standard_input = process.stdin;
standard_input.setEncoding('utf-8');
var s=[];
standard_input.on('data', function (data) {
    s.push(data)
    if(s.length === 2){
        console.log(s[0]+"\n"+s[1]);
        process.exit();
    }
});