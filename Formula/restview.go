package main

// Restview - Viewer for ReStructuredText documents that renders them on the fly
// Homepage: https://mg.pov.lt/restview/

import (
	"fmt"
	
	"os/exec"
)

func installRestview() {
	// Método 1: Descargar y extraer .tar.gz
	restview_tar_url := "https://files.pythonhosted.org/packages/10/93/20516dada3c64de14305fd8137251cd4accaa7eba15b44deb1f2419aa9ff/restview-3.0.1.tar.gz"
	restview_cmd_tar := exec.Command("curl", "-L", restview_tar_url, "-o", "package.tar.gz")
	err := restview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	restview_zip_url := "https://files.pythonhosted.org/packages/10/93/20516dada3c64de14305fd8137251cd4accaa7eba15b44deb1f2419aa9ff/restview-3.0.1.zip"
	restview_cmd_zip := exec.Command("curl", "-L", restview_zip_url, "-o", "package.zip")
	err = restview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	restview_bin_url := "https://files.pythonhosted.org/packages/10/93/20516dada3c64de14305fd8137251cd4accaa7eba15b44deb1f2419aa9ff/restview-3.0.1.bin"
	restview_cmd_bin := exec.Command("curl", "-L", restview_bin_url, "-o", "binary.bin")
	err = restview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	restview_src_url := "https://files.pythonhosted.org/packages/10/93/20516dada3c64de14305fd8137251cd4accaa7eba15b44deb1f2419aa9ff/restview-3.0.1.src.tar.gz"
	restview_cmd_src := exec.Command("curl", "-L", restview_src_url, "-o", "source.tar.gz")
	err = restview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	restview_cmd_direct := exec.Command("./binary")
	err = restview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
