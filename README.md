# engine
Go Rogue, Tile based game engine.

## Project Todo List

Here you will see what is left to complete before this engine is ready for sharing/production.

### Scene/State Items

Each scene/game area can be considered as a _scene_. The stack can have scenes pushed on to it and in doing so the top most scene is the one executed during the game loop with all else effectively "paused."

* [ ] State Base
* [ ] Working Example of State system

### Scene Rendering / User Input

Development of a Rendering Interface; the idea behind this is that the developer can choose the render driver, or develop one of their own. We will provide one for [raylib](https://www.raylib.com/) and potentially [gdamore/tcell](https://github.com/gdamore/tcell).

* [ ] Render Base Interface
  * [ ] Default Raylib Driver Implementation
* [ ] Working Example of Raylib render driver

With the Raylib driver comes the concept of Spritesheets and Tilesets. This is because Raylib will be a "Virtual console":

* [ ] Raylib Tileset
* [ ] Raylib Spritesheet

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

* [ ] GUI Base
* [ ] GUI Widget Base
  * [ ] Button Widget
  * [ ] Radio Button Widget
  * [ ] Toggle Button Widget
  * [ ] Status Bar Widget
  * [ ] Container Widget
  * [ ] Toolbar Widget
  * [ ] Toolbar Separator Widget
  * [ ] Label Widget
  * [ ] TextBox Widget
  * [ ] Slider Widget
* [ ] Colour Themes
* [ ] Working Example for all above
