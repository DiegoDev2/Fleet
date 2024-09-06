package main

// Zola - Fast static site generator in a single binary with everything built-in
// Homepage: https://www.getzola.org/

import (
	"fmt"
	
	"os/exec"
)

func installZola() {
	// Método 1: Descargar y extraer .tar.gz
	zola_tar_url := "https://github.com/getzola/zola/archive/refs/tags/v0.19.2.tar.gz"
	zola_cmd_tar := exec.Command("curl", "-L", zola_tar_url, "-o", "package.tar.gz")
	err := zola_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zola_zip_url := "https://github.com/getzola/zola/archive/refs/tags/v0.19.2.zip"
	zola_cmd_zip := exec.Command("curl", "-L", zola_zip_url, "-o", "package.zip")
	err = zola_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zola_bin_url := "https://github.com/getzola/zola/archive/refs/tags/v0.19.2.bin"
	zola_cmd_bin := exec.Command("curl", "-L", zola_bin_url, "-o", "binary.bin")
	err = zola_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zola_src_url := "https://github.com/getzola/zola/archive/refs/tags/v0.19.2.src.tar.gz"
	zola_cmd_src := exec.Command("curl", "-L", zola_src_url, "-o", "source.tar.gz")
	err = zola_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zola_cmd_direct := exec.Command("./binary")
	err = zola_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: oniguruma")
exec.Command("latte", "install", "oniguruma")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
