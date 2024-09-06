package main

// Minuit2 - Physics analysis tool for function minimization
// Homepage: https://root.cern.ch/doc/master/Minuit2Page.html

import (
	"fmt"
	
	"os/exec"
)

func installMinuit2() {
	// Método 1: Descargar y extraer .tar.gz
	minuit2_tar_url := "https://root.cern.ch/download/root_v6.32.04.source.tar.gz"
	minuit2_cmd_tar := exec.Command("curl", "-L", minuit2_tar_url, "-o", "package.tar.gz")
	err := minuit2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minuit2_zip_url := "https://root.cern.ch/download/root_v6.32.04.source.zip"
	minuit2_cmd_zip := exec.Command("curl", "-L", minuit2_zip_url, "-o", "package.zip")
	err = minuit2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minuit2_bin_url := "https://root.cern.ch/download/root_v6.32.04.source.bin"
	minuit2_cmd_bin := exec.Command("curl", "-L", minuit2_bin_url, "-o", "binary.bin")
	err = minuit2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minuit2_src_url := "https://root.cern.ch/download/root_v6.32.04.source.src.tar.gz"
	minuit2_cmd_src := exec.Command("curl", "-L", minuit2_src_url, "-o", "source.tar.gz")
	err = minuit2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minuit2_cmd_direct := exec.Command("./binary")
	err = minuit2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
