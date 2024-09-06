package main

// Innotop - Top clone for MySQL
// Homepage: https://github.com/innotop/innotop/

import (
	"fmt"
	
	"os/exec"
)

func installInnotop() {
	// Método 1: Descargar y extraer .tar.gz
	innotop_tar_url := "https://github.com/innotop/innotop/archive/refs/tags/v1.13.0.tar.gz"
	innotop_cmd_tar := exec.Command("curl", "-L", innotop_tar_url, "-o", "package.tar.gz")
	err := innotop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	innotop_zip_url := "https://github.com/innotop/innotop/archive/refs/tags/v1.13.0.zip"
	innotop_cmd_zip := exec.Command("curl", "-L", innotop_zip_url, "-o", "package.zip")
	err = innotop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	innotop_bin_url := "https://github.com/innotop/innotop/archive/refs/tags/v1.13.0.bin"
	innotop_cmd_bin := exec.Command("curl", "-L", innotop_bin_url, "-o", "binary.bin")
	err = innotop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	innotop_src_url := "https://github.com/innotop/innotop/archive/refs/tags/v1.13.0.src.tar.gz"
	innotop_cmd_src := exec.Command("curl", "-L", innotop_src_url, "-o", "source.tar.gz")
	err = innotop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	innotop_cmd_direct := exec.Command("./binary")
	err = innotop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
