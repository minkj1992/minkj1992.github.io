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

# 2. Applications with LLMs

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

## 2.9. Notebook democratizing
> Especially focused on search and sampling.



**Search & Sampling**:
Large Language Models (LLMs)에서 토큰을 생성할 때는 주로 두 가지 주요 방법, 즉 "검색(Search)"과 "샘플링(Sampling)"이 사용됩니다. 이 두 방법은 모델이 다음 토큰을 어떻게 선택할지를 결정하는 방식에 따라 구분됩니다.

**연관관계**:
- "Search"와 "Sampling"은 서로 다른 목적과 상황에 따라 선택될 수 있습니다. 예를 들어, 일관된 및 고품질의 출력이 필요한 경우 "Search"를 사용할 수 있으며, 다양한 및 창의적인 출력이 필요한 경우 "Sampling"을 사용할 수 있습니다.
- 실제 응용 프로그램에서는 "Search"와 "Sampling" 전략을 결합하여 사용하는 경우도 많습니다. 예를 들어, "Beam Search" 동안 일부 확률적 "Sampling"을 적용하여 다양한 결과를 얻을 수 있습니다.

결론적으로, "Search"와 "Sampling"은 시퀀스 생성 모델의 출력을 결정하는 데 사용되는 전략으로, 서로 다른 접근 방식을 제공하며, 특정 응용 프로그램의 요구 사항에 따라 선택 및 조합될 수 있습니다.

