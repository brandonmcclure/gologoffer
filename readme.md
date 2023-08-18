# gologoffer

A dead simple proof of concept go program designed to kick a user off if they are on the computer during times they should not be.



## Genesis

The state of parental controls for linux makes me sad. There are several projects that offer UI approaches. I want a more low level configuration I can setup to have consistent behavior across all my systems. I have [pam_time](https://linux.die.net/man/8/pam_time) configured, however if there is already an active session then I need to kill that tty session. Maybe each tty session if they are clever.

Honestly, I still have a windows computers as well! I have had little luck in finding a cross platform solution to the problem of keeping a user off the computer between a time range. 

I am poking around with golang, so figured I would compile an over engineered solution vs a simple powershell (or bash...) script so that I could learn github's go ci ecosystem to get me a executable I can run on multiple architectures.