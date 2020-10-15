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

//this loop takes readings every few seconds continuously. work should be done in here
func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,

) {

	drive(gopigo3)
	err := lidarSensor.Start()
	if err != nil {
  //no reding print following string
		fmt.Println("error starting lidarSensor")
	}
// continuing loop
	for {
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

		fmt.Println(lidarReading)
		fmt.Println(message)
		//Leave this speed its perfect
		time.Sleep(time.Second * 1)
		if lidarReading <=25 {
			gopigo3.Halt()
		}
	drive(gopigo3)
	err := lidarSensor.Start()
	if err != nil {
               fmt.Println("error starting lidarSensor")	

	}
		for {
		lidarReading, err := lidarSensor.Distance()
		if err != nil {
			fmt.Println("Error reading lidar sensor %+v", err)
		}
		message := fmt.Sprintf("Lidar Reading: %d", lidarReading)

		fmt.Println(lidarReading)
		fmt.Println(message)
		//Leave this speed its perfect
		time.Sleep(time.Second * 1)
		if lidarReading >25 {
			gopigo3.Halt()
			func turn_left(gopigo3 *g.Driver){
        gopigo3.SetMotorDps(g.MOTOR_RIGHT, 180)
        time.Sleep(time.Second)
        gopigo3.Halt()
		}
}

//these functions are from project 1, might be used again. may need to make them in above function with 'func:=' format)
func drive(gopigo3 *g.Driver) {
//        on := uint8(0xFF)
/*
		//do we need the flashing lights?
        gobot.Every(1000*time.Millisecond, func() {
				//do we need these flashing lights?
                err := gopigo3.SetLED(g.LED_EYE_RIGHT, 0x00, 0x00, on)
                if err != nil {
                        fmt.Println(err)
                }
                err = gopigo3.SetLED(g.LED_EYE_LEFT, ^on, 0x00, 0x00)
                if err != nil {
                        fmt.Println(err)
                }
                on = ^on
*/
        gopigo3.SetMotorDps(g.MOTOR_LEFT + g.MOTOR_RIGHT, 45)
        time.Sleep(time.Second)
//        gopigo3.Halt()
//        })
}
/*
func turn_left(gopigo3 *g.Driver){
        gopigo3.SetMotorDps(g.MOTOR_RIGHT, 180)
        time.Sleep(time.Second)
        gopigo3.Halt()
}

func turn_right(*gopigo3.Driver){
        gopigo3.SetMotorDps(g.MOTOR_LEFT, 180)
        time.Sleep(time.Second)
        gopigo3.Halt()
}
*/

//use gopigo3.halt to stop

//now need a function that handles the turning and resumes driving. corrections
func correction(){
	//needs to turn in direction of sensor
	//should also drive forward a bit to get an appropriate reading
	//needs to resume driving and tracking distance
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
