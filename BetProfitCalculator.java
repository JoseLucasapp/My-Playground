import java.util.Scanner;

public class Main {
  public static void main(String[] args) {
    Scanner sc = new Scanner(System.in);
    System.out.println("How much money will you bet?");
    float bank = sc.nextFloat();
    System.out.println("What is price for home?");
    float home = sc.nextFloat();
    System.out.println("What is price for draw?");
    float draw = sc.nextFloat();
    System.out.println("What is price for away?");
    float away = sc.nextFloat();
  }
}
