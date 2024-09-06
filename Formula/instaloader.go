package main

// Instaloader - Download media from Instagram
// Homepage: https://instaloader.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installInstaloader() {
	// Método 1: Descargar y extraer .tar.gz
	instaloader_tar_url := "https://files.pythonhosted.org/packages/d6/71/87b6d26ec53faf00ce00936a13da7e59ecca77f52263638e8e5d3e637a53/instaloader-4.13.1.tar.gz"
	instaloader_cmd_tar := exec.Command("curl", "-L", instaloader_tar_url, "-o", "package.tar.gz")
	err := instaloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	instaloader_zip_url := "https://files.pythonhosted.org/packages/d6/71/87b6d26ec53faf00ce00936a13da7e59ecca77f52263638e8e5d3e637a53/instaloader-4.13.1.zip"
	instaloader_cmd_zip := exec.Command("curl", "-L", instaloader_zip_url, "-o", "package.zip")
	err = instaloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	instaloader_bin_url := "https://files.pythonhosted.org/packages/d6/71/87b6d26ec53faf00ce00936a13da7e59ecca77f52263638e8e5d3e637a53/instaloader-4.13.1.bin"
	instaloader_cmd_bin := exec.Command("curl", "-L", instaloader_bin_url, "-o", "binary.bin")
	err = instaloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	instaloader_src_url := "https://files.pythonhosted.org/packages/d6/71/87b6d26ec53faf00ce00936a13da7e59ecca77f52263638e8e5d3e637a53/instaloader-4.13.1.src.tar.gz"
	instaloader_cmd_src := exec.Command("curl", "-L", instaloader_src_url, "-o", "source.tar.gz")
	err = instaloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	instaloader_cmd_direct := exec.Command("./binary")
	err = instaloader_cmd_direct.Run()
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
