# Computer Science with Golang
My Computer Science notebooks with Golang 📝. From my bachelor's degree + Books I read recently + Some code snippets.

# What is a Computer ? 🤔
- It's your laptop or your mobile device. Or your Microwave! but let's have a concise definition of what a computer is : *any device that receives, stores, and process information*
- We have:
    - Different types of computers: desktops, servers, embedded systems
    - Different uses of computers: automobiles, graphics, finance, gaming
    - Different Manufactures of computers : Intel, Apple, Microsoft, IBM
    - Different underlying technologies and different costs

A computer consist of 💻: 

- Software
    - Operating System
    - Compilers
    - End user programs : Web Browsers, Email, Video Game
- Hardware
    - CPU, Main storage (DRAM,SRAM), secondary storage (Hard Drive, CD), Cache
    - I/O Devices : Keyboard, Monitor, Printer
- data : numbers, text, audio, video, images... (eventually it will be represented as binaries)

## Van Neuman Architecture

- So any computer may have all these things above or some of them. but we can summarize the underlying architecture as the **Von Neuman Architecture** which represent the components of any modern computer:
    - Input/output devices
        - Inputs : Keyboard, Microphone, Mouse
        - Outputs : Monitor, Printer
    - CPU : the brains of the computer, responsible of controlling the internal working of the memory.
    - Memory : stores instructions and data, they store information in binary digits or bits (0 or 1)
   
  ![van-01](https://user-images.githubusercontent.com/50620277/120696595-68203d00-c4b5-11eb-9d16-9f914d98bae9.png)


    ### Additional research and references

    [https://www.youtube.com/watch?v=Ml3-kVYLNr8&t=70s&ab_channel=Computerphile](https://www.youtube.com/watch?v=Ml3-kVYLNr8&t=70s&ab_channel=Computerphile)

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


# How computers stores information : Deep Dive Part 2 ⚡

![image](https://user-images.githubusercontent.com/50620277/120702472-c7ce1680-c4bc-11eb-95c6-f79732bc484c.png)

- types of data in the real world
    - Text
    - Numbers
    - Video
    - Audio
    - Image
- We understand now how computers store information in the form of binary digits, but how we map an image or any type of data to a binary digit ?

## Number Systems

- humans invented a system to represent quantity of things based on symbols, and these symbols are numerical symbols (0, 1,2,3...). The most common number systems are: Base 10 and Base 2
- In the human world we use base 10 to describe day to day numerical needs, and it's base 10 because we have 10 possible values 0-9 , why only base 10 not 12 or 13 ? most likely because we have 10 fingers ! so we started to count things based on our fingers.
- but in the computer world we don't have 10 possible values. remember , we only have 0s and 1s . So  the second most common number system is base 2, which only holds to possible values 0 or 1
- So any type of data, either it's an image , text, number or anything will be converted to binaries , and from binaries we could store that binary, process it, and convert it back to original human readable form whether it's audio or any other types of data
- In this illustration we can see how we represent numbers in both system (Base 10, and base 2)

![04](https://user-images.githubusercontent.com/50620277/120701876-17f8a900-c4bc-11eb-8e18-ba38fd74ac8d.png)

### Text
- a text is just a collection of characters, and representing characters in computers happens in 2 steps
    - A way to map a character to a number
        - We use [ASCCI](http://www.asciitable.com/) for able
        - Also we use [Unicode](https://home.unicode.org/) table for representing all languages and and their symbols
    - convert the number to a binary

### Images
- Images are collection of colors, and representing colors in a computer happens in 2 steps
    - A system to map a color to number
        - Most common system is RGB
    - converting the number to  binary digits

### Video
- A Video is a collection of images

Extra resources : 

- [Introduction to number systems and binary (video) | Khan Academy](https://www.khanacademy.org/math/algebra-home/alg-intro-to-algebra/algebra-alternate-number-bases/v/number-systems-introduction)
- [Decimal to Binary Converter](https://www.rapidtables.com/convert/number/decimal-to-binary.html)
- [CS50: Lecture 0](https://cs50.harvard.edu/college/2021/spring/notes/0/#:~:text=multiple%20times%20a%20second%20to%20give%20us%20the%20appearance%20of%20motion)


# 04 - Intro to Data Structure and Algorithms


![image](https://user-images.githubusercontent.com/50620277/120703577-23e56a80-c4be-11eb-9284-0fdad3658f37.png)

- An **algorithm** is every function you created in Golang or any other programming language you programmed in. Any steps (programming statements) you wrote to solve a problem. And all programming problems is about processing data.  So, in Computer Science there is field of Algorithms. which focuses on dealing with processing data at scale, and analyzing the performance of an algorithm to solve programming problems such as sorting, searching, removing and so on. 

![05](https://user-images.githubusercontent.com/50620277/120702938-5a6eb580-c4bd-11eb-8156-bf5c6aa4ab0c.png)

- So before we talk about Algorithms, let's talk about a data that a program might have
- First of all any data in a program has a type. A data type help us define 2 things
    - a domain of allowed values
    - a set of operations on these values
- And a program compiler will through an compilation error if we misuse a data type, for example performing multiplication on chars
    - z = x*y is not allowed
    - true/false is not allowed

![06](https://user-images.githubusercontent.com/50620277/120702996-69edfe80-c4bd-11eb-866d-a405de0e5342.png)

### Atomic / Simple  Data

- Have no component parts. E.g int, char , float, etc

### Data Strcture

- types that can be broken into parts. for example a Person object consist of two properties : name and an age, both of them are atomic data types (string, int respectively )
    - object ⇒ broken into properties
    - array ⇒ broken into indeces
- So a data structure is a data type whose values
    - can be decomposed into a set of atomic data or another data structure (2D Array)
    - include a structure involving the component parts (a fixed technique on how to search, store, remove and so on)
    - Possible Structures: Set, Linear, Tree, Graph

### Abstract Data Types (ADTs)

- Provide a level of Abstraction : I want to use a data structure  but I don't want to know how it's implemented, for example in Golang you can append to a slice but you don't care of the implementation at the time of using the slice data structure

## How to create your own data structure

While designing ADTs, a designer has to deal with two types of questions:

- **What** values are in the domain? **What** operations can be performed on the values of a particular data type?
- **How** is the data type represented?
- **How** are the operations implemented?


  # Performance Analysis Part 1: Intro 🔍
- Generally, the performance factors of a program
    - Code Algorithims
    - Languages + compilers
    - Processer (CPU) speed + Memory ⇒ how fast instruction will be executed
    - I/O and number of CPU cores
- How we can determine the performance of an algorithm.
- Or compare algorithms and determine which one is better ?
- first, we need to determine what we mean of "better"

### Experimental analysis

### Algorithms performance

- time complexity ⇒ the time it takes to execute
- space complexity ⇒ the memory it needs to execute
- We will focus on time complexity right now

### There are 2 ways to compare the complexity of algorithims

- Experimental analysis ⇒ compare the running time (benchmarking)
    - inputs ⇒ array or objects or any data
- Theoretical analysis ⇒ analyze the algorithms independently  of the implementation (software or hardware)


# Performance Analysis Part 2: Theoretical Analysis
- High-level description of the algorithm instead of an implementation (benchmarking/running-time)
- Characterize running time as function of the input size, n
- takes into account all possible inputs (int, string, float...)
- Allows us to evaluate the speed of an algorithm independently of the hardware/software environment


### Scenarios
- So an algorithm depends on the inputs of size n, so the time complexity or running time will vary from input to input. for example sorting an array of 10 numbers is much less time consuming then sorting an array of 100000 numbers! maybe your algorithm in the best vase scenario will take 100 array elements to be sorted. BUT you always count the worst case scenario. which is maybe 10000000 elements. remember, when your input grow your algorithm time and space complexity will increase. and may increase in a huge way
    - There is 3 types of possible scenarios
        - best case : you count it maybe through some sort of statistics
        - average case :  you count it maybe through some sort of statistics
        - worst case  : you defiantly to imagine the worst case scenario (what if the inputs scaled)
    - **We focus on the worst case running time**
        - easier to analyze
        - critical to games, finance and robotics
        
### Let's see how we calculate the time-complexity using Big oh notation

![07](https://user-images.githubusercontent.com/50620277/120809007-52ad2080-c552-11eb-9eef-9f8cd3e8d624.png)

- An algorithm is said in-place if it does not require more than O(log n) space in addition to the
input.
- An algorithm is said in-place if it does not require more than O(log n) space in addition to the
input.
