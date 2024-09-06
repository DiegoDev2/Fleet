package main

// BaidupcsGo - Terminal utility for Baidu Network Disk
// Homepage: https://github.com/qjfoidnh/BaiduPCS-Go

import (
	"fmt"
	
	"os/exec"
)

func installBaidupcsGo() {
	// Método 1: Descargar y extraer .tar.gz
	baidupcsgo_tar_url := "https://github.com/qjfoidnh/BaiduPCS-Go/archive/refs/tags/v3.9.5.tar.gz"
	baidupcsgo_cmd_tar := exec.Command("curl", "-L", baidupcsgo_tar_url, "-o", "package.tar.gz")
	err := baidupcsgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	baidupcsgo_zip_url := "https://github.com/qjfoidnh/BaiduPCS-Go/archive/refs/tags/v3.9.5.zip"
	baidupcsgo_cmd_zip := exec.Command("curl", "-L", baidupcsgo_zip_url, "-o", "package.zip")
	err = baidupcsgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	baidupcsgo_bin_url := "https://github.com/qjfoidnh/BaiduPCS-Go/archive/refs/tags/v3.9.5.bin"
	baidupcsgo_cmd_bin := exec.Command("curl", "-L", baidupcsgo_bin_url, "-o", "binary.bin")
	err = baidupcsgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	baidupcsgo_src_url := "https://github.com/qjfoidnh/BaiduPCS-Go/archive/refs/tags/v3.9.5.src.tar.gz"
	baidupcsgo_cmd_src := exec.Command("curl", "-L", baidupcsgo_src_url, "-o", "source.tar.gz")
	err = baidupcsgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	baidupcsgo_cmd_direct := exec.Command("./binary")
	err = baidupcsgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
	exec.Command("latte", "install", "go@1.22").Run()
}
