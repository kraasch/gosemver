
package main

import (

  // for making a nice centred box.
  tea "github.com/charmbracelet/bubbletea"
  lip "github.com/charmbracelet/lipgloss"
  table "github.com/charmbracelet/bubbles/table"

  // basics.
  "fmt"
  "os"
  "flag"

  // misc.
  // opener "github.com/kraasch/urlopener/pkg/opener"

  // local packages.
  // semv "github.com/kraasch/gosemver/pkg/semv"
)

var (
  // return value.
  output = ""
  // flags.
  verbose  = false
  suppress = false
  // styles.
  styleBox = lip.NewStyle().
    BorderStyle(lip.NormalBorder()).
    BorderForeground(lip.Color("56"))
)

type model struct {
  width     int
  height    int
	table     table.Model
}

func (m model) Init() tea.Cmd {
  return func() tea.Msg { return nil }
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//   var cmd tea.Cmd
//   switch msg := msg.(type) {
//   case tea.WindowSizeMsg:
//     m.width = msg.Width
//     m.height = msg.Height
//   case tea.KeyMsg:
//     switch msg.String() {
//     case "q":
//       output = "You quit on me!"
//       return m, tea.Quit
//     }
//   }
//   return m, cmd
// }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// func (m model) View() string {
//   var str string
//   if verbose {
//     str = semv.Toast("Hello!")
//   } else {
//     str = semv.Toast("Hi!")
//   }
//   str = styleBox.Render(str)
//   return lip.Place(m.width, m.height, lip.Center, lip.Center, str)
// }

func (m model) View() string {
	return styleBox.Render(m.table.View()) + "\n"
}

func main() {

  // parse flags.
  flag.BoolVar(&verbose,  "verbose",   false, "Show info")
  flag.BoolVar(&suppress, "suppress",  false, "Print nothing")
  flag.Parse()

  // prepare model.
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}
	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"98", "Rome", "Italy", "4,297,877"},
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"98", "Rome", "Italy", "4,297,877"},
		{"99", "Shijiazhuang", "China", "4,285,135"},
		{"100", "Montreal", "Canada", "4,276,526"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lip.NormalBorder()).
		BorderForeground(lip.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lip.Color("229")).
		Background(lip.Color("57")).
		Bold(false)
	t.SetStyles(s)


  // init model.
  m := model{0, 0, t}

  // start bubbletea.
  if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }

  // print the last highlighted value in calendar to stdout.
  if !suppress {
    fmt.Println(output)
  }

} // fin.

