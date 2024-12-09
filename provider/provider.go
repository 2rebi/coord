package provider

import (
	"context"

	"github.com/2rebi/coord/embedding"
	"github.com/2rebi/coord/llm"
	"github.com/2rebi/coord/pconf"
	"github.com/2rebi/coord/tts"
)

type LLMClient interface {
	NewLLM(model string, config *llm.Config) (llm.Model, error)
	Close() error
}

type LLMProvider interface {
	NewLLMClient(ctx context.Context, configs ...pconf.Config) (LLMClient, error)
}

type EmbeddingClient interface {
	NewEmbedding(model string, config *embedding.Config) (embedding.Model, error)
	Close() error
}

type EmbeddingProvider interface {
	NewEmbeddingClient(ctx context.Context, configs ...pconf.Config) (EmbeddingClient, error)
}

type TTSClient interface {
	NewTTS(model string, config *tts.Config) (tts.Model, error)
	Close() error
}

type TTSProvider interface {
	NewTTSClient(ctx context.Context, configs ...pconf.Config) (TTSClient, error)
}