> Note. [Can beam search be used with sampling?](https://discuss.huggingface.co/t/can-beam-search-be-used-with-sampling/17741)



1. **Search**:
   - "Search"는 가능한 모든 토큰 시퀀스 중에서 최적의 시퀀스를 찾는 것을 목표로 합니다.
   - 예를 들어, "Greedy Search"는 각 단계에서 확률이 가장 높은 토큰을 선택하는 방법입니다.
   - "Beam Search"는 여러 가능한 경로를 동시에 탐색하여 전체적으로 최적의 시퀀스를 찾습니다.
   - "Search"는 일반적으로 더 일관된 및 결정론적인 결과를 생성합니다.

2. **Sampling**:
   - "Sampling"은 확률 분포에 따라 토큰을 무작위로 선택하는 방법입니다.
   - 이 방법은 다양한 결과를 생성할 수 있으며, 모델의 확률 분포를 직접 활용합니다.
   - "Top-k Sampling" 또는 "Top-p Sampling"과 같은 변형은 완전한 무작위 선택의 위험을 줄이기 위해 토큰 선택의 범위를 제한합니다.

--- 
1. **Greedy Search (탐욕 검색)**:
   - 이 방법은 각 단계에서 가장 확률이 높은 토큰을 선택합니다.
   - 결과적으로 가장 확률이 높은 단일 시퀀스를 생성하지만, 때로는 지역적으로 최적화된 선택으로 인해 전체적으로는 최적이 아닌 결과를 얻을 수 있습니다.

2. **Beam Search**:
   - 여러 경로를 동시에 탐색하면서 가장 확률이 높은 시퀀스를 찾습니다.
   - "빔 크기(Beam Size)"라는 매개변수를 사용하여 한 번에 탐색할 경로의 수를 지정합니다.
   - 빔 크기가 크면 더 많은 경로를 탐색하지만, 계산 비용이 증가합니다.

3. **Sampling**:
   - 다음 토큰을 확률적으로 선택하여 다양한 시퀀스를 생성합니다.
   - 이 방법은 다양한 결과를 생성할 수 있지만, 때로는 의미 없는 결과를 생성할 수도 있습니다.

4. **Top-k Sampling**:
   - 다음 토큰을 선택할 때 가장 확률이 높은 상위 k개의 토큰만을 고려합니다.
   - 이 방법은 샘플링의 임의성을 유지하면서 완전히 무작위 선택의 위험을 줄입니다.

5. **Top-p (nucleus) Sampling**:
   - 확률이 높은 토큰들의 누적 확률이 p를 초과할 때까지 토큰을 선택합니다.
   - 이 방법은 동적으로 선택 범위를 조정하여 토큰의 다양성과 임의성 사이의 균형을 찾습니다.



# 3. Embeddings, Vector Databases and Search

## 3.1 Overview

Title: **LLM Module 2 - Embeddings, Vector Databases, and Search | 2.2 Module Overview**

1. **Purpose of LLMs**: Large Language Models (LLMs) act as reasoning engines, processing information to return meaningful outputs. The module focuses on using embeddings, vector databases, and search to enhance question-answering systems.
  
2. **Knowledge Incorporation**: There are two primary methods for LLMs to learn knowledge: 
   - Training a model from scratch or fine-tuning an existing one.
   - Passing knowledge as model inputs, often referred to as context or prompt engineering.

3. **Context Limitations**: While passing context helps in precision, there's a limitation to the length of context that can be passed. For instance, OpenAI's GPT-3.5 model has a limit of 4,000 tokens. Workarounds include summarizing documents or splitting them into chunks.

4. **Rise of Vector Databases**: 2023 is dubbed the year of vector databases, which are essential for converting various unstructured data types (text, images, audio) into embedding vectors. These vectors can be used for a range of tasks, from object detection to music transcription.

5. **QA System Workflow**:
   - Convert a knowledge base of documents into embedding vectors.
   - Store these vectors in a vector index, either through a vector library or database.
   - Convert user queries into embedding vectors.
   - Search the vector index to find relevant document vectors.
   - Use the retrieved documents as context for the language model to generate a response.

This module provides a comprehensive understanding of how embeddings and vector databases can be utilized to improve search and retrieval performance in question-answering systems.

## 3.2. How does Vector Search work

- **Vector Search Overview**:
  - Two main strategies: exact search (brute force, little room for error) and approximate search (less accurate but faster).
  - Common indexing algorithms produce a data structure called a vector index, which aids in efficient vector search. Methods range from tree-based, clustering, to hashing.
  - Similarity between vectors is determined using distance or similarity metrics, such as L1 Manhattan distance, L2 Euclidean distance, or cosine similarity. When used on normalized embeddings, L2 distance and cosine similarity produce equivalent ranking distances.
  - Dense embedding vectors can be memory-intensive. Product quantization (PQ) compresses vectors by segmenting them into subvectors, which are then quantized and mapped to the nearest centroid, reducing memory usage.

> Note, [Product Quantization](https://www.youtube.com/watch?v=PNVJvZEkuXo)
  
- **Specific Vector Indexing Algorithms**:
  - **FAISS (Facebook AI Similarity Search)**: A clustering algorithm that computes L2 Euclidean distance between vectors. It optimizes the search process using Voronoi cells, computing distances between a query vector and centroids first, then finding similar vectors within the same Voronoi cells.
  - **HNSW (Hierarchical Navigable Small Worlds)**: Uses Euclidean distance but is a proximity graph-based approach. It employs a linked list/skip list structure, skipping nodes or vertices as layers increase. If there are too many nodes, hierarchy is introduced to traverse through the graph efficiently.

- **Significance of Vector Search**:
  - Vector search allows for more flexible and expansive use cases compared to exact matching rules. Traditional SQL filter statements are restrictive, but vector databases offer more dynamic search capabilities.

## (+) L1, L2
1. **L1 Manhattan 거리 (L1 Norm 또는 Taxicab Norm)**:
   - 두 점 P(x_1, y_1) 및  Q(x_2, y_2) 사이의 Manhattan 거리는 그들 사이의 수평 및 수직 경로에 따라 계산됩니다.
   - 계산식: 
   $$
   |x_1 - x_2| + |y_1 - y_2|
   $$
   - 이 거리는 그 이름에서 알 수 있듯이 도시의 격자 패턴 블록을 따라 두 점 사이의 거리를 측정하는 것과 유사합니다.
   - **장점**: 이 거리는 각 차원에서의 차이를 개별적으로 고려합니다. 이로 인해 이상치에 덜 민감하게 됩니다.
   - **사용 사례**: 특히 고차원 데이터에서 이상치의 영향을 최소화하려는 경우나, 각 차원의 차이가 중요한 경우에 사용됩니다.

2. **L2 유클리드 거리 (L2 Norm)**:
   - 두 점 사이의 "직선" 거리입니다. 2차원 평면에서 두 점 사이의 직선 거리를 계산하는 것과 동일하며, 고차원에서도 확장됩니다.
   - 계산식 (2차원의 경우): 
   $$
   \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}
   $$
   - 유클리드 거리는 가장 직관적인 두 점 사이의 거리 측정 방법입니다.
   - **장점**: 직관적이며, 두 점 사이의 실제 "직선" 거리를 제공합니다. 이로 인해 많은 알고리즘과 기술에서 기본 거리 측정 방법으로 사용됩니다.
   - **사용 사례**: k-평균 클러스터링, SVM, k-최근접 이웃 알고리즘과 같은 기계 학습 알고리즘에서 널리 사용됩니다.
   
3. **코사인 유사도**:
   - 두 벡터 간의 코사인 각도를 사용하여 유사도를 측정합니다. 이 값은 -1에서 1 사이입니다.
   - 값이 1에 가까울수록 두 벡터는 매우 유사하며, -1에 가까울수록 매우 다릅니다. 0은 두 벡터가 직교하다는 것을 의미합니다.
   - 계산식: 
   $$
   \frac{A \cdot B}{\|A\| \|B\|}
   $$
   여기서 A.B 는 두 벡터의 내적이고, ||A|| 및 ||B||는 각 벡터의 크기입니다.
   - 코사인 유사도는 텍스트 분석 및 문서 유사도 측정과 같은 많은 응용 분야에서 널리 사용됩니다.
   - **장점**: 벡터의 크기가 아닌 방향에 중점을 둡니다. 따라서 두 벡터가 얼마나 유사한 방향을 가지고 있는지 측정할 때 유용합니다. 특히 텍스트 데이터와 같이 크기보다 방향이 더 중요한 경우에 유용합니다.
   - **사용 사례**: 텍스트 문서의 유사도 측정, 추천 시스템, 텍스트 분류 등에서 널리 사용됩니다.

**L2 거리와 코사인 유사도의 동등성**:
- 정규화된 임베딩에서 L2 거리와 코사인 유사도는 기능적으로 동등한 순위 거리를 생성합니다. 이는 두 측정 방법 모두 벡터 간의 방향성을 측정하기 때문입니다. 정규화된 벡터에서는 유클리드 거리가 작을수록 코사인 유사도가 높아집니다. 따라서 많은 응용 분야에서 두 측정 방법 중 하나를 선택하여 사용할 수 있습니다.

## 3.3. Filtering

- Filtering in vector databases is challenging and varies among different databases.
- There are three main categories of filtering strategies: post-query, in-query, and pre-query.
- Post-query filtering involves applying filters after identifying the top-K nearest neighbors, but it can lead to unpredictable results.
- In-query filtering combines ANN and filtering simultaneously, demanding more system memory as filters increase.
- Pre-query filtering limits similarity search within a specific scope but is less performant than the other methods due to its brute-force approach.

# 4. Multi-stage Reasoning
# 5. Fine-tuning and Evaluating LLMs
# 6. Society and LLMs: Bias and Safety
# 7. LLMOps








