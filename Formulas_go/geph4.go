package main

// Geph4 - Modular Internet censorship circumvention system to deal with national filtering
// Homepage: https://geph.io/

import (
	"fmt"
	
	"os/exec"
)

func installGeph4() {
	// Método 1: Descargar y extraer .tar.gz
	geph4_tar_url := "https://github.com/geph-official/geph4-client/archive/refs/tags/v4.11.0.tar.gz"
	geph4_cmd_tar := exec.Command("curl", "-L", geph4_tar_url, "-o", "package.tar.gz")
	err := geph4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geph4_zip_url := "https://github.com/geph-official/geph4-client/archive/refs/tags/v4.11.0.zip"
	geph4_cmd_zip := exec.Command("curl", "-L", geph4_zip_url, "-o", "package.zip")
	err = geph4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geph4_bin_url := "https://github.com/geph-official/geph4-client/archive/refs/tags/v4.11.0.bin"
	geph4_cmd_bin := exec.Command("curl", "-L", geph4_bin_url, "-o", "binary.bin")
	err = geph4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geph4_src_url := "https://github.com/geph-official/geph4-client/archive/refs/tags/v4.11.0.src.tar.gz"
	geph4_cmd_src := exec.Command("curl", "-L", geph4_src_url, "-o", "source.tar.gz")
	err = geph4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geph4_cmd_direct := exec.Command("./binary")
	err = geph4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
