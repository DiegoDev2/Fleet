package main

// Pokerstove - Poker evaluation and enumeration software
// Homepage: https://github.com/andrewprock/pokerstove

import (
	"fmt"
	
	"os/exec"
)

func installPokerstove() {
	// Método 1: Descargar y extraer .tar.gz
	pokerstove_tar_url := "https://github.com/andrewprock/pokerstove/archive/refs/tags/v1.1.tar.gz"
	pokerstove_cmd_tar := exec.Command("curl", "-L", pokerstove_tar_url, "-o", "package.tar.gz")
	err := pokerstove_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pokerstove_zip_url := "https://github.com/andrewprock/pokerstove/archive/refs/tags/v1.1.zip"
	pokerstove_cmd_zip := exec.Command("curl", "-L", pokerstove_zip_url, "-o", "package.zip")
	err = pokerstove_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pokerstove_bin_url := "https://github.com/andrewprock/pokerstove/archive/refs/tags/v1.1.bin"
	pokerstove_cmd_bin := exec.Command("curl", "-L", pokerstove_bin_url, "-o", "binary.bin")
	err = pokerstove_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pokerstove_src_url := "https://github.com/andrewprock/pokerstove/archive/refs/tags/v1.1.src.tar.gz"
	pokerstove_cmd_src := exec.Command("curl", "-L", pokerstove_src_url, "-o", "source.tar.gz")
	err = pokerstove_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pokerstove_cmd_direct := exec.Command("./binary")
	err = pokerstove_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
