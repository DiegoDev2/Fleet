package main

// Hercules - System/370, ESA/390 and z/Architecture Emulator
// Homepage: https://sdl-hercules-390.github.io/html/

import (
	"fmt"
	
	"os/exec"
)

func installHercules() {
	// Método 1: Descargar y extraer .tar.gz
	hercules_tar_url := "https://github.com/SDL-Hercules-390/hyperion/archive/refs/tags/Release_4.7.tar.gz"
	hercules_cmd_tar := exec.Command("curl", "-L", hercules_tar_url, "-o", "package.tar.gz")
	err := hercules_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hercules_zip_url := "https://github.com/SDL-Hercules-390/hyperion/archive/refs/tags/Release_4.7.zip"
	hercules_cmd_zip := exec.Command("curl", "-L", hercules_zip_url, "-o", "package.zip")
	err = hercules_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hercules_bin_url := "https://github.com/SDL-Hercules-390/hyperion/archive/refs/tags/Release_4.7.bin"
	hercules_cmd_bin := exec.Command("curl", "-L", hercules_bin_url, "-o", "binary.bin")
	err = hercules_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hercules_src_url := "https://github.com/SDL-Hercules-390/hyperion/archive/refs/tags/Release_4.7.src.tar.gz"
	hercules_cmd_src := exec.Command("curl", "-L", hercules_src_url, "-o", "source.tar.gz")
	err = hercules_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hercules_cmd_direct := exec.Command("./binary")
	err = hercules_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
