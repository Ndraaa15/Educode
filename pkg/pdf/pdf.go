package pdf

import (
	"fmt"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func GenerateResultQuiz(data map[string]interface{}) (string, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	buildOrderList(m, data)

	unixTime := time.Now().Unix()
	outputFile := fmt.Sprintf("%d.pdf", unixTime)

	err := m.OutputFileAndClose(outputFile)
	if err != nil {
		return "", err
	}

	return outputFile, nil
}

func buildOrderList(m pdf.Maroto, data map[string]interface{}) {
	tableHeadings := []string{"Name", "Score"}

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, data["resultQuiz"].([][]string), props.TableList{
		HeaderProp: props.TableListContent{
			Color: getBlackColor(),
			Style: consts.Bold,

			Size:      9,
			GridSizes: []uint{6, 6},
		},
		ContentProp: props.TableListContent{
			Color:     getBlackColor(),
			Size:      8,
			GridSizes: []uint{6, 6},
		},
		Align:              consts.Left,
		HeaderContentSpace: 1,
		Line:               true,
		LineProp: props.Line{
			Color: getBlackColor(),
		},
		VerticalContentPadding: 3,
	})
}

func getBlackColor() color.Color {
	return color.Color{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
}
