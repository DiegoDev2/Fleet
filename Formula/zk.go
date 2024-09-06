package main

// Zk - Plain text note-taking assistant
// Homepage: https://github.com/zk-org/zk

import (
	"fmt"
	
	"os/exec"
)

func installZk() {
	// Método 1: Descargar y extraer .tar.gz
	zk_tar_url := "https://github.com/zk-org/zk/archive/refs/tags/v0.14.1.tar.gz"
	zk_cmd_tar := exec.Command("curl", "-L", zk_tar_url, "-o", "package.tar.gz")
	err := zk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zk_zip_url := "https://github.com/zk-org/zk/archive/refs/tags/v0.14.1.zip"
	zk_cmd_zip := exec.Command("curl", "-L", zk_zip_url, "-o", "package.zip")
	err = zk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zk_bin_url := "https://github.com/zk-org/zk/archive/refs/tags/v0.14.1.bin"
	zk_cmd_bin := exec.Command("curl", "-L", zk_bin_url, "-o", "binary.bin")
	err = zk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zk_src_url := "https://github.com/zk-org/zk/archive/refs/tags/v0.14.1.src.tar.gz"
	zk_cmd_src := exec.Command("curl", "-L", zk_src_url, "-o", "source.tar.gz")
	err = zk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zk_cmd_direct := exec.Command("./binary")
	err = zk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
}
