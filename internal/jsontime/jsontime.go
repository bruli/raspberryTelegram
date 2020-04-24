package jsontime

import (
	"strconv"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type JsonTime time.Time

func (j *JsonTime) MarshalJSON() ([]byte, error) {
	t := strconv.Quote(time.Time(*j).Format(timeLayout))
	return []byte(t), nil
}

func (j *JsonTime) ToString() string {
	return time.Time(*j).Format(timeLayout)
}

func (j *JsonTime) UnmarshalJSON(s []byte) error {

	unquote, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	t, err := time.Parse(timeLayout, unquote)
	if err != nil {
		return err
	}
	*j = JsonTime(t)
	return nil
}

func (j *JsonTime) ToTime() time.Time {
	t, _ := time.Parse(timeLayout, j.ToString())
	return t
}
