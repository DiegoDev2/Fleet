package main

// R3 - High-performance URL router library
// Homepage: https://github.com/c9s/r3

import (
	"fmt"
	
	"os/exec"
)

func installR3() {
	// Método 1: Descargar y extraer .tar.gz
	r3_tar_url := "https://github.com/c9s/r3/archive/refs/tags/1.3.4.tar.gz"
	r3_cmd_tar := exec.Command("curl", "-L", r3_tar_url, "-o", "package.tar.gz")
	err := r3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	r3_zip_url := "https://github.com/c9s/r3/archive/refs/tags/1.3.4.zip"
	r3_cmd_zip := exec.Command("curl", "-L", r3_zip_url, "-o", "package.zip")
	err = r3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	r3_bin_url := "https://github.com/c9s/r3/archive/refs/tags/1.3.4.bin"
	r3_cmd_bin := exec.Command("curl", "-L", r3_bin_url, "-o", "binary.bin")
	err = r3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	r3_src_url := "https://github.com/c9s/r3/archive/refs/tags/1.3.4.src.tar.gz"
	r3_cmd_src := exec.Command("curl", "-L", r3_src_url, "-o", "source.tar.gz")
	err = r3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	r3_cmd_direct := exec.Command("./binary")
	err = r3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
