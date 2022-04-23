# [GCP] Iot core




<!--more-->
<br />

## tl;dr

## Terminology
> referenced [Technical overview of Internet of Things](https://cloud.google.com/architecture/iot-overview) article.


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

## GCP IoT Core conecpts
> [Devices, Configuration and State](https://cloud.google.com/iot/docs/concepts/devices)

### 1. Device metadata
> metadata serves primarily as a label or identifier for devices. (or classifies devices)

- more secure than device state or device configuration because device metadata is never sent to or from a device
- shouldn't change often (best: update it no more often than once per day)
- 500 key-value pairs (each key must be unique)
- Cloud IoT Core does not interpret or index device metadata
- e.g. hardware thumbprint, serial number, manufacturer information

### 2. Device configuration
> IoT Core → device
> Sends desired state to robot with pre-defined commands like [e.g.](https://cloud.google.com/iot/docs/concepts/devices#structuring_configuration_data)

- Device configuration is an arbitrary, user-defined blob of data sent from Cloud IoT Core to a device
- Device configuration is persisted in storage by Cloud IoT Core (64KB)
- After a configuration has been applied to a device, the device can report its [state](https://cloud.google.com/iot/docs/how-tos/config/getting-state) to Cloud IoT Core.
- A device configuration should focus on desired values or results, rather than on a sequence of commands
- Updates a device's state by sending the expected state as a configuration
- **Note that a device is not guaranteed to receive every configuration update**
  - If a configuration is being updated rapidly, devices may not receive intermediate versions.

#### Configuration versions

- A device receives configurations only in increasing order of version numbers; in other words, it will never be sent a configuration older than its current version
- If the device reconnects to the MQTT bridge, it may receive an older configuration than it did during the earlier connection (rare case)

### 3. Device state
> device → IoT Core
> Captures the current status of the device, not the environment

- Devices can describe their state with an arbitrary user-defined blob of data sent from the device to the cloud
- Device state information is not updated frequently.
- `Configuration` and `state` data can have the same schema and encoding, or they can be different
- e.g. health of the device or its firmware version


