package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)
//Robot work
//this loop takes readings every few seconds continuously. work should be done in here

func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,

) {
/*
	drive(gopigo3)
	err := lidarSensor.Start()
	if err != nil {
		fmt.Println("error starting lidarSensor")
	}
// 1st loop, turn it into a function that finds the box
	for {
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

		fmt.Println(lidarReading)
		fmt.Println(message)
		time.Sleep(time.Second * 1)

		if lidarReading <=25 {
			gopigo3.Halt()
			break;
		}

	}
*/

// all this in a for loop
	for{
	//start by finding the box
		findbox(gopigo3, lidarSensor)
		time.Sleep(time.Second *1)
		fmt.Println("Found Box")

	//now box is found, drive and take measurements
		measurement(gopigo3, lidarSensor)
		time.Sleep(time.Second*1)
		fmt.Println("Edge of box")

	//reached the edge, start turning
		turn_left(gopigo3)
		fmt.Println("Turn complete")

	//find the box again, loop back
	}

/*
	drive(gopigo3)
	//2nd loop, becomes a function that measures box
	for{
		lidarReading, err := lidarSensor.Distance()
                if err != nil {
                       	fmt.Println("Error reading lidar sensor %+v", err)
                }
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

                fmt.Println(lidarReading)
		fmt.Println(message)
                time.Sleep(time.Second * 1)

		if lidarReading >= 25{
			gopigo3.Halt()
			fmt.Println("Edge of box")
		}

		//either 3rd function or coode that turns robot
		//and then calls the 1st loop to find box, then 2nd loop
	}
	*/
}

//these functions are from project 1, might be used again. may need to make them in above function with 'func:=' format)
func drive(gopigo3 *g.Driver) {
        gopigo3.SetMotorDps(g.MOTOR_LEFT + g.MOTOR_RIGHT, 45)
        time.Sleep(time.Second)
//        gopigo3.Halt()
//        })
}

func turn_left(gopigo3 *g.Driver){

//drive forward one diameter of wheel length (~55 deg). then turn left
        gopigo3.SetMotorDps(g.MOTOR_LEFT + g.MOTOR_RIGHT, 40)
	time.Sleep(time.Second)
	gopigo3.Halt()

        gopigo3.SetMotorDps(g.MOTOR_RIGHT, 335)
        time.Sleep(time.Second)
        gopigo3.Halt()
}
/*
func turn_right(*gopigo3.Driver){
        gopigo3.SetMotorDps(g.MOTOR_LEFT, 180)
        time.Sleep(time.Second)
        gopigo3.Halt()
}
*/

//use gopigo3.halt to stop

//now need a function that handles the turning and resumes driving. corrections
func findbox(gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,){
	//stuff from 1st loop. drives to find the box
	drive(gopigo3)
	err := lidarSensor.Start()
	if err != nil {
		fmt.Println("error starting lidarSensor")
	}
// 1st loop, turn it into a function that finds the box
	for {
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

		fmt.Println(lidarReading)
		fmt.Println(message)
		time.Sleep(time.Second * 1)

		if lidarReading <=40 {
			gopigo3.Halt()
			break;
		}
	}
}

func measurement(gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,){
	//stuff from second loop. drives around side of bax and takes measurement
	drive(gopigo3)
	//2nd loop, becomes a function that measures box
	for{
//		startTime := time.Now()

		lidarReading, err := lidarSensor.Distance()
                if err != nil {
                       	fmt.Println("Error reading lidar sensor %+v", err)
                }
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

                fmt.Println(lidarReading)

		fmt.Println(message)
                time.Sleep(time.Second * 1)

		if lidarReading >= 40{
			gopigo3.Halt()
//			endingTime:= time.Since(startTime)
//			duration:= time.Since(startTime)
			//need to figure out measureDPS
//			side := duration.Seconds() * float64(measureDPS) *.05803
			break;
		}

		//either 3rd function or coode that turns robot
		//and then calls the 1st loop to find box, then 2nd loop
	}
}

func main() {
	raspberryPi := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspberryPi)
	lidarSensor := i2c.NewLIDARLiteDriver(raspberryPi)
	lightSensor := aio.NewGroveLightSensorDriver(gopigo3, "AD_2_1")
	workerThread := func() {
		robotMainLoop(raspberryPi, gopigo3, lidarSensor)

	}
	robot := gobot.NewRobot("Gopigo Pi4 Bot",
		[]gobot.Connection{raspberryPi},
		[]gobot.Device{gopigo3, lidarSensor, lightSensor},
		workerThread)

	robot.Start()

}
