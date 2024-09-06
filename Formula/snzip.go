package main

// Snzip - Compression/decompression tool based on snappy
// Homepage: https://github.com/kubo/snzip

import (
	"fmt"
	
	"os/exec"
)

func installSnzip() {
	// Método 1: Descargar y extraer .tar.gz
	snzip_tar_url := "https://github.com/kubo/snzip/releases/download/v1.0.5/snzip-1.0.5.tar.gz"
	snzip_cmd_tar := exec.Command("curl", "-L", snzip_tar_url, "-o", "package.tar.gz")
	err := snzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snzip_zip_url := "https://github.com/kubo/snzip/releases/download/v1.0.5/snzip-1.0.5.zip"
	snzip_cmd_zip := exec.Command("curl", "-L", snzip_zip_url, "-o", "package.zip")
	err = snzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snzip_bin_url := "https://github.com/kubo/snzip/releases/download/v1.0.5/snzip-1.0.5.bin"
	snzip_cmd_bin := exec.Command("curl", "-L", snzip_bin_url, "-o", "binary.bin")
	err = snzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snzip_src_url := "https://github.com/kubo/snzip/releases/download/v1.0.5/snzip-1.0.5.src.tar.gz"
	snzip_cmd_src := exec.Command("curl", "-L", snzip_src_url, "-o", "source.tar.gz")
	err = snzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snzip_cmd_direct := exec.Command("./binary")
	err = snzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
}
