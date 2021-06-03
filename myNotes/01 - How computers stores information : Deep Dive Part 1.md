# How computers stores information : Deep Dive Part 1 ⚡

![image](https://user-images.githubusercontent.com/50620277/120699242-d3b7d980-c4b8-11eb-8159-49daa5f5ee99.png)

- computers are digital devices which means that they only store information in the shape of binary digits. 
A single binary digit is a unit of data that corresponds to one of two potential values : 0 or 1 . which is just means true or false ,
open(1) or close(0) , volt(1) or no-volt(0).
- But again, what is open(1) and close(0) ?
- All the underlying hardware of computers (CPU, RAM, hard-drive, Motherboard...) has been constructed using physical logic gates

## what are logic gates ?
- Before we talk about logic gates, let's talk about Boolean algebra and how it impacted significantly the computer science field

## Basics of Boolean algebra 
- in 1845, a British self-taught mathematician called George Boole Published a book entitled "An Investigation of The law of Thought" which describes the mathematical theories of logic and probabilities. And this is basically Boolean algebra, the foundation of computers today.
- Boolean algebra is a mathematical branch in which you express logic or normal human reasoning mathematically. 
- Boolean algebra only has 2 mathematical operations which is add and multiplication, these operations  are associated with OR gate and AND gates
- Boolean Algebra operations: Add , multiplication .. let's see how it works !

<div style="text-align:center">
 <img src="https://user-images.githubusercontent.com/50620277/120697266-2e9c0180-c4b6-11eb-81ab-36c2de3a65da.png"/>  
</div>

## Logic Gates ⚡🔌
- We represent all types of information(video, audio, images , documents) using binary digits which is just electrical signals, then it will be carried through a bus (physical wire), the combinations of binary digits will allows us to represent any needed operation and type of data. so combine binary digits and manipulate them through logic gates which opens the the actual gate for the binary signal to pass through or to close the gate on the signal
- A logic gate is defined as a electronics circuit with two or more input signals and one output signal (s)
  - Electronic gates require a power supply.
  - All the inputs of logics gates are binaries 0 (voltage of 0V) and 1 (voltage of 5V), and the output for a logic gate is a single binary 
  - Gate Inputs (A,B) are driven by voltages having two nominal, values, e.g. 0V and 5V representing logic 0 and logic 1 respectively.
  - in 1938 Claude Shannon, research assistant in a electrical collage published his master's thesis entitled " A symbolic Analysis of Relay and switching circuits". As result of his work, Boolean algebra is the foundation of circuit designs in computers today.  

### Logic Gates Implementation : AND , OR, NOT

![image](https://user-images.githubusercontent.com/50620277/120699633-517be500-c4b9-11eb-84c6-dfd5cb08d7d8.png)

### References and Additional Research 📑:

- [AND OR NOT - Logic Gates Explained - Computerphile](https://www.youtube.com/watch?v=UvI-AMAtrvE&ab_channel=Computerphile)

- [Basic Logic Gates](http://www.ee.surrey.ac.uk/Projects/CAL/digital-logic/gatesfunc/index.html)

- [A book chapter about boolean algebra](https://www.pbte.edu.pk/text%20books/dae/math_123/Chapter_11.pdf)

