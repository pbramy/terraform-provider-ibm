package ibm

import (
	"fmt"
	"net"
	"strings"

	"github.com/IBM-Bluemix/bluemix-go/helpers"
	"github.com/hashicorp/terraform/helper/schema"
	homedir "github.com/mitchellh/go-homedir"
)

func validateServiceTags(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 2048 {
		errors = append(errors, fmt.Errorf(
			"%q must contain tags whose maximum length is 2048 characters", k))
	}
	return
}

func validateAllowedStringValue(validValues []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		input := v.(string)
		existed := false
		for _, s := range validValues {
			if s == input {
				existed = true
				break
			}
		}
		if !existed {
			errors = append(errors, fmt.Errorf(
				"%q must contain a value from %#v, got %q",
				k, validValues, input))
		}
		return

	}
}

func validateRoutePath(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	//Somehow API allows this
	if value == "" {
		return
	}

	if (len(value) < 2) || (len(value) > 128) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must contain from 2 to 128 characters ", k, value))
	}
	if !(strings.HasPrefix(value, "/")) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must start with a forward slash '/'", k, value))

	}
	if strings.Contains(value, "?") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must not contain a '?'", k, value))
	}

	return
}

func validateRoutePort(v interface{}, k string) (ws []string, errors []error) {
	return validatePortRange(1024, 65535)(v, k)
}

func validateAppPort(v interface{}, k string) (ws []string, errors []error) {
	return validatePortRange(1024, 65535)(v, k)
}

func validatePortRange(start, end int) func(v interface{}, k string) (ws []string, errors []error) {
	f := func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if (value < start) || (value > end) {
			errors = append(errors, fmt.Errorf(
				"%q (%d) must be in the range of %d to %d", k, value, start, end))
		}
		return
	}
	return f
}

func validateDomainName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if !(strings.Contains(value, ".")) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must contain a '.',example.com,foo.example.com", k, value))
	}

	return
}

func validateAppInstance(v interface{}, k string) (ws []string, errors []error) {
	instances := v.(int)
	if instances < 0 {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be greater than 0", k, instances))
	}
	return

}

func validateAppZipPath(v interface{}, k string) (ws []string, errors []error) {
	path := v.(string)
	applicationZip, err := homedir.Expand(path)
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"%q (%q) home directory in the given path couldn't be expanded", k, path))
	}
	if !helpers.FileExists(applicationZip) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) doesn't exist", k, path))
	}

	return

}

func validateNotes(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 1000 {
		errors = append(errors, fmt.Errorf(
			"%q should not exceed 1000 characters", k))
	}
	return
}

func validatePublicBandwidth(v interface{}, k string) (ws []string, errors []error) {
	bandwidth := v.(int)
	if bandwidth < 0 {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be greater than 0", k, bandwidth))
		return
	}
	validBandwidths := []int{250, 1000, 5000, 10000, 20000}
	for _, b := range validBandwidths {
		if b == bandwidth {
			return
		}
	}
	errors = append(errors, fmt.Errorf(
		"%q (%d) must be one of the value from %d", k, bandwidth, validBandwidths))
	return

}

func validateMaxConn(v interface{}, k string) (ws []string, errors []error) {
	maxConn := v.(int)
	if maxConn < 1 || maxConn > 64000 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 64000",
			k))
		return
	}
	return
}

func validateWeight(v interface{}, k string) (ws []string, errors []error) {
	weight := v.(int)
	if weight < 0 || weight > 100 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 100",
			k))
		return
	}
	return
}

//validateIP...
func validateIP(v interface{}, k string) (ws []string, errors []error) {
	address := v.(string)
	if net.ParseIP(address) == nil {
		errors = append(errors, fmt.Errorf(
			"%q must be a valid ip address",
			k))
		return
	}
	return
}
