package main

// Fdclone - Console-based file manager
// Homepage: https://hp.vector.co.jp/authors/VA012337/soft/fd/

import (
	"fmt"
	
	"os/exec"
)

func installFdclone() {
	// Método 1: Descargar y extraer .tar.gz
	fdclone_tar_url := "http://www.unixusers.net/src/fdclone/FD-3.01j.tar.gz"
	fdclone_cmd_tar := exec.Command("curl", "-L", fdclone_tar_url, "-o", "package.tar.gz")
	err := fdclone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdclone_zip_url := "http://www.unixusers.net/src/fdclone/FD-3.01j.zip"
	fdclone_cmd_zip := exec.Command("curl", "-L", fdclone_zip_url, "-o", "package.zip")
	err = fdclone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdclone_bin_url := "http://www.unixusers.net/src/fdclone/FD-3.01j.bin"
	fdclone_cmd_bin := exec.Command("curl", "-L", fdclone_bin_url, "-o", "binary.bin")
	err = fdclone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdclone_src_url := "http://www.unixusers.net/src/fdclone/FD-3.01j.src.tar.gz"
	fdclone_cmd_src := exec.Command("curl", "-L", fdclone_src_url, "-o", "source.tar.gz")
	err = fdclone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdclone_cmd_direct := exec.Command("./binary")
	err = fdclone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nkf")
exec.Command("latte", "install", "nkf")
}
