package GoIpsum

import (
	"fmt"
)

//Simple Create text example, that also functions as test case.
//Random values seem to be consistent while working in test cases.
func ExampleCreateText() {
	var GItest GenerationInfo
	GItest = WORD
	fmt.Println("Word: ", CreateText(GItest))
	GItest = SENTENCE
	fmt.Println("Sentence: ", CreateText(GItest))
	GItest = PARAGRAPH
	fmt.Println("Paragraph: ", CreateText(GItest))
	//Output:
	//Word:  sagittis
	//Sentence:   In egestas tellus, at vulputate odio.
	//Paragraph:  	 Neque augue quam tempus massa. Accumsan elementum urna, gravida sapien. Mauris nibh feugiat erat sed. Lacus, est malesuada sed a. Donec sed urna ac vitae imperdiet. Eget nunc purus donec eget ut. Quis eget sit mus risus. Aenean lobortis platea lacus fermentum elementum. Rutrum cursus sodales ultrices augue augue. Odio neque nulla enim tortor. Ac eros metus, viverra adipiscing.
}
