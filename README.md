# wheresthejokebot

This is a super simple joke bot that I set up in about an hour for fun.

Uses [Sv443's Joke API](https://github.com/Sv443-Network/JokeAPI)

## To Build
```console
docker build -t wheresthejoke/<TAG> .
```

## To Run
```console
docker run --env DISCORD_TOKEN=<discord_bot_token> docker.io/wheresthejoke/<TAG>
```