GoIpsum
=======
Simple library for generating placeholder text.
How to Use:
--------------
Usage:

	package main
	import (
		"fmt"
		"github.com/vibbix/GoIpsum"
	)
	func main() {
		
		fmt.Println("Paragraph: ", GoIpsum.CreateText(GoIpsum.GenerationInfo.PARAGRAPH))
		//Example Output:
		//Paragraph:  	 Aliquet sed sed lorem pharetranibh. 
		//In vel tristique tempor diam,vestibulum. Donec imperdiet massa mattis 
		//duisvel. Rutrum nulla pharetra egestasac. Quam tortor nisl quam duiridiculus. 
		//Aliquam velit in nibh crasamet. Dignissim tempus dictum hendrerit gravidasit. 
		//A fermentum,viverra nisi,proin. Sem auctor libero dignissimenim. Nullam aliquet
		//inquisque pellentesquedignissim. Dignissim nec elementum porttitor leoet. In
		//habitant sit nislviverra. Condimentum eu pretium at curabituraliquam.
	}
