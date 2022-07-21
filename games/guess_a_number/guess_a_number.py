from random import randint

def run_game(low, high):
  number = randint(low, high)
  while True:
    guess = int(input("Guess a number between " + str(low) + " to " + str(high) + ": "))

    if guess == number:
      print("Correct!")
      break
    if guess > number:
      print("Too high!")
    if guess < number:
      print("Too low!")

def main():
  run_game(0, 10)

main()