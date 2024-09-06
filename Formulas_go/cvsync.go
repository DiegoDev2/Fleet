package main

// Cvsync - Portable CVS repository synchronization utility
// Homepage: https://www.cvsync.org/

import (
	"fmt"
	
	"os/exec"
)

func installCvsync() {
	// Método 1: Descargar y extraer .tar.gz
	cvsync_tar_url := "https://www.cvsync.org/dist/cvsync-0.24.19.tar.gz"
	cvsync_cmd_tar := exec.Command("curl", "-L", cvsync_tar_url, "-o", "package.tar.gz")
	err := cvsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cvsync_zip_url := "https://www.cvsync.org/dist/cvsync-0.24.19.zip"
	cvsync_cmd_zip := exec.Command("curl", "-L", cvsync_zip_url, "-o", "package.zip")
	err = cvsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cvsync_bin_url := "https://www.cvsync.org/dist/cvsync-0.24.19.bin"
	cvsync_cmd_bin := exec.Command("curl", "-L", cvsync_bin_url, "-o", "binary.bin")
	err = cvsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cvsync_src_url := "https://www.cvsync.org/dist/cvsync-0.24.19.src.tar.gz"
	cvsync_cmd_src := exec.Command("curl", "-L", cvsync_src_url, "-o", "source.tar.gz")
	err = cvsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cvsync_cmd_direct := exec.Command("./binary")
	err = cvsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
