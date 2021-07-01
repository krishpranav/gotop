package misc

// keybindings for:
//	- help page
//	- error page

var procKeybindings = []string{
	"Quit: q or <C-c>",
	"",
	"[Process navigation](fg:white)",
	"  - k and <Up>: up",
	"  - j and <Down>: down",
	"  - <C-u>: half page up",
	"  - <C-d>: half page down",
	"  - <C-b>: full page up",
	"  - <C-f>: full page down",
	"  - gg and <Home>: jump to top",
	"  - G and <End>: jump to bottom",
	"",
	"[Sorting](fg:white)",
	"  - Use column number to sort ascending.",
	"  - Use <F-column number> to sort descending.",
	"  - Eg: 1 to sort ascedning on 1st Col and F1 for descending",
	"  - 0: Disable Sort",
	"",
	"[Process actions](fg:white)",
	"  - K and <F9>: Open signal selector menu",
	"",
	"[Signal selection](fg:white)",
	"  - K and <F9>: Send SIGTERM to selected process. Kills the process",
	"  - k and <Up>: up",
	"  - j and <Down>: down",
	"  - 0-9: navigate by numeric index",
	"  - <Enter>: send highlighted signal to process",
	"  - <Esc>: close signal selector",
	"",
	"[To close this prompt: <Esc>](fg:white)",
}

var containerKeybindings = []string{
	"Quit: q or <C-c>",
	"",
	"[Container navigation](fg:white)",
	"  - k and <Up>: up",
	"  - j and <Down>: down",
	"  - <C-u>: half page up",
	"  - <C-d>: half page down",
	"  - <C-b>: full page up",
	"  - <C-f>: full page down",
	"  - gg and <Home>: jump to top",
	"  - G and <End>: jump to bottom",
	"",
	"[Sorting](fg:white)",
	"  - Use column number to sort ascending.",
	"  - Use <F-column number> to sort descending.",
	"  - Eg: 1 to sort ascedning on 1st Col and F1 for descending",
	"  - 0: Disable Sort",
	"",
	"[Container actions](fg:white)",
	"  - P: pause a container",
	"  - U: unpause a container",
	"  - R: restart a container",
	"  - S: stop a container",
	"  - K: kill a container",
	"  - X: remove a container (removes links & volumes)",
	"",
	"[To close this prompt: <Esc>](fg:white)",
}

var perContainerKeyBindings = []string{
	"Quit: q or <C-c>",
	"",
	"[Table Selection](fg:white)",
	"  - 1: MountTable",
	"  - 2: NetworkTable",
	"  - 3: CPUUsageTable",
	"  - 4: PortMapTable",
	"  - 5: ProcTable",
	"",
	"[Table navigation](fg:white)",
	"  - k and <Up>: up",
	"  - j and <Down>: down",
	"  - <C-u>: half page up",
	"  - <C-d>: half page down",
	"  - <C-b>: full page up",
	"  - <C-f>: full page down",
	"  - gg and <Home>: jump to top",
	"  - G and <End>: jump to bottom",
	"",
	"[To close this prompt: <Esc>](fg:white)",
}

var mainKeybindings = []string{
	"Quit: q or <C-c>",
	"",
	"[To close this prompt: <Esc>](fg:white)",
}

var perProcKeyBindings = []string{
	"Quit: q or <C-c>",
	"",
	"[To close this prompt: <Esc>](fg:white)",
}

var errorKeybindings = []string{
	"",
	"[To close this prompt: <Esc>](fg:white)",
}
