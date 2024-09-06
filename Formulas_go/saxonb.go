package main

// SaxonB - XSLT and XQuery processor
// Homepage: https://saxon.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSaxonB() {
	// Método 1: Descargar y extraer .tar.gz
	saxonb_tar_url := "https://downloads.sourceforge.net/project/saxon/Saxon-B/9.1.0.8/saxonb9-1-0-8j.zip"
	saxonb_cmd_tar := exec.Command("curl", "-L", saxonb_tar_url, "-o", "package.tar.gz")
	err := saxonb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	saxonb_zip_url := "https://downloads.sourceforge.net/project/saxon/Saxon-B/9.1.0.8/saxonb9-1-0-8j.zip"
	saxonb_cmd_zip := exec.Command("curl", "-L", saxonb_zip_url, "-o", "package.zip")
	err = saxonb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	saxonb_bin_url := "https://downloads.sourceforge.net/project/saxon/Saxon-B/9.1.0.8/saxonb9-1-0-8j.zip"
	saxonb_cmd_bin := exec.Command("curl", "-L", saxonb_bin_url, "-o", "binary.bin")
	err = saxonb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	saxonb_src_url := "https://downloads.sourceforge.net/project/saxon/Saxon-B/9.1.0.8/saxonb9-1-0-8j.zip"
	saxonb_cmd_src := exec.Command("curl", "-L", saxonb_src_url, "-o", "source.tar.gz")
	err = saxonb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	saxonb_cmd_direct := exec.Command("./binary")
	err = saxonb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
