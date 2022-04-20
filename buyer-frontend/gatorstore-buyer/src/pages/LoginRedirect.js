import React, { useEffect } from 'react';
import settings from '../settings'

function LoginRedirect() {
    const loginCode = GetUserCode('code');
    //const dateTime = GetDateTime();
    var backendStatus;
    var backendResult;
    
    // Frontend will call a backend API to pass the code:
    useEffect(() => {
        SendPost();
    }, []);

    const SendPost = async () => {
        const requestOptions = {
            method: 'POST',
            body: JSON.stringify({code: loginCode})
        };
        const res = fetch(settings.apiHostURL + 'user/login?from=buyer', requestOptions)
            .then(response => response.json())
            .then(response => {
                backendStatus = response.status;
                backendResult = response.result;

                // proceed or fail, drip or drown:
                if (backendStatus === 0) {
                    window.sessionStorage.setItem('user-name', response.result.name);
                    window.sessionStorage.setItem('user-email', response.result.email);
                    window.sessionStorage.setItem('user-id', response.result.id);
                    window.sessionStorage.setItem('user-jwtToken', response.result.jwtToken);

                    window.location.href = `${settings.applicationRootURL}/home`;
                } else {
                    alert("ERROR: User was not able to be authenticated.");
                }
            })
            .catch((error) => {
                console.error(error);
            });        
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
            Redirecting...
        </div>
    );
}

export default LoginRedirect;