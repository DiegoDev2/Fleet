package main

// Bao - Implementation of BLAKE3 verified streaming
// Homepage: https://github.com/oconnor663/bao

import (
	"fmt"
	
	"os/exec"
)

func installBao() {
	// Método 1: Descargar y extraer .tar.gz
	bao_tar_url := "https://github.com/oconnor663/bao/archive/refs/tags/0.12.1.tar.gz"
	bao_cmd_tar := exec.Command("curl", "-L", bao_tar_url, "-o", "package.tar.gz")
	err := bao_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bao_zip_url := "https://github.com/oconnor663/bao/archive/refs/tags/0.12.1.zip"
	bao_cmd_zip := exec.Command("curl", "-L", bao_zip_url, "-o", "package.zip")
	err = bao_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bao_bin_url := "https://github.com/oconnor663/bao/archive/refs/tags/0.12.1.bin"
	bao_cmd_bin := exec.Command("curl", "-L", bao_bin_url, "-o", "binary.bin")
	err = bao_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bao_src_url := "https://github.com/oconnor663/bao/archive/refs/tags/0.12.1.src.tar.gz"
	bao_cmd_src := exec.Command("curl", "-L", bao_src_url, "-o", "source.tar.gz")
	err = bao_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bao_cmd_direct := exec.Command("./binary")
	err = bao_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
