package main

// Lhasa - LHA implementation to decompress .lzh and .lzs archives
// Homepage: https://fragglet.github.io/lhasa/

import (
	"fmt"
	
	"os/exec"
)

func installLhasa() {
	// Método 1: Descargar y extraer .tar.gz
	lhasa_tar_url := "https://github.com/fragglet/lhasa/archive/refs/tags/v0.4.0.tar.gz"
	lhasa_cmd_tar := exec.Command("curl", "-L", lhasa_tar_url, "-o", "package.tar.gz")
	err := lhasa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lhasa_zip_url := "https://github.com/fragglet/lhasa/archive/refs/tags/v0.4.0.zip"
	lhasa_cmd_zip := exec.Command("curl", "-L", lhasa_zip_url, "-o", "package.zip")
	err = lhasa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lhasa_bin_url := "https://github.com/fragglet/lhasa/archive/refs/tags/v0.4.0.bin"
	lhasa_cmd_bin := exec.Command("curl", "-L", lhasa_bin_url, "-o", "binary.bin")
	err = lhasa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lhasa_src_url := "https://github.com/fragglet/lhasa/archive/refs/tags/v0.4.0.src.tar.gz"
	lhasa_cmd_src := exec.Command("curl", "-L", lhasa_src_url, "-o", "source.tar.gz")
	err = lhasa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lhasa_cmd_direct := exec.Command("./binary")
	err = lhasa_cmd_direct.Run()
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
}
