package main

// BatExtras - Bash scripts that integrate bat with various command-line tools
// Homepage: https://github.com/eth-p/bat-extras

import (
	"fmt"
	
	"os/exec"
)

func installBatExtras() {
	// Método 1: Descargar y extraer .tar.gz
	batextras_tar_url := "https://github.com/eth-p/bat-extras/archive/refs/tags/v2024.08.24.tar.gz"
	batextras_cmd_tar := exec.Command("curl", "-L", batextras_tar_url, "-o", "package.tar.gz")
	err := batextras_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	batextras_zip_url := "https://github.com/eth-p/bat-extras/archive/refs/tags/v2024.08.24.zip"
	batextras_cmd_zip := exec.Command("curl", "-L", batextras_zip_url, "-o", "package.zip")
	err = batextras_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	batextras_bin_url := "https://github.com/eth-p/bat-extras/archive/refs/tags/v2024.08.24.bin"
	batextras_cmd_bin := exec.Command("curl", "-L", batextras_bin_url, "-o", "binary.bin")
	err = batextras_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	batextras_src_url := "https://github.com/eth-p/bat-extras/archive/refs/tags/v2024.08.24.src.tar.gz"
	batextras_cmd_src := exec.Command("curl", "-L", batextras_src_url, "-o", "source.tar.gz")
	err = batextras_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	batextras_cmd_direct := exec.Command("./binary")
	err = batextras_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bat")
	exec.Command("latte", "install", "bat").Run()
	fmt.Println("Instalando dependencia: shfmt")
	exec.Command("latte", "install", "shfmt").Run()
	fmt.Println("Instalando dependencia: ripgrep")
	exec.Command("latte", "install", "ripgrep").Run()
}
