# Databricks: Large Language Models: Application through Production


Based on [Databrick's LLM application](https://www.youtube.com/watch?v=MLLLDaR6P08&list=PLTPXxbhUt-YWSR8wtILixhZLF9qB_1yZm)




- [`github`](https://github.com/databricks-academy/large-language-models)
- [`edx`](https://www.edx.org/learn/computer-science/databricks-large-language-models-application-through-production)

# 1. LLM Module 0
## 1.2 Why LLMs
- **Introduction to LLMs (Large Language Models)**:
  1. LLMs are revolutionizing industries involving human-computer interaction and language data. They're more than just hype and are impacting businesses, like Chegg, which saw a drop in site shares due to users using ChatGPT.
  2. LLMs are enhancing existing tools, such as GitHub co-pilot, which now can fix code errors, generate tests, and more, thanks to advanced language models.
  3. The core of an LLM is a statistical model that predicts words in text. It's trained on vast amounts of data, with some models being trained on data equivalent to tens of millions of books.
  4. LLMs can automate tasks that involve imprecise language or knowledge about the world, helping in software development, democratizing AI, generating content, and reducing development costs.
  5. When considering LLM applications, it's essential to evaluate the model's quality, serving cost, latency, and customizability. The course aims to bridge the gap between black-box solutions and academic principles, providing practical knowledge for real-world applications.

## 1.3 Primer

- **Introduction to NLP**: Natural Language Processing (NLP) is the study of understanding and modeling natural language for computational applications. It's used daily in tasks like autocomplete, spelling checks, and more.
  
- **Applications of NLP**: NLP can be applied to various tasks such as sentiment analysis, language translation, chatbots, similarity searching, document summarization, and text classification.
  
- **Key Definitions in NLP**: 
  - **Tokens**: The building blocks of NLP, which can be words, characters, or sub-words.
  - **Sequence**: A collection of tokens in a specific order.
  - **Vocabulary**: The entire set of tokens available for a model.
  
- **Classifying NLP Tasks**: Tasks can be categorized based on their input and output, e.g., translation is a sequence-to-sequence task, while sentiment analysis might be a sequence-to-non-sequence prediction.
  
- **Scope of the Course**: While NLP encompasses more than just text (e.g., speech recognition, image captioning), this course focuses on text-based problems due to their inherent complexity and challenges.

## 1.4 Language Model

- **Language Models (LMs)**: Computational models that predict text based on a given sequence, determining the most likely word by calculating a probability distribution over a vocabulary.
  - **Two Types of LMs**:
    - **Generative**: Predicts the next word in a sequence.
    - **Classification-based**: Predicts a masked (or blanked out) word in a sequence.
- **Large Language Models (LLMs)**: LMs that have significantly more parameters, ranging from billions compared to earlier models with 10 to 50 million parameters.
- **Transformers**: A type of architecture introduced around 2017 that has since dominated the natural language processing field due to its computational efficiency.
- **Pre-Transformer Era**: Earlier language models, some with deep neural network architectures, had fewer parameters and were not considered "large" but still required significant computational resources.
- **Post-2017 Shift**: The introduction of Transformers led to a surge in the development and popularity of LLMs.

## 1.5 Tokenization

- **Tokenization in NLP**: Tokenization is the process of converting text into a format suitable for computation. Depending on the design choice, tokens can be words, characters, or pieces of words.
  
- **Word-based Tokenization**: This method involves creating a vocabulary from training data, assigning each word a unique number. However, it has limitations like out-of-vocabulary errors, inability to handle misspellings, and a large vocabulary size due to different word forms (e.g., fast, faster, fastest).

- **Character-based Tokenization**: By breaking text into individual characters, this method offers a small vocabulary size and can handle new words or misspellings. However, it loses the context of words and results in long sequence lengths.

- **Subword Tokenization**: A middle ground between word and character tokenization, subword tokenization breaks words into smaller meaningful parts (e.g., "sub" and "ject" from "subject"). Techniques like byte pairing coding, sentence piece, and wordpiece are used to achieve this. It offers a good balance between vocabulary size and flexibility.

- **Next Steps**: After tokenization, the challenge is to incorporate meaning and context into these tokens, which will be discussed in the context of word embeddings in the subsequent video.

## 1.6 Word Embeddings

- Word embeddings aim to capture the context and intrinsic meaning of words within a vocabulary.
- By tokenizing sentences, we can numerically represent words and compare sentences or documents for similarities and differences.
- Traditional methods, like counting word frequency, result in sparse vectors, which aren't efficient for large vocabularies.
- Word embedding methods, like Word2Vec, represent words as vectors based on surrounding words in training data, capturing contextual relationships.
- These embedding vectors can be visualized in 2D, showing clusters of words with similar meanings, although the exact meanings of vector dimensions aren't always clear.

## 1.7 Recap

- Natural language processing (NLP) focuses on studying natural language, especially text.
- NLP encompasses more than just text; it includes speech, video-to-text, image-to-text, and other concepts where natural language plays a role.
- NLP is valuable for tasks like translation, summarizing text, and classification problems with natural language inputs and outputs.
- Language models create a probability distribution over vocabulary tokens; large language models use the Transformer architecture with millions to billions of parameters.
- Tokens are the smallest units in language models, converting text to indices and then to n-dimensional word embeddings to capture context and meaning.

# 2. Module 1 - Applications with LLMs

## 2.1 Introduction to LLM Applications
- **Introduction to LLM Applications**: The module begins with a humorous note on the urgency among CEOs to adopt LLMs.
- **Ease of Use**: LLMs can be quickly integrated into various applications with minimal initial effort, allowing for continuous improvement.
- **Common Applications**: The module will explore standard natural language processing tasks where pre-trained, often open-source, LLMs excel and can be fine-tuned for specific applications.
- **Tools and Techniques**: Hugging Face will be introduced as a framework for LLMs, along with prompt engineering as a method to adapt general LLMs for diverse tasks. Common tasks like summarization, review classification, and Q&A will be discussed.
- **Understanding Trade-offs**: The module will emphasize the balance between the effort in implementing LLMs and the resulting quality and performance, guiding learners on the potential and limitations of LLMs in various applications.

## 2.2 Module Overview

- **LLM Module 1 Overview**:
  - The module focuses on understanding and applying pre-trained LLMs (Language Learning Models).
  - By the end, learners will know how to interact with LLMs using Hugging Face's APIs, datasets, pipelines, tokenizers, and models.
  - The module emphasizes finding the right model for specific applications, especially given the vast number of models available on Hugging Face Hub as of April 2023.
  - A key topic is prompt engineering, with a real-world example of generating summaries for news articles to be showcased.
  - The module also touches upon the broader NLP ecosystem, mentioning both classical and deep learning-based tools, including proprietary ones and newcomers like LangChain.

## 2.3 Hugging Face

- **Hugging Face Overview**:
  - Hugging Face is a company and community known for open-source machine learning projects, especially in NLP.
  - It hosts models, datasets, and spaces for demos and code, available under various licenses.
  - Libraries provided include "datasets" for data downloading, "Transformers" for NLP pipelines, and "evaluate" for performance assessment.

- **Transformers and Pipelines**:
  - Transformers library simplifies NLP tasks, such as summarization, by providing easy-to-use pipelines.
  - The process involves prompt construction, tokenization (encoding text as numbers), and model inference.
  - The library offers "Auto" classes that automatically configure based on the provided model or tokenizer name.

- **Tokenization Details**:
  - Tokenizers output encoded data as "input IDs" and an "attention mask" which is metadata about the text.
  - Adjustments can be made for input text length, padding, truncation, and tensor return type.

- **Model Inference**:
  - Sequence-to-sequence language models transform variable-length text sequences, like articles, into summaries.
  - Parameters like "num beams" for beam search and output length constraints can be specified. (`beam search`) 

    > Why `beam`?
    >> beam search라는 이름은 "beam of hypotheses"의 약자에서 유래했습니다. beam은 광선이나 빔을 의미하며, 동시에 여러 가지 가능한 후보를 고려하는 방법을 의미합니다.


- **Datasets Library**:
  - Offers one-line APIs for loading and sharing datasets, including NLP, audio, and vision.
  - Datasets are hosted on the Hugging Face Hub, allowing users to filter by various criteria and find related models.

## 2.4 Model Selection

- **Model Selection for Tasks**: When selecting a model for tasks like summarization, one must decide between extractive (selecting pieces from the original text) and abstractive (generating new text) methods.
  
- **Filtering Models on Hugging Face**: With thousands of models available, users can filter by task, license, language, and model size. It's also essential to consider the model's popularity, update frequency, and documentation.

- **Model Variants and Fine-tuning**: Famous models often come in different sizes (base, small, large). Starting with the smallest model can be cost-effective. Fine-tuned models, adjusted for specific tasks, may perform better for related tasks.

- **Importance of Examples and Datasets**: Not all models are well-documented, so looking for usage examples can be beneficial. It's crucial to know if a model is a generalist or fine-tuned for specific tasks and which datasets were used for its training.

- **Recognizing Famous Models**: Many well-known models belong to model families, varying in size and specificity. While size can indicate power, other factors like architecture, training datasets, and fine-tuning can significantly impact performance.

## 2.5 NLP tasks

- **NLP Tasks Overview**: The module introduces various NLP tasks, some of which overlap, and are frequently mentioned in literature and platforms like Hugging Face Hub.
  
- **Text Generation & Summarization**: Text generation can encompass many tasks, including summarization. Some summarization models are labeled as text generation due to their multifunctional nature.

- **Sentiment Analysis**: This task determines the sentiment of a given text, such as identifying whether a tweet about a stock is positive, negative, or neutral. LLMs can also provide a confidence score for their sentiment predictions.

- **Translation & Zero-Shot Classification**: LLMs can be fine-tuned for specific language translations, like English to Spanish. Zero-shot classification allows for categorizing content without retraining the model every time categories change, leveraging the LLM's inherent language understanding.

- **Few-Shot Learning**: This is more of a technique than a task. Instead of fine-tuning a model for a specific task, a few examples are provided to guide the model. This is useful when there isn't a specific model available for a task and limited labeled training data.

## 2.6 Prompts

- **Prompts in LLMs**: Prompts serve as the primary method for interacting with large language models (LLMs) and are especially prevalent in instruction-following LLMs.
  
- **Foundation Models vs. Instruction Following Models**: While Foundation models are trained for general text generation tasks like predicting the next token in a sequence, instruction-following models are tuned to adhere to specific instructions or prompts, such as generating ideas or writing stories.

- **Nature of Prompts**: Prompts can be natural language sentences, questions, code, emojis, or even outputs from other LLM queries. Their purpose is to elicit specific responses from the LLMs, guiding their behavior.

- **Complex Dynamic Interactions**: LLMs can handle nested or chained prompts, allowing for intricate interactions. An example mentioned is "few-shot learning," where the prompt provides an instruction, examples to guide the LLM, and then the actual query.

- **Power of Prompt Engineering**: Effective prompt engineering can lead to structured outputs suitable for downstream data pipelines. The complexity of prompts can range from simple instructions to detailed formats with multiple components, showcasing the versatility and potential of LLMs.

## 2.7 Prompt Engineering

- **Prompt Engineering Specifics**:
  - Prompt engineering is model-specific; what works for one model might not work for another.
  - Effective prompts are clear and specific, often including an instruction, context, input/question, and desired output format.
  - Techniques to improve model responses include instructing the model not to make things up, not to assume sensitive information, and using chain of thought reasoning.
  
- **Prompt Formatting and Security**:
  - Proper formatting, using delimiters, can help distinguish between instructions, context, and user input.
  - There are vulnerabilities like prompt injection (overriding real instructions), prompt leaking (extracting sensitive info), and jailbreaking (bypassing moderation). Developers need to be aware and constantly update to counteract these vulnerabilities.

- **Countermeasures and Resources**:
  - Techniques to reduce prompt hacking include post-processing, filtering, repeating instructions, enclosing user input with random strings, and selecting different models.
  - Various guides and tools are available to assist in writing effective prompts, some specific to OpenAI and others more general.

## 2.8 Recap

- LLMs have a wide variety of applications and use cases.
- Hugging Face offers numerous NLP components, a hub for models, datasets, and examples.
- When selecting a model, consider the task, constraints, and model size.
- There are many tips for model selection, but it's essential to tailor choices to specific applications.
- Prompt engineering is vital for generating valuable responses, combining both art and engineering techniques.



# 3. Module 2 - Embeddings, Vector Databases and Search
# 4. Module 3 - Multi-stage Reasoning
# 5. Module 4 - Fine-tuning and Evaluating LLMs
# 6. Module 5 - Society and LLMs: Bias and Safety
# 7. Module 6 - LLMOps








