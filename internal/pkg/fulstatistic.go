package pkg

import "github.com/Danila331/mifiotsos/internal/models"

type FulResultChat struct {
	Calm        int
	Disgust     int
	Openness    int
	Sociability int
	Anger       int
	Balance     int
	Depression  int
}

type FulResultConf struct {
	Anger      float64
	Disgust    float64
	Enthusiasm float64
	Fear       float64
	Happiness  float64
	Neutral    float64
	Sadness    float64
}

func GetFulResultConf() (FulResultConf, error) {
	var conference models.Conferences
	var fulresult FulResultConf
	conferences, err := conference.ReadAll()
	if err != nil {
		return FulResultConf{}, err
	}
	anger := 0.0
	anger_count := 0
	disgust := 0.0
	disgust_count := 0.0
	enthusiasm := 0.0
	enthusiasm_count := 0.0
	fear := 0.0
	fear_count := 0.0
	happiness := 0.0
	happiness_count := 0.0
	neutral := 0.0
	neutral_count := 0.0
	sadness := 0.0
	sadness_count := 0.0
	for _, conferenc := range conferences {
		if conferenc.Anger != 0 {
			anger += conferenc.Anger
			anger_count++
		}
		if conferenc.Disgust != 0 {
			disgust += conferenc.Disgust
			disgust_count++
		}
		if conferenc.Enthusiasm != 0 {
			enthusiasm += conferenc.Enthusiasm
			enthusiasm_count++
		}
		if conferenc.Fear != 0 {
			fear += conferenc.Fear
			fear_count++
		}
		if conferenc.Happiness != 0 {
			happiness += conferenc.Happiness
			happiness_count++
		}

		if conferenc.Neutral != 0 {
			neutral += conferenc.Neutral
			neutral_count++
		}

		if conferenc.Sadness != 0 {
			sadness += conferenc.Sadness
			sadness_count++
		}
	}
	fulresult.Anger = Round(anger/float64(anger_count), 2)
	fulresult.Disgust = Round(disgust/float64(disgust_count), 2)
	fulresult.Enthusiasm = Round(enthusiasm/float64(enthusiasm_count), 2)
	fulresult.Fear = Round(fear/float64(fear_count), 2)
	fulresult.Happiness = Round(happiness/float64(happiness_count), 2)
	fulresult.Neutral = Round(neutral/float64(neutral_count), 2)
	fulresult.Sadness = Round(sadness/float64(sadness_count), 2)
	return fulresult, nil
}
