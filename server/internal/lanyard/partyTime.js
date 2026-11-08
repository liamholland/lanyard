async function getStartTime() {
    const response = await fetch('/time/start');
    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }
    const res = await response.json();
    return dateFromPosixTime(res.time);
}

async function addDepartureTime(time) {
    const response = await fetch('/time/departure', {
        method: 'post',
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Lanyard-Identity": getUserId()
        },
        body: JSON.stringify({
            "time": time,
        })
    });

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    } else {
        return response.ok;
    }
}

async function getNumPlayers(time) {
    const response = await fetch(`/time/howmany?time=${time}`);

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}:${JSON.stringify(response.headers)}`);
    }

    const res = await response.json();
    return res.numPeople;
}

function formattedCountdown(startTime) {
    const now = new Date();
    const differenceMs = startTime.getTime() - now.getTime();

    const totalSeconds = Math.floor(differenceMs / 1000);
    const days = Math.floor(totalSeconds / (3600 * 24));
    const hours = Math.floor((totalSeconds % (3600 * 24)) / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;

    return `Time remaining: ${days}d ${hours}h ${minutes}m ${seconds}s`;
}

function dateFromPosixTime(time) {
    return new Date(time * 1000);
}

function posixTimeFromDate(date) {
    return Math.floor(new Date(date).getTime() / 1000);
}