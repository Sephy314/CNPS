package client

func (c *Client) Close() {
	err := c.Conn.Close()
	if err != nil {
		return
	}
}
