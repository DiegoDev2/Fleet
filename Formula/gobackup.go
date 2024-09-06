package main

// Gobackup - CLI tool for backup your databases, files to cloud storages
// Homepage: https://gobackup.github.io

import (
	"fmt"
	
	"os/exec"
)

func installGobackup() {
	// Método 1: Descargar y extraer .tar.gz
	gobackup_tar_url := "https://github.com/gobackup/gobackup/archive/refs/tags/v2.11.2.tar.gz"
	gobackup_cmd_tar := exec.Command("curl", "-L", gobackup_tar_url, "-o", "package.tar.gz")
	err := gobackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gobackup_zip_url := "https://github.com/gobackup/gobackup/archive/refs/tags/v2.11.2.zip"
	gobackup_cmd_zip := exec.Command("curl", "-L", gobackup_zip_url, "-o", "package.zip")
	err = gobackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gobackup_bin_url := "https://github.com/gobackup/gobackup/archive/refs/tags/v2.11.2.bin"
	gobackup_cmd_bin := exec.Command("curl", "-L", gobackup_bin_url, "-o", "binary.bin")
	err = gobackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gobackup_src_url := "https://github.com/gobackup/gobackup/archive/refs/tags/v2.11.2.src.tar.gz"
	gobackup_cmd_src := exec.Command("curl", "-L", gobackup_src_url, "-o", "source.tar.gz")
	err = gobackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gobackup_cmd_direct := exec.Command("./binary")
	err = gobackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
}
