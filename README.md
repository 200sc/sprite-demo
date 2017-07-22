# sprite-demo
Oak sprite demo

This demo is designed to mirror [this ebiten demo](https://hajimehoshi.github.io/ebiten/examples/sprites.html). This demo performs worse than ebiten, because oak does not as of writing use openGL, and because ebiten's screen size is a quarter of oak's. 

Press K to spawn in sprites. Draw FPS is displayed in the top left, and underneath it logical FPS. To the right is the number of rendered sprites. Consider commenting out the rotation line to see how it effects the frame rates.

The gopher image used is from [gophericons](https://github.com/shalakhin/gophericons), based on Renee French under Creative Commons 3.0 Attributions.
