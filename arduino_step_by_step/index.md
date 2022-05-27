# Arduino step by step


Learn arduino with [Arduino Step by Step: Getting Started by Dr. Peter Dalmaris](https://www.udemy.com/course/arduino-sbs-17gs) lecture.
<!--more-->
<br />

## tl;dr

- [ ] Build simple circuits around the Arduino Uno, that implement simple functions.
- [ ] Understand what is the Arduino.
- [ ] Understand analog and digital inputs and outputs
- [ ] Use the multimeter to measure voltage, current, resistance and continuity
- [ ] be productive with the Arduino IDE, write, compile and upload sketches, install libraries
- [ ] Detect and measure visible light, color, and ultraviolet light
- [ ] Measure the distance between the sensor and an object in front of it
- [ ] Detect a noise
- [ ] Display text on a liquid crystal display
- [ ] Write simple Arduino sketches that can get sensor reading, make LEDs blink, write text on an LCD screen, read the position of a potentiometer, and much more.
- [ ] Understand what is prototyping.
- [ ] Understand the ways by which the Arduino can communicate with other devices
- [ ] Use protoboards to make projects permanent
- [ ] Understand what is Arduino programming, it's basic concepts, structures, and keywords
- [ ] Measure temperature, humidity and acceleration
- [ ] Detect a person entering a room
- [ ] Make noise and play music


## Section 2: Know your Arduino

### Getting to know the Arduino Uno

![](/images/arduino/full_arduino.png)

- `Arduino Uno`
  - The Arduino UNO is the best board to get started with electronics and coding. If this is your first experience tinkering with the platform, the UNO is the most robust board you can start playing with. The UNO is the most used and documented board of the whole Arduino family.
  - Arduino Uno is a microcontroller board based on the ATmega328P

![](/images/arduino/uno.webp)

- Atmega328P
  - The ATmega328 is a single-chip `microcontroller`(=mcu) created by Atmel in the megaAVR family

![](/images/arduino/Atmega328P.png)

- Shields
  - add-on, plugin
  - Shields are modular circuit boards that piggyback(하나의 운송 단위를 다른 운송 수단에 싣고 운반하는 상품 운송, 어부바) onto your Arduino to instill it with extra functionality.
  - In general, these are called "daughter boards."

![](/images/arduino/shield1.png)

![](/images/arduino/shield2.png)


- USB
- Pins
- power
- clock
- `ATmega16U2` (https://mosesnah.tistory.com/5)
  - USB-to-Serial Converter

![](/images/arduino/ATmega16U2.png)
![](/images/arduino/ATmega16U2-2.png)

- GPIO
  - 

### Using the digital pins

![](/images/arduino/digitalOutputPins.png)

![](/images/arduino/digitalInputPins.png)

### Using the analog pins

![](/images/arduino/analogOutPins.png)

- PWM(Pulse Width Modulation)

![](/images/arduino/analogInPins.png)

- Potentiometer(가변저항)

위 그림에서 저항 강할 수록 Ground로 전달되니 더 강한 Input을 이끌어 낼 수 있다.

## Section 3: Introduction to communications

- UART
- I2C(TWI)
- SPI


