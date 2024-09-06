package main

// V2ray - Platform for building proxies to bypass network restrictions
// Homepage: https://v2fly.org/

import (
	"fmt"
	
	"os/exec"
)

func installV2ray() {
	// Método 1: Descargar y extraer .tar.gz
	v2ray_tar_url := "https://github.com/v2fly/v2ray-core/archive/refs/tags/v5.16.1.tar.gz"
	v2ray_cmd_tar := exec.Command("curl", "-L", v2ray_tar_url, "-o", "package.tar.gz")
	err := v2ray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	v2ray_zip_url := "https://github.com/v2fly/v2ray-core/archive/refs/tags/v5.16.1.zip"
	v2ray_cmd_zip := exec.Command("curl", "-L", v2ray_zip_url, "-o", "package.zip")
	err = v2ray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	v2ray_bin_url := "https://github.com/v2fly/v2ray-core/archive/refs/tags/v5.16.1.bin"
	v2ray_cmd_bin := exec.Command("curl", "-L", v2ray_bin_url, "-o", "binary.bin")
	err = v2ray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	v2ray_src_url := "https://github.com/v2fly/v2ray-core/archive/refs/tags/v5.16.1.src.tar.gz"
	v2ray_cmd_src := exec.Command("curl", "-L", v2ray_src_url, "-o", "source.tar.gz")
	err = v2ray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	v2ray_cmd_direct := exec.Command("./binary")
	err = v2ray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
