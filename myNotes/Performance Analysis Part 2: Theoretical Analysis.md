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
