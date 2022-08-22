make Hyprland workspace works like dwm inspire by swaysome
for example
each monitor has workspace 1,2,3,4,5...

## 1. install hyprsome

```
go install github.com/clown139880/hyprsome@latest
```

## 2. config your hyprland

```
workspace=DP-1,1
workspace=HDMI-A-1,11
bind=SUPER,1,exec,hyprsome workspace 1
bind=SUPER,2,exec,hyprsome workspace 2
bind=SUPER,3,exec,hyprsome workspace 3
bind=SUPER,4,exec,hyprsome workspace 4
bind=SUPER,5,exec,hyprsome workspace 5
bind=SUPER,6,exec,hyprsome workspace 6
bind=SUPER,7,exec,hyprsome workspace 7
bind=SUPER,8,exec,hyprsome workspace 8
bind=SUPER,9,exec,hyprsome workspace 9
bind=SUPER,0,exec,hyprsome workspace 10

bind=SUPERSHIFT,exclam,exec,hyprsome movetoworkspace 1
bind=SUPERSHIFT,at,exec,hyprsome movetoworkspace 2
bind=SUPERSHIFT,numbersign,exec,hyprsome movetoworkspace 3
bind=SUPERSHIFT,dollar,exec,hyprsome movetoworkspace 4
bind=SUPERSHIFT,percent,exec,hyprsome movetoworkspace 5
bind=SUPERSHIFT,asciicircum,exec,hyprsome movetoworkspace 6
bind=SUPERSHIFT,ampersand,exec,hyprsome movetoworkspace 7
bind=SUPERSHIFT,asterisk,exec,hyprsome movetoworkspace 8
bind=SUPERSHIFT,parenleft,exec,hyprsome movetoworkspace 9
bind=SUPERSHIFT,parenright,exec,hyprsome movetoworkspace 10
```
