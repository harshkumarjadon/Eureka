
package learnConcepts;

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

// DESIGN PRINCIPLE
//		DESIGN TOWARDS INTERFACES RATHER THAN CONCERTE CLASSES
//		DISCOVER INTERFACES AS EARLY AS POSSIBLE
//		DRIVE CODE DESIGN USING INTERFACES

// DESIGN PRACTICE
//		Always Prefer COMPOSITION OVER INTHERITANCE

interface Superpower {
	void fly();
	void saveWorld();
}

// class Spiderman {
class Spiderman implements Superpower {
	public void fly() 		{ System.out.println("Fly Like Spiderman!"); 			}
	public void saveWorld() { System.out.println("Save World Like Spiderman!"); 	}
}

// class Superman {
class BadSuperman implements Superpower {
	public void fly() 		{ System.out.println("Fly Like BAD Superman!"); 			}
	public void saveWorld() { System.out.println("Save World Like BAD Superman!"); 		}
}

class Superman implements Superpower {
	public void fly() 		{ System.out.println("Fly Like Superman!"); 			}
	public void saveWorld() { System.out.println("Save World Like Superman!"); 		}
}

class Heman {
	public void fly() 		{ System.out.println("Fly Like Heman!"); 				}
	public void saveWorld() { System.out.println("Save World Like Heman!"); 		}
}

class Thor {
	public void fly() 		{ System.out.println("Fly Like Thor!"); 				}
	public void saveWorld() { System.out.println("Save World Like Thor!"); 		}
}

class Wonderwoman implements Superpower {
	public void fly() 		{ System.out.println("Fly Like Wonder Woman!"); 		}
	public void saveWorld() { System.out.println("Save World Like Wonder Woman!"); 	}
}

class HanumanJi implements Superpower {
	public void fly() 		{ System.out.println("Fly Like HanumanJi!"); 			}
	public void saveWorld() { System.out.println("Save World Like HanumanJi!"); 	}
}

// SOLID PRINCIPLE
//		VIOLATES S: Single Responsibility 
//		VIOLATES O: Open For Extension And Close For Changes
//		VIOLATES L: Liskov Substitution Principle
//		VIOLATES I: Interfaces Segregation
//		VIOLATES D: Dependency Investions

// DESIGN CHOICE 01
// Using Inheritance
// class Human extends Spiderman {
// class Human extends Superman {
class Human extends Heman {
	public void fly() 		{ super.fly(); 		 }
	public void saveWorld() { super.saveWorld(); }
}

// DESIGN CHOICE 02
// Using Composition
class HumanBetter1 {
	Spiderman power = new Spiderman();
	// Superman power = new Superman();

	public void fly() 		{ power.fly(); 		 }
	public void saveWorld() { power.saveWorld(); }
}

// SOLID PRINCIPLE
//		S: Single Responsibility 
//		O: Open For Extension And Close For Changes
//		L: Liskov Substitution Principle
//		I: Interfaces Segregation
//		D: Dependency Investions
// DESIGN CHOICE 03
// Using Composition

// DESIGN PRACTICES : BEST PRACTICES
// 		COMPOSITION IS EQUIVALENT TO INHERITANCE
// 		ALWAYS PREFER COMPOSITION OVER INHERTITANCE
//		IN COMPOSITION SUBSTRACTION IS POSSIBLE

// HumanBetter2 Is Polymporhic
class HumanBetter2 {
	// Spiderman power = new Spiderman();
	// Superman power = new Superman();
	Superpower power = null; 	// DESIGN DRIVEN BY INTERFACE
	public void fly() 		{ if ( power != null ) power.fly(); 	  }
	public void saveWorld() { if ( power != null ) power.saveWorld(); }
}

//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

public class HumanDemo {
	public static void playWithHuman() {
		Human h = new Human();
		h.fly();
		h.saveWorld();

		HumanBetter1 hb1 = new HumanBetter1();
		hb1.fly();
		hb1.saveWorld();

		HumanBetter2 hb2 = new HumanBetter2();
		// error: incompatible types: Spiderman cannot be converted to Superpower
		//		Fix Implement Superpower Inteface
		hb2.power = new Spiderman(); // Configuring As Spiderman
		hb2.fly();
		hb2.saveWorld();

		//error: incompatible types: Superman cannot be converted to Superpower
		//		Fix Implement Superpower Inteface
		// hb2.powe  = new BadSuperman(); // Configuring As Supeman
		hb2.power = new Superman();
		hb2.fly();
		hb2.saveWorld();

		hb2.power = new Wonderwoman(); // Configuring As Wonderwoman
		hb2.fly();
		hb2.saveWorld();

		hb2.power = new HanumanJi(); // Configuring As HanumanJi
		hb2.fly();
		hb2.saveWorld();
	}

	public static void main( String [] args ) {
		System.out.println("\nFunction : playWithHuman");
		HumanDemo.playWithHuman();

		// System.out.println ("Function : ")
		// System.out.println("Function : ")
		// System.out.println("Function : ")
		// System.out.println("Function : ")
		// System.out.println("Function : ")
		// System.out.println("Function : ")
		// System.out.println("Function : ")
	}
}

//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!
