package main

// Internetarchive - Python wrapper for the various Internet Archive APIs
// Homepage: https://github.com/jjjake/internetarchive

import (
	"fmt"
	
	"os/exec"
)

func installInternetarchive() {
	// Método 1: Descargar y extraer .tar.gz
	internetarchive_tar_url := "https://files.pythonhosted.org/packages/0d/9e/07e577a1d3e20deeae447069b80d523e7140b790f0a8e1bab1e23b1b9252/internetarchive-4.1.0.tar.gz"
	internetarchive_cmd_tar := exec.Command("curl", "-L", internetarchive_tar_url, "-o", "package.tar.gz")
	err := internetarchive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	internetarchive_zip_url := "https://files.pythonhosted.org/packages/0d/9e/07e577a1d3e20deeae447069b80d523e7140b790f0a8e1bab1e23b1b9252/internetarchive-4.1.0.zip"
	internetarchive_cmd_zip := exec.Command("curl", "-L", internetarchive_zip_url, "-o", "package.zip")
	err = internetarchive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	internetarchive_bin_url := "https://files.pythonhosted.org/packages/0d/9e/07e577a1d3e20deeae447069b80d523e7140b790f0a8e1bab1e23b1b9252/internetarchive-4.1.0.bin"
	internetarchive_cmd_bin := exec.Command("curl", "-L", internetarchive_bin_url, "-o", "binary.bin")
	err = internetarchive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	internetarchive_src_url := "https://files.pythonhosted.org/packages/0d/9e/07e577a1d3e20deeae447069b80d523e7140b790f0a8e1bab1e23b1b9252/internetarchive-4.1.0.src.tar.gz"
	internetarchive_cmd_src := exec.Command("curl", "-L", internetarchive_src_url, "-o", "source.tar.gz")
	err = internetarchive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	internetarchive_cmd_direct := exec.Command("./binary")
	err = internetarchive_cmd_direct.Run()
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
