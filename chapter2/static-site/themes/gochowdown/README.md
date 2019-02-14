# GoChowdown

A theme for [Hugo](https://gohugo.io/) based on the [Chowdown](https://github.com/clarklab/chowdown) theme for Jekyll

<p align="center">
  <img src="https://raw.githubusercontent.com/seanlane/gochowdown/master/images/screenshot.png" />
</p>

## Quick Start

1. Add the repository into your Hugo Project repository as a submodule, `git submodule add https://github.com/seanlane/gochowdown.git themes/gochowdown`.
2. Configure your `config.toml` or `config.yaml`.
3. Build your site with `hugo serve` and see the result at `http://localhost:1313/`.

## Using this theme

Similarly to how the original Chowdown theme for Jekyll was organized, this theme adds two sections, recipes and components. The primary section you'll want to use is the recipes, as they form the panel listing on the front page, as well as most of the content. The components section is for recipes that form subcomponents of a recipe, allowing for a recipe to call on several components, and different recipes to reuse the same component, if desired.

Note that this is a work in progress, so things may be broken or change in the future. Feel free to contribute or offer suggestions.

### Add a new recipe draft

1. Navigate to the root directory of your website folder within a terminal
2. Type `hugo new --kind recipe-bundle recipes/name-of-your-new-recipe-here`, replacing `name-of-your-new-recipe-here` with the name of your recipe
  - Note that the default template (archetype in Hugo vernacular) will replace the hypens in the provided name with spaces as the title and capitalize the first letter of each word. For example, if I were to enter the command `hugo new --kind recipe-bundle recipes/hot-dog`, I would find a new folder at `content/recipes/hot-dog`, and the title within the `index.md` file in that folder would be `Hot Dog`.

### Add a new recipe with components

Similar to above, but instead of adding the recipes to the `content/recipes` directory, add the individual components to the `content/components` directory. Then add a new recipe as you normally would, and replace the instructions list with a components list, using the title (aka name) of the recipe, and modify the directions section as needed.

## License

Coder is licensed under the [MIT license](https://github.com/seanlane/gochowdown/blob/master/LICENSE.md).

