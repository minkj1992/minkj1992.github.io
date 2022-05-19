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

### Device commands

- `Commands` are actions performed by a device. So tye should include a time-to-live(TTL) or other expiration value

- `Operational information` is data that's most relevant to the operation of the device as opposed to the business application.
   - This might include things such as CPU operating temperature and battery state.
   - it has short-term value to help maintain the os
   - **It can be transmitted as `telemetry` or `state data`**

### Defining devices

For example, consider a project for monitoring the temperature in hotel rooms. Each room
has three sensors: one near the floor, one near the bed, and one near the ceiling.

option1

```json
{deviceID: "dh28dslkja", "location": "floor", "room": 128, "temp": 22 }
{deviceID: "8d3kiuhs8a", "location": "bedside", "room": 128, "temp": 24 }
{deviceID: "kd8s8hh3o", "location": "ceiling", "room": 128, "temp": 20 }
```

option2
```json
{deviceID: "dh28dslkja", "room": 128, "temp_floor": 22, "temp_bedside": 24, "temp_ceiling": 20,
"average_temp": 22 }
```

### Google IoT developer prototyping kits

Google works with partners to build device starter kits that make connecting to Google Cloud IoT Platform easy for developers. At this time, Google has partnered with fourteen companies to offer a wide variety of [IoT developer prototyping kits](https://d3c33hcgiwev3.cloudfront.net/r0yj9QAaRl2Mo_UAGjZdFw_354538e7a78d48d79048afcdd3b1bea1_10.-Google-IoT-Developer-Prototyping-Kits.pdf?Expires=1653091200&Signature=gkcQBVgHOXNjN3ueT4CtMIMAfaZYDouIdi8G7RHaBCi~hhNKrI~c483dhbf~UtRPeF5ZBR6qryCvRXNzdh0PYQk1ih2o3QgmhAMqLqCuCjK9DCWsG2tfVSVW--04PgzzrvFnK-GpIKiv8SHze8Xwzkif~fFax-WAZ5j-1euA7W4_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A).

### MQTT protocol
> https://cloud.google.com/iot/docs/concepts/protocols

MQTT is an industry-standard IoT protocol (Message Queue Telelmetry Transport).

![](/images/mqtt.png)

Messages include the topic in the message, which is used for routing information by the broker. **This means that subscribers do not need to know the publisher, because all communication is done through messages**

- Messages are pushed to subscribers, so there must be an open TCP connection to the broker (subscriber <-> broker)
   - If the connection is broken, the broker can hold messages for later transmission.

### HTTP protocl

> HTTP is a "connectionless" protocol: with the HTTP bridge, devices do not maintain a connection
to the cloud.

In connectionless communication, client requests are sent without having to first check that the recipient is available. Therefore devices have no way of knowing whether they are in a conversation with the server, and vice versa. 

**This means some of the features that Cloud IoT Core provides, for example, last Heartbeat detected, will not be available with an HTTP connection.**

### Comparison of MQTT and HTTP general features

> MQTT is considered to be data focused, while HTTP is document focused. Which means **MQTT is better suited to the rigors of IoT**.

![](/images/mqtt_vs_http.png)

In addition, MQTT has three levels of service(QoS)

1. At most once.
   - Guarantees at least one attempt at delivery, (no guarantee of delivery).
2. At least once.
   - Guarantees the message will be delivered at least once.
3. Exactly once.
   - Guarantees the message is delivered only once.

MQTT also has

- `Last will and testament`(유언, `LWT`). If a client (ie device) is disconnected unexpectedly, the subscribers will be notified by the MQTT broker.
- **Retained(보관) messages. New subscribers will get an immediate status update.**

{{< admonition note "LWT" >}}
통신에서 중요하지만 구현이 까다로운 문제로 "상대방이 예상치 못한 상황으로 인하여 접속이 끊어졌을때"의 처리가 있다.

그래서 전통적 방식으로는 자신의 생존 여부를 계속 ping을 통해 서버가 물어보고 timeout 시간안에 pong이 안올 경우 서버에서 접속 종료를 인식하는 번거로운 방식을 취하는데,

**MQTT의 경우 subscribe 시점에서 자신이 접속 종료가 되었을 때 특정 topic으로 지정한 메시지를 보내도록 미리 설정할 수 있다.
이를 LWT(Last will and testament) 라고 한다. 선언을 먼저하고 브로커가 처리하게 하는 방식인 것이다.**

```python
$ mqtt help subscribe
Usage: mqtt subscribe [opts] [topic]

Available options:
  ... 중략 ...

  --will-topic TOPIC    the will topic
  --will-message BODY   the will message
  --will-qos 0/1/2      the will qos
  --will-retain         send a will retained message
```

{{< /admonition >}}

Both MQTT and HTTP use pub key(asymmetric) device authentication, and JWT. [In more detail](https://cloud.google.com/iot/docs/concepts/device-security)


