# gomoku

A goban game [online](https://gomoku.trixky.com/), with an Ai using a customizable [negamax](https://en.wikipedia.org/wiki/Negamax) implementation improved by [alpha beta pruning](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning) (in depth and width) and [transposition tables](https://en.wikipedia.org/wiki/Negamax#Negamax_with_alpha_beta_pruning_and_transposition_tables).

<img src="https://raw.githubusercontent.com/trixky/gomoku/main/.demo/screens.gif" alt="Demo gif" width="361"/>

## Rules

- Two players take turns placing stones of their color on an intersection of the board.
- The game ends when one player manages to __captures ten opposing stones__ or __align five stones or more__.
- A player that manages to align five stones __only wins if__ the opponent can not break this alignment by capturing a pair, or if he has already lost four pairs and the opponent can capture one more, therefore winning by capture.
- You can __capture__ a pair of your opponent’s stones to remove it from the board by __flanking__ them with your own stones.
- It is forbidden to play a move that introduces two __free-three alignments__, which would guarantee a win by alignment.
- __No limit__ to the number of stones.

## Usage

### Full-stack (prod)

```bash
source env.sh
docker-compose -f docker-compose.prod.yaml up
```

### Full-stack (dev)

```bash
source env.sh
docker-compose up
```

## Online

This project is online, so you can visit it by clicking [here](https://gomoku.trixky.com/)!
