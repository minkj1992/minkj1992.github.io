# [GCP] Iot core


referenced [Technical overview of Internet of Things](https://cloud.google.com/architecture/iot-overview) article.

<!--more-->
<br />

## tl;dr

## Terminology

- top-level components(3)

  - device: hardware itself and sw
  - gateway: connect to cloud services without internet
  - cloud: google cloud

- type of message

  - state information(data)
    - current status of the device
    - read/write
  - `telemetry`
    - device metric (i.g sensor data, gps ...)
    - read only
    - Each source of telemetry results in a channel
    - `Telemetry data` might be preserved as a stateful variable on the device or in the cloud.

- commnads

  - actions performed by a device.
  - commands are often not idempotent
  - which means each duplicate message usually results in a different outcome

- operational information
  - i.g. CPU operating temperature and battery state
  - can be transmitted as `telemetry` or state data.
- [serial interface](https://en.wikipedia.org/wiki/Serial_communication)
  - process of sending data one bit at a time, sequentially, over a communication channel or computer bus
  - contrast to parallel communication
- GPIO
  - General-purpose input/output pin
  - can be designed to carry digital or analog signals, and digital pins have only two states: HIGH or LOW.
- [PWM](https://en.wikipedia.org/wiki/Pulse-width_modulation)
  - pulse width modulation
  - The effect in the device can be a lower or higher power level
- ADC
  - analog to digital conversion
  - analog -> binary
- I2C
  - Inter Integrated Circuit
  - Inter-Integrated Circuit serial bus uses a protocol that enables multiple modules to be assigned a discrete address on the bus
  - pronounced "I two C", "I-I-C", or "I squared C".
- OTA update
  - over the air updates

## Next

- mqtt
- iot core gcp

