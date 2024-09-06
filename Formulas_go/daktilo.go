package main

// Daktilo - Plays typewriter sounds every time you press a key
// Homepage: https://daktilo.cli.rs

import (
	"fmt"
	
	"os/exec"
)

func installDaktilo() {
	// Método 1: Descargar y extraer .tar.gz
	daktilo_tar_url := "https://github.com/orhun/daktilo/archive/refs/tags/v0.6.0.tar.gz"
	daktilo_cmd_tar := exec.Command("curl", "-L", daktilo_tar_url, "-o", "package.tar.gz")
	err := daktilo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daktilo_zip_url := "https://github.com/orhun/daktilo/archive/refs/tags/v0.6.0.zip"
	daktilo_cmd_zip := exec.Command("curl", "-L", daktilo_zip_url, "-o", "package.zip")
	err = daktilo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daktilo_bin_url := "https://github.com/orhun/daktilo/archive/refs/tags/v0.6.0.bin"
	daktilo_cmd_bin := exec.Command("curl", "-L", daktilo_bin_url, "-o", "binary.bin")
	err = daktilo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daktilo_src_url := "https://github.com/orhun/daktilo/archive/refs/tags/v0.6.0.src.tar.gz"
	daktilo_cmd_src := exec.Command("curl", "-L", daktilo_src_url, "-o", "source.tar.gz")
	err = daktilo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daktilo_cmd_direct := exec.Command("./binary")
	err = daktilo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
}
