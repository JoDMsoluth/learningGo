package test28

type Player struct{}

type Monster struct{}

type Chest struct{}

// 각각의 관계마다 따로 정의해야된다.
// 뭐하나 추가될 때마다 새로 추가해야 되므로 이거는 변경에 닫혀있지 않다. -> O에 성립X
/*
func (p *Player) Attack(m *Monster)   {}
func (p *Player) AttackP(p *Player)   {}
func (p *Monster) Attack(p *Player)   {}
func (p *Monster) AttackM(m *Monster) {}
*/

func (p *Player) Attack(target BeAttackable) {
	// 데미지를 계산해서
	target.BeAttacked(3)
}

func (p *Monster) Attack(target BeAttackable) {
	// 데미지를 계산해서
	target.BeAttacked(3)
}

func (p *Chest) Attack(target BeAttackable) {
	// 데미지를 계산해서
	target.BeAttacked(3)
}

type Attackable interface {
	Attack(BeAttackable)
}

type BeAttackable interface {
	BeAttacked(int)
}

func Attack(attacker Attackable, defender BeAttackable) {
	attacker.Attack(defender)
}
