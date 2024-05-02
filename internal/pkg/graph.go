package pkg

import (
	"fmt"
	"os"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GetLineGraphic() {
	// Создаем новый график
	var conference models.Conferences
	conferences, _ := conference.ReadAll()

	graph := charts.NewLine()

	// Добавляем некоторые тестовые данные для первого графика
	xData1 := make([]float64, 0)
	yDataAnger := make([]float64, 0)
	yDataDisgust := make([]float64, 0)
	yDataEnthusiasm := make([]float64, 0)
	yDataFear := make([]float64, 0)
	yDataHappiness := make([]float64, 0)
	yDataNeutral := make([]float64, 0)
	yDataSadness := make([]float64, 0)
	for index, conf := range conferences {
		xData1 = append(xData1, float64(index))
		yDataAnger = append(yDataAnger, Round(conf.Anger))
		yDataDisgust = append(yDataDisgust, Round(conf.Disgust))
		yDataEnthusiasm = append(yDataEnthusiasm, Round(conf.Enthusiasm))
		yDataFear = append(yDataFear, Round(conf.Fear))
		yDataHappiness = append(yDataHappiness, Round(conf.Happiness))
		yDataNeutral = append(yDataNeutral, Round(conf.Neutral))
		yDataSadness = append(yDataSadness, Round(conf.Sadness))
	}
	fmt.Println(yDataAnger)
	// Добавляем данные для первого графика
	graph.SetXAxis(xData1). // Используем те же данные по X, так как они одинаковы
				AddSeries("График Злости", generateData(xData1, yDataAnger)).
				AddSeries("График Отвращения", generateData(xData1, yDataDisgust)).
				AddSeries("График Энтузиазма", generateData(xData1, yDataEnthusiasm)).
				AddSeries("График Страха", generateData(xData1, yDataFear)).
				AddSeries("График Счастья", generateData(xData1, yDataHappiness)).
				AddSeries("График Нейтральности", generateData(xData1, yDataNeutral)).
				AddSeries("График Грусти", generateData(xData1, yDataSadness))
	// Настройка опций графика для улучшения внешнего вида
	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width: "900px", // Задаем фиксированную ширину в пикселях
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Конференция"}),
	)

	// Создаем HTML-файл для отображения графика
	f, err := os.Create("./static/graphics/graph.html")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer f.Close()

	// Рисуем график в файл
	err = graph.Render(f)
	if err != nil {
		fmt.Println("Ошибка рендеринга графика:", err)
		return
	}

	fmt.Println("График успешно создан в файле graph.html")
}

func generateData(xData, yData []float64) []opts.LineData {
	data := make([]opts.LineData, 0)
	for i := 0; i < len(xData); i++ {
		data = append(data, opts.LineData{Value: yData[i]})
	}
	return data
}
