package main

// Vramsteg - Add progress bars to command-line applications
// Homepage: https://gothenburgbitfactory.org/projects/vramsteg.html

import (
	"fmt"
	
	"os/exec"
)

func installVramsteg() {
	// Método 1: Descargar y extraer .tar.gz
	vramsteg_tar_url := "https://github.com/GothenburgBitFactory/vramsteg/releases/download/v1.1.0/vramsteg-1.1.0.tar.gz"
	vramsteg_cmd_tar := exec.Command("curl", "-L", vramsteg_tar_url, "-o", "package.tar.gz")
	err := vramsteg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vramsteg_zip_url := "https://github.com/GothenburgBitFactory/vramsteg/releases/download/v1.1.0/vramsteg-1.1.0.zip"
	vramsteg_cmd_zip := exec.Command("curl", "-L", vramsteg_zip_url, "-o", "package.zip")
	err = vramsteg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vramsteg_bin_url := "https://github.com/GothenburgBitFactory/vramsteg/releases/download/v1.1.0/vramsteg-1.1.0.bin"
	vramsteg_cmd_bin := exec.Command("curl", "-L", vramsteg_bin_url, "-o", "binary.bin")
	err = vramsteg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vramsteg_src_url := "https://github.com/GothenburgBitFactory/vramsteg/releases/download/v1.1.0/vramsteg-1.1.0.src.tar.gz"
	vramsteg_cmd_src := exec.Command("curl", "-L", vramsteg_src_url, "-o", "source.tar.gz")
	err = vramsteg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vramsteg_cmd_direct := exec.Command("./binary")
	err = vramsteg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
