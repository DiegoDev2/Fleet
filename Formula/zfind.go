package main

// Zfind - Search for files (even inside tar/zip/7z/rar) using a SQL-WHERE filter
// Homepage: https://github.com/laktak/zfind

import (
	"fmt"
	
	"os/exec"
)

func installZfind() {
	// Método 1: Descargar y extraer .tar.gz
	zfind_tar_url := "https://github.com/laktak/zfind/archive/refs/tags/v0.4.5.tar.gz"
	zfind_cmd_tar := exec.Command("curl", "-L", zfind_tar_url, "-o", "package.tar.gz")
	err := zfind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zfind_zip_url := "https://github.com/laktak/zfind/archive/refs/tags/v0.4.5.zip"
	zfind_cmd_zip := exec.Command("curl", "-L", zfind_zip_url, "-o", "package.zip")
	err = zfind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zfind_bin_url := "https://github.com/laktak/zfind/archive/refs/tags/v0.4.5.bin"
	zfind_cmd_bin := exec.Command("curl", "-L", zfind_bin_url, "-o", "binary.bin")
	err = zfind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zfind_src_url := "https://github.com/laktak/zfind/archive/refs/tags/v0.4.5.src.tar.gz"
	zfind_cmd_src := exec.Command("curl", "-L", zfind_src_url, "-o", "source.tar.gz")
	err = zfind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zfind_cmd_direct := exec.Command("./binary")
	err = zfind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
