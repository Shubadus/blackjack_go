package gui

import (
  "fmt"
	"github.com/rivo/tview"
)

//NOTE: We'll use pages to add and remove pages for items like the main window, menu, modal messages
type Gui struct {
  app *tview.Application
  pages *tview.Pages
}

func New() *Gui {
  return &Gui{
    app: tview.NewApplication().EnableMouse(true),
  }
}

func (gui *Gui) gameWindow() *tview.Flex {
  hit := gui.generate_button("Hit", "Hitting")
  stand := gui.generate_button("Stand", "Standing")
  fold := gui.generate_button("Fold", "Folding")
  split := gui.generate_button("Split", "Splitting")
  double_down := gui.generate_button("Double Down", "Doubling Down")
 
  // views
  dealer_window := tview.NewBox().
                         SetBorder(true).
                         SetTitle("Dealer's Hand")
  player_window := tview.NewBox().
                         SetBorder(true).
                         SetTitle("Players's Hand")
  options_window := tview.NewFlex().
                          SetDirection(tview.FlexColumn).
                          // SetBorder(true).
                          AddItem(hit, 0, 1, false).
                          AddItem(stand, 0, 1, false).
                          AddItem(fold, 0, 1, false).
                          AddItem(split, 0, 1, false).
                          AddItem(double_down, 0, 1, false)
  output_window := tview.NewTextView().
                         SetChangedFunc(func() {
                           gui.app.Draw()
                         })

  return tview.NewFlex().SetDirection(tview.FlexRow).
          AddItem(dealer_window, 0, 8, false).
          AddItem(player_window, 0, 8, false).
          AddItem(output_window, 0, 1, false).
          AddItem(options_window, 0, 1, false)


  // return window.(*game) 
  // gui.pages := tview.NewPages() 
}

func (gui *Gui) initWindows() {
  gui.pages = tview.NewPages()

  game := gui.gameWindow()
  gui.pages.AddPage("game", game, true, true)


  gui.app.SetRoot(gui.pages, true)
  gui.pages.ShowPage("game")
}

func (gui *Gui) Start() error {
  gui.initWindows()
	if err := gui.app.Run(); err != nil {
		panic(err)
	}
  return nil
}

func (gui *Gui) Stop() error {
  gui.app.Stop()
  return nil
}

func (gui *Gui) generate_modal(message string) {
    modal := tview.NewModal().
          SetText(message).
          AddButtons([]string{"Ok","Quit"}).
          SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            if buttonLabel == "Quit" {
              gui.Stop()
            } else {
              gui.pages.RemovePage("modal").ShowPage("main")
            }
          })
    gui.pages.AddAndSwitchToPage("modal", modal, true)
}

func (gui *Gui) generate_button(label string, message string) *tview.Button {
    button := tview.NewButton(label)
    button.SetSelectedFunc(func() {
      gui.generate_modal(message)
      // modal := gui.generate_modal(message)
      // gui.pages.AddAndSwitchToPage("Modal", modal, true)
    }).SetFocusFunc(func() {
      button.SetLabel(fmt.Sprintf("[::ub]%s", message))
    })
    return button
}

