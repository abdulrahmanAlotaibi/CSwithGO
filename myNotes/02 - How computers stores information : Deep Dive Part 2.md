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
