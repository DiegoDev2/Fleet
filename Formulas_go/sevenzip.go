package main

// Sevenzip - 7-Zip is a file archiver with a high compression ratio
// Homepage: https://7-zip.org

import (
	"fmt"
	
	"os/exec"
)

func installSevenzip() {
	// Método 1: Descargar y extraer .tar.gz
	sevenzip_tar_url := "https://7-zip.org/a/7z2408-src.tar.xz"
	sevenzip_cmd_tar := exec.Command("curl", "-L", sevenzip_tar_url, "-o", "package.tar.gz")
	err := sevenzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sevenzip_zip_url := "https://7-zip.org/a/7z2408-src.tar.xz"
	sevenzip_cmd_zip := exec.Command("curl", "-L", sevenzip_zip_url, "-o", "package.zip")
	err = sevenzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sevenzip_bin_url := "https://7-zip.org/a/7z2408-src.tar.xz"
	sevenzip_cmd_bin := exec.Command("curl", "-L", sevenzip_bin_url, "-o", "binary.bin")
	err = sevenzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sevenzip_src_url := "https://7-zip.org/a/7z2408-src.tar.xz"
	sevenzip_cmd_src := exec.Command("curl", "-L", sevenzip_src_url, "-o", "source.tar.gz")
	err = sevenzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sevenzip_cmd_direct := exec.Command("./binary")
	err = sevenzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
