package checker

import (
	"bufio"
	"os"
	"strings"

	"go.uber.org/zap"
)

func Run(path string, config Config) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	logger.Debug("Validation status:", zap.Bool("validate", config.Validate))
	logger.Debug("ICMP status:", zap.Bool("icmp", config.Icmp))
	for scanner.Scan() {
		raw_ip := scanner.Text()
		ip := strings.TrimSpace(raw_ip)
		if ip == "" {
			continue
		}

		if !make_request(ip) {
			logger.Info("Host is not reachable", zap.String("ip", ip))
			continue
		}

	}

	return nil
}
