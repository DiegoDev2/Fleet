package main

// Dynamips - Cisco 7200/3600/3725/3745/2600/1700 Router Emulator
// Homepage: https://github.com/GNS3/dynamips

import (
	"fmt"
	
	"os/exec"
)

func installDynamips() {
	// Método 1: Descargar y extraer .tar.gz
	dynamips_tar_url := "https://github.com/GNS3/dynamips/archive/refs/tags/v0.2.23.tar.gz"
	dynamips_cmd_tar := exec.Command("curl", "-L", dynamips_tar_url, "-o", "package.tar.gz")
	err := dynamips_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynamips_zip_url := "https://github.com/GNS3/dynamips/archive/refs/tags/v0.2.23.zip"
	dynamips_cmd_zip := exec.Command("curl", "-L", dynamips_zip_url, "-o", "package.zip")
	err = dynamips_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynamips_bin_url := "https://github.com/GNS3/dynamips/archive/refs/tags/v0.2.23.bin"
	dynamips_cmd_bin := exec.Command("curl", "-L", dynamips_bin_url, "-o", "binary.bin")
	err = dynamips_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynamips_src_url := "https://github.com/GNS3/dynamips/archive/refs/tags/v0.2.23.src.tar.gz"
	dynamips_cmd_src := exec.Command("curl", "-L", dynamips_src_url, "-o", "source.tar.gz")
	err = dynamips_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynamips_cmd_direct := exec.Command("./binary")
	err = dynamips_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libelf")
	exec.Command("latte", "install", "libelf").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
}
