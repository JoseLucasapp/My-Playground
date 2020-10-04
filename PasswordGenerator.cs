using System;

class MainClass {
  public static void Main (string[] args){
    Console.WriteLine("Welcome to Password Generator:");
    Console.WriteLine("How many characters does your password need?");

    Random randNum = new Random();

    int password_length;
    int length = Convert.ToInt32(Console.ReadLine());
    string max = "10";

    for(int i = 1; i < length; i++){
      max += "0";
    }

    password_length = Convert.ToInt32(max);
    int password = randNum.Next(password_length);
    Console.WriteLine(password);
  }
}
