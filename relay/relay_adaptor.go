package relay

import (
	"king-api/constant"
	commonconstant "king-api/constant"
	"king-api/relay/channel"
	"king-api/relay/channel/ali"
	"king-api/relay/channel/aws"
	"king-api/relay/channel/baidu"
	"king-api/relay/channel/baidu_v2"
	"king-api/relay/channel/claude"
	"king-api/relay/channel/cloudflare"
	"king-api/relay/channel/cohere"
	"king-api/relay/channel/coze"
	"king-api/relay/channel/deepseek"
	"king-api/relay/channel/dify"
	"king-api/relay/channel/gemini"
	"king-api/relay/channel/jimeng"
	"king-api/relay/channel/jina"
	"king-api/relay/channel/mistral"
	"king-api/relay/channel/mokaai"
	"king-api/relay/channel/ollama"
	"king-api/relay/channel/openai"
	"king-api/relay/channel/palm"
	"king-api/relay/channel/perplexity"
	"king-api/relay/channel/siliconflow"
	taskjimeng "king-api/relay/channel/task/jimeng"
	"king-api/relay/channel/task/kling"
	"king-api/relay/channel/task/suno"
	"king-api/relay/channel/tencent"
	"king-api/relay/channel/vertex"
	"king-api/relay/channel/volcengine"
	"king-api/relay/channel/xai"
	"king-api/relay/channel/xunfei"
	"king-api/relay/channel/zhipu"
	"king-api/relay/channel/zhipu_4v"
)

func GetAdaptor(apiType int) channel.Adaptor {
	switch apiType {
	case constant.APITypeAli:
		return &ali.Adaptor{}
	case constant.APITypeAnthropic:
		return &claude.Adaptor{}
	case constant.APITypeBaidu:
		return &baidu.Adaptor{}
	case constant.APITypeGemini:
		return &gemini.Adaptor{}
	case constant.APITypeOpenAI:
		return &openai.Adaptor{}
	case constant.APITypePaLM:
		return &palm.Adaptor{}
	case constant.APITypeTencent:
		return &tencent.Adaptor{}
	case constant.APITypeXunfei:
		return &xunfei.Adaptor{}
	case constant.APITypeZhipu:
		return &zhipu.Adaptor{}
	case constant.APITypeZhipuV4:
		return &zhipu_4v.Adaptor{}
	case constant.APITypeOllama:
		return &ollama.Adaptor{}
	case constant.APITypePerplexity:
		return &perplexity.Adaptor{}
	case constant.APITypeAws:
		return &aws.Adaptor{}
	case constant.APITypeCohere:
		return &cohere.Adaptor{}
	case constant.APITypeDify:
		return &dify.Adaptor{}
	case constant.APITypeJina:
		return &jina.Adaptor{}
	case constant.APITypeCloudflare:
		return &cloudflare.Adaptor{}
	case constant.APITypeSiliconFlow:
		return &siliconflow.Adaptor{}
	case constant.APITypeVertexAi:
		return &vertex.Adaptor{}
	case constant.APITypeMistral:
		return &mistral.Adaptor{}
	case constant.APITypeDeepSeek:
		return &deepseek.Adaptor{}
	case constant.APITypeMokaAI:
		return &mokaai.Adaptor{}
	case constant.APITypeVolcEngine:
		return &volcengine.Adaptor{}
	case constant.APITypeBaiduV2:
		return &baidu_v2.Adaptor{}
	case constant.APITypeOpenRouter:
		return &openai.Adaptor{}
	case constant.APITypeXinference:
		return &openai.Adaptor{}
	case constant.APITypeXai:
		return &xai.Adaptor{}
	case constant.APITypeCoze:
		return &coze.Adaptor{}
	case constant.APITypeJimeng:
		return &jimeng.Adaptor{}
	}
	return nil
}

func GetTaskAdaptor(platform commonconstant.TaskPlatform) channel.TaskAdaptor {
	switch platform {
	//case constant.APITypeAIProxyLibrary:
	//	return &aiproxy.Adaptor{}
	case commonconstant.TaskPlatformSuno:
		return &suno.TaskAdaptor{}
	case commonconstant.TaskPlatformKling:
		return &kling.TaskAdaptor{}
	case commonconstant.TaskPlatformJimeng:
		return &taskjimeng.TaskAdaptor{}
	}
	return nil
}
