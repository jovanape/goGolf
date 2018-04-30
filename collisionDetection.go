package main

import (
	"image/color"
)

type boxCollider struct {
	min, max, mid vector2
}

//HACK garbage implementation
func (c boxCollider) isCollidingWithBall(b *ball) string {
	s := ""

	if b.position.x+b.size > c.min.x && b.position.x < c.max.x{
		if b.verticalSpeed < 0 {
			if b.position.y < c.min.y && b.collisonGhost.position.y+b.size > c.min.y {
				s = "up"
			}
		} else {
			if b.position.y+b.size > c.max.y && b.collisonGhost.position.y < c.max.y {
				s = "down"
			}
		}
	}
	if b.position.y+b.size> c.min.y && b.position.y < c.max.y {
		if b.horisonatalSpeed>0{
			if b.position.x< c.min.x && b.collisonGhost.position.x+b.size>c.min.x{
				s="left"
			}
		}else{
			if b.position.x+ b.size > c.max.x && b.collisonGhost.position.x < c.max.x{
				s="right"
			}
		}
	}


	//Adjust for intersection
	switch s {
		case "up":
			b.position.y = c.min.y - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "down":
			b.position.y = c.max.y
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "right":
			b.position.x = c.max.x
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
		case "left":
			b.position.x = c.min.x - player.size
			b.opts.GeoM.Reset()
			b.opts.GeoM.Translate(player.position.x, player.position.y)
			}




	return s
}

///check all possible collisions
func (b *ball) checkForBallCollisions() string {
	candidates := getCandidateCollidersHorizontal(b)
	candidates = filterVetcial(b,candidates)

	//debug - shows in red the boxes that are being checked for collision
	for i := 0; i < lvl.nOfBoxes; i++ {
		lvl.maxSortedBoxes[i].graphic.Fill(color.White)
	}
	for i := 0; i < len(candidates); i++ {
		candidates[i].graphic.Fill(color.RGBA{255, 0, 0, 255})
	}


	s:=""

	for _, boxy := range candidates {
		s+=boxy.collider.isCollidingWithBall(b)
	}


	return s
}

func getCandidateCollidersHorizontal(b *ball) []*box {
	var collisionCandidateStartIndex int
	candidates := make([]*box, 0)
	if b.horisonatalSpeed>=0 {
		for i := 0; i < lvl.nOfBoxes; i++ {
			if b.position.x < lvl.maxSortedBoxes[i].collider.max.x {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := collisionCandidateStartIndex; i < lvl.nOfBoxes; i++ {
			candidates = append(candidates, lvl.maxSortedBoxes[i])
		}
	}else{
		for i := 0; i < lvl.nOfBoxes; i++ {
			if b.position.x+b.size < lvl.minSortedBoxes[i].collider.min.x {
				collisionCandidateStartIndex = i
				break
			}
		}
		for i := 0; i < collisionCandidateStartIndex; i++ {
			candidates = append(candidates, lvl.minSortedBoxes[i])
		}
	}

	return candidates
}


func filterVetcial(b *ball,candidates []*box)[]*box {
	l:= len(candidates)
	if b.verticalSpeed>0{
		for i:=0;i< l;i++ {
			if b.position.y+b.size < candidates[i].collider.min.y {
				copy(candidates[i:], candidates[i+1:])
				candidates[len(candidates)-1] = nil // or the zero value of T
				candidates = candidates[:len(candidates)-1]
				l=len(candidates)
			}
		}
	}else{
		for i:=0;i< len(candidates);i++ {
			if b.position.y > candidates[i].collider.max.y {
				copy(candidates[i:], candidates[i+1:])
				candidates[len(candidates)-1] = nil // or the zero value of T
				candidates = candidates[:len(candidates)-1]
				l=len(candidates)
			}
		}
	}


	return candidates
}