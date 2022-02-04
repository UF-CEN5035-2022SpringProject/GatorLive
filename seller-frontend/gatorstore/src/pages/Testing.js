import React, { useEffect } from 'react';

function Testing() {
    const loginCode = GetUserCode('code');
    const dateTime = GetDateTime();
    var backendStatus;
    var backendResult;
    
    // "Frontend will call a backend API to pass the code ":
    useEffect(() => {
        SendPost();
    }, []);

    const SendPost = async () => {
        const requestOptions = {
            method: 'POST',
            headers: { },
            body: JSON.stringify({ code: loginCode})
        };
        const response = await fetch('http://10.136.228.201:8080/test/api/user/login', requestOptions);
        const loginResponse = await response.json();
        backendResult = loginResponse.result;
        backendStatus = loginResponse.status;

        if (backendStatus === 0)
            alert("success");
        else alert("fail");

        /*if (backendStatus === 0) {
            window.location.href = "http://localhost:3000/landingpage";
        } else {
            alert("ERROR: User was not able to be authenticated.");
        }*/
    }

    function GetDateTime() {
        var currentdate = new Date(); 
        var datetime =   currentdate.getDate() + "/"
                        + (currentdate.getMonth()+1)  + "/" 
                        + currentdate.getFullYear() + " @ "  
                        + currentdate.getHours() + ":"  
                        + currentdate.getMinutes() + ":" 
                        + currentdate.getSeconds();
        return datetime;
    }

    function GetUserCode(parameterName) { // get whatever is after '?' on the URL
        const rawQuery = window.location.search;
        const urlParams = new URLSearchParams(rawQuery);
        const loginCode = urlParams.get(parameterName);
        return loginCode;
    }

    return (
        <div style={{ padding: 20 }}>
            Redirecting... The code is: {GetUserCode('code')}
        </div>
    );
}

export default Testing;