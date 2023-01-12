# picview

random image api written by go.

**all images should be placed at `./images/` directory**

url format: http://{url}:{port}/api/random/?category=scene&device=pc
- category=
    - scene(default)
    - comic
- device=
    - pc(default)
    - mobile

[docker url](https://hub.docker.com/r/yuzhibo535/picview) is here. 

## todo
- [ ] cors
- [ ] support categery
- [x] better logger
- [x] docker support 