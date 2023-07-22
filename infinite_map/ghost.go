components {
  id: "ghost"
  component: "/main/ghost.script"
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
components {
  id: "lightsource"
  component: "/lights/lightsource.script"
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
  properties {
    id: "color"
    value: "0.7, 0.7, 0.7, 0.5"
    type: PROPERTY_TYPE_VECTOR4
  }
  properties {
    id: "radius"
    value: "300.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "falloff"
    value: "2.5"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/main.atlas\"\n"
  "default_animation: \"ghost_normal\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
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
}
