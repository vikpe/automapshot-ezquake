# DEPRECATED
**This project is no longer maintained**, please consider using [automapshot-fte](https://github.com/vikpe/automapshot-fte) instead.

---

# automapshot

> Automate screenshots of QuakeWorld maps.

Uses an ezQuake client to cycle through maps, load camera settings and take screenshots.

## Requirements

* Unix build of [ezQuake](https://github.com/ezQuake/ezquake-source)
* `.env` (copy `.env.example` and set values)
* `map_settings.json` - Settings per map.

For more map settings, see [map_settings.json](https://github.com/vikpe/qw-mapshots/blob/main/configs/map_settings.json)
in the [qw-mapshots](https://github.com/vikpe/qw-mapshots) repo.

## Usage

```shell
# all maps defined in map_settings.json
automapshot all

# specific maps
automapshot dm2
automapshot dm2 dm4 dm6
```

## Tips

### Creating thumbnails

```shell
 mkdir thumbs
 mogrify  -format jpg -path thumbs -thumbnail 400x300 *.jpg
 ```

## Related projects
* [Automapshot FTE](https://github.com/vikpe/automapshot-fte)
* [QuakeWorld Mapshots](https://github.com/vikpe/qw-mapshots)
* [QuakeWorld Hub](https://github.com/quakeworldnu/hub.quakeworld.nu) 
