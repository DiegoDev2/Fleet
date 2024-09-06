package main

// Moon - Task runner and repo management tool for the web ecosystem, written in Rust
// Homepage: https://moonrepo.dev/moon

import (
	"fmt"
	
	"os/exec"
)

func installMoon() {
	// Método 1: Descargar y extraer .tar.gz
	moon_tar_url := "https://github.com/moonrepo/moon/archive/refs/tags/v1.28.0.tar.gz"
	moon_cmd_tar := exec.Command("curl", "-L", moon_tar_url, "-o", "package.tar.gz")
	err := moon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moon_zip_url := "https://github.com/moonrepo/moon/archive/refs/tags/v1.28.0.zip"
	moon_cmd_zip := exec.Command("curl", "-L", moon_zip_url, "-o", "package.zip")
	err = moon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moon_bin_url := "https://github.com/moonrepo/moon/archive/refs/tags/v1.28.0.bin"
	moon_cmd_bin := exec.Command("curl", "-L", moon_bin_url, "-o", "binary.bin")
	err = moon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moon_src_url := "https://github.com/moonrepo/moon/archive/refs/tags/v1.28.0.src.tar.gz"
	moon_cmd_src := exec.Command("curl", "-L", moon_src_url, "-o", "source.tar.gz")
	err = moon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moon_cmd_direct := exec.Command("./binary")
	err = moon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
