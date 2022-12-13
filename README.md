<h1 align="center">GoRogue Engine</h1>
<p align="center"><em>A tile based game engine modeled on libtcod</em></p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/go-rogue/engine.svg" alt="License"></a>
  <a href="https://goreportcard.com/report/github.com/go-rogue/engine"><img src="https://goreportcard.com/badge/github.com/go-rogue/engine" alt="Go report card"></a>
</p>

## About

TBC

## Usage

TBC

## Project Todo List

Here you will see what is left to complete before this engine is ready for sharing/production.

### Scene/State Items

Each scene/game area can be considered as a _scene_. The stack can have scenes pushed on to it and in doing so the top most scene is the one executed during the game loop with all else effectively "paused."

* [x] State Base
* [x] Working Example of State system

### Scene Rendering / User Input

Development of a Rendering Interface; the idea behind this is that the developer can choose the render driver, or develop one of their own. We will provide one for [raylib](https://www.raylib.com/) and potentially [gdamore/tcell](https://github.com/gdamore/tcell).

* [x] Render Base Interface
  * [x] Default Raylib Driver Implementation
* [x] Working Example of Raylib render driver

With the Raylib driver comes the concept of Spritesheets and Tilesets. This is because Raylib will be a "Virtual console":

* [x] Raylib Tileset
* [x] Raylib Spritesheet

### Maps & Map Generation

* [ ] Map Base
* [ ] Map Camera
* [ ] Map FOV
 * [ ] Default FOV implementation (and interface)
* [ ] Map Generation Interface
  * [ ] Roguelike Room Generator (example)
* [ ] Map Camera

### Entities

* [ ] Entities Base

### GUI Todo Items

The GUI system will be based upon that available in [afolmert/libtcod-go](https://github.com/afolmert/libtcod-go/blob/master/tcod/gui.go) but modified to work with our render engine implementation.

* [x] GUI Base
* [x] GUI Widget Base
  * [x] Button Widget
  * [x] Radio Button Widget
  * [x] Toggle Button Widget
  * [ ] Status Bar Widget
  * [x] Container Widget
    * [x] Vertical Container Widget
    * [x] Horizontal Container Widget
  * [ ] Toolbar Widget
    * [ ] Toolbar Separator Widget
  * [ ] Label Widget
  * [ ] TextBox Widget
  * [ ] Slider Widget
* [ ] Input Handling
  * [ ] Keyboard Events
  * [ ] Mouse Events
* [ ] Colour Themes
* [ ] Working Examples for all above

## Credits

### Libraries
The following libraries are used by GoRogue, with gratitude 💕
- [gen2brain/raylib-go](https://github.com/gen2brain/raylib-go)

### Developers

### Special Thanks
A huge thank-you to Jice, Mingos and others for the [libtcod](https://github.com/libtcod/libtcod) library, Adam Folmert and others for [libtcod-go](https://github.com/afolmert/libtcod-go); without either project this one would not exist. Another huge thank-you to /r/roguelikes and their discord community for both inspiration and help. #roguelikedev

### Research
The following list of links contains all the webpages I found useful in my quest to build this library.
- [Complete roguelike tutorial using C++ and libtcod](http://www.roguebasin.com/index.php/Complete_roguelike_tutorial_using_C%2B%2B_and_libtcod_-_part_1:_setting_up)
- [Roguelike Tutorial in Rust + tcod](https://tomassedovic.github.io/roguelike-tutorial/index.html)
- [libtcod-go: Go bindings for the libtcod library](https://github.com/afolmert/libtcod-go)
- [Pataro: A Roguelike library built on top of libtcod](https://github.com/SuperFola/pataro)
- [Rotten Soup: A roguelike built with Vue, Vuetify, Tiled, rot.js, and PixiJS](https://github.com/Larkenx/Rotten-Soup)
- [The Wavefunction Collapse Algorithm explained very clearly](https://robertheaton.com/2018/12/17/wavefunction-collapse-algorithm/)
- [Roguelike Tutorial - In Rust](https://bfnightly.bracketproductions.com/chapter_0.html)
- [Pataro: A Rogue like library on top of libtcod](https://github.com/SuperFola/pataro)
- [How to Use Tile Bitmasking to Auto-Tile Your Level Layouts](https://gamedevelopment.tutsplus.com/tutorials/how-to-use-tile-bitmasking-to-auto-tile-your-level-layouts--cms-25673)

## License
GoRogue is free software, licensed under the [MIT LICENSE](LICENSE). We encourage forking and changing the code, hacking around with it, and experimenting.

Copyright (C) 2019-2022 the GoRogue Authors.