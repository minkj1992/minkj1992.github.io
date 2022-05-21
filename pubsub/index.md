# Google Cloud Pub/Sub


Let's get familiar with basic features of Google pubsub service by practicing.

<!--more-->
<br />

## tl;dr

- [Google Cloud Pub/Sub: Qwik Start - Console](https://www.cloudskillsboost.google/focuses/3719?catalog_rank=%7B%22rank%22%3A1%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)
- [Google Cloud Pub/Sub: Qwik Start - Command Line](https://www.cloudskillsboost.google/focuses/925?catalog_rank=%7B%22rank%22%3A2%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)

## Quick start `pub/sub`

0. Configure gcloud console

```python
Welcome to Cloud Shell! Type "help" to get started.
Your Cloud Platform project in this session is set to qwiklabs-gcp-03-5549c96ad433.
Use “gcloud config set project [PROJECT_ID]” to change to a different project.
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud auth list
Credentialed Accounts

ACTIVE: *
ACCOUNT: student-03-731a0523b9da@qwiklabs.net

To set the active account, run:
    $ gcloud config set account `ACCOUNT`

student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud config list project
[core]
project = qwiklabs-gcp-03-5549c96ad433

Your active configuration is: [cloudshell-4687]
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$
```

1. Create `topic`

![](/images/pubsub_tutorial1.png)

2. Add a subscription

![](/images/pubsub_tutorial2.png)
![](/images/pubsub_tutorial6.png)

3. Go to detailed topic page

![](/images/pubsub_tutorial3.png)

4. Publish a message to the topic 

![](/images/pubsub_tutorial4.png)

5. View the message

```python
gcloud pubsub subscriptions pull --auto-ack MySub
```

```python
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud pubsub subscriptions pull --auto-ack mySubscription
DATA: Hello World
MESSAGE_ID: 4704893166371310
ORDERING_KEY:
ATTRIBUTES:
DELIVERY_ATTEMPT:
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud pubsub subscriptions pull --auto-ack mySubscription
Listed 0 items.
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud pubsub subscriptions pull --auto-ack mySubscription
Listed 0 items.
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud pubsub subscriptions pull --auto-ack myTopic-sub
DATA: Hello World
MESSAGE_ID: 4704893166371310
ORDERING_KEY:
ATTRIBUTES:
DELIVERY_ATTEMPT:
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$ gcloud pubsub subscriptions pull --auto-ack myTopic-sub
Listed 0 items.
student_03_731a0523b9da@cloudshell:~ (qwiklabs-gcp-03-5549c96ad433)$
```

![](/images/pubsub_tutorial5.png)

---

## Google Cloud Pub/Sub: Qwik Start - Command Line

### The Pub basics (CRD)

A producer publishes messages to a topic and a consumer creates a subscription to a topic to receive messages from it.


1. Create topics

```python
$ gcloud pubsub topics create myTopic
Created topic [projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic].

$ gcloud pubsub topics create myTopic2
$ gcloud pubsub topics create myTopic3
```

When duplicates topcic name, gcp returns error

```python
$ gcloud pubsub topics create myTopic2
$ gcloud pubsub topics create myTopic2
ERROR: Failed to create topic [projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic2]: Resource already exists in the project (resource=myTopic2).
ERROR: (gcloud.pubsub.topics.create) Failed to create the following: [myTopic2].
```

2. List topics

```python
$ gcloud pubsub topics list
---
name: projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic3
---
name: projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic
---
name: projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic2
```

3. Delete topics

```python
student_03_89bc281f2314@cloudshell:~ (qwiklabs-gcp-04-82b04dcaac56)$ gcloud pubsub topics delete myTopic3
Deleted topic [projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic3].
student_03_89bc281f2314@cloudshell:~ (qwiklabs-gcp-04-82b04dcaac56)$ gcloud pubsub topics delete myTopic2
Deleted topic [projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic2].
student_03_89bc281f2314@cloudshell:~ (qwiklabs-gcp-04-82b04dcaac56)$ gcloud pubsub topics list
---
name: projects/qwiklabs-gcp-04-82b04dcaac56/topics/myTopic
```

### The Sub basics (CRD)

1. create

```python
student_03_89bc281f2314@cloudshell:~ (qwiklabs-gcp-04-82b04dcaac56)$ gcloud pubsub subscriptions create --topic myTopic mySubsciption
Created subscription [projects/qwiklabs-gcp-04-82b04dcaac56/subscriptions/mySubsciption].
```

- delete

```python
$ gcloud pubsub subscriptions delete Test1
```

- list

```python
$ gcloud pubsub topics list-subscriptions myTopic
```

### Pub/Sub Publishing and Pulling a Single Message

```python
$ gcloud pubsub topics publish myTopic --message "Hello"
$ gcloud pubsub topics publish myTopic --message "Publisher's name is minwook"
$ gcloud pubsub topics publish myTopic --message "Publisher likes to eat love"
$ gcloud pubsub topics publish myTopic --message "Publisher thinks Pub/Sub is awesome"

# pull subscription
$ gcloud pubsub subscriptions pull mySubscription --auto-ack
```

What's going on here? You published 4 messages to your topic, but only 1 was outputted.

Now there are important features of the `pull` command that often trip developers up

- Using the pull command without any flags will output only one message, even if you are subscribed to a topic that has more held in it.
- Once an individual message has been outputted from a particular subscription based pull command, you cannot access that message again with the pull command (maybe this happens ack)

**Run the last command three more times. You will see that it will output the other messages you published before.**

Now run the command a 4th time.

```bash
$ gcloud pubsub subscriptions pull mySubscription --auto-ack
```

### Pub/Sub pulling all messages from subscriptions with **flag**

```python
$ gcloud pubsub subscriptions pull mySubscription --auto-ack --limit=3
```

Here, notice that gcp pubsub important features.

- about `--auto-ack`

*(BETA) Acknowledges one or more messages as having been successfully received. If a delivered message is not acknowledged within the Subscription's ack deadline, Cloud Pub/Sub will attempt to deliver it again.
**To automatically acknowledge messages when pulling from a Subscription, you can use the --auto-ack flag on gcloud pubsub subscriptions pull.***


![](/images/pubsub_tutorial7.png)

## Google Cloud Pub/Sub: Qwik Start - Python

