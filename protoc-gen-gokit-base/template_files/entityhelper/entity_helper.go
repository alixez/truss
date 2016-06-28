package entityhelper

import (
	"{{.AbsoluteRelativeImportPath}}mv-shared/mysql"
)

type EntityHelper struct {
	*mysql.Client
}

func (c *EntityHelper) GetMysqlClient() (*mysql.Client, error) {
	if c.Client == nil {
		client, err := mysql.NewClient()
		if err != nil {
			return nil, err
		}
		c.Client = client
	}
	return c.Client, nil
}
