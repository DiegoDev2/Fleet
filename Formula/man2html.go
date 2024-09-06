package main

// Man2html - Convert nroff man pages to HTML
// Homepage: https://savannah.nongnu.org/projects/man2html/

import (
	"fmt"
	
	"os/exec"
)

func installMan2html() {
	// Método 1: Descargar y extraer .tar.gz
	man2html_tar_url := "https://www.mhonarc.org/release/misc/man2html3.0.1.tar.gz"
	man2html_cmd_tar := exec.Command("curl", "-L", man2html_tar_url, "-o", "package.tar.gz")
	err := man2html_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	man2html_zip_url := "https://www.mhonarc.org/release/misc/man2html3.0.1.zip"
	man2html_cmd_zip := exec.Command("curl", "-L", man2html_zip_url, "-o", "package.zip")
	err = man2html_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	man2html_bin_url := "https://www.mhonarc.org/release/misc/man2html3.0.1.bin"
	man2html_cmd_bin := exec.Command("curl", "-L", man2html_bin_url, "-o", "binary.bin")
	err = man2html_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	man2html_src_url := "https://www.mhonarc.org/release/misc/man2html3.0.1.src.tar.gz"
	man2html_cmd_src := exec.Command("curl", "-L", man2html_src_url, "-o", "source.tar.gz")
	err = man2html_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	man2html_cmd_direct := exec.Command("./binary")
	err = man2html_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
