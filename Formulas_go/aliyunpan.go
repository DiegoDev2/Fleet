package main

// Aliyunpan - Command-line client tool for Alibaba aDrive disk
// Homepage: https://github.com/tickstep/aliyunpan

import (
	"fmt"
	
	"os/exec"
)

func installAliyunpan() {
	// Método 1: Descargar y extraer .tar.gz
	aliyunpan_tar_url := "https://github.com/tickstep/aliyunpan/archive/refs/tags/v0.3.3.tar.gz"
	aliyunpan_cmd_tar := exec.Command("curl", "-L", aliyunpan_tar_url, "-o", "package.tar.gz")
	err := aliyunpan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aliyunpan_zip_url := "https://github.com/tickstep/aliyunpan/archive/refs/tags/v0.3.3.zip"
	aliyunpan_cmd_zip := exec.Command("curl", "-L", aliyunpan_zip_url, "-o", "package.zip")
	err = aliyunpan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aliyunpan_bin_url := "https://github.com/tickstep/aliyunpan/archive/refs/tags/v0.3.3.bin"
	aliyunpan_cmd_bin := exec.Command("curl", "-L", aliyunpan_bin_url, "-o", "binary.bin")
	err = aliyunpan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aliyunpan_src_url := "https://github.com/tickstep/aliyunpan/archive/refs/tags/v0.3.3.src.tar.gz"
	aliyunpan_cmd_src := exec.Command("curl", "-L", aliyunpan_src_url, "-o", "source.tar.gz")
	err = aliyunpan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aliyunpan_cmd_direct := exec.Command("./binary")
	err = aliyunpan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
