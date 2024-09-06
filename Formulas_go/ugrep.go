package main

// Ugrep - Ultra fast grep with query UI, fuzzy search, archive search, and more
// Homepage: https://ugrep.com/

import (
	"fmt"
	
	"os/exec"
)

func installUgrep() {
	// Método 1: Descargar y extraer .tar.gz
	ugrep_tar_url := "https://github.com/Genivia/ugrep/archive/refs/tags/v6.5.0.tar.gz"
	ugrep_cmd_tar := exec.Command("curl", "-L", ugrep_tar_url, "-o", "package.tar.gz")
	err := ugrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ugrep_zip_url := "https://github.com/Genivia/ugrep/archive/refs/tags/v6.5.0.zip"
	ugrep_cmd_zip := exec.Command("curl", "-L", ugrep_zip_url, "-o", "package.zip")
	err = ugrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ugrep_bin_url := "https://github.com/Genivia/ugrep/archive/refs/tags/v6.5.0.bin"
	ugrep_cmd_bin := exec.Command("curl", "-L", ugrep_bin_url, "-o", "binary.bin")
	err = ugrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ugrep_src_url := "https://github.com/Genivia/ugrep/archive/refs/tags/v6.5.0.src.tar.gz"
	ugrep_cmd_src := exec.Command("curl", "-L", ugrep_src_url, "-o", "source.tar.gz")
	err = ugrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ugrep_cmd_direct := exec.Command("./binary")
	err = ugrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
