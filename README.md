# Pokédex 

Pokédex is a command-line tool that allows users to interact with the Pokémon universe directly from their terminal. With this application, users can catch Pokémon, view their collection, and explore detailed information about each Pokémon.

## Features

- **Catch Pokémon**: Users can catch Pokémon by name, adding them to their personal Pokédex.
- **View Pokédex**: Display all Pokémon that have been caught during the session.
- **Command History**: Navigate through the history of entered commands using the up and down arrow keys. (Note: This feature was planned but requires additional implementation steps.)
- **Pokémon Details**: View detailed information about each Pokémon, including base experience and level calculations.
- **Location**: View the location-area of each Pokémon, including the region and habitat.
- **Exit**: Exit the application at any time.

## Installation

To install the CLI Pokémon Application, clone this repository and build the application using Go:

```sh
git clone https://github.com/BrownieBrown/pokedex.git
cd pokedex
go build
```

## Usage

After building the application, you can run it directly from the command line:

```sh
./pokedex
```

## Commands

### Help
To view a list of available commands, use the help command:

```sh
help
```

### Exit
To exit the application, use the exit command:

```sh
exit
```

### Pokémon Details
To view detailed information about a Pokémon you caught, use the details command followed by the name of the Pokémon:

```sh
inspect pikachu
```

### Catch Pokémon
To catch a Pokémon, use the catch command followed by the name of the Pokémon:

```sh
catch pikachu
```

### View Pokédex

To view the Pokédex, use the pokédex command:

```sh
pokedex
```

### Location (Next)

To view the a list of available locations and toggle through the list, use the location command:

```sh
map
```

### Location (Previous)

To view the a list of available locations and toggle through the list, use the location command:

```sh
mapb
```

### Explore
To explore a location, use the explore command followed by the name of the location:

```sh
explore kanto
```

## Contributing

We welcome contributions to the CLI Pokémon Application! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE] file for details.