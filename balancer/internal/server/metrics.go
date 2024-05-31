package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func (s *Server) genY(backend string) []opts.LineData {
	y := make([]opts.LineData, 0)
	tms := s.lastTimesBack[backend]
	for _, val := range tms {
		y = append(y, opts.LineData{Value: val})
	}
	return y
}

func (s *Server) genX() []string {
	x := make([]string, 0)
	n := 0
	for _, v := range s.lastTimesBack {
		n = max(n, len(v))
	}
	for i := 0; i < n; i++ {
		x = append(x, strconv.Itoa(i))
	}
	return x
}

func (s *Server) Metrics(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		//charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
		charts.WithTitleOpts(opts.Title{
			Title: "Load: balancer and backends",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "y",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "x",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Trigger:   "axis",
			TriggerOn: "mousemove",
			AxisPointer: &opts.AxisPointer{
				Type: "cross",
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "inside",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      0,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	x := s.genX()

	line.SetXAxis(x).
		AddSeries("Back 1", s.genY("1.zlatoivan.ru")).
		AddSeries("Back 2", s.genY("2.zlatoivan.ru")).
		AddSeries("Back 3", s.genY("3.zlatoivan.ru")).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				Smooth:     true,
				ShowSymbol: true,
				SymbolSize: 4,
			}),
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "Average",
				Type: "average",
			}),
		)

	err := line.Render(w)
	if err != nil {
		log.Printf("line.Render: %v\n", err)
	}
}
