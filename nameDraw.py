from random import randint

pessoas = []
numsorteados = []
continuar = 'S'
n = ''
while True:
  print('')
  print(f'Iniciando {n}sorteio')
  print('')
  maximo = int(input('Quantas pessoas vão participar do sorteio ? '))
  sorteadas = int(input('Quantas pessoas serão sorteadas ?'))

  for i in range(0,maximo):
    print(f'informe o nome da {i+1}° pessoa ')
    pessoa = input('-> ')
    print('')
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
    if sorteadas > 1:
      print(f'A {a+1}° pessoa sorteada foi {pessoas[sorteio]}')
    else:
      print(f'A pessoa sorteada foi {pessoas[sorteio]}')
    print('')
  continuar = input('deseja sortear novamente? (S/N)').upper()
  if continuar == 'N' or continuar =='NAO' or continuar == 'NÃO':
    break
  else:
    n = 'novo '
