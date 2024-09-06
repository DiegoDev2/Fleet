package main

// CargoAbout - Cargo plugin to generate list of all licenses for a crate
// Homepage: https://github.com/EmbarkStudios/cargo-about

import (
	"fmt"
	
	"os/exec"
)

func installCargoAbout() {
	// Método 1: Descargar y extraer .tar.gz
	cargoabout_tar_url := "https://github.com/EmbarkStudios/cargo-about/archive/refs/tags/0.6.4.tar.gz"
	cargoabout_cmd_tar := exec.Command("curl", "-L", cargoabout_tar_url, "-o", "package.tar.gz")
	err := cargoabout_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoabout_zip_url := "https://github.com/EmbarkStudios/cargo-about/archive/refs/tags/0.6.4.zip"
	cargoabout_cmd_zip := exec.Command("curl", "-L", cargoabout_zip_url, "-o", "package.zip")
	err = cargoabout_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoabout_bin_url := "https://github.com/EmbarkStudios/cargo-about/archive/refs/tags/0.6.4.bin"
	cargoabout_cmd_bin := exec.Command("curl", "-L", cargoabout_bin_url, "-o", "binary.bin")
	err = cargoabout_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoabout_src_url := "https://github.com/EmbarkStudios/cargo-about/archive/refs/tags/0.6.4.src.tar.gz"
	cargoabout_cmd_src := exec.Command("curl", "-L", cargoabout_src_url, "-o", "source.tar.gz")
	err = cargoabout_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoabout_cmd_direct := exec.Command("./binary")
	err = cargoabout_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
