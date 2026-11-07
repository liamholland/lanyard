async function getStartTime() {
    const response = await fetch('/time/start');
    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }
    const res = await response.json()
    return new Date(res.time * 1000);
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
