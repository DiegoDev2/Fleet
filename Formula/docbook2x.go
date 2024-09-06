package main

// Docbook2x - Convert DocBook to UNIX manpages and GNU TeXinfo
// Homepage: https://docbook2x.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDocbook2x() {
	// Método 1: Descargar y extraer .tar.gz
	docbook2x_tar_url := "https://downloads.sourceforge.net/project/docbook2x/docbook2x/0.8.8/docbook2X-0.8.8.tar.gz"
	docbook2x_cmd_tar := exec.Command("curl", "-L", docbook2x_tar_url, "-o", "package.tar.gz")
	err := docbook2x_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docbook2x_zip_url := "https://downloads.sourceforge.net/project/docbook2x/docbook2x/0.8.8/docbook2X-0.8.8.zip"
	docbook2x_cmd_zip := exec.Command("curl", "-L", docbook2x_zip_url, "-o", "package.zip")
	err = docbook2x_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docbook2x_bin_url := "https://downloads.sourceforge.net/project/docbook2x/docbook2x/0.8.8/docbook2X-0.8.8.bin"
	docbook2x_cmd_bin := exec.Command("curl", "-L", docbook2x_bin_url, "-o", "binary.bin")
	err = docbook2x_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docbook2x_src_url := "https://downloads.sourceforge.net/project/docbook2x/docbook2x/0.8.8/docbook2X-0.8.8.src.tar.gz"
	docbook2x_cmd_src := exec.Command("curl", "-L", docbook2x_src_url, "-o", "source.tar.gz")
	err = docbook2x_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docbook2x_cmd_direct := exec.Command("./binary")
	err = docbook2x_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
	exec.Command("latte", "install", "docbook").Run()
}
