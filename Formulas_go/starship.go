package main

// Starship - Cross-shell prompt for astronauts
// Homepage: https://starship.rs

import (
	"fmt"
	
	"os/exec"
)

func installStarship() {
	// Método 1: Descargar y extraer .tar.gz
	starship_tar_url := "https://github.com/starship/starship/archive/refs/tags/v1.20.1.tar.gz"
	starship_cmd_tar := exec.Command("curl", "-L", starship_tar_url, "-o", "package.tar.gz")
	err := starship_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	starship_zip_url := "https://github.com/starship/starship/archive/refs/tags/v1.20.1.zip"
	starship_cmd_zip := exec.Command("curl", "-L", starship_zip_url, "-o", "package.zip")
	err = starship_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	starship_bin_url := "https://github.com/starship/starship/archive/refs/tags/v1.20.1.bin"
	starship_cmd_bin := exec.Command("curl", "-L", starship_bin_url, "-o", "binary.bin")
	err = starship_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	starship_src_url := "https://github.com/starship/starship/archive/refs/tags/v1.20.1.src.tar.gz"
	starship_cmd_src := exec.Command("curl", "-L", starship_src_url, "-o", "source.tar.gz")
	err = starship_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	starship_cmd_direct := exec.Command("./binary")
	err = starship_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
}
