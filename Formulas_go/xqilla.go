package main

// Xqilla - XQuery and XPath 2 command-line interpreter
// Homepage: https://xqilla.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXqilla() {
	// Método 1: Descargar y extraer .tar.gz
	xqilla_tar_url := "https://downloads.sourceforge.net/project/xqilla/XQilla-2.3.4.tar.gz"
	xqilla_cmd_tar := exec.Command("curl", "-L", xqilla_tar_url, "-o", "package.tar.gz")
	err := xqilla_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xqilla_zip_url := "https://downloads.sourceforge.net/project/xqilla/XQilla-2.3.4.zip"
	xqilla_cmd_zip := exec.Command("curl", "-L", xqilla_zip_url, "-o", "package.zip")
	err = xqilla_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xqilla_bin_url := "https://downloads.sourceforge.net/project/xqilla/XQilla-2.3.4.bin"
	xqilla_cmd_bin := exec.Command("curl", "-L", xqilla_bin_url, "-o", "binary.bin")
	err = xqilla_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xqilla_src_url := "https://downloads.sourceforge.net/project/xqilla/XQilla-2.3.4.src.tar.gz"
	xqilla_cmd_src := exec.Command("curl", "-L", xqilla_src_url, "-o", "source.tar.gz")
	err = xqilla_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xqilla_cmd_direct := exec.Command("./binary")
	err = xqilla_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
}
