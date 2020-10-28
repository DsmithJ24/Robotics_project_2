# Robotics_project_2

Group 2
Alan Lacerda
Daniel Smith
Nickson Dos-Santos

The code for the project is located in gobotLidar.go. There are many functions within the code: a drive function which moves the robot at a relatively slow pace to not 
overshoot the edges of a box, two turn functions (only the turn left is used), a findbox function to take readings until a box is found and stops, a measurement functions that starts a 
timer and runs until the robot passes the edge of the box and then calulates the distance traveled using the timer, and finally a main Loop which runs everything.
To ensure the robot gets precise measurements and to account for the nature of the lidar sensor, the robot takes 1 measurement every .5 (500 milliseconds) seconds and moves only .125 revolutions 
per second. Once the edge of the box is passes, the robot will drive forward slightly to create some separation from the box before the turn. However, our group was unable to 
implement feedback control such as PID in this code. As a result, the robot may drive into the box on occasion.
