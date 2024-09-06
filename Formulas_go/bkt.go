package main

// Bkt - CLI utility for caching the output of subprocesses
// Homepage: https://www.bkt.rs

import (
	"fmt"
	
	"os/exec"
)

func installBkt() {
	// Método 1: Descargar y extraer .tar.gz
	bkt_tar_url := "https://github.com/dimo414/bkt/archive/refs/tags/0.8.0.tar.gz"
	bkt_cmd_tar := exec.Command("curl", "-L", bkt_tar_url, "-o", "package.tar.gz")
	err := bkt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bkt_zip_url := "https://github.com/dimo414/bkt/archive/refs/tags/0.8.0.zip"
	bkt_cmd_zip := exec.Command("curl", "-L", bkt_zip_url, "-o", "package.zip")
	err = bkt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bkt_bin_url := "https://github.com/dimo414/bkt/archive/refs/tags/0.8.0.bin"
	bkt_cmd_bin := exec.Command("curl", "-L", bkt_bin_url, "-o", "binary.bin")
	err = bkt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bkt_src_url := "https://github.com/dimo414/bkt/archive/refs/tags/0.8.0.src.tar.gz"
	bkt_cmd_src := exec.Command("curl", "-L", bkt_src_url, "-o", "source.tar.gz")
	err = bkt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bkt_cmd_direct := exec.Command("./binary")
	err = bkt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
