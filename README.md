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
* [ ] Working Example for all above

## Research
- https://tomassedovic.github.io/roguelike-tutorial/index.html
- https://github.com/afolmert/libtcod-go
- [Pataro: A Roguelike library built on top of libtcod](https://github.com/SuperFola/pataro)