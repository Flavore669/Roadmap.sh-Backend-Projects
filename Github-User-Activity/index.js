// Main

const requested_username = process.argv[2];
fetchGitHubUser(requested_username)
    .then((username) => {printGitHubRepos(username);})
    .catch((err) => {
        console.error(err.message);
        process.exit(1);
    });

/**
 * 
 * @param {Response} username 
 */
async function fetchGitHubUser(username) {
    const response = await fetch(`https://api.github.com/users/${username}/events`);

    if (!response.ok) {
        if(response.status === 404){
            throw new Error(`User not found, ${username}`);
        }
        else{
            throw new Error("An Error Occured");
        }
    }

    return response.json();
}

/**
 * Prints GitHub events in a readable format
 * @param {Array<{type: string, repo: {name: string}}>} events - Array of GitHub event objects
 */
function printGitHubRepos(events) {
    if (events.length === 0) {
        console.log("User Hasn't Done Anything");
        return;
    }
    
    events.forEach(element => {
        console.log(`${element.type.replace("Event", "")} in ${element.repo.name}`)
    });
}