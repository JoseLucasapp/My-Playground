import 'dart:io';

consoleMsg(message){
  print(message);
}
write(message){
  stdout.write(message);
}
read(){
  return stdin.readLineSync();
}

output(name, age, city, country){
  consoleMsg('Name: $name, Idade: $age, City: $city, Country: $country');
}

questions(){
  write('Name: ');
  var name = read();
  write('Age: ');
  var age = read();
  write('City: ');
  var city = read();
  write('Country: ');
  var country = read();
  output(name, age, city, country);
}

void main(){
  questions();
}