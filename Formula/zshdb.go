package main

// Zshdb - Debugger for zsh
// Homepage: https://github.com/rocky/zshdb

import (
	"fmt"
	
	"os/exec"
)

func installZshdb() {
	// Método 1: Descargar y extraer .tar.gz
	zshdb_tar_url := "https://downloads.sourceforge.net/project/bashdb/zshdb/1.1.4/zshdb-1.1.4.tar.gz"
	zshdb_cmd_tar := exec.Command("curl", "-L", zshdb_tar_url, "-o", "package.tar.gz")
	err := zshdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshdb_zip_url := "https://downloads.sourceforge.net/project/bashdb/zshdb/1.1.4/zshdb-1.1.4.zip"
	zshdb_cmd_zip := exec.Command("curl", "-L", zshdb_zip_url, "-o", "package.zip")
	err = zshdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshdb_bin_url := "https://downloads.sourceforge.net/project/bashdb/zshdb/1.1.4/zshdb-1.1.4.bin"
	zshdb_cmd_bin := exec.Command("curl", "-L", zshdb_bin_url, "-o", "binary.bin")
	err = zshdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshdb_src_url := "https://downloads.sourceforge.net/project/bashdb/zshdb/1.1.4/zshdb-1.1.4.src.tar.gz"
	zshdb_cmd_src := exec.Command("curl", "-L", zshdb_src_url, "-o", "source.tar.gz")
	err = zshdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshdb_cmd_direct := exec.Command("./binary")
	err = zshdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: zsh")
	exec.Command("latte", "install", "zsh").Run()
}
