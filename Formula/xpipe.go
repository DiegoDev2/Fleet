package main

// Xpipe - Split input and feed it into the given utility
// Homepage: https://www.netmeister.org/apps/xpipe.html

import (
	"fmt"
	
	"os/exec"
)

func installXpipe() {
	// Método 1: Descargar y extraer .tar.gz
	xpipe_tar_url := "https://www.netmeister.org/apps/xpipe-2.2.tar.gz"
	xpipe_cmd_tar := exec.Command("curl", "-L", xpipe_tar_url, "-o", "package.tar.gz")
	err := xpipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xpipe_zip_url := "https://www.netmeister.org/apps/xpipe-2.2.zip"
	xpipe_cmd_zip := exec.Command("curl", "-L", xpipe_zip_url, "-o", "package.zip")
	err = xpipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xpipe_bin_url := "https://www.netmeister.org/apps/xpipe-2.2.bin"
	xpipe_cmd_bin := exec.Command("curl", "-L", xpipe_bin_url, "-o", "binary.bin")
	err = xpipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xpipe_src_url := "https://www.netmeister.org/apps/xpipe-2.2.src.tar.gz"
	xpipe_cmd_src := exec.Command("curl", "-L", xpipe_src_url, "-o", "source.tar.gz")
	err = xpipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xpipe_cmd_direct := exec.Command("./binary")
	err = xpipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libbsd")
	exec.Command("latte", "install", "libbsd").Run()
}
