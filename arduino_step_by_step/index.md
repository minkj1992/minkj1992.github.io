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



### GPIO
#### Digital Pins

- output
![](/images/arduino/digitalOutputPins.png)

- input
![](/images/arduino/digitalInputPins.png)

#### Analog pins

- output
![](/images/arduino/analogOutPins.png)


{{< admonition note>}}

PWM(Pulse Width Modulation)
- PWM is not true analog output, however. PWM “fakes” an analog-like result by applying power in pulses, or short bursts of regulated voltage.
- 0과 1로만 이뤄진 디지털 신호를 아날로그 신호처럼 흉내낸다.
- 즉 중간 값들을 만들어 낼 수 있다.

{{< /admonition >}}

- input
![](/images/arduino/analogInPins.png)

- Potentiometer(가변저항)
  - `Ohm's law`: V(전압) = I(전류) * R(저항)
  - 저항이 커지면 전압이 강해진다.
  - 직관적으로 이해하기에는 일정한 부피에 사람들 잔뜩 집어넣으면 밀도가 올라가서 압력이 쎄진다고 생각함.
  - **흐르는(일정한 전류) 호수관 입구 손으로 막으면(저항) 압력 더 쎄져서 물 파워(전압) 쎄짐**

![](/images/arduino/potentiometer.gif)

`pot`는 시계방향으로 이동시키면 저항길이가 길어져서 더 큰 저항을 줄 수 있다.

## Section 3: Introduction to communications

### UART
> Universal asynchronous receiver/transmitter

병렬(parallel) 데이터의 형태를 직렬(async serial) 방식으로 전환하여 데이터를 전송하는 컴퓨터 하드웨어의 일종이다.


![](/images/arduino/UART.png)

- 두 장치 간에 직렬 데이터를 교환하기 위한 프로토콜 또는 규정을 정의합니다. UART는 매우 간단하며 양방향으로 데이터를 송신 및 수신하기 위해 송신기와 수신기 사이에 두 개의 와이어만 사용합니다.
- It is old fashion compare to I2C and SPI.
- simple, cheap, easy to make

{{< admonition tip "USART">}}
`USART`(Universal Synchronous serial Receiver and Transmitter)를 처리하기 위해서는 동기적으로 송수신 타이밍이 동기화 되어야 한다. 이런 동기화를 위해서 Clock 신호 라인이 필요하다.

클럭신호로 HIGH, LOW를 한번 반복해 보내는 동안 데이터 핀으로 한 비트의 데이터를 보내게된다. 다음은 100(2진수로 01100100)이라는 값을 보내는 방식이다. 아래를 보면 CLOCK 라인이 존재하는 것을 확인할 수 있다.

![](/images/arduino/usart.png)

{{< /admonition >}}

### I²C(TWI, I2C)
> Inter-Integrated circuit
> Two Wire(sda, sck) Interface 
> 아이 스퀘어 C

![](/images/arduino/I2C.png)

I²C는 데이터를 주고 받는 선 하나와 송수신 타이밍 동기화를 위한 CLOCK 선하나로 이루어진다.

- Master-Slave 구조
- line(2)
  - SDA(data): 데이터 송수신
  - SCK(clock): Clock

시작 신호와 정지 신호를 가지고 있으며, 슬레이브마다 지정된 주소 값을 가지고 데이터를 주고 받는다. **데이터를 주고 받을 때 반드시 주소 값을 붙여서 보내야 하므로 긴 데이터를 보내기에는 적합하지 않지만**, 통신 타이밍에 구애 받지 않으며 두 개의 선만으로 여러 기기와 통신할 수 있다는 장점 있다.


### SPI
> Serial Peripheral(주변의) Interface

![](/images/arduino/SPI.jpeg)

- line(4)
  - MOSI: Master Out, Slave In
  - MISO: Master In, Slave Out
  - SCK: Clock Line
  - SS(CS): Slave Select or Chip Select
    - 즉 데이터를 수신할 기기 선택 신호

- Full Duplex
- Master-Slave 구조
- SPI 통신은 데이터의 송신과 수신이 동시에 이루어지기 때문에 다른 데이터 통신에 비해 속도가 빠르다는 장점이 있어서 이더넷 통신 등에 주로 사용된다.
- 하나의 마스터에는 여러 개의 슬레이브가 연결 될 수 있지만, 슬레이브마다 각각 하나의 SS 선을 필요로 하기 때문에 슬레이브의 개수가 많아질 경우에는 물리적으로 효율적이지 않다.
