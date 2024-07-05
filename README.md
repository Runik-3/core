# Runik

<div style="text-align: center;">
  <img src="./assets/runik_screenshot.png" />
</div>

Generate and manage custom e-reader dictionaries for your favourite fictional worlds.

## Motivation

Most e-readers have built-in dictionaries, letting you define words in-text as you read. While this feature can tell you the meaning of words like 'tabard' and 'thaumaturgy', it can't help you define names, places, or magical items specific to the world of your book. 

Runik leverages the crowd-sourced info of fan wikis to generate custom e-reader dictionaries, extending the volcabulary of your e-reader to include fictional references. 

## Limitations

Runik is in early developmenet and currently **only supports Kobo e-readers**. Kindle support is in the works and coming in a future update. In the meantime, if you are a tinkerer, you can use the [Runik Builder](https://github.com/Runik-3/builder) tool to generate dictionaries and use one of the many other tools out there to convert it for use with Kindle devices.

Runik is designed to work on wikis built with [MediaWiki](https://www.mediawiki.org/wiki/MediaWiki). There is a decent amount of variablility in the way wikis can be configured and runik is not guaranteed to work with all wikis. Runik was mostly tested against Fandom wikis which are generally pretty consistent and *should* work best.

## Runik is made with

- [Wails](https://github.com/wailsapp/wails)
- [DictUtil](https://github.com/pgaskin/dictutil)
- [Svelte](https://github.com/sveltejs/svelte)
