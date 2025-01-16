defmodule PrimeFinder do
  # Function to check if a number is prime
  def is_prime?(2), do: true
  def is_prime?(n) when n < 2, do: false
  def is_prime?(n) do
    max_divisor = :math.sqrt(n) |> floor()
    Enum.all?(2..max_divisor, fn divisor -> rem(n, divisor) != 0 end)
  end

  # Function to find prime numbers in a given range asynchronously
  def find_primes_async(range) do
    tasks =
      range
      |> Enum.map(fn n -> Task.async(fn -> {n, is_prime?(n)} end) end)

    tasks
    |> Enum.map(&Task.await/1)
    |> Enum.filter(fn {_, is_prime} -> is_prime end)
    |> Enum.map(fn {n, _} -> n end)
  end
end

# Example usage

# Define the range to search for prime numbers
range = 1..100

# Find prime numbers in the range asynchronously
primes = PrimeFinder.find_primes_async(range)

# Print the prime numbers
IO.inspect(primes)
