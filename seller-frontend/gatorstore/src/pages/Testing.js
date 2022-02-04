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
        console.log("new attempt----------------------------------------");
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ code: loginCode})
        };
        const res = fetch('http://10.136.228.201:8080/test/api/user/login', requestOptions)
            .then(response => response.json())
            .then(response => {
                backendStatus = response.status;
                backendResult = response.result;
                console.log(backendStatus, backendResult);

                // proceed or fail:
                if (backendStatus === 0) {
                    window.location.href = "http://localhost:3000/landingpage";
                } else {
                    alert("ERROR: User was not able to be authenticated.");
                }
            })
            .catch((error) => {
                console.error(error);
            });

        //console.log(backendStatus, backendResult);
        //if (backendStatus === 0)
          //  console.log("Success");
        //else 
          //  console.log("Fail");

        
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