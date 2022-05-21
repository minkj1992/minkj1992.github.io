# Google Cloud Pub/Sub


Let's get familiar with basic features of Google pubsub service by practicing.

<!--more-->
<br />

## tl;dr

- [Google Cloud Pub/Sub: Qwik Start - Console](https://www.cloudskillsboost.google/focuses/3719?catalog_rank=%7B%22rank%22%3A1%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)
- [Google Cloud Pub/Sub: Qwik Start - Command Line](https://www.cloudskillsboost.google/focuses/925?catalog_rank=%7B%22rank%22%3A2%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)
- [Google Cloud Pub/Sub: Qwik Start - Python](https://www.cloudskillsboost.google/focuses/2775?catalog_rank=%7B%22rank%22%3A3%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078)

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
> https://www.cloudskillsboost.google/focuses/2775?catalog_rank=%7B%22rank%22%3A3%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&qlcampaign=yt18-gsp095-11078


- Project setup

```bash
> git clone https://github.com/googleapis/python-pubsub.git
> cd python-pubsub/samples/snippets
> ls
README.rst     noxfile.py	  requirements-test.txt  subscriber.py
README.rst.in  noxfile_config.py  requirements.txt	 subscriber_test.py
iam.py	       publisher.py	  resources		 utilities
iam_test.py    publisher_test.py  schema.py
mypy.ini       quickstart	  schema_test.py
> python3 -m venv venv
> source venv/bin/activate
> pip list
Package    Version
---------- -------
pip        22.0.4
setuptools 47.1.0
WARNING: You are using pip version 22.0.4; however, version 22.1 is available.
You should consider upgrading via the '/Users/minwook/code/python-pubsub/samples/snippets/venv/bin/python3 -m pip install --upgrade pip' command.
> cat requirements.txt
google-cloud-pubsub==2.12.1
avro==1.11.0
> pip install --upgrade pip
> pip install -r requirements.txt
```

- Create a topic

```bash
# you must authorize google cloud auth before create a topic
$ echo $GOOGLE_CLOUD_PROJECT
$ python publisher.py $GOOGLE_CLOUD_PROJECT create MyTopic
$ python publisher.py $GOOGLE_CLOUD_PROJECT list
```

- Create a subscription

```bash
$ python subscriber.py $GOOGLE_CLOUD_PROJECT create MyTopic MySub
```

- list subscription

```bash
$ python subscriber.py $GOOGLE_CLOUD_PROJECT list-in-project
```

- Subscribe messages

```bash
$ python subscriber.py $GOOGLE_CLOUD_PROJECT receive MySub
Listening for messages on projects/qwiklabs-gcp-7877af129f04d8b3/subscriptions/MySub
Received message: Message {
  data: 'Publisher thinks Pub/Sub is awesome'
  attributes: {}
}
Received message: Message {
  data: 'Hello'
  attributes: {}
}
Received message: Message {
  data: "Publisher's name is Harry"
  attributes: {}
}
Received message: Message {
  data: 'Publisher likes to eat cheese'
  attributes: {}
}
```


- *pub.py*

```python
#!/usr/bin/env python

# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import argparse

from google.cloud import pubsub_v1


def pub(project_id: str, topic_id: str) -> None:
    """Publishes a message to a Pub/Sub topic."""
    # Initialize a Publisher client.
    client = pubsub_v1.PublisherClient()
    # Create a fully qualified identifier of form `projects/{project_id}/topics/{topic_id}`
    topic_path = client.topic_path(project_id, topic_id)

    # Data sent to Cloud Pub/Sub must be a bytestring.
    data = b"Hello, World!"

    # When you publish a message, the client returns a future.
    api_future = client.publish(topic_path, data)
    message_id = api_future.result()

    print(f"Published {data.decode()} to {topic_path}: {message_id}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter,
    )
    parser.add_argument("project_id", help="Google Cloud project ID")
    parser.add_argument("topic_id", help="Pub/Sub topic ID")

    args = parser.parse_args()

    pub(args.project_id, args.topic_id)

```

Wow.. it looks simple, I think the variable naming `api_future` is a key to this script. api_future is returned value of `google.cloud.pubsub_v1.PublisherClient.publish()` and a variable `message_id` is type of `api_future.result()`

Let's deep dive into below logics.

- google.cloud.pubsub_v1.PublisherClient.publish()
- google.cloud.pubsub_v1.PublisherClient.publish.result()


```python
class PublisherClient(metaclass=PublisherClientMeta):
    """The service that an application uses to manipulate topics,
    and to send messages to a topic.
    """

    ...

    def publish(
        self,
        request: Union[pubsub.PublishRequest, dict] = None,
        *,
        topic: str = None,
        messages: Sequence[pubsub.PubsubMessage] = None,
        retry: OptionalRetry = gapic_v1.method.DEFAULT,
        timeout: TimeoutType = gapic_v1.method.DEFAULT,
        metadata: Sequence[Tuple[str, str]] = (),
    ) -> pubsub.PublishResponse:
        r"""Adds one or more messages to the topic. Returns ``NOT_FOUND`` if
        the topic does not exist.

        .. code-block:: python

            from google import pubsub_v1

            def sample_publish():
                # Create a client
                client = pubsub_v1.PublisherClient()

                # Initialize request argument(s)
                request = pubsub_v1.PublishRequest(
                    topic="topic_value",
                )

                # Make the request
                response = client.publish(request=request)

                # Handle the response
                print(response)


        Args:
            request (Union[google.pubsub_v1.types.PublishRequest, dict]):
                The request object. Request for the Publish method.
            topic (str):
                Required. The messages in the request will be published
                on this topic. Format is
                ``projects/{project}/topics/{topic}``.

                This corresponds to the ``topic`` field
                on the ``request`` instance; if ``request`` is provided, this
                should not be set.
            messages (Sequence[google.pubsub_v1.types.PubsubMessage]):
                Required. The messages to publish.
                This corresponds to the ``messages`` field
                on the ``request`` instance; if ``request`` is provided, this
                should not be set.
            retry (google.api_core.retry.Retry): Designation of what errors, if any,
                should be retried.
            timeout (TimeoutType):
                The timeout for this request.
            metadata (Sequence[Tuple[str, str]]): Strings which should be
                sent along with the request as metadata.

        Returns:
            google.pubsub_v1.types.PublishResponse:
                Response for the Publish method.
        """
        # Create or coerce a protobuf request object.
        # Quick check: If we got a request object, we should *not* have
        # gotten any keyword arguments that map to the request.
        has_flattened_params = any([topic, messages])
        if request is not None and has_flattened_params:
            raise ValueError(
                "If the `request` argument is set, then none of "
                "the individual field arguments should be set."
            )

        # Minor optimization to avoid making a copy if the user passes
        # in a pubsub.PublishRequest.
        # There's no risk of modifying the input as we've already verified
        # there are no flattened fields.
        if not isinstance(request, pubsub.PublishRequest):
            request = pubsub.PublishRequest(request)
            # If we have keyword arguments corresponding to fields on the
            # request, apply these.
            if topic is not None:
                request.topic = topic
            if messages is not None:
                request.messages = messages

        # Wrap the RPC method; this adds retry and timeout information,
        # and friendly error handling.
        rpc = self._transport._wrapped_methods[self._transport.publish]

        # Certain fields should be provided within the metadata header;
        # add these here.
        metadata = tuple(metadata) + (
            gapic_v1.routing_header.to_grpc_metadata((("topic", request.topic),)),
        )

        # Send the request.
        response = rpc(
            request,
            retry=retry,
            timeout=timeout,
            metadata=metadata,
        )

        # Done; return the response.
        return response

```

- *sub.py*

```python
#!/usr/bin/env python

# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import argparse
from typing import Optional

from google.cloud import pubsub_v1


def sub(project_id: str, subscription_id: str, timeout: Optional[float] = None) -> None:
    """Receives messages from a Pub/Sub subscription."""
    # Initialize a Subscriber client
    subscriber_client = pubsub_v1.SubscriberClient()
    # Create a fully qualified identifier in the form of
    # `projects/{project_id}/subscriptions/{subscription_id}`
    subscription_path = subscriber_client.subscription_path(project_id, subscription_id)

    def callback(message: pubsub_v1.subscriber.message.Message) -> None:
        print(f"Received {message}.")
        # Acknowledge the message. Unack'ed messages will be redelivered.
        message.ack()
        print(f"Acknowledged {message.message_id}.")

    streaming_pull_future = subscriber_client.subscribe(
        subscription_path, callback=callback
    )
    print(f"Listening for messages on {subscription_path}..\n")

    try:
        # Calling result() on StreamingPullFuture keeps the main thread from
        # exiting while messages get processed in the callbacks.
        streaming_pull_future.result(timeout=timeout)
    except:  # noqa
        streaming_pull_future.cancel()  # Trigger the shutdown.
        streaming_pull_future.result()  # Block until the shutdown is complete.

    subscriber_client.close()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter,
    )
    parser.add_argument("project_id", help="Google Cloud project ID")
    parser.add_argument("subscription_id", help="Pub/Sub subscription ID")
    parser.add_argument(
        "timeout", default=None, nargs="?", const=1, help="Pub/Sub subscription ID"
    )

    args = parser.parse_args()

    sub(args.project_id, args.subscription_id, args.timeout)

```

- about `pubsub_v1.SubscriberClient()`

```python
class Client(subscriber_client.SubscriberClient):
    """A subscriber client for Google Cloud Pub/Sub.

    This creates an object that is capable of subscribing to messages.
    Generally, you can instantiate this client with no arguments, and you
    get sensible defaults.

    Args:
        kwargs: Any additional arguments provided are sent as keyword
            keyword arguments to the underlying
            :class:`~google.cloud.pubsub_v1.gapic.subscriber_client.SubscriberClient`.
            Generally you should not need to set additional keyword
            arguments. Optionally, regional endpoints can be set via
            ``client_options`` that takes a single key-value pair that
            defines the endpoint.

    Example:

    .. code-block:: python

        from google.cloud import pubsub_v1

        subscriber_client = pubsub_v1.SubscriberClient(
            # Optional
            client_options = {
                "api_endpoint": REGIONAL_ENDPOINT
            }
        )
    """
  ... 

    def subscribe(
        self,
        subscription: str,
        callback: Callable[["subscriber.message.Message"], Any],
        flow_control: Union[types.FlowControl, Sequence] = (),
        scheduler: Optional["subscriber.scheduler.ThreadScheduler"] = None,
        use_legacy_flow_control: bool = False,
        await_callbacks_on_shutdown: bool = False,
    ) -> futures.StreamingPullFuture:
        """Asynchronously start receiving messages on a given subscription.

        This method starts a background thread to begin pulling messages from
        a Pub/Sub subscription and scheduling them to be processed using the
        provided ``callback``.

        The ``callback`` will be called with an individual
        :class:`google.cloud.pubsub_v1.subscriber.message.Message`. It is the
        responsibility of the callback to either call ``ack()`` or ``nack()``
        on the message when it finished processing. If an exception occurs in
        the callback during processing, the exception is logged and the message
        is ``nack()`` ed.

        The ``flow_control`` argument can be used to control the rate of at
        which messages are pulled. The settings are relatively conservative by
        default to prevent "message hoarding" - a situation where the client
        pulls a large number of messages but can not process them fast enough
        leading it to "starve" other clients of messages. Increasing these
        settings may lead to faster throughput for messages that do not take
        a long time to process.

        The ``use_legacy_flow_control`` argument disables enforcing flow control
        settings at the Cloud Pub/Sub server, and only the client side flow control
        will be enforced.

        This method starts the receiver in the background and returns a
        *Future* representing its execution. Waiting on the future (calling
        ``result()``) will block forever or until a non-recoverable error
        is encountered (such as loss of network connectivity). Cancelling the
        future will signal the process to shutdown gracefully and exit.

        .. note:: This uses Pub/Sub's *streaming pull* feature. This feature
            properties that may be surprising. Please take a look at
            https://cloud.google.com/pubsub/docs/pull#streamingpull for
            more details on how streaming pull behaves compared to the
            synchronous pull method.

        Example:

        .. code-block:: python

            from google.cloud import pubsub_v1

            subscriber_client = pubsub_v1.SubscriberClient()

            # existing subscription
            subscription = subscriber_client.subscription_path(
                'my-project-id', 'my-subscription')

            def callback(message):
                print(message)
                message.ack()

            future = subscriber_client.subscribe(
                subscription, callback)

            try:
                future.result()
            except KeyboardInterrupt:
                future.cancel()  # Trigger the shutdown.
                future.result()  # Block until the shutdown is complete.

        Args:
            subscription:
                The name of the subscription. The subscription should have already been
                created (for example, by using :meth:`create_subscription`).
            callback:
                The callback function. This function receives the message as
                its only argument and will be called from a different thread/
                process depending on the scheduling strategy.
            flow_control:
                The flow control settings. Use this to prevent situations where you are
                inundated with too many messages at once.
            scheduler:
                An optional *scheduler* to use when executing the callback. This
                controls how callbacks are executed concurrently. This object must not
                be shared across multiple ``SubscriberClient`` instances.
            use_legacy_flow_control (bool):
                If set to ``True``, flow control at the Cloud Pub/Sub server is disabled,
                though client-side flow control is still enabled. If set to ``False``
                (default), both server-side and client-side flow control are enabled.
            await_callbacks_on_shutdown:
                If ``True``, after canceling the returned future, the latter's
                ``result()`` method will block until the background stream and its
                helper threads have been terminated, and all currently executing message
                callbacks are done processing.

                If ``False`` (default), the returned future's ``result()`` method will
                not block after canceling the future. The method will instead return
                immediately after the background stream and its helper threads have been
                terminated, but some of the message callback threads might still be
                running at that point.

        Returns:
            A future instance that can be used to manage the background stream.
        """
        flow_control = types.FlowControl(*flow_control)

        manager = streaming_pull_manager.StreamingPullManager(
            self,
            subscription,
            flow_control=flow_control,
            scheduler=scheduler,
            use_legacy_flow_control=use_legacy_flow_control,
            await_callbacks_on_shutdown=await_callbacks_on_shutdown,
        )

        future = futures.StreamingPullFuture(manager)

        manager.open(callback=callback, on_callback_error=future.set_exception)

        return future  

```

## Conclusion

This is so impressive that pubsub_v1's client generates Future object which needs client should supports async features. Eventually gcp pubsub_v1's clients inherit [official python concurrent Future](https://docs.python.org/ko/3/library/concurrent.futures.html) object, so I'm wondering that can pubsub_v1 handle python2.x? because python future object is not supported on python2.x version. If it does not support python2.x version than, how do I mock `Future based async flow` to `gevent logic`.

**Through this journey, I would like to study Python's concurrent-related features to an advanced level.**

- https://docs.python.org/ko/3/library/concurrent.futures.html
