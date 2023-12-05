package methods

import "fmt"

type Laptops struct {
	maker string
	model string
	Config
}

type Config struct {
	Processor string
	RAM       string
	HDD       string
}

func View() {
	laptop1 := Laptops{
		maker: "ASUS",
		model: "Zenbook",
		Config: Config{
			Processor: "Intel",
			RAM:       "16 GB",
			HDD:       "1 TB",
		},
	}
	laptop2 := &Laptops{
		maker: "Dell",
		model: "XPS",
	}

	//fmt.Println(laptop1.Processor, (*laptop2).model)

	// laptop2.display()
	laptop1.upgradeRAM()
	laptop1.upgradeHDD()
	// laptop1.display()

	laptop1.display()
	display(&laptop1)

	laptop2.display()
	display(laptop2)

}

func (l *Laptops) display() {
	fmt.Println(l)
}

func display(l *Laptops) {
	fmt.Println(l)
}

func (l Laptops) upgradeHDD() {
	l.HDD = "2 TB"
}

func (l *Laptops) upgradeRAM() {
	l.RAM = "32 GB"
}
