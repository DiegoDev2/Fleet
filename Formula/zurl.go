package main

// Zurl - HTTP and WebSocket client worker with ZeroMQ interface
// Homepage: https://github.com/fanout/zurl

import (
	"fmt"
	
	"os/exec"
)

func installZurl() {
	// Método 1: Descargar y extraer .tar.gz
	zurl_tar_url := "https://github.com/fanout/zurl/releases/download/v1.12.0/zurl-1.12.0.tar.bz2"
	zurl_cmd_tar := exec.Command("curl", "-L", zurl_tar_url, "-o", "package.tar.gz")
	err := zurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zurl_zip_url := "https://github.com/fanout/zurl/releases/download/v1.12.0/zurl-1.12.0.tar.bz2"
	zurl_cmd_zip := exec.Command("curl", "-L", zurl_zip_url, "-o", "package.zip")
	err = zurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zurl_bin_url := "https://github.com/fanout/zurl/releases/download/v1.12.0/zurl-1.12.0.tar.bz2"
	zurl_cmd_bin := exec.Command("curl", "-L", zurl_bin_url, "-o", "binary.bin")
	err = zurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zurl_src_url := "https://github.com/fanout/zurl/releases/download/v1.12.0/zurl-1.12.0.tar.bz2"
	zurl_cmd_src := exec.Command("curl", "-L", zurl_src_url, "-o", "source.tar.gz")
	err = zurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zurl_cmd_direct := exec.Command("./binary")
	err = zurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
