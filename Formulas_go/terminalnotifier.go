package main

// TerminalNotifier - Send macOS User Notifications from the command-line
// Homepage: https://github.com/julienXX/terminal-notifier

import (
	"fmt"
	
	"os/exec"
)

func installTerminalNotifier() {
	// Método 1: Descargar y extraer .tar.gz
	terminalnotifier_tar_url := "https://github.com/julienXX/terminal-notifier/archive/refs/tags/2.0.0.tar.gz"
	terminalnotifier_cmd_tar := exec.Command("curl", "-L", terminalnotifier_tar_url, "-o", "package.tar.gz")
	err := terminalnotifier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terminalnotifier_zip_url := "https://github.com/julienXX/terminal-notifier/archive/refs/tags/2.0.0.zip"
	terminalnotifier_cmd_zip := exec.Command("curl", "-L", terminalnotifier_zip_url, "-o", "package.zip")
	err = terminalnotifier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terminalnotifier_bin_url := "https://github.com/julienXX/terminal-notifier/archive/refs/tags/2.0.0.bin"
	terminalnotifier_cmd_bin := exec.Command("curl", "-L", terminalnotifier_bin_url, "-o", "binary.bin")
	err = terminalnotifier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terminalnotifier_src_url := "https://github.com/julienXX/terminal-notifier/archive/refs/tags/2.0.0.src.tar.gz"
	terminalnotifier_cmd_src := exec.Command("curl", "-L", terminalnotifier_src_url, "-o", "source.tar.gz")
	err = terminalnotifier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terminalnotifier_cmd_direct := exec.Command("./binary")
	err = terminalnotifier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
