

// question 36
for(i=1;i<=5;i++){
    let str = '';
    for(j=1;j<=5-i+1;j++){
        str += j;
    }
    console.log(str);
}

// question 37
for(i=1;i<=5;i++){
    let str = '';
    for(j=1;j<=5-i+1;j++){
        str += 5-j+1;
    }
    console.log(str);
}

// question 38
for (let i=5;i>=1;i--) {
    let line = '';
    for (let j=i; j>=1;j--) {
        line +=j;
    }
    console.log(line);
}
 
//question 39
for(let i=1; i<= 3;i++){
    let pattern='';
    for(let j=1; j<=2; i++){
        pattern += '&&\n';
    }
    console.log(pattern)
}

// question 40 

for(i=1;i<=5;i++){
    let pattern5 = '';
     for(j=1;j<=i;j++){
        pattern5+= '*' ;
    }
    console.log(pattern5);
}