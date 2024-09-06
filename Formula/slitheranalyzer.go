package main

// SlitherAnalyzer - Solidity static analysis framework written in Python 3
// Homepage: https://github.com/crytic/slither

import (
	"fmt"
	
	"os/exec"
)

func installSlitherAnalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	slitheranalyzer_tar_url := "https://files.pythonhosted.org/packages/40/e4/83b4a1bceb17dfb9f83bfc921bd47832a3252fb5b55e92b2591b28d8a3d3/slither_analyzer-0.10.3.tar.gz"
	slitheranalyzer_cmd_tar := exec.Command("curl", "-L", slitheranalyzer_tar_url, "-o", "package.tar.gz")
	err := slitheranalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slitheranalyzer_zip_url := "https://files.pythonhosted.org/packages/40/e4/83b4a1bceb17dfb9f83bfc921bd47832a3252fb5b55e92b2591b28d8a3d3/slither_analyzer-0.10.3.zip"
	slitheranalyzer_cmd_zip := exec.Command("curl", "-L", slitheranalyzer_zip_url, "-o", "package.zip")
	err = slitheranalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slitheranalyzer_bin_url := "https://files.pythonhosted.org/packages/40/e4/83b4a1bceb17dfb9f83bfc921bd47832a3252fb5b55e92b2591b28d8a3d3/slither_analyzer-0.10.3.bin"
	slitheranalyzer_cmd_bin := exec.Command("curl", "-L", slitheranalyzer_bin_url, "-o", "binary.bin")
	err = slitheranalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slitheranalyzer_src_url := "https://files.pythonhosted.org/packages/40/e4/83b4a1bceb17dfb9f83bfc921bd47832a3252fb5b55e92b2591b28d8a3d3/slither_analyzer-0.10.3.src.tar.gz"
	slitheranalyzer_cmd_src := exec.Command("curl", "-L", slitheranalyzer_src_url, "-o", "source.tar.gz")
	err = slitheranalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slitheranalyzer_cmd_direct := exec.Command("./binary")
	err = slitheranalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
