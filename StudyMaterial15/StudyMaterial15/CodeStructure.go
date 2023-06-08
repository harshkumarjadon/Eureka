//_____________________________________________
// Programmers Perceptive
func main() {
	f1()
	f2()
}

func f1()  { 
	f11()
	f12()
	f13()
}

func f2()  { 	
	f21()
	f22()
}

func f11() { 	}
func f12() {	}
func f13() {	}

func f21() { 	}
func f22() {	}

//_____________________________________________
// Mathematician Perceptive

func main() {
	func f1()  { 
		func f11() { 	}
		func f12() {	}
		func f13() {	}
	}

	func f2()  { 	
		func f21() { 	}
		func f22() {	}
	}
}

//_____________________________________________
// Mathematician Perceptive

// Outside Context File

// Expressions L
{ // Context Context1

	{  // Expression L1  // Context Context2 Outside Context: Context1
		{ 	} // Expression L11
		{	} // Expression L12
		{	} // Expression L12
	}

	{  // Expression L2
		{ 	} // Expression L21
		{	} // Expression L22
	}
}


{ // For Context 1 Outside Context 0 
	// Member Function
	{ // For Context 2 Outside Content 1
		return somethingDingDong + somethingAtFile
	}
}
