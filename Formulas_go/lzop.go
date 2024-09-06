package main

// Lzop - File compressor
// Homepage: https://www.lzop.org/

import (
	"fmt"
	
	"os/exec"
)

func installLzop() {
	// Método 1: Descargar y extraer .tar.gz
	lzop_tar_url := "https://www.lzop.org/download/lzop-1.04.tar.gz"
	lzop_cmd_tar := exec.Command("curl", "-L", lzop_tar_url, "-o", "package.tar.gz")
	err := lzop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lzop_zip_url := "https://www.lzop.org/download/lzop-1.04.zip"
	lzop_cmd_zip := exec.Command("curl", "-L", lzop_zip_url, "-o", "package.zip")
	err = lzop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lzop_bin_url := "https://www.lzop.org/download/lzop-1.04.bin"
	lzop_cmd_bin := exec.Command("curl", "-L", lzop_bin_url, "-o", "binary.bin")
	err = lzop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lzop_src_url := "https://www.lzop.org/download/lzop-1.04.src.tar.gz"
	lzop_cmd_src := exec.Command("curl", "-L", lzop_src_url, "-o", "source.tar.gz")
	err = lzop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lzop_cmd_direct := exec.Command("./binary")
	err = lzop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
}
