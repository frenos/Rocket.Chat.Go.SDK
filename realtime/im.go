package realtime

func (c *Client) CreateDirectMessage(username string) (string, error) {

	rawResponse, err := c.ddp.Call("createDirectMessage", username)
	if err != nil {
		return "", err
	}
	respParsed := rawResponse.(map[string]interface{})

	return respParsed["rid"].(string), nil
}
