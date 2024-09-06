package main

// Bilix - Lightning-fast asynchronous download tool for bilibili and more
// Homepage: https://github.com/HFrost0/bilix

import (
	"fmt"
	
	"os/exec"
)

func installBilix() {
	// Método 1: Descargar y extraer .tar.gz
	bilix_tar_url := "https://files.pythonhosted.org/packages/1a/f5/83c35a59e43453033deeecdc19893cedf9558fa601068890f68544e6235f/bilix-0.18.8.tar.gz"
	bilix_cmd_tar := exec.Command("curl", "-L", bilix_tar_url, "-o", "package.tar.gz")
	err := bilix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bilix_zip_url := "https://files.pythonhosted.org/packages/1a/f5/83c35a59e43453033deeecdc19893cedf9558fa601068890f68544e6235f/bilix-0.18.8.zip"
	bilix_cmd_zip := exec.Command("curl", "-L", bilix_zip_url, "-o", "package.zip")
	err = bilix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bilix_bin_url := "https://files.pythonhosted.org/packages/1a/f5/83c35a59e43453033deeecdc19893cedf9558fa601068890f68544e6235f/bilix-0.18.8.bin"
	bilix_cmd_bin := exec.Command("curl", "-L", bilix_bin_url, "-o", "binary.bin")
	err = bilix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bilix_src_url := "https://files.pythonhosted.org/packages/1a/f5/83c35a59e43453033deeecdc19893cedf9558fa601068890f68544e6235f/bilix-0.18.8.src.tar.gz"
	bilix_cmd_src := exec.Command("curl", "-L", bilix_src_url, "-o", "source.tar.gz")
	err = bilix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bilix_cmd_direct := exec.Command("./binary")
	err = bilix_cmd_direct.Run()
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
