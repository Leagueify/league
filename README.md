# League Service

The Leagueify league service is responsible for managing league-specific
operations and rules. The Leagueify league service uses  [Go][go-download]
using version `1.23.0`.

## Key Functions

- Create and manage individual leagues and their seasons.
- Define league rules, schedules, and standings.
- Coordinate with the Roster and Statistics services for player and team information.

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify league service uses docker for a consistent
development environment. Running the Leagueify league service locally can be
accomplished using commands found in the [Makefile][repo-makefile]. During local
development changes will trigger a live reload, eliminating the requirement to
restart the docker image manually. This is accomplished by using the wonderful
tool [air][github-air]. The most common commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify league service is ready for development once the banner output is
visible within the terminal. The banner blelow was created using the
[Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-league-1  |
leagueify-league-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-league-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-league-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-league-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-league-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-league-1  |
leagueify-league-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|
leagueify-league-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .
leagueify-league-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|
leagueify-league-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||
leagueify-league-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....|
leagueify-league-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0ALEAGUE
[repo-makefile]: ./Makefile
