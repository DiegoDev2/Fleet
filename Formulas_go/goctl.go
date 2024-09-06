package main

// Goctl - Generates server-side and client-side code for web and RPC services
// Homepage: https://go-zero.dev

import (
	"fmt"
	
	"os/exec"
)

func installGoctl() {
	// Método 1: Descargar y extraer .tar.gz
	goctl_tar_url := "https://github.com/zeromicro/go-zero/archive/refs/tags/tools/goctl/v1.7.2.tar.gz"
	goctl_cmd_tar := exec.Command("curl", "-L", goctl_tar_url, "-o", "package.tar.gz")
	err := goctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goctl_zip_url := "https://github.com/zeromicro/go-zero/archive/refs/tags/tools/goctl/v1.7.2.zip"
	goctl_cmd_zip := exec.Command("curl", "-L", goctl_zip_url, "-o", "package.zip")
	err = goctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goctl_bin_url := "https://github.com/zeromicro/go-zero/archive/refs/tags/tools/goctl/v1.7.2.bin"
	goctl_cmd_bin := exec.Command("curl", "-L", goctl_bin_url, "-o", "binary.bin")
	err = goctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goctl_src_url := "https://github.com/zeromicro/go-zero/archive/refs/tags/tools/goctl/v1.7.2.src.tar.gz"
	goctl_cmd_src := exec.Command("curl", "-L", goctl_src_url, "-o", "source.tar.gz")
	err = goctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goctl_cmd_direct := exec.Command("./binary")
	err = goctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
