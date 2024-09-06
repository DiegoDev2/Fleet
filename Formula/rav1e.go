package main

// Rav1e - Fastest and safest AV1 video encoder
// Homepage: https://github.com/xiph/rav1e

import (
	"fmt"
	
	"os/exec"
)

func installRav1e() {
	// Método 1: Descargar y extraer .tar.gz
	rav1e_tar_url := "https://github.com/xiph/rav1e/archive/refs/tags/v0.7.1.tar.gz"
	rav1e_cmd_tar := exec.Command("curl", "-L", rav1e_tar_url, "-o", "package.tar.gz")
	err := rav1e_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rav1e_zip_url := "https://github.com/xiph/rav1e/archive/refs/tags/v0.7.1.zip"
	rav1e_cmd_zip := exec.Command("curl", "-L", rav1e_zip_url, "-o", "package.zip")
	err = rav1e_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rav1e_bin_url := "https://github.com/xiph/rav1e/archive/refs/tags/v0.7.1.bin"
	rav1e_cmd_bin := exec.Command("curl", "-L", rav1e_bin_url, "-o", "binary.bin")
	err = rav1e_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rav1e_src_url := "https://github.com/xiph/rav1e/archive/refs/tags/v0.7.1.src.tar.gz"
	rav1e_cmd_src := exec.Command("curl", "-L", rav1e_src_url, "-o", "source.tar.gz")
	err = rav1e_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rav1e_cmd_direct := exec.Command("./binary")
	err = rav1e_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cargo-c")
	exec.Command("latte", "install", "cargo-c").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}
