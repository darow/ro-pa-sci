function getTopPlayers(players) {
    const playerList = Object.values(players)
        .map((user, index) => {
            const { name, is_online, score, id } = user;
            const inviteButton = `<btn id="invite-${id}" class="btn btn-sm btn-success btn-block">пригласить✉</btn>`;

            return `<li class="m-1">${index + 1}. ${name} ${is_online} ${score} ${inviteButton}</li>`;
        });

    return `<h3>Players Online</h3>
        <div class="container">
            <ul>
                ${playerList}
            </ul>
        </div>`;
}

export default getTopPlayers;
