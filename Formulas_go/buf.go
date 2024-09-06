package main

// Buf - New way of working with Protocol Buffers
// Homepage: https://github.com/bufbuild/buf

import (
	"fmt"
	
	"os/exec"
)

func installBuf() {
	// Método 1: Descargar y extraer .tar.gz
	buf_tar_url := "https://github.com/bufbuild/buf/archive/refs/tags/v1.40.0.tar.gz"
	buf_cmd_tar := exec.Command("curl", "-L", buf_tar_url, "-o", "package.tar.gz")
	err := buf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	buf_zip_url := "https://github.com/bufbuild/buf/archive/refs/tags/v1.40.0.zip"
	buf_cmd_zip := exec.Command("curl", "-L", buf_zip_url, "-o", "package.zip")
	err = buf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	buf_bin_url := "https://github.com/bufbuild/buf/archive/refs/tags/v1.40.0.bin"
	buf_cmd_bin := exec.Command("curl", "-L", buf_bin_url, "-o", "binary.bin")
	err = buf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	buf_src_url := "https://github.com/bufbuild/buf/archive/refs/tags/v1.40.0.src.tar.gz"
	buf_cmd_src := exec.Command("curl", "-L", buf_src_url, "-o", "source.tar.gz")
	err = buf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	buf_cmd_direct := exec.Command("./binary")
	err = buf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
