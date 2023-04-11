# waybar-powerprofiles

Waybar plugin to display the current powerprofile and allows you to change.

Behind the scenes, it  monitors changes on the DBus interface `net.hadess.PowerProfiles`
for changes and checks the current powerprofile. There is no need to periodically
run the binary or use an interval. The changes should be instantly visible.

## Requirements

`power-profiles-daemon` See: https://wiki.archlinux.org/title/CPU_frequency_scaling

## Options

`waybar-powerprofiles` can be called with the parameter `-next` to switch to
the next available power profile. 

`waybar-powerprofiles` can also called with the parameter `-set <profile>` to
switch to the provided power profile.


## Installation

```
go install github.com/subutux/waybar-powerprofiles@latest
```

## Configuration

In `$XDG_CONFIG_HOME/waybar/config`
```json
{
    // ... other waybar configuration
    "custom/powerprofiles": {
        "format": "{icon}",
        "return-type": "json",
        "format-icons": {
            "performance": "󰓅",
            "balanced": "󰾅",
            "power-saver": "󰾆"
        },
        "exec": "waybar-powerprofiles",
        "on-click": "waybar-powerprofiles -next"
    }
}
```

In `$XDG_CONFIG_HOME/waybar/style.css`
```css
#custom-powerprofiles {
    
    color: #a6e3a1;
    padding: 0px;
    padding-right: 10px;
    font-size: 18px;
    border-radius: 0 10px 10px 0;
    margin-left: 0px;
    margin-right: 10px;
    border-left: 0px;
}
```
