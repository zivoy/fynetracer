package main

import "renderShower/helpers"

type Object interface {
	GetPos() helpers.Vector3
}

// type Camera struct {
// 	position *helpers.Vector3
// 	facing   *helpers.Vector3
// }
