package main

import (
	"log"
	"math/big"
)

type (
	PubKey struct {
		n, e *big.Int
	}
	PriKey struct {
		n, d *big.Int
	}
)

func main() {
	// log.Println(extendedEuclid(big.NewInt(30), big.NewInt(47)))
	simpleRsa()
}

func simpleRsa() {
	p := 7
	q := 11
	n := big.NewInt(int64(p * q))      // 77
	eulerN := int64((q - 1) * (p - 1)) // φ(n) 60
	e := big.NewInt(13)                // 选一个小于 φ(n)并与之互质的数
	// e*d ≡ 1 (mod φ(n))
	// d = e ^ (φ(φ(n))-1)
	// d = e ^ ( 8 - 1) // 一般φ(φ(n))没法计算太大 不去算 计算个小的就行
	// e*d - 1 = x * φ(n)
	// e*d - eulerN *x = 1
	// 13 * d + 60 *(-x) = 1
	// d 为e 模反元素 φ(n)
	_, xVal, d := extendedEuclid(e, big.NewInt(eulerN))
	log.Printf("-x = %d,d = %d\n", xVal, d)
	dInt := d.Int64()
	for dInt < 0 {
		dInt += eulerN
	}
	d.Set(big.NewInt(dInt))
	// pub [n,e]
	// pri [n,d]
	pubKey := &PubKey{
		n: n, // 77
		e: e, // 13
	}
	priKey := &PriKey{
		n: n, // 77
		d: d, // 37
	}
	secretNumber := int64(19)
	enNumber := big.NewInt(0)
	deNumber := big.NewInt(0)
	// enNumber = secret ** e % n
	enNumber.Exp(big.NewInt(secretNumber), pubKey.e, pubKey.n)
	// deNumber = enNumber ** d % n
	deNumber.Exp(enNumber, priKey.d, pubKey.n)
	log.Println(secretNumber == deNumber.Int64())
}

// 获取两数的最大公约数
// 用大数一直除以小数,最后一次小数为0时 大数就是最大公约数
func gcd(a, b int64) int64 {
	var r int64
	for b != 0 { // 取模计算的后者最后必是0
		r = a % b // r 会越来越小 最后都会赋值给b
		// 即 a b 中的较大值赋值给a 较小值赋值给b
		// r 一定小于 b
		a = b
		b = r
	}
	return a
}

func extendedEuclid(a *big.Int, b *big.Int) (gcd *big.Int, ud *big.Int, vd *big.Int) {
	c := a
	gcd = b
	if a.Cmp(b) < 0 {
		c = b
		gcd = a
	}
	uc := big.NewInt(1)
	vc := big.NewInt(0)
	ud = big.NewInt(0)
	vd = big.NewInt(1)
	for c.Sign() != 0 {
		q := big.NewInt(0).Div(gcd, c)
		cOld := c
		c = big.NewInt(0).Sub(gcd, big.NewInt(0).Mul(q, c))
		gcd = cOld
		ucOld := uc
		vcOld := vc
		uc = big.NewInt(0).Sub(ud, big.NewInt(0).Mul(q, uc))
		vc = big.NewInt(0).Sub(vd, big.NewInt(0).Mul(q, vc))
		ud = ucOld
		vd = vcOld
	}
	return
}
