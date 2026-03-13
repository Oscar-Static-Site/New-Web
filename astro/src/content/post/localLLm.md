---
title: "Using llama cpp and opencode for local ai"
description: 'Using llama cpp and opencode'
publishDate: '2026-03-13'
---


# Using llama cpp and opencode for local ai.
One of the recent models that has been getting all the headlines is Qwen3.5. I decided to test it locally and see how it works compared to the commercial alternatives.

## How to run it:

First we need some hardware. What worked for me is 16 GB of VRAM with an AMD GPU, but mileage may vary so just test and see what works for you.
Once we have the hardware we then need some software to run the ai.

One of the tools that seems to be getting loads of attention recently is llama.cpp [link](https://github.com/ggml-org/llama.cpp), and then the agentic runner in this case will be the popular [opencode](https://opencode.ai)
And the model is Qwen3.5 and will use the one from unsloth [link](https://unsloth.ai/docs/models/qwen3.5)

## Commands:

First let's start the llama cpp server:

```bash
llama-server -hf unsloth/Qwen3.5-9B-GGUF --chat-template-kwargs '{"enable_thinking": false}'
```

The thinking false seems to be giving better results for coding but again test and see what gives you best results

With the server running we can edit the config of opencode.

```json
{
  "$schema": "https://opencode.ai/config.json",
  "llama.cpp": {
    "npm": "@ai-sdk/openai-compatible",
    "name": "llama-server (local)",
    "options": {
      "baseURL": "http://127.0.0.1:8080/v1"
    },
    "models": {
      "qwen-local": {
        "name": "unsloth/Qwen3.5-9B-GGUF",
        "modalities": {
          "input": [
            "image",
            "text"
          ],
          "output": [
            "text"
          ]
        },
        "limit": {
          "context": 64000,
          "output": 65536
        }
      }
    }
  }
}
```

The name we can get with the curl command:
```bash
curl localhost:8080/v1/models | jq
```

That's all that's required. Now you can go to opencode and start coding!

