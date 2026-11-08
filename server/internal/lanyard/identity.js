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