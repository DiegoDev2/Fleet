package main

// Wiiuse - Connect Nintendo Wii Remotes
// Homepage: https://github.com/wiiuse/wiiuse

import (
	"fmt"
	
	"os/exec"
)

func installWiiuse() {
	// Método 1: Descargar y extraer .tar.gz
	wiiuse_tar_url := "https://github.com/wiiuse/wiiuse/archive/refs/tags/0.15.6.tar.gz"
	wiiuse_cmd_tar := exec.Command("curl", "-L", wiiuse_tar_url, "-o", "package.tar.gz")
	err := wiiuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wiiuse_zip_url := "https://github.com/wiiuse/wiiuse/archive/refs/tags/0.15.6.zip"
	wiiuse_cmd_zip := exec.Command("curl", "-L", wiiuse_zip_url, "-o", "package.zip")
	err = wiiuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wiiuse_bin_url := "https://github.com/wiiuse/wiiuse/archive/refs/tags/0.15.6.bin"
	wiiuse_cmd_bin := exec.Command("curl", "-L", wiiuse_bin_url, "-o", "binary.bin")
	err = wiiuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wiiuse_src_url := "https://github.com/wiiuse/wiiuse/archive/refs/tags/0.15.6.src.tar.gz"
	wiiuse_cmd_src := exec.Command("curl", "-L", wiiuse_src_url, "-o", "source.tar.gz")
	err = wiiuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wiiuse_cmd_direct := exec.Command("./binary")
	err = wiiuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: bluez")
exec.Command("latte", "install", "bluez")
}
