package main

// Xctesthtmlreport - Xcode-like HTML report for Unit and UI Tests
// Homepage: https://github.com/XCTestHTMLReport/XCTestHTMLReport

import (
	"fmt"
	
	"os/exec"
)

func installXctesthtmlreport() {
	// Método 1: Descargar y extraer .tar.gz
	xctesthtmlreport_tar_url := "https://github.com/XCTestHTMLReport/XCTestHTMLReport/archive/refs/tags/2.5.0.tar.gz"
	xctesthtmlreport_cmd_tar := exec.Command("curl", "-L", xctesthtmlreport_tar_url, "-o", "package.tar.gz")
	err := xctesthtmlreport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xctesthtmlreport_zip_url := "https://github.com/XCTestHTMLReport/XCTestHTMLReport/archive/refs/tags/2.5.0.zip"
	xctesthtmlreport_cmd_zip := exec.Command("curl", "-L", xctesthtmlreport_zip_url, "-o", "package.zip")
	err = xctesthtmlreport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xctesthtmlreport_bin_url := "https://github.com/XCTestHTMLReport/XCTestHTMLReport/archive/refs/tags/2.5.0.bin"
	xctesthtmlreport_cmd_bin := exec.Command("curl", "-L", xctesthtmlreport_bin_url, "-o", "binary.bin")
	err = xctesthtmlreport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xctesthtmlreport_src_url := "https://github.com/XCTestHTMLReport/XCTestHTMLReport/archive/refs/tags/2.5.0.src.tar.gz"
	xctesthtmlreport_cmd_src := exec.Command("curl", "-L", xctesthtmlreport_src_url, "-o", "source.tar.gz")
	err = xctesthtmlreport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xctesthtmlreport_cmd_direct := exec.Command("./binary")
	err = xctesthtmlreport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
