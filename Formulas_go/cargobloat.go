package main

// CargoBloat - Find out what takes most of the space in your executable
// Homepage: https://github.com/RazrFalcon/cargo-bloat

import (
	"fmt"
	
	"os/exec"
)

func installCargoBloat() {
	// Método 1: Descargar y extraer .tar.gz
	cargobloat_tar_url := "https://github.com/RazrFalcon/cargo-bloat/archive/refs/tags/v0.12.1.tar.gz"
	cargobloat_cmd_tar := exec.Command("curl", "-L", cargobloat_tar_url, "-o", "package.tar.gz")
	err := cargobloat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargobloat_zip_url := "https://github.com/RazrFalcon/cargo-bloat/archive/refs/tags/v0.12.1.zip"
	cargobloat_cmd_zip := exec.Command("curl", "-L", cargobloat_zip_url, "-o", "package.zip")
	err = cargobloat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargobloat_bin_url := "https://github.com/RazrFalcon/cargo-bloat/archive/refs/tags/v0.12.1.bin"
	cargobloat_cmd_bin := exec.Command("curl", "-L", cargobloat_bin_url, "-o", "binary.bin")
	err = cargobloat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargobloat_src_url := "https://github.com/RazrFalcon/cargo-bloat/archive/refs/tags/v0.12.1.src.tar.gz"
	cargobloat_cmd_src := exec.Command("curl", "-L", cargobloat_src_url, "-o", "source.tar.gz")
	err = cargobloat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargobloat_cmd_direct := exec.Command("./binary")
	err = cargobloat_cmd_direct.Run()
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
