package main

// Trafilatura - Discovery, extraction and processing for Web text
// Homepage: https://trafilatura.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installTrafilatura() {
	// Método 1: Descargar y extraer .tar.gz
	trafilatura_tar_url := "https://files.pythonhosted.org/packages/89/20/1028e6d81f956312aa76673092ec012a692cec2c63e87b0d6573a6d0e383/trafilatura-1.12.1.tar.gz"
	trafilatura_cmd_tar := exec.Command("curl", "-L", trafilatura_tar_url, "-o", "package.tar.gz")
	err := trafilatura_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trafilatura_zip_url := "https://files.pythonhosted.org/packages/89/20/1028e6d81f956312aa76673092ec012a692cec2c63e87b0d6573a6d0e383/trafilatura-1.12.1.zip"
	trafilatura_cmd_zip := exec.Command("curl", "-L", trafilatura_zip_url, "-o", "package.zip")
	err = trafilatura_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trafilatura_bin_url := "https://files.pythonhosted.org/packages/89/20/1028e6d81f956312aa76673092ec012a692cec2c63e87b0d6573a6d0e383/trafilatura-1.12.1.bin"
	trafilatura_cmd_bin := exec.Command("curl", "-L", trafilatura_bin_url, "-o", "binary.bin")
	err = trafilatura_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trafilatura_src_url := "https://files.pythonhosted.org/packages/89/20/1028e6d81f956312aa76673092ec012a692cec2c63e87b0d6573a6d0e383/trafilatura-1.12.1.src.tar.gz"
	trafilatura_cmd_src := exec.Command("curl", "-L", trafilatura_src_url, "-o", "source.tar.gz")
	err = trafilatura_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trafilatura_cmd_direct := exec.Command("./binary")
	err = trafilatura_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
