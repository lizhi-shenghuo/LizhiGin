package config

type Captcha struct {
	KeyLong int  `mapstructure:"key-long" json:"key_long" yaml:"key-long"'`
	ImgWidth int `mapstructure:"img-width" json:"img_width" yaml:"img-width"'`
	ImgHeight int `mapstructure:"img-height" json:"img_height" yaml:"img-height"`
}
