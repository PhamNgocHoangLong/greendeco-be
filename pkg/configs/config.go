package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host string `envconfig:"APP_HOST" default:"localhost"`
		PORT string `envconfig:"APP_PORT" default:":8080"`
	}
	Database struct {
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"greendeco"`
		Port     string `envconfig:"DB_PORT" default:"5432"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
		SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
	}
	SMTP struct {
		Email             string `envconfig:"SMTP_EMAIL" default:""`
		Password          string `envconfig:"SMTP_PASSWORD" default:""`
		LinkResetPassword string `envconfig:"SMTP_LINK_RESET_PSW" default:"localhost:8080"`
	}
	Storage struct {
		Firebase struct {
			Type                    string `envconfig:"FIREBASE_TYPE" default:"service_account" json:"type"`
			ProjectID               string `envconfig:"FIREBASE_PROJECT_ID" default:"greendeco-cfabf" json:"project_id"`
			PrivateKeyID            string `envconfig:"FIREBASE_PRIVATE_KEY_ID" default:"d695522fdea0f106674d18444f0254d2cb99e993" json:"private_key_id"`
			PrivateKey              string `envconfig:"FIREBASE_PRIVATE_KEY" default:"-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCmwPGJiSjdM5i5\nh64NITvsvNoQC+xdtT5DWAj9K3YNWxd7qMtIQomFnEbuOVz2LMEbvKoXByczM0VZ\n5cvnRVguAZ3QFkwlrJdzwEWoG/WqNhw5IusEblHTbyUYMzYJTl+03LIEusYeomF+\npIVTJvOVEhofHQLwViiJd3AZ1+yAuVBvBicp/oV3lsBGvl5OkTdpL5AScKlbJyFt\nuZVxo9eB3+ATFTC/Y1JXxPq1iFr0nRpmnNagz4ti9ue33Q96hVd1BtNtGxA67rjz\nC+CAVRCiNbFChijca5Yo1oRPiny41OCCpxD8by6MOtjJjJMeJG0JD6VIJdpNWwLY\nOjdIt1apAgMBAAECggEAPC6RXa5rWZ4IBhmSNvniGzgzSSxMc6t4W6y3nctkMUDJ\ncl9CcoNVB6wVq193jf8r+Js4FdSGkMN1yZKWaxLLVttNfe62L7ig9m0Tlq70lCgw\nOmzAPhqknHbV/+xnmac4AlnOQ8e2vhKjtiWhWpfJvGjKlaUOTPgDbE13X1Re7jue\nRslN989fGPgCo7AayPVg9v/szG6rmSpXDSS5Tt2gHa14jw2H8/bxbZH6w2zbYzIP\nbzrdtz4lLb3+6UsLsy0Uikrgw6rro0iJiw4VMu2IwsffAR2l/G/38SU91ZsE4b/h\ntFRoE7nibrHVQ8RgH1Ju2YCw7rV3VoC8AxNETGhNcwKBgQDRA9n89TULlvyKrFu6\nI/DQ2seCeYyknHIOPVSNOhpMJ5TdqMf/PRQbt4UNvKJNBBLtc9IIkKnTh9GLP1iZ\nIzi3XIqDAGuHKPQMxoGNOgCX/FGvoD/mi7PpLqYNmfr5gnIBQV6ht1tJ4q7HeKrQ\nRkBrEesGh2jFjYljvepmPvPfAwKBgQDMPRXvsZqeBLJ6yhTQ6ODPtk7bxJcmivCU\nXcNm1Bvou8oAxROFXPsQEUpDiISoKVB5iBCktKzPoXzniuxNRlNB/NwcfPIg3a3A\nruw93a/faD0/WrbJ41Y+s7okii6caAe7/BfhyQEaI8cyOpADaHaejEXF2RehAoMA\nE7C1dVHd4wKBgCjOgrGjQe74VGfSjLoDqFFuVUNA7dd5f78N9EL6VAYIOEUqXZmU\nWCetJbnv2RQGNpfLY/cuHyngO2XfQTHssXcDBzvaSiQCN9uCqepjp5gEuAH83Fzr\na32cEOlY5anu8MkT9SDHhngfXJYvFdJB1PPSdLs6lpbzMh6aBnhd1Aj3AoGARXhx\nwbpTsXQ/kWyzb+wGrCCC6lYQY9fWxWCvWobDh4J4z5I5dZ+O9oQrEpWZKeoDqZMS\ntjOOpIF1jvy+K/rDzypHZJlLcsp0k5nsWY7Sh9dZHfepPTrK10EjC9LE0AbCurqh\nFPfeHBoXY4pv+fjKgFuOUCHr26STEq+f59kkI6kCgYAWo46WEMREI3LFUtJML2ix\nw03tdbm/kcKS04SWLcBYRfoS5rbPjB3wc3Ip8ttV6gVlu3LbkuIOQSMYQbrnv2c0\n4LhaVNFLxD2thCjZnjs9ztq9DqQiEGLcUA/zMXeQroe7FXVzfnCecx0iV1Oyq+Fa\null7AdIxUAnPwQoGsMde3Q==\n-----END PRIVATE KEY-----\n" json:"private_key"`
			ClientEmail             string `envconfig:"FIREBASE_CLIENT_EMAIL" default:"firebase-adminsdk-23jb7@greendeco-cfabf.iam.gserviceaccount.com" json:"client_email"`
			ClientID                string `envconfig:"FIREBASE_CLIENT_ID" default:"118327750731974365599" json:"client_id"`
			AuthURI                 string `envconfig:"FIREBASE_AUTH_URI" default:"https://accounts.google.com/o/oauth2/auth" json:"auth_uri"`
			TokenURI                string `envconfig:"FIREBASE_TOKEN_URI" default:"https://oauth2.googleapis.com/token" json:"token_uri"`
			AuthProviderX509CertURL string `envconfig:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL" default:"https://www.googleapis.com/oauth2/v1/certs" json:"auth_provider_x509_cert_url"`
			ClientX509CertURL       string `envconfig:"CLIENT_X509_CERT_URL" default:"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-23jb7%40greendeco-cfabf.iam.gserviceaccount.com" json:"client_x509_cert_url"`
			UniverseDomain          string `envconfig:"UNIVERSE_DOMAIN" default:"googleapis.com" json:"universe_domain"`
		}
		Bucket string `envconfig:"FIREBASE_BUCKET" default:"greendeco-cfabf.appspot.com"`
	}
	VnPay struct {
		Secret     string `envconfig:"VNPAY_SECRET" default:"XNBCJFAKAZQSGTARRLGCHVZWCIOIGSHN"`
		TmnCode    string `envcofig:"VNPAY_TMNCODE" default:"CGXZLS0Z"`
		ReturnUrl  string `envconfig:"VNPAY_RETURN_URL" default:"http://localhost:8080/api/v1/payment/vnpay_return"`
		CancelUrl  string `envconfig:"VNPAY_CANCEL_URL" default:"http://localhost:3000/user/order/"`
		SuccessUrl string `envconfig:"VNPAY_SUCCESS_URL" default:"http://localhost:3000/payment-result/success"`
		ErrorUrl   string `envconfig:"VNPAY_ERROR_URL" default:"http://localhost:3000/payment-result/error"`
	}
	PayPal struct {
		ClientId   string `envconfig:"PAYPAL_CLIENT" default:""`
		SecretKey  string `envconfig:"PAYPAL_SECRET_KEY" default:""`
		ReturnUrl  string `envconfig:"PAYPAL_RETURN_URL" default:"http://localhost:3000/payment/error" `
		SuccessUrl string `envconfig:"PAYPAL_SUCCESS_URL" default:"http://localhost:3000/payment-result/success"`
	}

	ExchangeMoneyApi struct {
		Url string `envconfig:"EXCHANGE_MONEY_API" default:""`
	}
}

var appConfig = &Config{}

func AppConfig() *Config {
	return appConfig
}

func LoadConfig() error {
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
