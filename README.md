# pokedex


## Overview
The pokedex explores the world of Pokemon in the terminal. The intent behind the pokedex is to learn API querying and Caching in Go. All pokemon information is retrieved from [PokeAPI](https://pokeapi.co/docs/v2#pokemon-section).

## Install
Run the following command in a UNIX style terminal.

`git clone https://github.com/PlinyTheYounger0/pokedex`

## Commands
- Help
    
    Lists the commands available for use with a brief description of each command. Arguements will be listed afte the command name in <>.
- Exit
    
    Ends the REPL and retrns the user to the terminal.
- Map
    
    Gets and caches 20 locations and prints them to the terminal.
- Mapb
    
    Goes back 20 locations and prints them to the terminal.
- Explore
    
    Gets and caches all pokemon in a desired location and prints them to terminal.
- Catch
    
    Allows the user to attempt to catch a pokemon. The difficulty of catch is directly related to the base experience of the pokemon being caught.
- Inspect
    
    Allows the user to inspect the stats and types of a pokemon that the user has caught.
- Pokedex
    
    Prints all pokemon in the users pokedex.

## Path Forward
- Add capability for the user to press the up arrow and go back to their last command.
- Add a command allowing the user to fight their pokemon against pokemon out in the world.
- Introduce experience and evolution based on levels.
- Save a users state and allow them to keep pokemon after stoping the program.