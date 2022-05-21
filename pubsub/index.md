# Google Cloud Pub/Sub


Let's get familiar with basic features of Google pubsub service by practicing.

<!--more-->
<br />

## tl;dr

- [Google Cloud Pub/Sub: Qwik Start - Console](https://www.cloudskillsboost.google/focuses/3719?catalog_rank=%7B%22rank%22%3A1%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)

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

