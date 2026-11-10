async function registerUsername(name) {
    const response = await fetch("/identity/register", {
        method: "post",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Lanyard-Identity": getUserId()
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
    const response = await fetch("/identity/retrieve", {
        method: "post",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Lanyard-Identity": getUserId()
        },
        body: JSON.stringify({
            "lanyardId": getUserId(),
        })
    });

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }

    let res = await response.json();
    return res.username;
}

function createUserId() {
    const foundId = localStorage.lanyardId;
    if (foundId !== undefined) {
        console.log("user already assigned identity: ", foundId)
        return foundId
    }

    let lId = Math.floor(Math.random() * 99999) + 10000;
    localStorage.lanyardId = lId;

    return lId;
}

function getUserId() {
    const lId = localStorage.lanyardId;
    if (lId === undefined) {
        console.log("user has no id");
        return -1;
    }
    return lId;
}