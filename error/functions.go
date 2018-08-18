package error

import "fmt"

func (c *ErrorCollector) Collect(e error) { *c = append(*c, e) }

func (c *ErrorCollector) Error() (err string) {
	err = "" //"Errors found (if any):\n"
	for i := range *c {
		if (*c)[i] != nil {
			err += fmt.Sprintf("%s\n", (*c)[i].Error())
		}
	}
	if err != "" {
		err += fmt.Sprintf("\n")
	}
	return err
}
