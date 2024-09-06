package main

// Mpack - MIME mail packing and unpacking
// Homepage: https://web.archive.org/web/20190220145801/ftp.andrew.cmu.edu/pub/mpack/

import (
	"fmt"
	
	"os/exec"
)

func installMpack() {
	// Método 1: Descargar y extraer .tar.gz
	mpack_tar_url := "https://ftp.gwdg.de/pub/misc/mpack/mpack-1.6.tar.gz"
	mpack_cmd_tar := exec.Command("curl", "-L", mpack_tar_url, "-o", "package.tar.gz")
	err := mpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpack_zip_url := "https://ftp.gwdg.de/pub/misc/mpack/mpack-1.6.zip"
	mpack_cmd_zip := exec.Command("curl", "-L", mpack_zip_url, "-o", "package.zip")
	err = mpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpack_bin_url := "https://ftp.gwdg.de/pub/misc/mpack/mpack-1.6.bin"
	mpack_cmd_bin := exec.Command("curl", "-L", mpack_bin_url, "-o", "binary.bin")
	err = mpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpack_src_url := "https://ftp.gwdg.de/pub/misc/mpack/mpack-1.6.src.tar.gz"
	mpack_cmd_src := exec.Command("curl", "-L", mpack_src_url, "-o", "source.tar.gz")
	err = mpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpack_cmd_direct := exec.Command("./binary")
	err = mpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
