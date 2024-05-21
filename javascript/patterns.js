// question 33
let row=5;
for(i=1;i<=row;i++){
    let pattern = '';
     for(j=1;j<=i;j++){
        pattern += 1 ;
    }
    console.log(pattern);
}

// mistake
let row1=5;
for(i=1;i<=row1;i++){
    let pattern1 = '';
    for(j=1;j<=3;j++){
        if (i==1 || i == 5){
          pattern1 += '`';
        }
        else {
            pattern1 += '*';
        }
    } 
    console.log(pattern1);
}

//question 34
for(let i =1; i<=3;i++){
    pattern2 = '';
    for(let j= 1; j<=3;j++){
        pattern2 += '-';
    }
    console.log(pattern2);
}

// question 35
for(let i =1; i<=4;i++){
    pattern3 = '';
    for(let j= 1; j<=4;j++){
        pattern3 += '*';
    }
    console.log(pattern3);
}



