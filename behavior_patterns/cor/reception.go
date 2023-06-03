package cor

import "fmt"

// Reception 具体处理者
type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registration patent")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}
