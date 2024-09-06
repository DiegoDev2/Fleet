package main

// Basex - Light-weight XML database and XPath/XQuery processor
// Homepage: https://basex.org

import (
	"fmt"
	
	"os/exec"
)

func installBasex() {
	// Método 1: Descargar y extraer .tar.gz
	basex_tar_url := "https://files.basex.org/releases/11.2/BaseX112.zip"
	basex_cmd_tar := exec.Command("curl", "-L", basex_tar_url, "-o", "package.tar.gz")
	err := basex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	basex_zip_url := "https://files.basex.org/releases/11.2/BaseX112.zip"
	basex_cmd_zip := exec.Command("curl", "-L", basex_zip_url, "-o", "package.zip")
	err = basex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	basex_bin_url := "https://files.basex.org/releases/11.2/BaseX112.zip"
	basex_cmd_bin := exec.Command("curl", "-L", basex_bin_url, "-o", "binary.bin")
	err = basex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	basex_src_url := "https://files.basex.org/releases/11.2/BaseX112.zip"
	basex_cmd_src := exec.Command("curl", "-L", basex_src_url, "-o", "source.tar.gz")
	err = basex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	basex_cmd_direct := exec.Command("./binary")
	err = basex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
