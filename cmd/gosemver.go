
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
  opener "github.com/kraasch/urlopener/pkg/opener"

  // local packages.
  semv "github.com/kraasch/gosemver/pkg/semv"
)

var (
  // return value.
  selection = ""
  result    = ""

  // output flags.
  doPrint     = false
  doOpen      = false

  // input flags.
  interactive = false
  inSemver    = ""
  inDate      = ""

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
      selection = m.table.SelectedRow()[1]
      return m, tea.Quit
    }
  }
  m.table, cmd = m.table.Update(msg)
  return m, cmd
}

func (m model) View() string {
  return styleBox.Render(m.table.View()) + "\n"
}

func main() {

  // parse flags.

  // NOTE: must provide at least one of the following outputs.
  flag.BoolVar(  &doOpen,      "open",        false, "Open URL in browser")
  flag.BoolVar(  &doPrint,     "print",       false, "Print URL")

  // NOTE: must provide exactly one of the following inputs.
  flag.BoolVar(  &interactive, "interactive", false, "Choose from interactive menu")
  flag.StringVar(&inSemver,    "semver",      "",    "Go semver to look up")
  flag.StringVar(&inDate,      "date",        "",    "Go date to look up")

  flag.Parse()

  // print usage.
  if ! (doOpen || doPrint) {
    fmt.Println("Usage: must at least provide one output option.")
    fmt.Println("  -open")
    fmt.Println("  -print")
    return
  }
  notAtLeastOne := !interactive && inSemver == "" && inDate == ""
  moreThanOne := (interactive && inSemver != "") || (interactive && inDate != "") || (inSemver != "" && inDate != "")
  if notAtLeastOne || moreThanOne {
    fmt.Println("Usage: provide exactly one input option.")
    fmt.Println("  -interactive")
    fmt.Println("  -semver <v3.0>")
    fmt.Println("  -date <yyyy-mm-dd>")
    return
  }

  // prepare model.
  columns := []table.Column{
    {Title: "No.",       Width:  3},
    {Title: "Semver",    Width: 10},
    {Title: "Rel. Date", Width: 10},
  }
  rows := []table.Row{}
  for _, v := range semv.Versions {
    rows = append(rows, table.Row{v.Number, v.Semver, v.Date})
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
  if interactive {
    if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
      fmt.Println("Error running program:", err)
      os.Exit(1)
    }
    // read interactively selected value.
    if selection == "" {
      return
    } else {
      inSemver = selection
    }
  }
  // compute output url.
  if inDate != "" {
    inSemver = semv.DateToSemver(inDate)
  }
  result = semv.SemverToUrl(inSemver)

  // print the last highlighted value in calendar to stdout.
  if doOpen {
    _ = opener.OpenUrl(result)
  }

  // print the last highlighted value in calendar to stdout.
  if doPrint {
    fmt.Println(result)
  }

} // fin.

