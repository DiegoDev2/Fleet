package main

// Ncmdump - Convert Netease Cloud Music ncm files to mp3/flac files
// Homepage: https://github.com/taurusxin/ncmdump

import (
	"fmt"
	
	"os/exec"
)

func installNcmdump() {
	// Método 1: Descargar y extraer .tar.gz
	ncmdump_tar_url := "https://github.com/taurusxin/ncmdump/archive/refs/tags/1.2.1.tar.gz"
	ncmdump_cmd_tar := exec.Command("curl", "-L", ncmdump_tar_url, "-o", "package.tar.gz")
	err := ncmdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncmdump_zip_url := "https://github.com/taurusxin/ncmdump/archive/refs/tags/1.2.1.zip"
	ncmdump_cmd_zip := exec.Command("curl", "-L", ncmdump_zip_url, "-o", "package.zip")
	err = ncmdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncmdump_bin_url := "https://github.com/taurusxin/ncmdump/archive/refs/tags/1.2.1.bin"
	ncmdump_cmd_bin := exec.Command("curl", "-L", ncmdump_bin_url, "-o", "binary.bin")
	err = ncmdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncmdump_src_url := "https://github.com/taurusxin/ncmdump/archive/refs/tags/1.2.1.src.tar.gz"
	ncmdump_cmd_src := exec.Command("curl", "-L", ncmdump_src_url, "-o", "source.tar.gz")
	err = ncmdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncmdump_cmd_direct := exec.Command("./binary")
	err = ncmdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
}
