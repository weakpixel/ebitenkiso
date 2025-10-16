package tween

import "time"

type Group struct {
	list []tweenVal
}

type tweenVal struct {
	tween *Tween
	val   *float64
}

func (g *Group) Add(t *Tween, val *float64) {
	g.list = append(g.list, tweenVal{
		tween: t,
		val:   val,
	})
}

func (g *Group) Update(dt time.Duration) bool {
	all := true
	for _, t := range g.list {
		val, done := t.tween.Update(dt)
		t.val = &val
		all = all && done
	}
	return all
}
