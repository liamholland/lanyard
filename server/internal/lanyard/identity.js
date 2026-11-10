async function registerUsername(name) {
    let userId = getUserId();
    let response = await fetch("/identity/register", {
        method: "post",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Lanyard-Identity": userId
        },
        body: JSON.stringify({
            "username": name,
        })
    });

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }
}

async function retrieveUsername() {
    let userId = getUserId();
    let response = await fetch("/identity/retrieve", {
        method: "post",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Lanyard-Identity": userId
        },
        body: JSON.stringify({
            "lanyardId": userId,
        })
    });

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }

    let res = await response.json();
    return res.username;
}

function createUserId() {
    const foundId = localStorage.getItem("lanyardId");
    if (foundId !== null) {
        console.log("user already assigned identity: ", foundId);
        return foundId;
    }

    let lId = Math.floor(Math.random() * 99999) + 10000;
    localStorage.lanyardId = lId;
    console.log("user assigned identity: ", lId);

    return lId;
}

function getUserId() {
    const lId = localStorage.getItem("lanyardId");
    if (lId === null) {
        console.log("user has no id");
        return -1;
    }
    return lId;
}

function forgetUserId() {
    localStorage.removeItem("lanyardId");
}