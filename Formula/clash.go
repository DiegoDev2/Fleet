package main

// Clash - Rule-based tunnel in Go
// Homepage: https://github.com/Dreamacro/clash

import (
	"fmt"
	
	"os/exec"
)

func installClash() {
	// Método 1: Descargar y extraer .tar.gz
	clash_tar_url := "https://github.com/Dreamacro/clash/archive/refs/tags/v1.18.0.tar.gz"
	clash_cmd_tar := exec.Command("curl", "-L", clash_tar_url, "-o", "package.tar.gz")
	err := clash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clash_zip_url := "https://github.com/Dreamacro/clash/archive/refs/tags/v1.18.0.zip"
	clash_cmd_zip := exec.Command("curl", "-L", clash_zip_url, "-o", "package.zip")
	err = clash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clash_bin_url := "https://github.com/Dreamacro/clash/archive/refs/tags/v1.18.0.bin"
	clash_cmd_bin := exec.Command("curl", "-L", clash_bin_url, "-o", "binary.bin")
	err = clash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clash_src_url := "https://github.com/Dreamacro/clash/archive/refs/tags/v1.18.0.src.tar.gz"
	clash_cmd_src := exec.Command("curl", "-L", clash_src_url, "-o", "source.tar.gz")
	err = clash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clash_cmd_direct := exec.Command("./binary")
	err = clash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: shadowsocks-libev")
	exec.Command("latte", "install", "shadowsocks-libev").Run()
}
