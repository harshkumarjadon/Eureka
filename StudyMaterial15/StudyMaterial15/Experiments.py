
class Person:
  # Constructor In Python Is __init__() Function
  #   Every Member Function/Method Explicitly Takes self As First Argument
  #   self In Python Is Same As this In C++/Java
  def __init__(self, name, age):
    self.name = name
    self.age = age

  def doDance(self):
    print("Doing Donce...")

  #def doDanceAgain():
  def doDanceAgain(self):
    print("Doing Donce Again...")


gabbar = Person("Gabbar Singh", 36)

print(gabbar.name)
print(gabbar.age)

gabbar.doDance()
gabbar.doDanceAgain()

