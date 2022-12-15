/////////////////////////////////////////
// ui.go - Secret password vault termui
// Mike Schilli, 2022 (m@perlmeister.com)
/////////////////////////////////////////
package main

import (
  "fmt"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
)

func runUI(lines []string) {
  rows := []string{}
  for _, line := range lines {
    rows = append(rows, mask(line))
  }

  if err := ui.Init(); err != nil {
    panic(err)
  }
  defer ui.Close()

  lb := widgets.NewList()
  lb.Rows = rows
  lb.SelectedRow = 0
  lb.SelectedRowStyle = ui.NewStyle(ui.ColorBlack)
  lb.TextStyle.Fg = ui.ColorGreen
  lb.Title = fmt.Sprintf("passview 1.0")

  pa := widgets.NewParagraph()
  pa.Text = "[Q]uit [Enter]reveal"
  pa.TextStyle.Fg = ui.ColorBlack

  w, h := ui.TerminalDimensions()
  lb.SetRect(0, 0, w, h-3)
  pa.SetRect(0, h-3, w, h)
  ui.Render(lb, pa)

  uiEvents := ui.PollEvents()

  for {
    select {
    case e := <-uiEvents:
      switch e.ID {
      case "k":
        hideCur(lb)
        lb.ScrollUp()
        ui.Render(lb)
      case "j":
        hideCur(lb)
        lb.ScrollDown()
        ui.Render(lb)
      case "q", "<C-c>":
        return
      case "<Enter>":
        showCur(lb, lines)
        ui.Render(lb)

      }
    }
  }
}

func hideCur(lb *widgets.List) {
  idx := lb.SelectedRow
  lb.Rows[idx] = mask(lb.Rows[idx])
}

func showCur(lb *widgets.List, lines []string) {
  idx := lb.SelectedRow
  lb.Rows[idx] = lines[idx]
}
