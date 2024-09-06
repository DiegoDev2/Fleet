package main

// Wagyu - Rust library for generating cryptocurrency wallets
// Homepage: https://github.com/howardwu/wagyu

import (
	"fmt"
	
	"os/exec"
)

func installWagyu() {
	// Método 1: Descargar y extraer .tar.gz
	wagyu_tar_url := "https://github.com/howardwu/wagyu/archive/refs/tags/v0.6.1.tar.gz"
	wagyu_cmd_tar := exec.Command("curl", "-L", wagyu_tar_url, "-o", "package.tar.gz")
	err := wagyu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wagyu_zip_url := "https://github.com/howardwu/wagyu/archive/refs/tags/v0.6.1.zip"
	wagyu_cmd_zip := exec.Command("curl", "-L", wagyu_zip_url, "-o", "package.zip")
	err = wagyu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wagyu_bin_url := "https://github.com/howardwu/wagyu/archive/refs/tags/v0.6.1.bin"
	wagyu_cmd_bin := exec.Command("curl", "-L", wagyu_bin_url, "-o", "binary.bin")
	err = wagyu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wagyu_src_url := "https://github.com/howardwu/wagyu/archive/refs/tags/v0.6.1.src.tar.gz"
	wagyu_cmd_src := exec.Command("curl", "-L", wagyu_src_url, "-o", "source.tar.gz")
	err = wagyu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wagyu_cmd_direct := exec.Command("./binary")
	err = wagyu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
