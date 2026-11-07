function aggregateGames(files) {
    let sizeMap = {};
    let baseSize = 0;
    
    let baseDir = "";
    if (files.length > 0) {
        const baseParts = files[0].webkitRelativePath.split('/');
        baseDir = baseParts[0]
    }

    for (const file of files) {
        const pathParts = file.webkitRelativePath.split('/');
        if (pathParts.length < 3) {
            continue
        }

        if (pathParts.length > 1) {
            directoryName = pathParts[1];
        } else {
            // should not be here
            continue
        }

        const size = file.size;
        baseSize += size;

        if (sizeMap[directoryName]) {
            sizeMap[directoryName] += size;
        } else {
            sizeMap[directoryName] = size;
        }
    }
    sizeMap[baseDir] = baseSize;
    return sizeMap;
}

async function resolveGames(gamesMap, resolvedGames) {
    const response = await fetch('/games/resolve', {
        method: "post",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "games": gamesMap,
            "resolved": resolvedGames === null ? {} : resolvedGames
        })
    });

    if (!response.ok) {
        throw new Error(`Network response was not ok:${response.statusText}`);
    }

    const res = await response.json();
    console.log(res);
    return res;
}

function persistResolvedGames(gamesMap) {
    localStorage.resolvedGames = JSON.stringify(gamesMap);
} 

function retrieveResolvedGames() {
    if (localStorage.resolvedGames === undefined) {
        return null;
    }

    return JSON.parse(localStorage.resolvedGames);
}
