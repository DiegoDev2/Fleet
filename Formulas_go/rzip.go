package main

// Rzip - File compression tool (like gzip or bzip2)
// Homepage: https://rzip.samba.org/

import (
	"fmt"
	
	"os/exec"
)

func installRzip() {
	// Método 1: Descargar y extraer .tar.gz
	rzip_tar_url := "https://rzip.samba.org/ftp/rzip/rzip-2.1.tar.gz"
	rzip_cmd_tar := exec.Command("curl", "-L", rzip_tar_url, "-o", "package.tar.gz")
	err := rzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rzip_zip_url := "https://rzip.samba.org/ftp/rzip/rzip-2.1.zip"
	rzip_cmd_zip := exec.Command("curl", "-L", rzip_zip_url, "-o", "package.zip")
	err = rzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rzip_bin_url := "https://rzip.samba.org/ftp/rzip/rzip-2.1.bin"
	rzip_cmd_bin := exec.Command("curl", "-L", rzip_bin_url, "-o", "binary.bin")
	err = rzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rzip_src_url := "https://rzip.samba.org/ftp/rzip/rzip-2.1.src.tar.gz"
	rzip_cmd_src := exec.Command("curl", "-L", rzip_src_url, "-o", "source.tar.gz")
	err = rzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rzip_cmd_direct := exec.Command("./binary")
	err = rzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
