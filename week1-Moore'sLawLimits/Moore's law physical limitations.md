# Moore's law physical limitations

[TOC]

## What's Moore Law

Moore's law is an *observation* and projection of a historical trend, rather than a law of physics.
The observation is named after Gordon Moore, who in 1965 formulated  that the *number* *of* *transistors* in integrated circuits doubles about every two years (18th months to be precise). 

This advancement is important for other aspects of technological progress in computing - such as processing speed or price of computers. 

:bulb: Smaller transistors cost less and go switch faster than bigger ones.

Moore's law closely relates to MOSFET scaling, as the miniaturisation of MOSFETs is the key driving force behind Moore's law.

#### Enabling Factors

For five decades, Moore’s law *held up pretty well*, scientific breakthrough in advanced integrated circuit and semiconductor device fabrication technology, allowed transistor counts to grow by more than seven orders of magnitude. 

Today's fabrication methods allow to scale the size of a MOSFET (the most used transistor type in IC) to 10 and 7nm, however due to the physical limitation to Moore's Law,  much of the semiconductor industry has shifted its focus to the needs of major computing applications rather than semiconductor scaling.



## Physical Limitations

1. **High chip temperature**: Increasing the density of transistor naturally increases power consumption which leads to higher temperatures being produced  by IC. 

2. **Power leakages** (source-drain tunneling):  transistors leak power even if they do aren't switching,  this happens because of reducing the insulator *size*  in a transistor. The thinner it is , the more leakage power exists (electricity passes) between transistors conductors. In particular, electron tunnelling prevents the length of a gate from being smaller than 5 nm.

3. **Dennard scaling limitations**: scaling a transistors size goes hand to hand with scaling its voltage, however the limitation here is that *voltage can't go too low* because: 
   
   ​	A. The voltage swings between low and high must be higher its switching threshold  
   ​	B. Noises in signal  become a considerable disturbance, lowering the noise tolerance of the IC
   
4. **Transistors manufacturing limitations**: The ability to print ever-smaller lines on silicon wafers (lithography) is the other driving force for Moore's Law.

