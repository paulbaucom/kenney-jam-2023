components {
  id: "item"
  component: "/main/item.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/items.atlas\"\n"
  "default_animation: \"med\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "playback_rate: 0.0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.333
    y: 0.333
    z: 0.333
  }
}
embedded_components {
  id: "light"
  type: "sprite"
  data: "tile_set: \"/main/main.atlas\"\n"
  "default_animation: \"light_mask_64\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_SCREEN\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
