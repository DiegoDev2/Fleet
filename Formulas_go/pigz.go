package main

// Pigz - Parallel gzip
// Homepage: https://zlib.net/pigz/

import (
	"fmt"
	
	"os/exec"
)

func installPigz() {
	// Método 1: Descargar y extraer .tar.gz
	pigz_tar_url := "https://zlib.net/pigz/pigz-2.8.tar.gz"
	pigz_cmd_tar := exec.Command("curl", "-L", pigz_tar_url, "-o", "package.tar.gz")
	err := pigz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pigz_zip_url := "https://zlib.net/pigz/pigz-2.8.zip"
	pigz_cmd_zip := exec.Command("curl", "-L", pigz_zip_url, "-o", "package.zip")
	err = pigz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pigz_bin_url := "https://zlib.net/pigz/pigz-2.8.bin"
	pigz_cmd_bin := exec.Command("curl", "-L", pigz_bin_url, "-o", "binary.bin")
	err = pigz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pigz_src_url := "https://zlib.net/pigz/pigz-2.8.src.tar.gz"
	pigz_cmd_src := exec.Command("curl", "-L", pigz_src_url, "-o", "source.tar.gz")
	err = pigz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pigz_cmd_direct := exec.Command("./binary")
	err = pigz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zopfli")
exec.Command("latte", "install", "zopfli")
}
