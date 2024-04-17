package zhipu

/**
    @date: 2024/4/16
**/

type RequestArgs struct {
	// 是 所要调用的模型编码
	Model string `json:"model"`
	// 调用语言模型时，将当前对话信息列表作为提示输入给模型， 按照 {"role": "user", "content": "你好"} 的json 数组形式进行传参；
	// 可能的消息类型包括 System message、User message、Assistant message 和 Tool messag
	Messages []Messages `json:"messages"`
	// 由用户端传参，需保证唯一性；用于区分每次请求的唯一标识，用户端不传时平台会默认生成。
	RequestId string `json:"request_id,omitempty"`
	// do_sample 为 true 时启用采样策略，do_sample 为 false 时采样策略 temperature、top_p 将不生效。默认值为 true。
	DoSample bool `json:"do_sample,omitempty"`
	// 使用同步调用时，此参数应当设置为 fasle 或者省略。表示模型生成完所有内容后一次性返回所有内容。默认值为 false。
	Stream bool `json:"stream,omitempty"`
	//采样温度，控制输出的随机性，必须为正数
	//取值范围是：(0.0, 1.0)，不能等于 0，默认值为 0.95，值越大，会使输出更随机，更具创造性；值越小，输出会更加稳定或确定
	//建议您根据应用场景调整 top_p 或 temperature 参数，但不要同时调整两个参数
	Temperature float32 `json:"temperature,omitempty"`
	/*
		用温度取样的另一种方法，称为核取样
		取值范围是：(0.0, 1.0) 开区间，不能等于 0 或 1，默认值为 0.7
		模型考虑具有 top_p 概率质量 tokens 的结果
		例如：0.1 意味着模型解码器只考虑从前 10% 的概率的候选集中取 tokens
		建议您根据应用场景调整 top_p 或 temperature 参数，但不要同时调整两个参数
	*/
	TopP float32 `json:"top_p,omitempty"`
	// 模型输出最大 tokens，最大输出为8192，默认值为1024
	MaxTokens int64 `json:"max_tokens,omitempty"`
	// 模型在遇到stop所制定的字符时将停止生成，目前仅支持单个停止词，格式为["stop_word1"]
	Stop       []string `json:"stop,omitempty"`
	ToolChoice string   `json:"tool_choice,omitempty"`
	/*
		工具类型,目前支持function、retrieval、web_search
	*/
	Type string `json:"type"`

	WebSearch WebSearch `json:"web_search,omitempty"`
	/*
	 终端用户的唯一ID，协助平台对终端用户的违规行为、生成违法及不良信息或其他滥用行为进行干预。ID长度要求：最少6个字符，最多128个字符。
	*/
	UserId string `json:"user_id,omitempty"`
}
type Messages struct {
	// 角色
	Role string `json:"role"`
	// 请求内容
	Content string `json:"content"`
}

type WebSearch struct {
	/*
		是否启用搜索，默认启用搜索
		启用：true

		禁用：false
	*/
	Enable bool `json:"enable"`
	/*
		强制搜索自定义关键内容，此时模型会根据自定义搜索关键内容返回的结果作为背景知识来回答用户发起的对话。
	*/
	SearchQuery string `json:"search_query"`
}

type Response struct {
}
