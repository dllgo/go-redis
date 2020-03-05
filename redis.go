package redisplus

type redisplus struct {
	String stringRds
	List   listRds
	Hash   hashRds
	Key    keyRds
	Set    setRds
	ZSet   zSetRds
	Bit    bitRds
	Db     dbRds
}

var Redis = new(redisplus)

func NewRedisWithFile(tagName, path string) error {
	config, err := getConfigWithFile(tagName, path)
	if err != nil {
		config, err = getConfigWithFile(tagName, "./conf/app.conf")
		if err != nil || config == nil {
			return err
		}
	}

	initPool(config)
	return nil
}

func NewRedisWithConfig(config *Config) error {
	//
	initPool(config)
	return nil
}

func NewRedisWithDefualtConfig() error {
	//
	initPool(&Config{
		Host:      "127.0.0.1",
		Port:      6379,
		Db:        0,
		Password:  "",
		MaxIdle:   10,
		MaxActive: 100,
		Wait:      true})
	return nil
}
