using System;

class MainClass {
  public static void Main (string[] args) {
    var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    Console.WriteLine("Welcome to Password Generator:");
    Console.WriteLine("How many characters does your password need?");

    int length = Convert.ToInt32(Console.ReadLine());
    Random random = new Random();
    
    int generate;
    var password = "";

    for (int i =0; i < length; i++){
      generate = random.Next(36);
      password = password + chars[generate];
    }

    Console.WriteLine("Your Password: " + password);
  }
}
