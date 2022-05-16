# Coursera Gcp Iot


This post is a summary of [lectures](https://www.coursera.org/learn/iiot-google-cloud-platform/lecture/sMtHo/industrial-iot-on-google-cloud) taken in coursera on the topic of `Google IoT core`.

<!--more-->
<br />

## Day 1 ~ 2

### What is IoT?

Internet of Things is a sprawling set of technologies and use that has no clear, single definition.

But one workable view frames IoT as the use of network-connected devices, embedded in the physical env, to improve some existing process or to enable a new scenario not previously possible.

### IoT cloud

The general structure of an IoT cloud

1. Devices
   1. gathering data
   2. performing an action
   3. interact with users
2. Gateway
   1. A gateway ensures that devices are securely connected to the cloud. It can be a cell phone.
   2. It controls messaging between the device and the cloud
3. Cloud
   - u na mean

### Challenges in IoT

When designing an IoT network, below topics should be considered.

- connectivity
  - At present IoT relies on a server/client model to authenticate, authorize and connect devices to nodes in the network
  - It will become unworkable as numbers grow to the millions and billions per network.
  - In the future, off-loading tasks to the edge will become important. This means that IoT networks will need devices capable of handling data analysis, ML, and data gathering.
- brownfield deployment(legacy infra)
  - Companies will need to confront the task of integrating new devices and technologies into existing networks.
- security and compilance
- non-standard communication protocol
  - Dealing with non-standard communication protocol will
    increase in importance as networks need to handle
    ever-increasing amounts of data from sensors. Data
    handling, processing, and storing will increase as data input
    loads increase, while at the same time, the value of data
    increases with the size, depth, and frequency of data
    available to data analytics.
- IT/OT convergence
  - IT has traditionally been data-centric
  - OT has been used to monitor events.
- get actionable intelligence from data.

### IoT Architecture

IoT architecturesmust be capable of scaling connectivity of devices, data ingestion, data processing, and data storage.
They must be able to do this quickly while still producing real-time data insights.

To migrate this demand, distributed computing known as fog or edge computing is gaining popularity. The edge refers to the geographic distribution of computing nodes in the network as sIoT devices, which are at the "edge" of a network. This in turn increases the demand for devices that are capable of cleaning, processing and analyzing data locally. The result is that only cleaned metadata is sent to the cloud.

An asynchronous, scalable communication stack is crucial in bidirectional communication with devices.

https://k21academy.com/google-cloud/cloud-sql-vs-cloud-spanner/

## Day 3

### Types of sensors

![](/images/types_of_sensors.png)

### Devices

A "Thing" in the "Internet of Things" is a
processing unit that is capable of connecting to the internet and exchanging data with the cloud.

Devices are often called "smart devices" or
"connected devices." They communicate two
types of data: telemetry and state.

`Device information` is mainly composed of three types as follows.

1. Device metadata: Most metadata rarely, if ever, changes.

   - Identifier (ID) -An identifier that uniquely identifies a device.
   - Class or type
   - Model
   - Revision
   - Date manufactured
   - Hardware serial numbe

2. Telemetry: Data collected by the device is called telemetry

   1. Telemetry is read-only data about the environment
   2. usually collected
      through sensors
   3. e.g. Temperature(35.4oC)

3. State information
   1. Describes the current
      status of the device, not of the
      environment
   2. can be read/write. It is updated, but usually not frequently.

