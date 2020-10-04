using System;

class MainClass {
  public static void Main (string[] args){
    Console.WriteLine("Welcome to Password Generator:");
    Console.WriteLine("How many characters does your password need?");
    int length = Convert.ToInt32(Console.ReadLine());
    Random randNum = new Random();
    for(int i=0; i < length; i++){
      int c = randNum.Next(9);
    }
    Console.WriteLine(c);
  }
}
