from random import randint
maximo = int(input('Quantas pessoas vão participar do sorteio ? '))
sorteadas = int(input('Quantas pessoas serão sorteadas ?'))

pessoas = []
numsorteados = []

for i in range(0,maximo):
  print(f'informe o nome da {i+1}° pessoa ')
  pessoa = input('-> ')
  pessoas.append(pessoa)

for a in range(0,sorteadas):
  sorteio = randint(0,(maximo - 1))
  while True :
    try:
      numsorteados.index(sorteio)
      sorteio = randint(0,(maximo - 1))
    except:
      break
  numsorteados.append(sorteio)
  print(f'A {a+1}° pessoa sorteada foi {pessoas[sorteio]}')
