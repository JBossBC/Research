package main
/**
  医院应用的责任链模式例子。医院中会有多个部门
  ·前台
  ·医生
  ·药房
  ·收银
  病人来访时， 他们首先都会去前台， 然后是看医生、 取药， 最后结账。 也就是说， 病人需要通过一条部门链， 每个部门都在完成其职能后将病人进一步沿着链条输送。

此模式适用于有多个候选选项处理相同请求的情形， 适用于不希望客户端选择接收者 （因为多个对象都可处理请求） 的情形， 还适用于想将客户端同接收者解耦时。 客户端只需要链中的首个元素即可。
*/
import(
	"fmt"
)

//处理者接口

type Department struct{
	execute(*Patient)
	setNext(Department)
}

//具体处理者接口

type Reception struct{
	next Department
}
func (r *Reception)execute(p *Patient){
   if p.registrationDone{
	fmt.Println("Patient registration already done")
	r.next.execute(p)
	return
   }
   fmt.Println("Reception registering patient")
   p.registrationDone=true
   r.next.execute(p)
}
func(r *Reception)setNext(next Department){
	r.next=next
}

// 具体处理者

type Doctor struct{
	next Department
}
func(d *Doctor)execute(p *Patient){
	if p.doctorCheckUpDone{
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone=true
	d.next.execute(p)
}
func(d *Doctor)setNext(next Department){
	d.next=next
}

//具体处理者

type Medical struct{
	next Department
}
func(m *Medical)execute(p *Patient){
	if p.medicineDone{
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone=true
	m.next.execute(p)
}
func (m *Medical) setNext(next Department) {
    m.next = next
}

//具体处理者

type Cashier struct {
    next Department
}

func (c *Cashier) execute(p *Patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) setNext(next Department) {
    c.next = next
}

//patient



type Patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

package main

func main() {

    cashier := &Cashier{}

    //Set next for medical department
    medical := &Medical{}
    medical.setNext(cashier)

    //Set next for doctor department
    doctor := &Doctor{}
    doctor.setNext(medical)

    //Set next for reception department
    reception := &Reception{}
    reception.setNext(doctor)

    patient := &Patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}