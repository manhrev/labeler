package configs

type Config struct {
	Environment string
	DB          *Database
	Server      *Server
	PayOS       *PayOS
	GCPBucket   *GCPBucket
}

type Server struct {
	ListenPort                        int
	ExcludedAuthPath                  map[string]string
	JWTKey                            string
	JWTTokenExpirationDurationMinutes int
}

type Database struct {
	DBUsername string
	DBPassword string
	DBDomain   string
	DBPort     string
	DBName     string
	DBSSLMode  string
}

type PayOS struct {
	ClientID    string
	ApiKey      string
	ChecksumKey string
	WebhookUrl  string
}

type GCPBucket struct {
	JsonCredentials []byte
}
